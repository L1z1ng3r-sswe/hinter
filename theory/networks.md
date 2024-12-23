# Процесс запроса в браузере и детали сетевого взаимодействия

1. [Что происходит, когда вы вводите что-то в браузере](#что-происходит-когда-вы-вводите-что-то-в-браузере)
2. [Составные части HTTP-запроса](#составные-части-http-запроса)

---

## Что происходит, когда вы вводите что-то в браузере <a id="что-происходит-когда-вы-вводите-что-то-в-браузере"></a>

Например, `https://facebook.com:80/feed/videos?cat_id=1234&breed=scot#title` состоит из:
- **Протокол:** `https`
- **Домен/IP:** `facebook.com`
- **Порт:** `80`
- **Путь:** `/feed/videos`
- **Параметры запроса:** `cat_id=1234&breed=scot`
- **Якорь:** `#title` (ID HTML-элемента)

### Шаги:
1. **Проверка кеша:**
   - **Кеш браузера → Кеш ОС → Кеш настроенного DNS-сервера:**  
     Если не найдено, переходим к следующему шагу.
2. **Запрос к корневому DNS-серверу:**  
   Предоставляет информацию о серверах верхнего уровня (TLD).
3. **Запрос к серверу TLD:**  
   Например, `.com` для `example.com`. Указывает на авторитетный DNS-сервер для данного домена.
4. **Запрос к авторитетному DNS-серверу:**  
   Получает A/AAAA запись домена (IP-адрес).
5. **Кеширование на каждом уровне:**  
   Ускоряет обработку будущих запросов для того же домена.

---

## Составные части HTTP-запроса <a id="составные-части-http-запроса"></a>

HTTP-запрос состоит из:
- **Метод HTTP:** (например, GET, POST, PUT)
- **Запрос:** (например, путь и параметры URL)
- **Версия HTTP:** (например, HTTP/1.1 или HTTP/2)
- **Заголовки:** (например, `Host`, `User-Agent`, `Content-Type: application/json`)
- **Тело запроса:** Если метод POST/PUT/PATCH — может содержать данные формы, JSON или файлы.