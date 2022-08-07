# Book management

![action status](https://github.com/tuanlc/book-management/actions/workflows/ci.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/tuanlc/book-management?style=flat-square)](https://goreportcard.com/report/github.com/tuanlc/book-management)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/tuanlc/book-management)

## Motivation
I'm a Golang & gRPC newbie (August 2022). A good way to practice is working on a real opensource project, so I decided to initialie the book management project with these technologies. If you are a newbie as me, don't hesitate to join me in contribute to this project. If you are an experienced developer, I'm highly appreciate you if you can take a look to the project structure, codebase, etc and give me any feedback on mistakes, good practices, etc. Feel free to create [issues](https://github.com/tuanlc/book-management/issues) or join the project [discussion](https://github.com/tuanlc/book-management/discussions)

You can use this project as the template for the following technologies.

## Technologies
- Language: [Golang](https://github.com/golang/go)
- API: [gRPC](https://grpc.io/)
- ORM: [Gorm](https://pkg.go.dev/gorm.io/gorm)
- Database: [PostgreSQL](https://www.postgresql.org/)
- Hot reload: [Air](https://github.com/cosmtrek/air)

## Development
### Prequisite
- Install hot reload tool: https://github.com/cosmtrek/air

```sh
$ make deps
```

### Configure databse DNS
Set the environment varibale `DATABASE_DNS` to configure database DNS. For example:

```
$ DATABASE_DNS=host=localhost user=admin password=admin dbname=book-management port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh
```

### Run in dev mode with hot reload feature
```sh
$ make dev
```

## Production

### Run
```sh
$ make run
```

## License

[MIT](https://opensource.org/licenses/MIT)

Copyright (c) 2022-present, Tuan LE CONG
