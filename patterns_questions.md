# SOLID Principles

## List of Topics

1. [Single Responsibility Principle (SRP)](#single-responsibility-principle-srp)
2. [Open/Closed Principle (OCP)](#open-closed-principle-ocp)
3. [Liskov Substitution Principle (LSP)](#liskov-substitution-principle-lsp)
4. [Interface Segregation Principle (ISP)](#interface-segregation-principle-isp)
5. [Dependency Inversion Principle (DIP)](#dependency-inversion-principle-dip)



# Common Programming Principles

## List of Topics

1. [DRY (Don't Repeat Yourself)](#dry-dont-repeat-yourself)
2. [KISS (Keep It Simple, Stupid)](#kiss-keep-it-simple-stupid)
3. [YAGNI (You Aren't Gonna Need It)](#yagni-you-arent-gonna-need-it)



# Design Patterns

## List of Topics

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




# Transaction Isolation Levels and Problems

## List of Topics

1. [Read Uncommitted](#read-uncommitted)
2. [Read Committed](#read-committed)
3. [Repeatable Read](#repeatable-read)
4. [Serializable](#serializable)
5. [Problems Solved by Transaction Isolation Levels](#problems-solved-by-transaction-isolation-levels)




# Database Concepts

## List of Topics

1. [Normalization Levels](#normalization-levels)
2. [Replication](#replication)
3. [Sharding](#sharding)
4. [Partitioning](#partitioning)




# Channels Axioms

## List of Topics

1. [Unbuffered Channel](#unbuffered-channel)
2. [Buffered Channel](#buffered-channel)
3. [Close Channel](#close-channel)
4. [Send on Closed Channel](#send-on-closed-channel)
5. [Nil Channels](#nil-channels)
6. [Closing a Channel is Idempotent](#closing-a-channel-is-idempotent)
7. [Select and Channel Operations](#select-and-channel-operations)
8. [Range Through Channel](#range-through-channel)



# OOP Axioms

## List of Topics

1. [Encapsulation](#encapsulation)
2. [Abstraction](#abstraction)
3. [Inheritance (Composition over Inheritance in Go)](#inheritance-composition-over-inheritance-in-go)
4. [Polymorphism](#polymorphism)



---

## Unbuffered Channel <a id="unbuffered-channel"></a>

### Send-Receive Synchronization
A send on a channel blocks until a receiver is ready, and a receive blocks until a sender is ready.

```go
ch := make(chan int)
go func() {
    ch <- 42  // send blocks until there's a receiver
}()
val := <-ch   // receive blocks until there's a sender
```

---

## Buffered Channel <a id="buffered-channel"></a>

Sends on a buffered channel block only when the buffer is full. Receives block only when the buffer is empty.

```go
ch := make(chan int, 2)  // buffered channel with capacity 2
ch <- 1  // doesn't block
ch <- 2  // doesn't block
go func() {
    <-ch  // receive one value to free buffer space
}()
ch <- 3  // doesn't block now because buffer has space
```

---

## Close Channel <a id="close-channel"></a>

Закрытие канала сигнализирует о том, что по нему больше не будут отправляться значения.

---

## Send on Closed Channel <a id="send-on-closed-channel"></a>

Sending on a closed channel causes a runtime panic.

---

## Nil Channels <a id="nil-channels"></a>

Sending, receiving, or closing a nil channel blocks forever.

---

## Closing a Channel is Idempotent <a id="closing-a-channel-is-idempotent"></a>

Closing an already closed channel causes a panic, but it’s safe to call `close` exactly once.

---

## Select and Channel Operations <a id="select-and-channel-operations"></a>

The `select` statement blocks until at least one of its cases can proceed.

---

## Range Through Channel <a id="range-through-channel"></a>

The `select` statement blocks until the channel is closed.

---

## Encapsulation <a id="encapsulation"></a>

Controlled access to data through public methods (Go uses capitalized names for public visibility).

---

## Abstraction <a id="abstraction"></a>

Achieved through interfaces that hide implementation details and expose behavior.

---

## Inheritance (Composition over Inheritance in Go) <a id="inheritance-composition-over-inheritance-in-go"></a>

Go uses composition instead of classical inheritance by embedding structs within other structs.

---

## Polymorphism <a id="polymorphism"></a>

Achieved through interfaces that allow different types to be treated uniformly as the interface type.


---

## Single Responsibility Principle (SRP) <a id="single-responsibility-principle-srp"></a>

A class should have only one reason to change, meaning it should only have one job or responsibility.

---

## Open/Closed Principle (OCP) <a id="open-closed-principle-ocp"></a>

Software entities (classes, modules, functions, etc.) should be open for extension but closed for modification.

---

## Liskov Substitution Principle (LSP) <a id="liskov-substitution-principle-lsp"></a>

Objects of a superclass should be replaceable with objects of a subclass without affecting the correctness of the program.

---

## Interface Segregation Principle (ISP) <a id="interface-segregation-principle-isp"></a>

Clients should not be forced to depend on interfaces they do not use. Instead of one fat interface, many small, specific interfaces are preferred.

---

## Dependency Inversion Principle (DIP) <a id="dependency-inversion-principle-dip"></a>

High-level modules should not depend on low-level modules. Both should depend on abstractions. Abstractions should not depend on details. Details should depend on abstractions.

---
<!-- !________________________________ dry ___________________________________ -->

## DRY (Don't Repeat Yourself) <a id="dry-dont-repeat-yourself"></a>

A principle aimed at reducing the repetition of code by ensuring that functionality is defined only once. If you have similar code, consider refactoring it to a common function or module.

---
<!-- !________________________________ kiss ___________________________________ -->


## KISS (Keep It Simple, Stupid) <a id="kiss-keep-it-simple-stupid"></a>

A design principle stating that simplicity should be a key goal. Avoid unnecessary complexity in the code and aim for straightforward, easy-to-understand solutions.

---

<!-- !________________________________ yagni ___________________________________ -->


## YAGNI (You Aren't Gonna Need It) <a id="yagni-you-arent-gonna-need-it"></a>

A principle of extreme programming that encourages developers not to add functionality until it is absolutely necessary. Avoid overengineering and implementing features that are not immediately needed.

---

<!-- !________________________________ designing ___________________________________ -->


## Factory Pattern <a id="factory-pattern"></a>

The Factory Pattern provides a way to create objects without specifying the exact class of object that will be created. This pattern is useful when the exact type of the object can be determined at runtime.

---

## Singleton Pattern <a id="singleton-pattern"></a>

The Singleton Pattern ensures that a class has only one instance and provides a global point of access to that instance.

---

## Builder Pattern <a id="builder-pattern"></a>

The Builder Pattern separates the construction of a complex object from its representation, allowing the same construction process to create different representations.

---

## Observer Pattern <a id="observer-pattern"></a>

The Observer Pattern defines a one-to-many relationship between objects so that when one object changes state, all its dependents are notified and updated automatically.

---

## Strategy Pattern <a id="strategy-pattern"></a>

The Strategy Pattern allows selecting an algorithm's behavior at runtime. It defines a family of algorithms, encapsulates each one, and makes them interchangeable.

---

## Adapter Pattern <a id="adapter-pattern"></a>

The Adapter Pattern allows incompatible interfaces to work together. It acts as a bridge between two incompatible interfaces.

---

## Decorator Pattern <a id="decorator-pattern"></a>

The Decorator Pattern allows behavior to be added to an individual object, dynamically, without affecting the behavior of other objects from the same class.

---

## Facade Pattern <a id="facade-pattern"></a>

The Facade Pattern provides a simplified interface to a complex subsystem. It hides the complexities of the system and provides an interface to the client using which the client can access the system.

---

## Proxy Pattern <a id="proxy-pattern"></a>

The Proxy Pattern provides a surrogate or placeholder for another object to control access to it.

---

## Command Pattern <a id="command-pattern"></a>

The Command Pattern encapsulates a request as an object, thereby letting you parameterize clients with different requests, queue requests, and log requests.


---

<!-- !________________________________ transactions types ___________________________________ -->


## Read Uncommitted <a id="read-uncommitted"></a>

The lowest level of isolation where transactions can read uncommitted changes made by other transactions, potentially leading to issues like dirty reads.

---

## Read Committed <a id="read-committed"></a>

This level ensures that a transaction can only read data that has been committed by other transactions. It solves the problem of dirty reads but may still allow non-repeatable reads and phantom reads.

---

## Repeatable Read <a id="repeatable-read"></a>

At this level, a transaction is guaranteed to see the same data when reading the same row multiple times during the transaction, solving the issue of non-repeatable reads. However, it doesn't prevent phantom reads.

---

## Serializable <a id="serializable"></a>

The highest isolation level, where transactions are executed in a way that it appears as if they are executed serially. This level solves all problems, including phantom reads, at the cost of performance.

---

<!-- !_____________________ problems solved by iso level ___________________________ -->


## Problems Solved by Transaction Isolation Levels <a id="problems-solved-by-transaction-isolation-levels"></a>

### 1. Dirty Reads
Occurs when a transaction reads data that has been modified by another transaction but not yet committed. Solved by the **Read Committed** level and higher.

### 2. Non-Repeatable Reads
Happens when a transaction reads the same row multiple times and gets different results because another transaction modified the row in the meantime. Solved by the **Repeatable Read** level and higher.

### 3. Phantom Reads
Occurs when a transaction reads a set of rows that match a condition, and a second transaction inserts or deletes rows that would have matched the condition. This is solved by the **Serializable** isolation level.

### 4. Lost Updates
Occurs when two transactions simultaneously update the same row, and one of the updates is lost. Solved by **Repeatable Read** and **Serializable** levels.



<!-- !________________________________ database concepts ___________________________________ -->

---

## Normalization Levels <a id="normalization-levels"></a>

### First Normal Form (1NF)
Ensures that each column contains atomic values, and there are no repeating groups or arrays in a table.

### Second Normal Form (2NF)
Achieved when a database is in 1NF and all non-key attributes are fully dependent on the primary key.

### Third Normal Form (3NF)
Achieved when a database is in 2NF and all attributes are only dependent on the primary key, removing transitive dependency.

### Boyce-Codd Normal Form (BCNF)
A stricter version of 3NF where every determinant is a candidate key.

### Fourth Normal Form (4NF)
Achieved when a database is in BCNF and has no multi-valued dependencies.

### Fifth Normal Form (5NF)
Achieved when a database is in 4NF and there are no join dependencies.

---

## Replication <a id="replication"></a>

Replication is the process of copying and maintaining database objects, like tables, across multiple database servers to ensure redundancy and fault tolerance. Types of replication include:

1. **Master-Slave Replication**: One master database handles writes and updates, and slave databases replicate the data from the master.
2. **Master-Master Replication**: Both servers can handle read and write operations, replicating data across all masters.

---

## Sharding <a id="sharding"></a>

Sharding is a method of horizontal partitioning in databases where data is distributed across multiple servers. Each shard contains a subset of the total data, and queries are routed to the correct shard based on a shard key. This technique improves scalability by spreading the data load across multiple databases.

---

## Partitioning <a id="partitioning"></a>

Partitioning involves splitting a database into distinct sections to improve performance and manageability. Types of partitioning include:

1. **Horizontal Partitioning**: Dividing a table into rows, so each partition contains a subset of rows (similar to sharding).
2. **Vertical Partitioning**: Dividing a table into columns, grouping frequently accessed columns together.
3. **Range Partitioning**: Distributing data based on a range of values (e.g., dates or numeric ranges).
4. **List Partitioning**: Partitioning based on a list of discrete values (e.g., a list of countries or regions).