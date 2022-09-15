# Golang REST Example

### Stack
- Gin
- Gorm
- Mysql
- Swagger
### Manual Installation 
1. Clean Project 
    ```sh
    $ make clean
    ```

2. Run Project
    ```sh
    $ make run
    ```
3. Struct Project
    ```sh
    .
    ├── Makefile
    ├── README.md
    ├── apispec.yml
    ├── controller
    │   └── book.go
    ├── docs
    │   ├── docs.go
    │   ├── swagger.json
    │   └── swagger.yaml
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── model
    │   ├── book.go
    │   └── database.go
    ├── repository
    │   └── book.go
    ├── service
    │   └── book.go
    ├── shared
    │   ├── config
    │   │   └── config.go
    │   ├── deps.go
    │   ├── dto
    │   │   └── book.go
    │   ├── internal_const
    │   │   └── errors.go
    │   ├── mapper
    │   │   └── book.go
    │   └── middlewares
    │       └── middlewares.go
    └── utils
        ├── config
        │   └── config.go
        ├── context.go
        └── errors
            └── errors.go

    14 directories, 23 files
    ```
### Documentation
http://localhost:3000/swagger/index.html

### Contact
https://www.linkedin.com/in/aji-indra-jaya

License
----

MIT