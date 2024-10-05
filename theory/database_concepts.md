# Database Concepts

1. [Normalization Levels](#normalization-levels)
2. [Replication](#replication)
3. [Sharding](#sharding)
4. [Partitioning](#partitioning)

---

## Database Concepts Details

### Normalization Levels <a id="normalization-levels"></a>

- **1NF**: Atomic values, no repeating groups or arrays.
- **2NF**: No partial dependencies on a primary key.
- **3NF**: No transitive dependency on the primary key.
- **BCNF**: Every determinant is a candidate key.
- **4NF**: No multi-valued dependencies.
- **5NF**: No join dependencies.

---

### Replication <a id="replication"></a>

Replication is the process of copying and maintaining database objects across multiple servers to ensure redundancy.

---

### Sharding <a id="sharding"></a>

Sharding is a method of horizontal partitioning in databases where data is distributed across multiple servers.

---

### Partitioning <a id="partitioning"></a>

Partitioning involves splitting a database into distinct sections to improve performance and manageability.