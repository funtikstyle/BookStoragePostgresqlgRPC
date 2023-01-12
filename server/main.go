package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	//"github.com/jackc/pgx"
	"log"
	"net"
	"os"

	// "server/pgxpool"
	// "server/pgx/pgxpool"
	pb "BookStoragePostgresqlgRPC/proto"

	// "github.com/jackc/pgx"
	"BookStoragePostgresqlgRPC/config"
	"github.com/jackc/pgx/v5"
	// "github.com/jackc/pgx/v5"

	//"github.com/jackc/pgx/pgxpool"
	// "/user/local/go/jackc/pgx/pgxpool"
	// "pgx"
	// "server/pgxpool"

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
	// pb.list[ai] = book
	// ai++
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

func (s *server) DeleteBookStorage(ctx context.Context, id *pb.DeleteBookRequest) (*pb.DeleteBookReply, error) {
	_, err := s.connect.Query(
		context.Background(),
		fmt.Sprintf(
			"DELETE FROM book "+
				"WHERE \"id\" = %d", id.Id),
	)
	if err != nil {
		return nil, err

	}

	return &pb.DeleteBookReply{}, nil
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

func main() {

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.Login, config.Pass, config.IP, config.Port, config.DB)
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
