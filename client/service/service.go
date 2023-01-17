package service

import (
	pb "BookStoragePostgresqlgRPC/proto"
	"client/domain"
	"context"
	"errors"
	"log"
)

type service struct {
	ctx  context.Context
	conn pb.GreeterClient
	//bookStorage   domain.BookStorage
	//clientStorage domain.ClientStorage

}

func NewServiceBook(ctx context.Context, conn pb.GreeterClient) *service {
	return &service{
		ctx:  ctx,
		conn: conn,

		//bookStorage:   bookstorage,
		//clientStorage: clientstorage,
	}
}

func (s *service) TakeABookService(clientId int64, ids []int64) []int64 {
	//books, err := s.conn.TakeABookStorage(s.ctx, &pb.TakeABookRequest{ClientId: clientId, Ids: ids})
	//if err != nil {
	//	println(err.Error())
	//	errors.New("internal server error")
	//}

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
	//err := s.bookStorage.CreateBookStorage(book)

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

	//return s.bookStorage.GetBooksStorage()
}

func (s *service) GetBookService(id int64) (domain.Book, error) {
	book, _ := s.conn.GetBookStorage(s.ctx, &pb.GetBookRequest{ID: id})
	return protoToBookConversion(book.Book), nil
	//	book, err := s.bookStorage.GetBookStorage(id)
	//	return book, err
}

func (s *service) ReturnABookService(id int64) error {
	//s.conn.ReturnABookStorage(s.ctx, &pb.ReturnABookRequest{ID: id})
	//
	//return nil

	book, err := s.conn.GetBookStorage(s.ctx, &pb.GetBookRequest{ID: id})
	if err != nil {
		return err
	}

	if !book.Book.IsTaken {
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
	//err := s.bookStorage.DeleteBookStorage(id)
	if err != nil {
		return err
	}

	//
	return nil
}

func (s *service) UpdateBookService(id int64, book domain.Book) bool {
	request := bookToProtoConversion(book)
	reply, _ := s.conn.UpdateBookStorage(s.ctx, &pb.UpdateBookRequest{ID: id, Book: request})
	//s.bookStorage.UpdateBookStorage(id, book)

	return reply.StatusReply
}

func (s *service) CreateClientService(client domain.Client) error {
	request := clientToProtoConversion(client)
	_, err := s.conn.CreateClientStorage(s.ctx, &pb.CreateClientRequest{Client: request})
	if err != nil {
		return err
	}
	//err := s.clientStorage.CreateClientStorage(client)

	return nil

}

func (s *service) GetClientsService() map[int64]domain.Client {
	ClientList, _ := s.conn.GetClientsStorage(s.ctx, &pb.GetClientsRequest{})
	reply := protoToClientListConversion(ClientList.ClientList)
	return reply

}

func (s *service) DeleteClientService(id int64) error {
	//_, err := s.conn.DeleteClientStorage(s.ctx, &pb.DeleteClientRequest{ID: id})
	//if err != nil {
	//	return err
	//}
	//return nil

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
	//client, err := s.clientStorage.GetClientStorage(id)
	//
	//if err != nil {
	//	println(err.Error())
	//	return errors.New("internal server error")
	//}
	//
	//if client.ID == 0 {
	//	return errors.New("клиент не найден")
	//}
	//
	//haveBooks, err := s.bookStorage.StatusClientByBooks(id)
	//if err != nil {
	//	println(err.Error())
	//	return errors.New("internal server error")
	//}
	//if haveBooks {
	//	return errors.New("клиент не вернул книги")
	//}
	//
	//err = s.clientStorage.DeleteClientStorage(id)
	//
	//if err != nil {
	//	println(err.Error())
	//	return errors.New("internal server error")
	//}
	//
	//return nil
}

func (s *service) GetClientService(id int64) (domain.Client, bool) {
	reply, err := s.conn.GetClientStorage(s.ctx, &pb.GetClientRequest{ID: id})
	if err != nil {
		return domain.Client{}, false
	}
	client := protoToClientConversion(reply.Client)
	return client, true
	//client, _ := s.clientStorage.GetClientStorage(id)
	//
	//if client.ID == 0 {
	//	return client, false
	//}
	//
	//return client, true
}

func (s *service) UpdateClientService(id int64, client domain.Client) (bool, error) {
	request := clientToProtoConversion(client)
	_, err := s.conn.UpdateClientStorage(s.ctx, &pb.UpdateClientRequest{ID: id, Client: request})
	if err != nil {
		return false, err
	}
	return true, nil
	//clientOld, err := s.clientStorage.GetClientStorage(id)
	//
	//if clientOld.ID == 0 {
	//	return false, err
	//}
	//
	//error := s.clientStorage.UpdateClientStorage(id, client)
	//
	//return true, error
}

func (s *service) GetBooksByClientId(id int64) (map[int64]domain.Book, error) {
	reply, err := s.conn.GetBooksByClientIdStorage(s.ctx, &pb.GetBooksByClientIdRequest{ID: id})
	if err != nil {
		return nil, err
	}
	booeList := protoToBookListConversion(reply.BookList)
	return booeList, nil
	//booklist, err := s.bookStorage.GetBooksByClientId(id)
	//
	//return booklist, err
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
