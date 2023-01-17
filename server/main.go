package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"strconv"
	"strings"

	"log"
	"net"
	"os"

	pb "BookStoragePostgresqlgRPC/proto"

	"BookStoragePostgresqlgRPC/config"
	"github.com/jackc/pgx/v5"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", config.ServerPort, "The server port")
)

type server struct {
	pb.UnimplementedGreeterServer
	connect *pgx.Conn
}

func (s *server) CreateBookStorage(ctx context.Context, book *pb.CreateBookRequest) (*pb.CreateBookReply, error) {
	_, err := s.connect.Query(
		context.Background(),
		fmt.Sprintf(
			"INSERT INTO book (\"Author\", \"Title\")"+
				"VALUES ('%s', '%s')", book.Book.Author, book.Book.Title),
	)

	if err != nil {
		return nil, err
	}

	return &pb.CreateBookReply{}, nil
}

func (s *server) GetBooksStorage(ctx context.Context, in *pb.GetBooksRequest) (*pb.BooksReply, error) {
	rows, err := s.connect.Query(context.Background(), "SELECT * FROM book")
	if err != nil {
		return nil, err
	}

	answer := &pb.BooksReply{}
	answer.Books, err = rowsToBooksList(rows)

	return answer, err
}

func (s *server) DeleteBookStorage(ctx context.Context, request *pb.DeleteBookRequest) (*pb.DeleteBookReply, error) {
	_, err := s.connect.Query(
		context.Background(),
		fmt.Sprintf(
			"DELETE FROM book "+
				"WHERE \"id\" = %d", request.Id),
	)
	if err != nil {
		return nil, err

	}

	return &pb.DeleteBookReply{}, nil
}

func (s *server) GetBookStorage(ctx context.Context, request *pb.GetBookRequest) (*pb.GetBookReply, error) {
	reply := &pb.GetBookReply{}

	rows, err := s.connect.Query(
		context.Background(),
		fmt.Sprintf(
			"SELECT id, Author, Title, ClientID, isTaken"+
				"FROM book AS b WHERE b.id = %d", request.ID),
	)

	if err != nil {
		return reply, err
	}

	if rows.Next() {
		val, err := rows.Values()
		if err != nil {
			return reply, err
		}

		reply.Book = valuesToBook(val)
	}

	return reply, nil
}

func (s *server) UpdateBookStorage(ctx context.Context, request *pb.UpdateBookRequest) (*pb.UpdateBookReply, error) {
	book := request.Book
	id := request.ID
	_, err := s.connect.Query(
		context.Background(),
		fmt.Sprintf(
			"UPDATE book "+
				"SET \"Author\" = '%s', "+
				"\"Title\" = '%s', "+
				"\"ClientID\" = %d "+
				"WHERE \"id\" = %d", book.Author, book.Title, book.ClientID, id),
	)
	if err != nil {
		return &pb.UpdateBookReply{}, err
	}
	return &pb.UpdateBookReply{}, nil
}

func (s *server) StatusClientByBooks(ctx context.Context, request *pb.StatusClientByBooksRequest) (*pb.StatusClientByBooksReply, error) {
	id := request.ID

	rows, err := s.connect.Query(context.Background(),
		fmt.Sprintf("SELECT * FROM book WHERE \"ClientID\" = %d LIMIT 1", id))

	if err != nil {
		println(err.Error())

		return &pb.StatusClientByBooksReply{IsTaken: false}, err
	}

	if rows.Next() {
		return &pb.StatusClientByBooksReply{IsTaken: true}, nil
	}

	return &pb.StatusClientByBooksReply{IsTaken: false}, nil
}

func (s *server) GetBooksByClientId(ctx context.Context, request *pb.GetBooksByClientIdRequest) (*pb.GetBooksByClientIdReply, error) {
	id := request.ID

	rows, err := s.connect.Query(context.Background(),
		fmt.Sprintf("SELECT * FROM book WHERE \"ClientID\" = %d ", id))

	if err != nil {
		println(err.Error())

		return nil, err
	}

	books, err := rowsToBooksList(rows)

	if err != nil {
		println(err.Error())

		return nil, err
	}

	return &pb.GetBooksByClientIdReply{BookList: books}, nil
}

func (s *server) GetNotTakenBookByIds(ctx context.Context, request *pb.GetNotTakenBookByIdsRequest) (*pb.GetNotTakenBookByIdsReply, error) {
	ids := request.Ids
	idsString := make([]string, len(ids))

	for k, v := range ids {
		idsString[k] = strconv.FormatInt(v, 10)
	}

	rows, _ := s.connect.Query(context.Background(),
		fmt.Sprintf("SELECT * FROM book WHERE \"id\" IN (%s) AND \"isTaken\" = false", strings.Join(idsString, ",")))
	books, _ := rowsToBooksList(rows)

	return &pb.GetNotTakenBookByIdsReply{BookList: books}, nil
}

func (s *server) CreateClientStorage(ctx context.Context, request *pb.CreateClientRequest) (*pb.CreateClientReply, error) {
	client := request.Client

	_, err := s.connect.Query(
		context.Background(),
		fmt.Sprintf(
			"INSERT INTO client (\"Name\", \"Phone\")"+
				"VALUES ('%s', '%s')", client.Name, client.PhoneName),
	)
	if err != nil {

		return &pb.CreateClientReply{}, err
	}

	return &pb.CreateClientReply{}, nil
}

func (s *server) GetClientsStorage(ctx context.Context, request *pb.GetClientsRequest) (*pb.GetClientsReply, error) {

	rows, err := s.connect.Query(context.Background(), "SELECT * FROM client")

	if err != nil {

		return nil, err
	}

	answer := &pb.GetClientsReply{}
	answer.ClientList, err = rowsToClientsList(rows)

	return answer, err
}

func (s *server) DeleteClientStorage(ctx context.Context, request *pb.DeleteClientRequest) (*pb.DeleteClientReply, error) {
	id := request.ID

	_, err := s.connect.Query(
		context.Background(),
		fmt.Sprintf(
			"DELETE FROM client "+
				"WHERE \"id\" = %d", id),
	)

	if err != nil {

		return &pb.DeleteClientReply{}, err
	}

	return &pb.DeleteClientReply{}, nil
}

func (s *server) GetClientStorage(ctx context.Context, request *pb.GetClientRequest) (*pb.GetClientReply, error) {
	id := request.ID

	rows, err := s.connect.Query(context.Background(),
		fmt.Sprintf("SELECT \"id\", \"Name\", \"Phone\" FROM clientWHERE\"id\" = %d", id))

	if err != nil {
		return &pb.GetClientReply{}, err
	}

	client := &pb.Client{}

	if rows.Next() {
		val, err := rows.Values()
		if err != nil {

			return &pb.GetClientReply{}, err
		}
		client = valuesToClient(val)
	}

	return &pb.GetClientReply{Client: client, StatusReply: true}, nil
}

func (s *server) UpdateClientStorage(ctx context.Context, request *pb.UpdateClientRequest) (*pb.UpdateClientReply, error) {
	client := request.Client
	id := request.ID

	_, err := s.connect.Query(
		context.Background(),
		fmt.Sprintf(
			"UPDATE book "+
				"SET \"Name\" = '%s', "+
				"\"Phone\" = '%s', "+
				"WHERE \"id\" = %d", client.Name, client.PhoneName, id),
	)

	if err != nil {

		return &pb.UpdateClientReply{StatusReply: false}, err
	}

	return &pb.UpdateClientReply{StatusReply: true}, nil
}

func rowsToBooksList(rows pgx.Rows) (map[int64]*pb.Book, error) {
	booklist := map[int64]*pb.Book{}

	for rows.Next() {
		val, err := rows.Values()
		if err != nil {
			return nil, errors.New("error while iterating dataset")
		}

		book := valuesToBook(val)
		booklist[int64(book.ID)] = book

	}

	return booklist, nil
}

func rowsToClientsList(rows pgx.Rows) (map[int64]*pb.Client, error) {
	clientlist := map[int64]*pb.Client{}

	for rows.Next() {
		val, err := rows.Values()
		if err != nil {
			return nil, errors.New("error while iterating dataset")
		}

		client := valuesToClient(val)
		clientlist[int64(client.ID)] = client

	}

	return clientlist, nil
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

func valuesToClient(val []any) *pb.Client {
	id := val[0].(int32)
	Name := val[1].(string)
	PhoneName := val[2].(string)

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
