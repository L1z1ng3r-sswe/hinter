1. **Интерфейс** — занимает 16 байт: указатель на тип и указатель на значение. Тип содержит динамический массив методов.

**Как методы привязываются к структурам:**
```go
type Method struct {
	name    string           
	handler func(interface{}) 
}

type TypeDescriptor struct {
	name         string   
	valueMethods []Method 
	ptrMethods   []Method
}
```

2. Просто сравните наборы методов структур с наборами методов интерфейсов.