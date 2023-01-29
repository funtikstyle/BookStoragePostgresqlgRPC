package clientHandler

import (
	"client/domain"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type ClientHandler struct {
	service domain.ClientService
}

func NewClientHandler(clientService domain.ClientService) *ClientHandler {
	return &ClientHandler{service: clientService}
}

func (ch *ClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	newClient := domain.Client{}

	err := json.NewDecoder(r.Body).Decode(&newClient)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	err = ch.service.CreateClientService(newClient)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	response, err := json.Marshal(newClient)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	_, _ = w.Write(response)
}

func (ch *ClientHandler) GetClients(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(ch.service.GetClientsService())

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	_, _ = w.Write(response)
}

func (ch *ClientHandler) DeleteClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 0)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	err = ch.service.DeleteClientService(id)

	if err != nil {
		if err.Error() == "internal server error" {
			http.Error(w, "", http.StatusInternalServerError)

			return
		}

		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
}

func (ch *ClientHandler) GetClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 0)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	book, err := ch.service.GetClientService(id)

	if err != nil {
		http.Error(w, "Клиент отсутствует", http.StatusInternalServerError)

		return
	}

	response, err := json.Marshal(book)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	_, _ = w.Write(response)
}

func (ch *ClientHandler) UpdateClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 0)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	newClient := domain.Client{}
	err = json.NewDecoder(r.Body).Decode(&newClient)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	err = ch.service.UpdateClientService(id, newClient)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	//if !ok {
	//	http.Error(w, "Клиент отсутствует", http.StatusBadRequest)
	//
	//	return
	//}
}
