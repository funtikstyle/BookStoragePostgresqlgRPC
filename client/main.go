package main

import (
	"client/handler/clientHandler"
	pb "client/proto"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"time"

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
	//conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	//c := pb.NewGreeterClient()

	//ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	//ctx, cancel := context.WithCancel(context.Background())
	ctx := context.Background()
	//defer cancel()

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
		WriteTimeout: 45 * time.Second,
		ReadTimeout:  45 * time.Second,
	}

	log.Println("Client started...")
	log.Fatal(srv.ListenAndServe())
}
