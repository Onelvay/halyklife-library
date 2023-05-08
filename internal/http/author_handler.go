package http

import (
	"encoding/json"
	"github.com/Onelvay/halyklife-library/pkg/domain"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func (h *Handler) GetAuthorBooks(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	books, err := h.repo.GetAuthorBooks(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	js, err := json.Marshal(books)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := h.repo.GetAuthors()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	js, err := json.Marshal(authors)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	w.WriteHeader(http.StatusOK)
}
func (h *Handler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var author domain.Author
	err = json.Unmarshal(bytes, &author)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = h.repo.CreateAuthor(author)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
