# gdbc-postgresql
GDBC Postgresql Driver - It is based on [https://github.com/jackc/pgx](https://github.com/jackc/pgx)

[![Go Report Card](https://goreportcard.com/badge/github.com/go-gdbc/gdbc-postgresql)](https://goreportcard.com/report/github.com/go-gdbc/gdbc-postgresql)
[![codecov](https://codecov.io/gh/go-gdbc/gdbc-postgresql/branch/main/graph/badge.svg?token=AsVeTnBKU1)](https://codecov.io/gh/go-gdbc/gdbc-postgresql)
[![Build Status](https://travis-ci.com/go-gdbc/gdbc-postgresql.svg?branch=main)](https://travis-ci.com/go-gdbc/gdbc-postgresql)

# Usage
```go
dataSource, err := gdbc.GetDataSource("gdbc:postgresql://username:password@localhost:3000/testdb?sslmode=disable")
if err != nil {
    panic(err)
}

var connection *sql.DB
connection, err = dataSource.GetConnection()
if err != nil {
    panic(err)
}
```

Postgresql GDBC URL takes one of the following forms:

```
gdbc:postgresql://host:port/database-name?arg1=value1
gdbc:postgresql://host/database-name?arg1=value1
gdbc:postgresql:database-name?arg1=value1
gdbc:postgresql:?arg1=value1
gdbc:postgresql://username:password@host:port/database-name?arg1=value1
```

Default Values:
* **Host** : localhost
* **Port** : 5432
* **User** : postgres
* **Password** : 

Checkout [https://github.com/jackc/pgx](https://github.com/jackc/pgx) for arguments details.
