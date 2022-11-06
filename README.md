# Book management

![action status](https://github.com/tuanlc/book-management/actions/workflows/ci.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/tuanlc/book-management?style=flat-square)](https://goreportcard.com/report/github.com/tuanlc/book-management)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/tuanlc/book-management)

## Table Of Contents
- [Motivation](#motivation)
- [Technologies](#technologies)
- [Demonstration](#demonstration)
- [Development](#development)
- [Production](#production)
- [License](#license)

## Motivation
I'm a Golang & gRPC newbie (August 2022). A good way to practice is working on a real opensource project, so I decided to create & work on the book management project with these technologies. If you are a newbie as me, don't hesitate to join me in contribute to this project. If you are an experienced developer, I'm highly appreciate you if you can take a look to the project structure, codebase, etc and give me any feedback on mistakes, good practices, etc. Feel free to create [issues](https://github.com/tuanlc/book-management/issues) or join the project [discussion](https://github.com/tuanlc/book-management/discussions)

You can use this project as the template for the following technologies.

## Technologies
- Language: [Golang](https://github.com/golang/go)
- API: [gRPC](https://grpc.io/)
- ORM: [Gorm](https://pkg.go.dev/gorm.io/gorm)
- Database: [PostgreSQL](https://www.postgresql.org/)
- Hot reload: [Air](https://github.com/cosmtrek/air)

## Demonstration
You can experiment the project quickly on your local by using [docker compose](https://docs.docker.com/compose/).

### Setup & up services

```bash
$ docker-compose pull
$ docker-compose up -d
```

You can verify all service up without any error by:
```bash
$ docker-compose ps

               Name                             Command               State                    Ports                  
----------------------------------------------------------------------------------------------------------------------
book-management-server               ./main.out                       Up      0.0.0.0:8080->8080/tcp,:::8080->8080/tcp
book-management_database-manager_1   entrypoint.sh docker-php-e ...   Up      0.0.0.0:9000->8080/tcp,:::9000->8080/tcp
database                             docker-entrypoint.sh postgres    Up      0.0.0.0:5432->5432/tcp,:::5432->5432/tcp
```

Services:
- book-management-server: gRPC server is listening on port 8080
- database: Postgresql database service
- book-management_database-manager_1: Database management UI service

### Test
There is no gRPC client for the project yet. However, you can use Postman to test APIs. Thanks to [Postman](https://blog.postman.com/postman-now-supports-grpc/).

https://user-images.githubusercontent.com/7950991/183283953-e9c0d295-2763-477a-ad00-97126879e273.mp4

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

### Unit test
```sh
$ make test
```

## Production

### Run
```sh
$ make run
```

## License

[MIT](https://opensource.org/licenses/MIT)

Copyright (c) 2022-present, Tuan LE CONG
