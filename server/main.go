package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"server/config"
	pb "server/proto"
	"strconv"
	"strings"
)

var (
	port = flag.Int("port", config.ServerPort, "The server port")
)

type server struct {
	pb.UnimplementedGreeterServer
	connect *pgx.Conn
}

func (s *server) CreateBookStorage(ctx context.Context, book *pb.CreateBookRequest) (*pb.CreateBookReply, error) {
	reply := pb.CreateBookReply{}
	rows, err := s.connect.Query(
		context.Background(),
		fmt.Sprintf(
			"INSERT INTO book (\"Author\", \"Title\")"+
				"VALUES ('%s', '%s')", book.Book.Author, book.Book.Title),
	)
	defer rows.Close()

	if err != nil {
		log.Println(err)
		return &reply, errors.New("internal server error")
	}

	return &reply, nil
}

func (s *server) GetBooksStorage(ctx context.Context, in *pb.GetBooksRequest) (*pb.BooksReply, error) {
	reply := pb.BooksReply{}
	rows, err := s.connect.Query(context.Background(), "SELECT * FROM book")
	defer rows.Close()

	if err != nil {
		log.Println(err)
		return &reply, errors.New("internal server error")
	}

	reply.Books, err = rowsToBooksList(rows)

	return &reply, nil
}

func (s *server) GetBookStorage(ctx context.Context, request *pb.GetBookRequest) (*pb.GetBookReply, error) {
	reply := pb.GetBookReply{}

	rows, err := s.connect.Query(
		context.Background(),
		fmt.Sprintf(
			"SELECT b.id, b.\"Author\", b.\"Title\", b.\"ClientID\", b.\"isTaken\" "+
				"FROM book AS b WHERE b.id = %d", request.ID),
	)
	defer rows.Close()
	if err != nil {
		log.Println(err)

		return &reply, errors.New("internal server error")
	}

	if rows.Next() {
		val, err := rows.Values()
		if err != nil {
			log.Println(err)
			return &reply, errors.New("internal server error")
		}

		reply.Book = valuesToBook(val)
	}

	return &reply, nil
}

func (s *server) DeleteBookStorage(ctx context.Context, request *pb.DeleteBookRequest) (*pb.DeleteBookReply, error) {
	reply := pb.DeleteBookReply{}
	rows, err := s.connect.Query(
		context.Background(),
		fmt.Sprintf(
			"DELETE FROM book "+
				"WHERE \"id\" = %d", request.Id),
	)
	defer rows.Close()

	if err != nil {
		log.Println(err)
		return nil, errors.New("internal server error")
	}

	return &reply, nil
}

func (s *server) UpdateBookStorage(ctx context.Context, request *pb.UpdateBookRequest) (*pb.UpdateBookReply, error) {
	reply := pb.UpdateBookReply{}
	book := request.Book
	id := request.ID

	rows, err := s.connect.Query(
		context.Background(),
		fmt.Sprintf(
			"UPDATE book "+
				"SET \"Author\" = '%s', "+
				"\"Title\" = '%s', "+
				"\"ClientID\" = %d, "+
				"\"isTaken\" = '%t' "+
				"WHERE \"id\" = %d ", book.Author, book.Title, book.ClientID, book.IsTaken, id),
	)
	defer rows.Close()

	if err != nil {
		log.Println("internal server error")
		return &reply, errors.New("internal server error")
	}

	return &reply, nil
}

func (s *server) StatusClientByBooks(ctx context.Context, request *pb.StatusClientByBooksRequest) (*pb.StatusClientByBooksReply, error) {
	reply := pb.StatusClientByBooksReply{IsTaken: false}
	id := request.ID

	rows, err := s.connect.Query(context.Background(),
		fmt.Sprintf("SELECT * FROM book WHERE \"ClientID\" = %d LIMIT 1", id))
	defer rows.Close()

	if err != nil {
		log.Println(err)
		return &reply, errors.New("internal server error")
	}

	if rows.Next() {
		reply.IsTaken = true
		return &reply, nil
	}

	return &reply, nil
}

func (s *server) GetBooksByClientIdStorage(ctx context.Context, request *pb.GetBooksByClientIdRequest) (*pb.GetBooksByClientIdReply, error) {
	reply := pb.GetBooksByClientIdReply{}
	id := request.ID

	rows, err := s.connect.Query(context.Background(),
		fmt.Sprintf("SELECT * FROM book WHERE \"ClientID\" = %d ", id))
	defer rows.Close()

	if err != nil {
		log.Println(err)

		return nil, errors.New("internal server error")
	}

	books, err := rowsToBooksList(rows)

	if err != nil {
		log.Println(err)

		return &reply, errors.New("internal server error")
	}
	reply.BookList = books
	return &reply, nil
}

func (s *server) GetNotTakenBookByIds(ctx context.Context, request *pb.GetNotTakenBookByIdsRequest) (*pb.GetNotTakenBookByIdsReply, error) {
	reply := pb.GetNotTakenBookByIdsReply{}
	ids := request.Ids
	idsString := make([]string, len(ids))

	for k, v := range ids {
		idsString[k] = strconv.FormatInt(v, 10)
	}

	rows, _ := s.connect.Query(context.Background(),
		fmt.Sprintf("SELECT * FROM book WHERE \"id\" IN (%s) AND \"isTaken\" = false", strings.Join(idsString, ",")))
	defer rows.Close()

	books, _ := rowsToBooksList(rows)
	reply.BookList = books

	return &reply, nil
}

func (s *server) CreateClientStorage(ctx context.Context, request *pb.CreateClientRequest) (*pb.CreateClientReply, error) {
	reply := pb.CreateClientReply{}

	rows, err := s.connect.Query(
		context.Background(),
		fmt.Sprintf(
			"INSERT INTO client (\"Name\", \"Phone\") "+
				"VALUES ('%s', %d)", request.Client.Name, request.Client.PhoneName),
	)
	defer rows.Close()

	if err != nil {
		log.Println("internal server error")
		return &reply, errors.New("internal server error")
	}

	return &reply, nil
}

func (s *server) GetClientsStorage(ctx context.Context, request *pb.GetClientsRequest) (*pb.GetClientsReply, error) {
	reply := &pb.GetClientsReply{}

	rows, err := s.connect.Query(context.Background(), "SELECT * FROM client")
	defer rows.Close()

	if err != nil {
		log.Println(err)
		return reply, errors.New("internal server error")
	}

	reply.ClientList, err = rowsToClientsList(rows)

	return reply, err
}

func (s *server) DeleteClientStorage(ctx context.Context, request *pb.DeleteClientRequest) (*pb.DeleteClientReply, error) {
	reply := pb.DeleteClientReply{}
	id := request.ID

	rows, err := s.connect.Query(
		context.Background(),
		fmt.Sprintf(
			"DELETE FROM \"client\" "+
				"WHERE \"id\" = %d", id),
	)
	defer rows.Close()

	if err != nil {
		log.Println(err)
		return &reply, errors.New("internal server error")
	}

	return &reply, nil
}

func (s *server) GetClientStorage(ctx context.Context, request *pb.GetClientRequest) (*pb.GetClientReply, error) {
	reply := pb.GetClientReply{}
	id := request.ID

	rows, err := s.connect.Query(context.Background(),
		fmt.Sprintf("SELECT \"id\", \"Name\", \"Phone\" FROM client WHERE \"id\" = %d", id))
	defer rows.Close()

	if err != nil {
		log.Println(err)
		return &reply, errors.New("internal server error")
	}

	client := &pb.Client{}

	if rows.Next() {
		val, err := rows.Values()

		if err != nil {
			log.Println(err)
			return &reply, errors.New("internal server error")
		}

		client = valuesToClient(val)
	}
	reply.Client = client
	//reply.StatusReply = true
	return &reply, nil
}

func (s *server) UpdateClientStorage(ctx context.Context, request *pb.UpdateClientRequest) (*pb.UpdateClientReply, error) {
	reply := pb.UpdateClientReply{}
	client := request.Client
	id := request.ID

	queryString := fmt.Sprintf(
		"UPDATE client "+
			"SET \"Name\" = '%s', "+
			"\"Phone\" = %d "+
			"WHERE \"id\" = %d", client.Name, client.PhoneName, id)

	rows, err := s.connect.Query(
		context.Background(), queryString)
	defer rows.Close()

	if err != nil {
		//reply.StatusReply = false
		log.Println(err)
		return &reply, errors.New("internal server error")
	}
	//reply.StatusReply = true
	return &reply, nil
}

func rowsToBooksList(rows pgx.Rows) (map[int64]*pb.Book, error) {
	booklist := map[int64]*pb.Book{}

	for rows.Next() {
		val, err := rows.Values()
		if err != nil {
			log.Println(err)
			return nil, errors.New("error while iterating dataset")
		}

		book := valuesToBook(val)
		booklist[int64(book.ID)] = book
	}
	rows.Close()

	return booklist, nil
}

func valuesToBook(val []any) *pb.Book {
	id := val[0].(int32)
	Author := val[1].(string)
	Title := val[2].(string)
	var ClientID int64

	if val[3] != nil {
		ClientID = int64(val[3].(int32))
	}

	IsTaken := val[4].(bool)

	book := pb.Book{
		ID:       id,
		Author:   Author,
		Title:    Title,
		ClientID: ClientID,
		IsTaken:  IsTaken,
	}

	return &book
}

func rowsToClientsList(rows pgx.Rows) (map[int64]*pb.Client, error) {
	clientlist := map[int64]*pb.Client{}

	for rows.Next() {
		val, err := rows.Values()
		if err != nil {
			log.Println(err)
			return nil, errors.New("error while iterating dataset")
		}

		client := valuesToClient(val)
		clientlist[int64(client.ID)] = client

	}
	rows.Close()

	return clientlist, nil
}

func valuesToClient(val []any) *pb.Client {
	id := val[0].(int32)
	Name := val[1].(string)
	PhoneName := val[2].(int32)

	client := pb.Client{
		ID:        id,
		Name:      Name,
		PhoneName: PhoneName,
	}

	return &client
}

func main() {

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.DbLogin, config.DbPass, config.DbIP, config.DbPort, config.DB)
	dbpool, err := pgx.Connect(context.Background(), url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	defer dbpool.Close(context.Background())

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{connect: dbpool})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
