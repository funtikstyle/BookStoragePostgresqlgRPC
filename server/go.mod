module server

go 1.19

require (
	//BookStoragePostgresqlgRPC/config v0.0.0
	//BookStoragePostgresqlgRPC/proto v0.0.0
	github.com/jackc/pgx/v5 v5.2.0
	google.golang.org/grpc v1.52.0
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	golang.org/x/crypto v0.5.0 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/text v0.6.0 // indirect
	google.golang.org/genproto v0.0.0-20221118155620-16455021b5e6 // indirect
)

//replace BookStoragePostgresqlgRPC/config => ../config

//replace BookStoragePostgresqlgRPC/proto => ../proto
//replace server/proto => ../proto
