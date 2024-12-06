# Техники управления памятью в Go

1. [Цель](#цель)
2. [Кастомный пул на базе каналов](#кастомный-пул-на-базе-каналов)
   - [Реализация](#реализация-кастомного-пула)
   - [Недостатки](#недостатки-кастомного-пула)
3. [Sync.Pool](#sync-pool)
   - [Реализация](#реализация-sync-pool)
   - [Преимущества](#преимущества-sync-pool)
   - [Недостатки](#недостатки-sync-pool)
4. [Memory Arena с Sync.Pool](#memory-arena-с-sync-pool)
   - [Реализация](#реализация-memory-arena)
   - [Плюсы и минусы](#плюсы-и-минусы-memory-arena)
5. [Что выбрать?](#что-выбрать)

---

## Цель <a id="цель"></a>

Все описанные техники направлены на снижение частоты работы сборщика мусора (GC). Частые паузы GC (Stop-The-World, STW) происходят при превышении лимита памяти, что может негативно сказаться на производительности.

---

## Кастомный пул на базе каналов <a id="кастомный-пул-на-базе-каналов"></a>

### Реализация <a id="реализация-кастомного-пула"></a>

```go
type Resource interface {
	Clean()
	Cap() int
}

const contentCap = 10

type Pool struct {
	buff     chan Resource
	factory  func() (Resource, error)
	isClosed bool
	mu       sync.Mutex
}

func NewPool(buffSize int, factory func() (Resource, error)) (*Pool, error) {
	if buffSize <= 0 {
		return nil, errors.New("buffer size must be greater than 0")
	}

	return &Pool{
		buff:    make(chan Resource, buffSize),
		factory: factory,
	}, nil
}

func (p *Pool) Get() (Resource, error) {
	p.mu.Lock()
	if p.isClosed {
		p.mu.Unlock()
		return nil, errors.New("pool is closed")
	}
	p.mu.Unlock()

	select {
	case res, ok := <-p.buff:
		if ok {
			if res.Cap() > contentCap {
				res.Clean() 
			}
			return res, nil
		}
	default:
		return p.factory()
	}
}

func (p *Pool) Put(res Resource) error {
	if res.Cap() > contentCap {
		return errors.New("resource content exceeds capacity limit")
	}
	res.Clean()

	p.mu.Lock()
	defer p.mu.Unlock()

	if p.isClosed {
		return errors.New("cannot put resource into a closed pool")
	}

	select {
	case p.buff <- res:
		return nil
	default:
		return errors.New("pool is full, resource discarded")
	}
}

func (p *Pool) Close() {
	p.mu.Lock()
	defer p.mu.Unlock()

	if !p.isClosed {
		p.isClosed = true
		close(p.buff)
		for res := range p.buff {
			res.Clean()
		}
	}
}
```

### Недостатки <a id="недостатки-кастомного-пула"></a>

1. Необходимо заранее определить количество буферов.
2. Возврат переразмеренных (увеличенных) ресурсов может привести к неэффективному использованию памяти.

---

## Sync.Pool <a id="sync-pool"></a>

### Реализация <a id="реализация-sync-pool"></a>

```go
const contentCap = 10

var heapBuffPool = sync.Pool{
	New: func() any {
		b := make([]byte, 0, contentCap)
		return &b
	},
}

type SyncPool struct{}

func (p *SyncPool) GetBytes() *[]byte {
	return heapBuffPool.Get().(*[]byte)
}

func (p *SyncPool) PutBytes(b *[]byte) {
	if cap(*b) > contentCap {
		return
	}
	*b = (*b)[:0]
	heapBuffPool.Put(b)
}
```

### Преимущества <a id="преимущества-sync-pool"></a>

1. Буферы выделяются динамически по мере необходимости.
2. Более эффективен, чем кастомная реализация в большинстве случаев.

### Недостатки <a id="недостатки-sync-pool"></a>

1. Если включен GC, неиспользуемые объекты в пуле будут очищены. Непригоден для использования в качестве кеша.

---

## Memory Arena с Sync.Pool <a id="memory-arena-с-sync-pool"></a>

### Реализация <a id="реализация-memory-arena"></a>

```go
const contentCap = 10

var arenaBuffPool = &sync.Pool{
	New: func() interface{} {
		b := arena.MakeSlice[byte](a, 0, contentCap)
		return &b
	},
}

type ArenaSyncPool struct{}

func (p *ArenaSyncPool) GetBytes() *[]byte {
	return arenaBuffPool.Get().(*[]byte)
}

func (p *ArenaSyncPool) PutBytes(b *[]byte) {
	if cap(*b) > contentCap {
		return
	}
	*b = (*b)[:0]
	arenaBuffPool.Put(b)
}
```

### Плюсы и минусы <a id="плюсы-и-минусы-memory-arena"></a>

**Плюсы**:
- Работает в области памяти, не контролируемой GC.

**Минусы**:
1. Требуется знание числа буферов и их размеров.

---

## Что выбрать? <a id="что-выбрать"></a>

- **Выделение на стеке**: Если выделение памяти минимально, и вам не нужны пулы, можно обойтись стеком.
- **Sync.Pool с подходом на основе кучи**: Используйте, если выделение памяти велико и количество буферов неизвестно. Учтите, что неиспользуемая память может быть очищена GC.
- **Кастомный пул на каналах**: Подходит, если количество буферов и объем памяти известны заранее.
- **Sync.Pool с memory arena**: Идеально, если память нужна постоянно, размеры буферов практически одинаковы, а количество буферов неизвестно.