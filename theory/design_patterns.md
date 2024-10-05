# Design Patterns

1. [Factory Pattern](#factory-pattern)
2. [Singleton Pattern](#singleton-pattern)
3. [Builder Pattern](#builder-pattern)
4. [Observer Pattern](#observer-pattern)
5. [Strategy Pattern](#strategy-pattern)
6. [Adapter Pattern](#adapter-pattern)
7. [Decorator Pattern](#decorator-pattern)
8. [Facade Pattern](#facade-pattern)
9. [Proxy Pattern](#proxy-pattern)
10. [Command Pattern](#command-pattern)

---

## Design Patterns Details

### Factory Pattern <a id="factory-pattern"></a>

The Factory Pattern предоставляет возможность создавать объекты без указания точного класса создаваемого объекта.

The Factory Pattern provides a way to create objects without specifying the exact class of object that will be created. This pattern is useful when the exact type of the object can be determined at runtime.

```go
package main

import (
	"fmt"
	"math"
)

// Shape interface
type Shape interface {
	Area() float64
}

// Circle struct
type Circle struct {
	Radius float64
}

// Area method for Circle
func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Square struct
type Square struct {
	Side float64
}

// Area method for Square
func (s *Square) Area() float64 {
	return s.Side * s.Side
}

// ShapeFactory function (factory function)
func ShapeFactory(shapeType string, dimension float64) (Shape, error) {
	switch shapeType {
	case "circle":
		return &Circle{Radius: dimension}, nil
	case "square":
		return &Square{Side: dimension}, nil
	default:
		return nil, fmt.Errorf("invalid shape type")
	}
}

func main() {
	// Create a Circle
	shape1, err := ShapeFactory("circle", 5.0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Area of Circle: %.2f\n", shape1.Area())

	// Create a Square
	shape2, err := ShapeFactory("square", 4.0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Area of Square: %.2f\n", shape2.Area())
}
```

---

### Singleton Pattern <a id="singleton-pattern"></a>

The Singleton Pattern — это шаблон творческого проектирования, который гарантирует, что класс имеет только один экземпляр, и предоставляет глобальную точку доступа к этому экземпляру.

The Singleton pattern is a creational design pattern that ensures a class has only one instance and provides a global point of access to that instance.

```go
type Singleton struct {
	Value int
}

var instance *Singleton

var once sync.Once

func GetInstance() *Singleton {
	// The once.Do() ensures the following code runs only once
	once.Do(func() {
		instance = &Singleton{Value: 42}
		fmt.Println("Creating Singleton instance")
	})
	return instance
}
```

---

### Builder Pattern <a id="builder-pattern"></a>

The Builder Pattern отделяет построение сложного объекта от его представления, особенно полезен, когда у вас есть объект со множеством дополнительных параметров или сложными шагами инициализации.

The Builder Pattern separates the construction of a complex object from its representation, especially helpful when you have an object with many optional parameters or complex initialization steps.

```go
type House struct {
	Windows int
	Doors   int
	HasGarage  bool
	HasGarden  bool
	HasSwimmingPool bool
}

type HouseBuilder struct {
	house House
}

func NewHouseBuilder() *HouseBuilder {
	return &HouseBuilder{}
}

func (b *HouseBuilder) AddWindows(num int) *HouseBuilder {
	b.house.Windows = num
	return b
}

func (b *HouseBuilder) AddDoors(num int) *HouseBuilder {
	b.house.Doors = num
	return b
}

func (b *HouseBuilder) AddGarage() *HouseBuilder {
	b.house.HasGarage = true
	return b
}

func (b *HouseBuilder) AddGarden() *HouseBuilder {
	b.house.HasGarden = true
	return b
}

func (b *HouseBuilder) AddSwimmingPool() *HouseBuilder {
	b.house.HasSwimmingPool = true
	return b
}

func (b *HouseBuilder) Build() House {
	return b.house
}
```

---

### Observer Pattern <a id="observer-pattern"></a>

The Observer Pattern определяет связь «один ко многим» между объектами, поэтому, когда один объект меняет состояние, все его зависимые объекты уведомляются.

The Observer Pattern defines a one-to-many relationship between objects so that when one object changes state, all its dependents are notified.

```go
type Observer interface {
	Update(string)
}

type Subject interface {
	Register(Observer)
	Unregister(Observer)
	NotifyAll()
}

type ConcreteSubject struct {
	observers []Observer
	state     string
}

func (s *ConcreteSubject) Register(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *ConcreteSubject) Unregister(o Observer) {
	var index int
	for i, observer := range s.observers {
		if observer == o {
			index = i
			break
		}
	}
	s.observers = append(s.observers[:index], s.observers[index+1:]...)
}

func (s *ConcreteSubject) NotifyAll() {
	for _, observer := range s.observers {
		observer.Update(s.state)
	}
}

func (s *ConcreteSubject) UpdateState(state string) {
	s.state = state
	s.NotifyAll()
}

type ConcreteObserver struct {
	id string
}

func (o *ConcreteObserver) Update(state string) {
	fmt.Printf("Observer %s: Subject changed state to %s\n", o.id, state)
}
```

---

### Strategy Pattern <a id="strategy-pattern"></a>

The Strategy pattern - это шаблон поведенческого проектирования, который позволяет вам определить семейство алгоритмов, инкапсулировать каждый из них и сделать их взаимозаменяемыми. Шаблон стратегии позволяет алгоритму изменяться независимо от клиентов, которые его используют.

The Strategy pattern is a behavioral design pattern that allows you to define a family of algorithms, encapsulate each one, and make them interchangeable. The strategy pattern lets the algorithm vary independently from clients that use it.


```go
type DiscountStrategy interface {
	ApplyDiscount(float64) float64
}

type NoDiscount struct{}

func (nd *NoDiscount) ApplyDiscount(price float64) float64 {
	return price
}

type PercentageDiscount struct {
	Percentage float64
}

func (pd *PercentageDiscount) ApplyDiscount(price float64) float64 {
	return price * (1 - pd.Percentage/100)
}

type FixedDiscount struct {
	DiscountAmount float64
}

func (fd *FixedDiscount) ApplyDiscount(price float64) float64 {
	return price - fd.DiscountAmount
}

type PriceCalculator struct {
	strategy DiscountStrategy
}

func (pc *PriceCalculator) SetStrategy(strategy DiscountStrategy) {
	pc.strategy = strategy
}

func (pc *PriceCalculator) CalculatePrice(price float64) float64 {
	return pc.strategy.ApplyDiscount(price)
}

func main() {
	calculator := &PriceCalculator{}

	calculator.SetStrategy(&NoDiscount{})
	fmt.Printf("Price with no discount: $%.2f\n", calculator.CalculatePrice(100))

	calculator.SetStrategy(&PercentageDiscount{Percentage: 20})
	fmt.Printf("Price with 20%% discount: $%.2f\n", calculator.CalculatePrice(100))

	calculator.SetStrategy(&FixedDiscount{DiscountAmount: 15})
	fmt.Printf("Price with $15 discount: $%.2f\n", calculator.CalculatePrice(100))
}
```

---

### Adapter Pattern <a id="adapter-pattern"></a>

The Adapter Pattern позволяет несовместимым интерфейсам работать вместе.

The Adapter Pattern allows incompatible interfaces to work together.

```go
type JSONLogger struct{}

func (jl *JSONLogger) LogJSON(data map[string]string) {
	jsonData, _ := json.Marshal(data)
	fmt.Printf("Logging JSON: %s\n", string(jsonData))
}

type JSONLoggerAdapter struct {
	JSONLogger *JSONLogger
}

func (adapter *JSONLoggerAdapter) LogMessage(message string) {
	data := map[string]string{
		"message": message,
	}
	adapter.JSONLogger.LogJSON(data)
}

type Logger interface {
	LogMessage(string)
}
```

---

### Decorator Pattern <a id="decorator-pattern"></a>

The Decorator Pattern позволяет динамически добавлять поведение к отдельному объекту, не затрагивая поведение других объектов того же класса. (обеспечивает OCP)

The Decorator Pattern allows behavior to be added to an individual object, dynamically, without affecting the behavior of other objects from the same class. (ensures OCP)

```go
type Notifier interface {
	Send(string)
}

type ConcreteNotifier struct{}

func (n *ConcreteNotifier) Send(message string) {
	fmt.Println("Sending basic notification:", message)
}

type SMSDecorator struct {
	Notifier Notifier
}

func (d *SMSDecorator) Send(message string) {
	d.Notifier.Send(message)
	fmt.Println("Sending SMS notification:", message)
}

type EmailDecorator struct {
	Notifier Notifier
}

func (d *EmailDecorator) Send(message string) {
	d.Notifier.Send(message)
	fmt.Println("Sending Email notification:", message)
}

func main() {
	notifier := &ConcreteNotifier{}

	smsNotifier := &SMSDecorator{
		Notifier: notifier,
	}

	emailSMSNotifier := &EmailDecorator{
		Notifier: smsNotifier,
	}

	emailSMSNotifier.Send("Hello, World!")
```

---

### Facade Pattern <a id="facade-pattern"></a>

The Facade Pattern provides a simplified interface to a complex subsystem.

The Facade Pattern обеспечивает упрощенный интерфейс для сложной подсистемы.

```go
type TV struct{}

func (tv *TV) On() {
	fmt.Println("Turning on the TV...")
}

func (tv *TV) Off() {
	fmt.Println("Turning off the TV...")
}

type SoundSystem struct{}

func (ss *SoundSystem) On() {
	fmt.Println("Turning on the sound system...")
}

func (ss *SoundSystem) SetVolume(volume int) {
	fmt.Printf("Setting sound system volume to %d...\n", volume)
}

func (ss *SoundSystem) Off() {
	fmt.Println("Turning off the sound system...")
}

type StreamingService struct{}

func (ss *StreamingService) PlayMovie(movie string) {
	fmt.Printf("Playing the movie: %s...\n", movie)
}

func (ss *StreamingService) StopMovie() {
	fmt.Println("Stopping the movie...")
}

type HomeTheaterFacade struct {
	tv             *TV
	soundSystem    *SoundSystem
	streamingService *StreamingService
}

func NewHomeTheaterFacade(tv *TV, soundSystem *SoundSystem, streamingService *StreamingService) *HomeTheaterFacade {
	return &HomeTheaterFacade{
		tv:             tv,
		soundSystem:    soundSystem,
		streamingService: streamingService,
	}
}

func (htf *HomeTheaterFacade) StartMovie(movie string) {
	htf.tv.On()
	htf.soundSystem.On()
	htf.soundSystem.SetVolume(10)
	htf.streamingService.PlayMovie(movie)
}

func (htf *HomeTheaterFacade) StopMovie() {
	htf.streamingService.StopMovie()
	htf.soundSystem.Off()
	htf.tv.Off()
}

func main() {
	tv := &TV{}
	soundSystem := &SoundSystem{}
	streamingService := &StreamingService{}

	homeTheater := NewHomeTheaterFacade(tv, soundSystem, streamingService)

	homeTheater.StartMovie("Home Alone")

	homeTheater.StopMovie()
}
```

---

### Proxy Pattern <a id="proxy-pattern"></a>

The Proxy Pattern предоставляет суррогат или заполнитель для другого объекта для управления доступом к нему.

The Proxy Pattern provides a surrogate or placeholder for another object to control access to it.

```go
type File interface {
	Display()
}

type RealFile struct {
	filename string
}

func (f *RealFile) Display() {
	fmt.Println("Displaying file:", f.filename)
}

func NewRealFile(filename string) *RealFile {
	fmt.Println("Loading file:", filename)
	return &RealFile{filename: filename}
}

type ProxyFile struct {
	filename string
	realFile *RealFile
}

func (p *ProxyFile) Display() {
	if p.realFile == nil {
		p.realFile = NewRealFile(p.filename)
	}
	p.realFile.Display()
}

func NewProxyFile(filename string) *ProxyFile {
	return &ProxyFile{filename: filename}
}

func main() {
	file := NewProxyFile("example.txt")

	file.Display() 
	file.Display() 
}
```
---

### Command Pattern <a id="command-pattern"></a>

The Command Pattern инкапсулирует запрос как объект, тем самым позволяя параметризовать клиентов с различными запросами.

The Command Pattern encapsulates a request as an object, thereby letting you parameterize clients with different requests.

```go
type Command interface {
	Execute()
	Undo()
}

type Light struct {
	isOn bool
}

func (l *Light) TurnOn() {
	l.isOn = true
	fmt.Println("Light is turned ON")
}

func (l *Light) TurnOff() {
	l.isOn = false
	fmt.Println("Light is turned OFF")
}

type TurnOnCommand struct {
	light *Light
}

func (c *TurnOnCommand) Execute() {
	c.light.TurnOn()
}

func (c *TurnOnCommand) Undo() {
	c.light.TurnOff()
}

type TurnOffCommand struct {
	light *Light
}

func (c *TurnOffCommand) Execute() {
	c.light.TurnOff()
}

func (c *TurnOffCommand) Undo() {
	c.light.TurnOn()
}

type RemoteControl struct {
	command Command
}

func (r *RemoteControl) SetCommand(command Command) {
	r.command = command
}

func (r *RemoteControl) PressButton() {
	r.command.Execute()
}

func (r *RemoteControl) PressUndo() {
	r.command.Undo()
}

func main() {
	light := &Light{}

	turnOnCommand := &TurnOnCommand{light: light}
	turnOffCommand := &TurnOffCommand{light: light}

	remote := &RemoteControl{}

	remote.SetCommand(turnOnCommand)
	remote.PressButton() 
	remote.PressUndo()

	remote.SetCommand(turnOffCommand)
	remote.PressButton() 
	remote.PressUndo()
}
```