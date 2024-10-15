1. Строка занимает 16 байт: указатель на первый элемент базового массива и len.

2. Для вычисления длины рун можно использовать `utf8.RuneCountInString(str)`, `len([]rune())` или цикл `range`.

3.  
```go
var builder strings.Builder

// Записываем строки в builder
builder.WriteString("Hello, ")
builder.WriteString("World!")
builder.WriteString(" How are you?")

// Преобразуем содержимое builder в строку
result := builder.String()
```