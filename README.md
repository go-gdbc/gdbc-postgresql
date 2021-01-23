# gdbc-postgresql
GDBC Postgresql Driver - It is based on [https://github.com/jackc/pgx](https://github.com/jackc/pgx)

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
