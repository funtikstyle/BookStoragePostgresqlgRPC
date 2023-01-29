package domain

type Book struct {
	ID       int32
	Author   string
	Title    string
	ClientID int64
	IsTaken  bool
}

type Client struct {
	ID        int32
	Name      string
	PhoneName int32
}

type BookService interface {
	GetBookService(id int64) (Book, error)
	TakeABookService(id int64, ids []int64) []int64
	CreateBookService(book Book) error
	GetBooksService() (map[int64]Book, error)
	ReturnABookService(id int64) error
	DeleteBookService(id int64) error
	UpdateBookService(id int64, book Book) error
	GetBooksByClientId(id int64) (map[int64]Book, error)
}

type ClientService interface {
	CreateClientService(client Client) error
	GetClientsService() map[int64]Client
	DeleteClientService(id int64) error
	GetClientService(id int64) (Client, error)
	UpdateClientService(id int64, client Client) error
}
