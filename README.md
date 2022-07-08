# Book management

[![CircleCI](https://dl.circleci.com/status-badge/img/gh/tuanlc/book-management/tree/main.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/tuanlc/book-management/tree/main)
[![Go Report Card](https://goreportcard.com/badge/github.com/tuanlc/book-management?style=flat-square)](https://goreportcard.com/report/github.com/tuanlc/book-management)

## Technologies
- Database: PostgreSQL
- ORM: Gorm
- Language: Go
- Hot reload: Air

## Development
### Prequisite
- Install hot reload tool: https://github.com/cosmtrek/air

```sh
make deps
```

### Configure databse DNS
Set the environment varibale `DATABASE_DNS` to configure database DNS. For example:

```
$ DATABASE_DNS=host=localhost user=admin password=admin dbname=book-management port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh
```

### Run
```sh
$ make run
```

## License

[MIT](https://opensource.org/licenses/MIT)

Copyright (c) 2022-present, Tuan LE CONG