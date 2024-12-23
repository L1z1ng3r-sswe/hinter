# Аксиомы гарантии доставки TCP

1. [Трехэтапное установление соединения (Three-Way Handshake)](#three-way-handshake)  
4. [Проверка целостности данных (Error Checking)](#error-checking)  
5. [Управление потоком (Flow Control)](#flow-control)  
6. [Управление перегрузкой сети (Congestion Control)](#congestion-control)  

---

## Детали гарантии доставки TCP

### Трехэтапное установление соединения (Three-Way Handshake) <a id="three-way-handshake"></a>

Критически важная основа для надежной передачи данных.  
Перед началом передачи данных TCP устанавливает соединение между отправителем и получателем с помощью следующего процесса:

1. **SYN (Synchronize)**: Отправитель отправляет запрос на установление соединения.  
2. **SYN-ACK (Synchronize-Acknowledge)**: Получатель подтверждает запрос и сигнализирует о готовности.  
3. **ACK (Acknowledge)**: Отправитель подтверждает получение ответа, завершая установление соединения.  

Этот процесс гарантирует, что обе стороны согласны с начальными параметрами, такими как порядковые номера, и готовы к надежной передаче данных.

---

### Порядковые номера (Sequence Numbers) <a id="sequence-numbers"></a>

TCP назначает каждому пакету порядковый номер, что позволяет:

1. **Контроль порядка**: Данные собираются в правильной последовательности на стороне получателя.  
2. **Обнаружение потерь**: Пропущенные пакеты легко обнаружить по отсутствующим порядковым номерам.  

---

### Проверка целостности данных (Error Checking) <a id="error-checking"></a>

Каждый пакет в TCP содержит check-сумму (hash), которая используется для проверки целостности данных:

- Если данные были повреждены во время передачи, пакет отбрасывается.  
- TCP повторно отправляет пакет, пока он не будет успешно принят.  

---

### Управление потоком (Flow Control) <a id="flow-control"></a>

TCP использует механизмы управления потоком, чтобы предотвратить перегрузку получателя:

- Получатель сообщает отправителю о доступном размере буфера.  
- Отправитель регулирует скорость передачи данных, чтобы избежать переполнения буфера у получателя.  

---

### Управление перегрузкой сети (Congestion Control) <a id="congestion-control"></a>

TCP адаптирует скорость передачи данных в зависимости от состояния сети:

1. **Избежание перегрузки**: Отправитель снижает скорость передачи данных при обнаружении перегрузки.  
2. **Динамическая настройка**: Скорость передачи увеличивается постепенно, если сеть работает стабильно.  

Этот механизм помогает минимизировать потерю пакетов и обеспечивает стабильность передачи данных.
Представьте водопровод: TCP регулирует поток воды через краны (размер окна) в зависимости от того, как быстро вода проходит по трубам (сеть) и возвращается сигнал (ACK) о доставке. Если трубы перегружены, кран прикрывается; если трубы свободны, кран открывается шире.

Это позволяет TCP динамически адаптироваться к состоянию сети, сохраняя баланс между эффективностью и стабильностью передачи данных.