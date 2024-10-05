# Channels Axioms

1. [Unbuffered Channel](#unbuffered-channel)
2. [Buffered Channel](#buffered-channel)
3. [Close Channel](#close-channel)
4. [Send on Closed Channel](#send-on-closed-channel)
5. [Nil Channels](#nil-channels)
6. [Closing a Channel is Idempotent](#closing-a-channel-is-idempotent)
7. [Select and Channel Operations](#select-and-channel-operations)
8. [Range Through Channel](#range-through-channel)

---

## Channels Axioms Details

### Unbuffered Channel <a id="unbuffered-channel"></a>

A send on a channel blocks until a receiver is ready, and a receive blocks until a sender is ready.

---

### Buffered Channel <a id="buffered-channel"></a>

Sends on a buffered channel block only when the buffer is full. Receives block only when the buffer is empty.

---

### Close Channel <a id="close-channel"></a>

Закрытие канала сигнализирует о том, что по нему больше не будут отправляться значения.

---

### Send on Closed Channel <a id="send-on-closed-channel"></a>

Sending on a closed channel causes a runtime panic.

---

### Nil Channels <a id="nil-channels"></a>

Sending, receiving, or closing a nil channel blocks forever.

---

### Closing a Channel is Idempotent <a id="closing-a-channel-is-idempotent"></a>

Closing an already closed channel causes a panic, but it’s safe to call `close` exactly once.

---

### Select and Channel Operations <a id="select-and-channel-operations"></a>

The `select` statement blocks until at least one of its cases can proceed.

---

### Range Through Channel <a id="range-through-channel"></a>

The `range` statement blocks until the channel is closed.