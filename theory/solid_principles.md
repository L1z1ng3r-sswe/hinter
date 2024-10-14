# SOLID Principles

1. [Single Responsibility Principle (SRP)](#single-responsibility-principle-srp)
2. [Open/Closed Principle (OCP)](#open-closed-principle-ocp)
3. [Liskov Substitution Principle (LSP)](#liskov-substitution-principle-lsp)
4. [Interface Segregation Principle (ISP)](#interface-segregation-principle-isp)
5. [Dependency Inversion Principle (DIP)](#dependency-inversion-principle-dip)

---

## SOLID Principles Details

### Single Responsibility Principle (SRP) <a id="single-responsibility-principle-srp"></a>
  
Принцип единой ответственности (SRP) означает, что у класса должна быть только одна обязанность. Например, в случае с автомобилем, его основная ответственность заключается в выполнении всех действий, связанных с эксплуатацией автомобиля: вождение, включение сигналов и фар. Эти функции взаимосвязаны, так как относятся к работе автомобиля в целом.

Удаление таких компонентов, как фары или двигатель, не нарушает SRP, поскольку они относятся к одной общей обязанности — эксплуатации автомобиля. Пока класс "Автомобиль" отвечает за работу автомобиля, он соблюдает SRP.


The Single Responsibility Principle (SRP) states that a class should have only one responsibility. For example, in the case of a car, its main responsibility is performing actions related to driving: moving, signaling, and using headlights. These functions are related because they are all part of operating the car.

Removing components like headlights or the engine doesn’t break SRP because they all belong to the same responsibility—operating the car. As long as the "Car" class handles the car's overall behavior, it adheres to SRP.

---

### Open/Closed Principle (OCP) <a id="open-closed-principle-ocp"></a>
  
Программные объекты (классы, модули, функции и т. д.) должны быть открыты для расширения, но закрыты для модификации.

Например, у меня есть машина с фарами. Я хочу добавить в одну из фар новую функцию. Если я внесу изменения напрямую, фары могут перестать работать, и даже откат изменений может не решить проблему. Однако, если внести изменения снаружи, например, через внешний модуль или оболочку, фары будут работать, и внешняя часть легко удаляется без последствий.


Software objects (classes, modules, functions, etc.) should be open for extension but closed for modification.

For example, I have a car with headlights. I want to add a new feature to one of the headlights. If I change something directly, the headlights may stop working, and rolling back the changes may not fix the problem. However, if I make changes externally, like through a wrapper or module, the headlights will work, and the external part can be removed without consequences.

---

### Liskov Substitution Principle (LSP) <a id="liskov-substitution-principle-lsp"></a>
  
Принцип подстановки Барбары Лисков (Liskov Substitution Principle, LSP) заключается в том, что объекты подклассов должны заменять объекты базового класса без нарушения работы программы.

Пример: стандарт крепления VESA для мониторов. Допустим, у нас есть базовый класс `Monitor`, и несколько подклассов для разных типов мониторов. Контейнер VESA может крепить любой монитор, реализующий интерфейс крепления, и независимо от того, заменим мы стандартный монитор на другой, система будет работать корректно, что и соблюдает принцип LSP.


The Liskov Substitution Principle (LSP) states that objects of a subclass should be replaceable with objects of the base class without altering the correctness of the program.

Example: the VESA mount standard for monitors. Suppose we have a base class `Monitor` and several subclasses for different types of monitors. A VESA mount (container) can mount any monitor that implements the interface. Regardless of whether we swap the standard monitor with another one, the system will still work correctly, adhering to LSP.

---

### Interface Segregation Principle (ISP) <a id="interface-segregation-principle-isp"></a>
  
Клиентов не следует заставлять зависеть от интерфейсов, которые они не используют. Вместо одного большого интерфейса, лучше использовать множество маленьких, специфических интерфейсов.

Например, при проектировании интерфейсов для разных мониторов лучше создать отдельные интерфейсы для мониторов с поддержкой VESA и без неё, чтобы клиенты использовали только необходимые методы.


Clients should not be forced to depend on interfaces they do not use. Instead of one large interface, many small, specific interfaces are preferred.

For example, when designing interfaces for different monitors, it's better to create separate interfaces for monitors with VESA support and those without, so clients only use the methods they need.

---

### Dependency Inversion Principle (DIP) <a id="dependency-inversion-principle-dip"></a>
  
Модули высокого уровня не должны зависеть от модулей низкого уровня. Оба должны зависеть от абстракций. Абстракции не должны зависеть от деталей, наоборот, детали должны зависеть от абстракций.

Пример: если вы используете MySQL как базу данных, чтобы избежать зависимости от поставщика (vendor lock-in), лучше построить систему таким образом, чтобы можно было заменить MySQL на другую базу данных (например, PostgreSQL), не изменяя основной код. Это можно сделать, используя интерфейсы и абстракции.


High-level modules should not depend on low-level modules. Both should depend on abstractions. Abstractions should not depend on details; rather, details should depend on abstractions.

For example, if you're using MySQL as your database, to avoid vendor lock-in, you should design your system so that you can replace MySQL with another database (e.g., PostgreSQL) without changing the core code. This can be done by using interfaces and abstractions.