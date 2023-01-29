package service

import (
	"client/domain"
	pb "client/proto"
	"context"
	"errors"
	"log"
)

type service struct {
	ctx  context.Context
	conn pb.GreeterClient
}

func NewServiceBook(ctx context.Context, conn pb.GreeterClient) *service {
	return &service{
		ctx:  ctx,
		conn: conn,
	}
}

func (s *service) TakeABookService(clientId int64, ids []int64) []int64 {
	takeBooks, _ := s.conn.GetNotTakenBookByIds(s.ctx, &pb.GetNotTakenBookByIdsRequest{Ids: ids})
	var listId []int64

	for id, book := range takeBooks.BookList {
		log.Println(clientId)
		book.ClientID = clientId
		book.IsTaken = true

		_, err := s.conn.UpdateBookStorage(s.ctx, &pb.UpdateBookRequest{ID: id, Book: book})

		if err != nil {
			println(err.Error())
			errors.New("internal server error")
		}

		listId = append(listId, id)
	}

	return listId
}

func (s *service) CreateBookService(book domain.Book) error {
	_, err := s.conn.CreateBookStorage(s.ctx, &pb.CreateBookRequest{Book: bookToProtoConversion(book)})

	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetBooksService() (map[int64]domain.Book, error) {
	protoBooksReply, err := s.conn.GetBooksStorage(s.ctx, &pb.GetBooksRequest{})

	if err != nil {
		return nil, err
	}

	books := protoToBookListConversion(protoBooksReply.Books)

	return books, nil
}

func (s *service) GetBookService(id int64) (domain.Book, error) {
	book, err := s.conn.GetBookStorage(s.ctx, &pb.GetBookRequest{ID: id})

	if err != nil {
		return domain.Book{}, err
	}

	return protoToBookConversion(book.Book), nil
}

func (s *service) ReturnABookService(id int64) error {
	book, err := s.conn.GetBookStorage(s.ctx, &pb.GetBookRequest{ID: id})
	if err != nil {
		return err
	}

	if !book.Book.IsTaken {

		err.Error()
		return err
	}

	book.Book.ClientID = 0
	book.Book.IsTaken = false

	_, ok := s.conn.UpdateBookStorage(s.ctx, &pb.UpdateBookRequest{ID: id, Book: book.Book})

	if ok != nil {
		println(err.Error())
		errors.New("internal server error")

		return ok
	}

	return nil
}

func (s *service) DeleteBookService(id int64) error {
	_, err := s.conn.DeleteBookStorage(s.ctx, &pb.DeleteBookRequest{Id: id})

	if err != nil {
		return err
	}

	//
	return nil
}

func (s *service) UpdateBookService(id int64, book domain.Book) error {
	request := bookToProtoConversion(book)
	_, err := s.conn.UpdateBookStorage(s.ctx, &pb.UpdateBookRequest{ID: id, Book: request})
	if err != nil {
		return err
	}

	return nil
}

func (s *service) CreateClientService(client domain.Client) error {
	request := clientToProtoConversion(client)
	_, err := s.conn.CreateClientStorage(s.ctx, &pb.CreateClientRequest{Client: request})

	if err != nil {
		return err
	}

	return nil

}

func (s *service) GetClientsService() map[int64]domain.Client {
	clientList, _ := s.conn.GetClientsStorage(s.ctx, &pb.GetClientsRequest{})
	reply := protoToClientListConversion(clientList.ClientList)

	return reply
}

func (s *service) DeleteClientService(id int64) error {
	client, err := s.conn.GetClientStorage(s.ctx, &pb.GetClientRequest{ID: id})
	if err != nil {
		println(err.Error())
		return errors.New("internal server error")
	}
	if client.Client.ID == 0 {
		return errors.New("клиент не найден")
	}

	haveBooks, err := s.conn.StatusClientByBooks(s.ctx, &pb.StatusClientByBooksRequest{ID: id})
	if err != nil {
		println(err.Error())
		return errors.New("internal server error")
	}
	if haveBooks.IsTaken {
		return errors.New("клиент не вернул книги")
	}
	_, ok := s.conn.DeleteClientStorage(s.ctx, &pb.DeleteClientRequest{ID: id})
	if ok != nil {
		println(err.Error())

		return errors.New("internal server error")
	}

	return nil
}

func (s *service) GetClientService(id int64) (domain.Client, error) {
	reply, err := s.conn.GetClientStorage(s.ctx, &pb.GetClientRequest{ID: id})

	if err != nil {
		return domain.Client{}, err
	}

	client := protoToClientConversion(reply.Client)

	return client, nil
}

func (s *service) UpdateClientService(id int64, client domain.Client) error {
	request := clientToProtoConversion(client)
	_, err := s.conn.UpdateClientStorage(s.ctx, &pb.UpdateClientRequest{ID: id, Client: request})

	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetBooksByClientId(id int64) (map[int64]domain.Book, error) {
	reply, err := s.conn.GetBooksByClientIdStorage(s.ctx, &pb.GetBooksByClientIdRequest{ID: id})

	if err != nil {
		return nil, err
	}

	bookList := protoToBookListConversion(reply.BookList)

	return bookList, nil
}

func bookToProtoConversion(book domain.Book) *pb.Book {
	protoBook := pb.Book{
		ID:       book.ID,
		Author:   book.Author,
		Title:    book.Title,
		ClientID: book.ClientID,
		IsTaken:  book.IsTaken,
	}

	return &protoBook
}

func protoToBookConversion(book *pb.Book) domain.Book {
	domainBook := domain.Book{
		ID:       book.ID,
		Author:   book.Author,
		Title:    book.Title,
		ClientID: book.ClientID,
		IsTaken:  book.IsTaken,
	}
	return domainBook
}

func protoToBookListConversion(protoBookList map[int64]*pb.Book) map[int64]domain.Book {
	bookList := map[int64]domain.Book{}

	for key, val := range protoBookList {
		bookList[key] = protoToBookConversion(val)
	}

	return bookList

}

func bookListToProtoConversion(bookList map[int64]domain.Book) map[int64]*pb.Book {
	bookListProto := map[int64]*pb.Book{}
	for key, val := range bookList {
		bookListProto[key] = bookToProtoConversion(val)
	}
	return bookListProto
}

func clientToProtoConversion(client domain.Client) *pb.Client {
	protoClient := pb.Client{
		ID:        client.ID,
		Name:      client.Name,
		PhoneName: client.PhoneName,
	}

	return &protoClient
}

func protoToClientConversion(client *pb.Client) domain.Client {
	domainClient := domain.Client{
		ID:        client.ID,
		Name:      client.Name,
		PhoneName: client.PhoneName,
	}
	return domainClient
}

func protoToClientListConversion(protoClientList map[int64]*pb.Client) map[int64]domain.Client {
	clientList := map[int64]domain.Client{}

	for key, val := range protoClientList {
		clientList[key] = protoToClientConversion(val)
	}

	return clientList

}

func clientListToProtoConversion(clientList map[int64]domain.Client) map[int64]*pb.Client {
	clientListProto := map[int64]*pb.Client{}
	for key, val := range clientList {
		clientListProto[key] = clientToProtoConversion(val)
	}
	return clientListProto
}
