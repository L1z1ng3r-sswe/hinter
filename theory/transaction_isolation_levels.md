# Transaction Isolation Levels and Problems

1. [Read Uncommitted](#read-uncommitted)
2. [Read Committed](#read-committed)
3. [Repeatable Read](#repeatable-read)
4. [Serializable](#serializable)
5. [Problems Solved by Transaction Isolation Levels](#problems-solved-by-transaction-isolation-levels)

---

## Transaction Isolation Levels Details

### Read Uncommitted <a id="read-uncommitted"></a>

The lowest level of isolation where transactions can read uncommitted changes made by other transactions.

---

### Read Committed <a id="read-committed"></a>

This level ensures that a transaction can only read data that has been committed by other transactions.

---

### Repeatable Read <a id="repeatable-read"></a>

At this level, a transaction is guaranteed to see the same data when reading the same row multiple times.

---

### Serializable <a id="serializable"></a>

The highest isolation level, where transactions are executed in a way that it appears as if they are executed serially.

---

### Problems Solved by Transaction Isolation Levels <a id="problems-solved-by-transaction-isolation-levels"></a>

- **Dirty Reads**: Occurs when a transaction reads data that has been modified by another transaction but not yet committed.
- **Non-Repeatable Reads**: Happens when a transaction reads the same row multiple times and gets different results because another transaction modified the row in the meantime.
- **Phantom Reads**: Occurs when a transaction reads a set of rows that match a condition, and another transaction inserts or deletes rows that match the condition, leading to different results on subsequent reads.
- **Lost Updates**: Occurs when two transactions simultaneously update the same row, and one of the updates is lost.
