# go-book-management-api

Simple REST API for managing books using Go, GORM, and MySQL.

## Requirements

- Go 1.25+
- MySQL
- Docker Compose (optional)

## Run

1. Start MySQL (optional):

```powershell
docker compose up -d
```

2. Run the API:

```powershell
go run ./cmd/main
```

## Database config

The database connection is set in `pkg/config/app.go`:

```go
dsn := "bookuser:bookpassword@tcp(127.0.0.1:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
```

## API Endpoints

Base URL: `http://localhost:9010`

- `POST /book/` : create a book
- `GET /book/` : list books
- `GET /book/{id}` : get a book
- `PUT /book/{id}` : update a book
- `DELETE /book/{id}` : delete a book

## Example request

`POST /book/`

```json
{
  "name": "The Go Programming Language",
  "author": "Alan A. A. Donovan",
  "publication": "Addison-Wesley"
}
```

## Project structure

- `cmd/main/main.go` - application entry point
- `pkg/config/app.go` - database connection
- `pkg/controllers/book-contoller.go` - handlers
- `pkg/models/book.go` - book model
- `pkg/routes/bookstore-routes.go` - routes
- `pkg/utils/utils.go` - request body parser
- `docker-compose.yml` - MySQL service

## Notes

- GORM auto-creates the `books` table.
- Error handling is minimal; add validation for production.
