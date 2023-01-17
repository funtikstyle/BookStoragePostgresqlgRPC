# BookStoragePostgresqlgRPC

---

This service is used to store books in a Postgresql database with the server part located in a Docker container.

---
## Modules used:
[github.com/jackc/pgx](github.com/jackc/pgx)  
[google.golang.org/grpc](google.golang.org/grpc)


Example Usage
-----

server configuration:
```go
//config.go
package config

const (
	DbLogin = "database login"
	DBPass  = "database password"
	DbIP    = "database ip address"
	DbPort  = "database port"
	DB    = "database name"
	
	ServerPort = "The server port"
)
```
---

client configuration:
```go
//client/config.go
package config

const (
	IP   = "server IP address"
	Port = "server port"
)
```