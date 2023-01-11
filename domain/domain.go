package domain

type Book struct {
	ID       int32
	Author   string
	Title    string
	ClientID int64
	IsTaken  bool
}