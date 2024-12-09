# Интерфейсы Go и привязка методов

## 1. **Представление интерфейсов в памяти**
Интерфейс в Go занимает 16 байт: указатель на тип и указатель на значение. Тип содержит set методов.

### Как методы привязываются к структурам:
```go
type Method struct {
    name    string           // Имя метода
    handler func(interface{}) // Реализация метода
}

type TypeDescriptor struct {
    name         string   // Имя типа
    valueMethods []Method // Методы для значений
    ptrMethods   []Method // Методы для указателей
}
```

## 2. Сравнение наборов методов структур и интерфейсов

### Пример 1: Интерфейс `Doer` и структура `Object`
```go
func main() {
    doer := NewDoer()
    if doer == nil { // false, потому что val == nil && type == *Object
        log.Fatalln("Doer object is nil, terminating...")
    }

    doer.do()
}

type Doer interface {
    do()
}

type Object struct{}

func NewDoer() Doer {
    var obj *Object
    return &obj
}

func (o Object) do() {
    fmt.Println("Doer is doing something...")
}
```

### Объяснение:
- Программа **завершится с паникой**, потому что метод `func (o Object) do()` определен для значения, но интерфейс `Doer` требует привязки к указателю `*Object`.
- Исправления:
  1. Изменить метод `do()` на привязку к указателю:
     ```go
     func (o *Object) do() {
         fmt.Println("Doer is doing something...")
     }
     ```
  2. Альтернативно, определить `obj` как значение, а не указатель:
     ```go
     func NewDoer() Doer {
         var obj Object
         return &obj
     }
     ```

---

### Пример 2: Интерфейс `Worker` и структура `User`
```go
type Worker interface {
    Do()
}

type User struct{}

func (u *User) Do() { // Если не указать "*" результат будет: true true false паника
    fmt.Println("Doing...")
}

func main() {
    var w Worker
    fmt.Println(w == nil) // true

    var u *User
    fmt.Println(u == nil) // true

    w = u
    fmt.Println(w == nil) // false

    w.Do() // Doing...

    var u2 User // Не указатель
    fmt.Println(u2 == nil) // Ошибка компиляции: вызов value method через nil-указатель
}
```

### Объяснение:
1. `w == nil` возвращает `true`, потому что `w` — неинициализированный интерфейс.
2. `u == nil` возвращает `true`, потому что `u` — nil-указатель.
3. После присваивания `u` (nil-указатель) интерфейсу `w`, `w == nil` возвращает `false`, потому что интерфейс теперь содержит тип (`*User`), но значение остается `nil`.
4. Вызов `w.Do()` работает, потому что метод определен для `*User`.

### Важные замечания:
- Сравнение интерфейсов с `nil`:
  - Интерфейс считается `nil`, если **и тип, и значение равны `nil`**.
  - Если интерфейс содержит тип, но значение `nil`, он не считается `nil`.



```go
package main

import "fmt"

type Worker interface {
	Do()
}

type User struct{}

func (u User) Do() {
	fmt.Println("Doing...")
}

func main() {
	var w Worker
	fmt.Println(w == nil)

	var u = &User{}
	fmt.Println(u == nil)

	w = u
	fmt.Println(w == nil)

	w.Do()
}
```