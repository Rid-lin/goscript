# goscript

[goscript](https://github.com/rid-lin/goscript)

Python/Go script
    1. Must be added to CRON (runs every N minute)
    2. The script is located on WEB server /local/scripts and receives data from DB with SELECT query
    3. The script generates a static HTML web page from the given data
    4. Received data must be added as tags (`<p>`, `<div>`, etc.) to the `<body>` block
    5. Generated HTML page should be served by WEB server

Сценарий Python/Go
    1. Должен быть добавлен в CRON (запускается каждые N минут)
    2. Сценарий расположен на WEB-сервере /local/scripts и получает данные из БД с помощью SELECT-запроса
    3. Сценарий генерирует статическую веб-страницу HTML из заданных данных
    4. Полученные данные должны быть добавлены в виде тегов (`<p>`, `<div>` и т.д.) в блок `<body>`.
    5. Созданная HTML-страница должна обслуживаться WEB-сервером

## Usage

```bash

```
