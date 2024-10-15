# Transaction Isolation Levels and Problems

1. [Read Uncommitted](#read-uncommitted)
2. [Read Committed](#read-committed)
3. [Repeatable Read](#repeatable-read)
4. [Serializable](#serializable)
5. [Problems Solved by Transaction Isolation Levels](#problems-solved-by-transaction-isolation-levels)

---

## Transaction Isolation Levels Details

### Read Uncommitted <a id="read-uncommitted"></a>

Самый низкий уровень изоляции, при котором транзакции могут читать незафиксированные изменения, внесенные другими транзакциями.

---

### Read Committed <a id="read-committed"></a>

Этот уровень гарантирует, что транзакция может читать только данные, зафиксированные другими транзакциями.

---

### Repeatable Read <a id="repeatable-read"></a>

На этом уровне транзакция гарантированно увидит одни и те же данные при многократном чтении одной и той же строки.

---

### Serializable <a id="serializable"></a>

Самый высокий уровень изоляции, при котором транзакции выполняются так, как будто они выполняются последовательно.

---

### Problems Solved by Transaction Isolation Levels <a id="problems-solved-by-transaction-isolation-levels"></a>

- **Dirty Reads**: Происходит, когда транзакция считывает данные, которые были изменены другой транзакцией, но еще не зафиксированы.
- **Non-Repeatable Reads**: : Происходят, когда транзакция считывает одну и ту же строку несколько раз и получает разные результаты, потому что другая транзакция изменила строку в это время.
- **Phantom Reads**: Происходит, когда транзакция считывает набор строк, соответствующих условию, а другая транзакция вставляет или удаляет строки, соответствующие условию, что приводит к другим результатам при последующих чтениях.
- **Lost Updates**: Происходит, когда две транзакции одновременно обновляют одну и ту же строку, и одно из обновлений теряется.