# BookStoragePostgresqlgRPC

---

This service is used to store books in a Postgresql database with the server part located in a Docker container.

---
## Modules used:
[github.com/jackc/pgx](github.com/jackc/pgx)  
[google.golang.org/grpc](google.golang.org/grpc)


Example Usage
-----


```go
//config.go
package config

const (
	Login = "login"
	Pass  = "passwoed"
	IP    = "ip address"
	Port  = "port"
	DB    = "database"
)
```
