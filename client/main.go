package main

import (
	"client/handler/clientHandler"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"time"

	pb "BookStoragePostgresqlgRPC/proto"
	"client/config"
	"client/handler/bookHandler"
	"client/service"

	"github.com/gorilla/mux"
)

var url = fmt.Sprintf("%s:%s", config.GrpcIP, config.GrpcPort)
var addr = flag.String("addr", url, "the address to connect to")

func main() {

	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	serviceBook := service.NewServiceBook(ctx, c)
	BookHandler := bookHandler.NewBookHandler(serviceBook)
	ClientHandler := clientHandler.NewClientHandler(serviceBook)

	route := mux.NewRouter()

	route.HandleFunc("/book/list", BookHandler.GetBooks).Methods("GET")
	route.HandleFunc("/book/{id}", BookHandler.GetBook).Methods("GET")
	route.HandleFunc("/book/create", BookHandler.CreateBook).Methods("POST")
	route.HandleFunc("/book/{id}", BookHandler.UpdateBook).Methods("PUT")
	route.HandleFunc("/book/{id}", BookHandler.DeleteBook).Methods("DELETE")
	route.HandleFunc("/book/take/{clientId}", BookHandler.TakeABook).Methods("POST")
	route.HandleFunc("/book/return/{id}", BookHandler.ReturnABook).Methods("GET")
	route.HandleFunc("/client/booklist/{id}", BookHandler.GetBooksByClientId).Methods("GET")

	route.HandleFunc("/client/create", ClientHandler.CreateClient).Methods("POST")
	route.HandleFunc("/client/list", ClientHandler.GetClients).Methods("GET")
	route.HandleFunc("/client/{id}", ClientHandler.DeleteClient).Methods("DELETE")
	route.HandleFunc("/client/{id}", ClientHandler.GetClient).Methods("GET")
	route.HandleFunc("/client/{id}", ClientHandler.UpdateClient).Methods("PUT")

	http.Handle("/", route)

	urlServer := fmt.Sprintf("%s:%s", config.ServerIP, config.ServerPort)

	srv := &http.Server{
		Handler: route,
		Addr:    urlServer,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
