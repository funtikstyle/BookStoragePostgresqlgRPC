package bookHandler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"client/domain"
	"github.com/gorilla/mux"
)

type BookHandler struct {
	service domain.BookService
}

func NewBookHandler(bookService domain.BookService) *BookHandler {
	return &BookHandler{service: bookService}
}

func (bh *BookHandler) TakeABook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["clientId"], 10, 0)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	var dat map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&dat)

	if err != nil {
		log.Println(err)

		return
	}

	strs := dat["ids"].([]interface{})
	ids := []int64{}

	for _, val := range strs {
		ids = append(ids, int64(val.(float64)))
	}

	response, err := json.Marshal(bh.service.TakeABookService(id, ids))
	if err != nil {
		log.Println(err)

		return
	}
	_, _ = w.Write(response)
}

func (bh *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	newBook := domain.Book{}
	err := json.NewDecoder(r.Body).Decode(&newBook)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	err = bh.service.CreateBookService(newBook)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	response, err := json.Marshal(newBook.Author)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
	_, _ = w.Write(response)
}

func (bh *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := bh.service.GetBooksService()

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	response, err := json.Marshal(books)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	_, _ = w.Write(response)
}

func (bh *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 0)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	book, err := bh.service.GetBookService(id)

	if err != nil {
		log.Println(err)
		http.Error(w, "Книга отсутствует", http.StatusBadRequest)

		return
	}

	response, err := json.Marshal(book)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
	_, _ = w.Write(response)
}

func (bh *BookHandler) ReturnABook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 0)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	err = bh.service.ReturnABookService(id)

	if err != nil {
		log.Println(err)
		http.Error(w, "Книга отсутствует", http.StatusBadRequest)

		return
	}
}

func (bh *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 0)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	err = bh.service.DeleteBookService(id)

	if err != nil {
		log.Println(err)
		http.Error(w, "Книга отсутствует", http.StatusBadRequest)

		return
	}
}

func (bh *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 0)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	newBook := domain.Book{}
	err = json.NewDecoder(r.Body).Decode(&newBook)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	err = bh.service.UpdateBookService(id, newBook)

	if err != nil {
		http.Error(w, "Книга отсутствует", http.StatusBadRequest)

		return
	}
}

func (bh *BookHandler) GetBooksByClientId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 0)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	bookList, err := bh.service.GetBooksByClientId(id)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	response, err := json.Marshal(bookList)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
	_, _ = w.Write(response)
}
