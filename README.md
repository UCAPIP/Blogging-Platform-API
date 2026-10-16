# Blogging Platform API

🇬🇧 [English](#english) | 🇷🇺 [Русский](#русский)

---

## English

A simple RESTful API for a personal blogging platform, written in Go. It supports basic CRUD operations on blog posts and was built as a learning project to practice REST conventions, HTTP status codes, error handling and working with a database.

### Features

- Create, read, update and delete blog posts
- Request body validation with clear error messages
- Consistent JSON error responses
- Request body size limit (1 MB)
- SQLite storage, parameterized queries (protection against SQL injection)

### Tech stack

- Go 1.22+ (uses method and wildcard patterns in the standard `net/http` router)
- SQLite via [`modernc.org/sqlite`](https://pkg.go.dev/modernc.org/sqlite) (pure Go, no CGO required)

### Project structure

| File | Purpose |
|---|---|
| `main.go` | Opens the database, registers routes, starts the server |
| `handlers.go` | HTTP handlers, input validation |
| `repository.go` | All database access (queries, `Post` model) |
| `response.go` | JSON response and error helpers |

### Getting started

**1. Clone the repository and install dependencies**

```bash
git clone https://github.com/UCAPIP/Blogging-Platform-API
cd <your-repo-folder>
go mod tidy
```

**2. Create the `posts` table**

The server uses a SQLite file named `store.db` in the project folder. Create the table with the following schema (for example with the `sqlite3` CLI or DB Browser for SQLite):

```sql
CREATE TABLE IF NOT EXISTS posts (
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    title      TEXT NOT NULL,
    content    TEXT NOT NULL,
    category   TEXT NOT NULL,
    tags       TEXT NOT NULL DEFAULT '[]',
    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL
);
```

Tags are stored as a JSON array in a text column.

**3. Run the server**

```bash
go run .
```

The API is available at `http://localhost:8080`.

### API reference

#### Post object

```json
{
  "id": 1,
  "title": "My First Blog Post",
  "content": "This is the content of my first blog post.",
  "category": "Technology",
  "tags": ["Tech", "Programming"],
  "createdAt": "2021-09-01T12:00:00Z",
  "updatedAt": "2021-09-01T12:00:00Z"
}
```

`id`, `createdAt` and `updatedAt` are set by the server. The client sends only `title`, `content`, `category` and `tags`.

#### Endpoints

| Method | Path | Description | Success | Errors |
|---|---|---|---|---|
| `POST` | `/posts` | Create a post | `201 Created` + post | `400` |
| `GET` | `/posts` | Get all posts | `200 OK` + array | `500` |
| `GET` | `/posts/{id}` | Get one post | `200 OK` + post | `400`, `404` |
| `PUT` | `/posts/{id}` | Replace a post (all fields required) | `200 OK` + updated post | `400`, `404` |
| `DELETE` | `/posts/{id}` | Delete a post | `204 No Content` | `400`, `404` |

#### Validation rules

- `title`, `content` and `category` must not be empty (whitespace-only values are rejected)
- `tags` is an array of strings (may be omitted, it is stored as an empty array)
- Malformed JSON or a body larger than 1 MB results in `400 Bad Request`

#### Error format

All errors are returned as JSON:

```json
{
  "error": "validation failed",
  "details": "title is required; content is required"
}
```

`details` is optional.

#### Status codes

| Code | Meaning |
|---|---|
| `200` | OK |
| `201` | Post created |
| `204` | Post deleted, no body |
| `400` | Invalid id, invalid JSON or validation error |
| `404` | Post not found |
| `500` | Unexpected server error |

### Usage examples

**bash / macOS / Linux / Git Bash**

```bash
# Create
curl -i -X POST http://localhost:8080/posts \
  -H "Content-Type: application/json" \
  -d '{"title":"My First Blog Post","content":"Hello!","category":"Technology","tags":["Tech","Programming"]}'

# Get all
curl -i http://localhost:8080/posts

# Get one
curl -i http://localhost:8080/posts/1

# Update
curl -i -X PUT http://localhost:8080/posts/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Updated title","content":"New text","category":"Technology","tags":["Go"]}'

# Delete
curl -i -X DELETE http://localhost:8080/posts/1
```

**Windows PowerShell**

In PowerShell `curl` is an alias for `Invoke-WebRequest`, so use `curl.exe`. Quotes inside JSON are easiest to handle by putting the body into a file:

```powershell
# body.json contains: {"title":"My First Blog Post","content":"Hello!","category":"Technology","tags":["Tech","Programming"]}
curl.exe -i -X POST http://localhost:8080/posts -H "Content-Type: application/json" -d "@body.json"

curl.exe -i http://localhost:8080/posts
curl.exe -i http://localhost:8080/posts/1
curl.exe -i -X DELETE http://localhost:8080/posts/1
```

Or use the native cmdlet:

```powershell
Invoke-RestMethod http://localhost:8080/posts
```

### Roadmap

- [x] Search by term: `GET /posts?term=tech` (wildcard search in title, content and category)
- [ ] Create the table automatically on startup
- [ ] Automated tests (`net/http/httptest`)
- [ ] Configuration through environment variables (port, database path)
- [ ] Graceful shutdown

---

## Русский

Простой REST API для личной блог-платформы на Go. Поддерживает базовые CRUD-операции над постами. Проект учебный: на нём отрабатываются принципы REST, коды статусов HTTP, обработка ошибок и работа с базой данных.

### Возможности

- Создание, чтение, обновление и удаление постов
- Валидация тела запроса с понятными сообщениями об ошибках
- Единый JSON-формат ошибок
- Ограничение размера тела запроса (1 МБ)
- Хранение в SQLite, параметризованные запросы (защита от SQL-инъекций)

### Стек

- Go 1.22+ (используются шаблоны с методом и параметрами пути в стандартном роутере `net/http`)
- SQLite через [`modernc.org/sqlite`](https://pkg.go.dev/modernc.org/sqlite) (чистый Go, CGO не нужен)

### Структура проекта

| Файл | Назначение |
|---|---|
| `main.go` | Открывает базу, регистрирует маршруты, запускает сервер |
| `handlers.go` | HTTP-обработчики, валидация входных данных |
| `repository.go` | Весь доступ к базе (запросы, модель `Post`) |
| `response.go` | Помощники для JSON-ответов и ошибок |

### Быстрый старт

**1. Склонируй репозиторий и подтяни зависимости**

```bash
git clone <ссылка-на-репозиторий>
cd <папка-репозитория>
go mod tidy
```

**2. Создай таблицу `posts`**

Сервер использует файл SQLite с именем `store.db` в папке проекта. Создай таблицу по такой схеме (например, через консоль `sqlite3` или DB Browser for SQLite):

```sql
CREATE TABLE IF NOT EXISTS posts (
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    title      TEXT NOT NULL,
    content    TEXT NOT NULL,
    category   TEXT NOT NULL,
    tags       TEXT NOT NULL DEFAULT '[]',
    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL
);
```

Теги хранятся как JSON-массив в текстовой колонке.

**3. Запусти сервер**

```bash
go run .
```

API будет доступен по адресу `http://localhost:8080`.

### Справочник по API

#### Объект поста

```json
{
  "id": 1,
  "title": "My First Blog Post",
  "content": "This is the content of my first blog post.",
  "category": "Technology",
  "tags": ["Tech", "Programming"],
  "createdAt": "2021-09-01T12:00:00Z",
  "updatedAt": "2021-09-01T12:00:00Z"
}
```

Поля `id`, `createdAt` и `updatedAt` выставляет сервер. Клиент отправляет только `title`, `content`, `category` и `tags`.

#### Эндпоинты

| Метод | Путь | Описание | Успех | Ошибки |
|---|---|---|---|---|
| `POST` | `/posts` | Создать пост | `201 Created` + пост | `400` |
| `GET` | `/posts` | Получить все посты | `200 OK` + массив | `500` |
| `GET` | `/posts/{id}` | Получить один пост | `200 OK` + пост | `400`, `404` |
| `PUT` | `/posts/{id}` | Заменить пост целиком (все поля обязательны) | `200 OK` + обновлённый пост | `400`, `404` |
| `DELETE` | `/posts/{id}` | Удалить пост | `204 No Content` | `400`, `404` |

#### Правила валидации

- `title`, `content` и `category` не должны быть пустыми (строки из одних пробелов отклоняются)
- `tags` это массив строк (можно не передавать, тогда сохранится пустой массив)
- Некорректный JSON или тело больше 1 МБ даёт `400 Bad Request`

#### Формат ошибок

Все ошибки возвращаются в JSON:

```json
{
  "error": "validation failed",
  "details": "title is required; content is required"
}
```

Поле `details` необязательное.

#### Коды статусов

| Код | Значение |
|---|---|
| `200` | OK |
| `201` | Пост создан |
| `204` | Пост удалён, тела нет |
| `400` | Некорректный id, некорректный JSON или ошибка валидации |
| `404` | Пост не найден |
| `500` | Непредвиденная ошибка сервера |

### Примеры запросов

**bash / macOS / Linux / Git Bash**

```bash
# Создать
curl -i -X POST http://localhost:8080/posts \
  -H "Content-Type: application/json" \
  -d '{"title":"My First Blog Post","content":"Hello!","category":"Technology","tags":["Tech","Programming"]}'

# Получить все
curl -i http://localhost:8080/posts

# Получить один
curl -i http://localhost:8080/posts/1

# Обновить
curl -i -X PUT http://localhost:8080/posts/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Updated title","content":"New text","category":"Technology","tags":["Go"]}'

# Удалить
curl -i -X DELETE http://localhost:8080/posts/1
```

**Windows PowerShell**

В PowerShell `curl` это алиас на `Invoke-WebRequest`, поэтому пиши `curl.exe`. Кавычки внутри JSON проще всего не мучить, а положить тело запроса в файл:

```powershell
# в body.json лежит: {"title":"My First Blog Post","content":"Hello!","category":"Technology","tags":["Tech","Programming"]}
curl.exe -i -X POST http://localhost:8080/posts -H "Content-Type: application/json" -d "@body.json"

curl.exe -i http://localhost:8080/posts
curl.exe -i http://localhost:8080/posts/1
curl.exe -i -X DELETE http://localhost:8080/posts/1
```

Или родной командлет:

```powershell
Invoke-RestMethod http://localhost:8080/posts
```

### Планы

- [x] Поиск по слову: `GET /posts?term=tech` (поиск по подстроке в title, content и category)
- [ ] Автоматическое создание таблицы при старте
- [ ] Автотесты (`net/http/httptest`)
- [ ] Настройка через переменные окружения (порт, путь к базе)
- [ ] Аккуратное завершение сервера (graceful shutdown)
