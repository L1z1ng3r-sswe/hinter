# [What Happens When You Type Something in the Browser](#what-happens-when-you-type-something-in-the-browser)

2. An HTTP request consists of what?

3. how tcp handshake looks like?
---

## What Happens When You Type Something in the Browser <a id="what-happens-when-you-type-something-in-the-browser"></a>

f.g. https://facebook.com:80/feed/videos?=cat_id=1234&breed=scot#title - url is consists of protocol, domain / ip, path, params, anchor of the id of the html 
1. **Проверка кеша браузера → Проверка кеша ОС → Проверка кеша настроенного DNS-сервера**:  
   - DNS-запрос сначала проверяется в кеше браузера, затем в кеше операционной системы и, наконец, в кеше настроенного DNS-сервера (это может быть DNS-сервер вашего интернет-провайдера (ISP-internet service provider) или публичный DNS-сервер, такой как Google DNS (8.8.8.8) или Cloudflare (1.1.1.1)).

2. **Запрос к корневому DNS-серверу, если запись не найдена в кеше**:  
   - Если запись не найдена ни в одном из кешей, запрос отправляется к корневому DNS-серверу. Корневые серверы распределены по всему миру и предоставляют информацию о серверах, которые обслуживают домены верхнего уровня (TLD — top-level domain).

3. **Запрос к серверу TLD (например, .com для example.com)**:  
   - Корневой сервер направляет запрос к DNS-серверу, обслуживающему соответствующий домен верхнего уровня (например, .com для example.com). У каждого TLD обычно имеется несколько серверов для распределения нагрузки и обеспечения отказоустойчивости. Серверы TLD указывают на DNS-серверы, которые управляют конкретным доменом.

4. **Авторитетный DNS-сервер предоставляет IP-адрес для домена**:  
   - Авторитетный DNS-сервер содержит записи DNS для данного домена, которыми управляет владелец домена или хостинг-провайдер. Когда запрос достигает авторитетного сервера, он проверяет свои записи, чтобы найти A-запись (для IPv4) или AAAA-запись (для IPv6), которая указывает на IP-адрес сервера, к которому должен подключиться клиент.

5. **Кэширование на каждом уровне для оптимизации будущих запросов**:  
   - На каждом уровне (авторитетный сервер, TLD-сервер, корневой сервер, DNS-сервер, ОС, браузер) ответы могут кэшироваться, чтобы ускорить обработку последующих запросов к тому же домену.

---

An HTTP request consists of what?
HTTP method, Request, http version , headers (host (domain), user-agent (browser/app), content-type: application/json), body if put/post/patch -> form, json, file data.


How tcp handshake looks like?
three-way handshake (трехстороннее рукопожатие)

SYN (Synchronize) — Client to Server:

The client starts by sending a TCP packet with the SYN (synchronize) flag set. This packet indicates that the client wants to establish a connection and synchronize sequence numbers with the server.
The SYN packet contains an initial sequence number (ISN) chosen by the client.
SYN-ACK (Synchronize-Acknowledge) — Server to Client:

Upon receiving the SYN packet, the server responds with a SYN-ACK packet.
The SYN-ACK packet serves two purposes:
It acknowledges the client's SYN by sending an acknowledgment number, which is the client's ISN plus 1.
It also includes the server's own SYN, with its initial sequence number, to establish a connection in the opposite direction.
ACK (Acknowledge) — Client to Server:

Finally, the client sends an ACK packet back to the server.
This packet acknowledges the server's SYN by incrementing the server's ISN by 1, completing the handshake.

Client sends: SYN (Client ISN) 
Server replies: SYN-ACK (Server ISN, ACK = Client ISN + 1)
Client responds: ACK (ACK = Server ISN + 1)

