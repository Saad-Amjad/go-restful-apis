
## Go Restful Apis

### Installation

    1. go build main.go

    2. go run main.go

This repo uses [Echo](https://echo.labstack.com/), Go minimalist framework.

Opening sqlLite db shows a new database table called books.

And hitting localhost:8080/books returns the list of the books from the SQLite DB.

Took inspiration from [here](https://www.codementor.io/@packt/how-to-set-up-a-project-in-echo-mpo39w5zu).


### Features:

Server
Logger

SQLite Database connection

GET API: /books


### Todos

- Rest of the crud
- Folder structure
- Authentication
- Open API
- Dockerize

