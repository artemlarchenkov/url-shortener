# URL Shortener

Простой сервис для сокращения URL с возможностью создания кастомного alias и управления через HTTP API.

## Функции

* Сокращение длинных ссылок
* Генерация случайного alias, если не указан
* Получение оригинальной ссылки по alias
* Удаление ссылки по alias (через DELETE)
* Аутентификация Basic Auth для API

## Установка

1. Клонируем репозиторий:

```bash
git clone <url_репозитория>
cd url-shortener
```

2. Устанавливаем зависимости:

```bash
go mod tidy
```

3. Настраиваем конфигурацию (`config/local.yaml`):

```yaml
storage_path: "../../storage/storage.db"
http_server:
  address: "localhost:8082"
  timeout: 4s
  idle-timeout: 60s
  user: "myuser"
  password: "mypass"
```

4. Запускаем сервер:

```bash
go run cmd/url-shortener/main.go
```

## API

### POST /url

Создание новой короткой ссылки.

**Запрос:**

```json
{
  "url": "https://example.com",
  "alias": "myalias"  // необязательный
}
```

**Ответ:**

```json
{
  "status": "ok",
  "alias": "myalias"
}
```

### GET /{alias}

Переход на оригинальный URL по alias.

### DELETE /url/{alias}

Удаление ссылки по alias (требует Basic Auth).

## Тестирование

* Через **Postman**:

    * POST: `http://localhost:8082/url`
    * GET: `http://localhost:8082/{alias}`
    * DELETE: `http://localhost:8082/url/{alias}` (с Basic Auth)

## Логи

Логи пишутся через `slog`, включают request\_id и операции для удобной отладки.

## Примечания

* SQLite база (`storage/storage.db`) хранит все ссылки.
* Генерация alias по умолчанию — 6 символов.