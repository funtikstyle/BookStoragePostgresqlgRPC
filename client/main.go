package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"BookStoragePostgresqlgRPC/config"
	pb "BookStoragePostgresqlgRPC/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = flag.String("addr", config.GetIP(), "the address to connect to")

var id string

func main() {

	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	fmt.Scan(&id)
	fmt.Println(id)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// r, err := c.GetBooksStorage(ctx, &pb.GetBooksRequest{})
	// if err != nil {
	// 	log.Fatalf("could not greet: %v", err)
	// }
	// fmt.Println(r.Books)

	book := pb.Book{Author: "Vasya", Title: "1234567"}
	_, err = c.CreateBookStorage(ctx, &pb.CreateBookRequest{Book: &book})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// answer, _ := strconv.ParseInt(id, 10, 64)
	// bookDelete := pb.DeleteBookRequest{Id: answer}
	// _, err = c.DeleteBookStorage(ctx, &bookDelete)
	// if err != nil {
	// 	log.Fatalf("could not greet: %v", err)
	// }

}
