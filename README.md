# Student REST API with Gorilla Mux

A CRUD REST API in Go using **gorilla/mux** for routing and **GORM** on **PostgreSQL**.

## Endpoints

| Method | Path | Action |
|---|---|---|
| `POST` | `/students` | Create a student |
| `GET` | `/students` | List all students |
| `GET` | `/students/{id}` | Get one student |
| `PUT` | `/students/{id}` | Update a student |
| `DELETE` | `/students/{id}` | Delete a student |

## Files

- `main.go`: routes and server
- `handler.go`: request handlers
- `model.go`: the `Student` model
- `database.go`: database connection

## Run

```sh
export DATABASE_URL="host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
go run .
```

The server starts on `:8080`. If `DATABASE_URL` is not set, it uses a local PostgreSQL with the
default settings.
