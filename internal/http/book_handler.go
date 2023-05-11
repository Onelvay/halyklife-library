package http

import (
	"encoding/json"
	"github.com/Onelvay/halyklife-library/pkg/domain"
	"io/ioutil"
	"net/http"
	"sort"
)

func (h *Handler) GetBooks(w http.ResponseWriter, r *http.Request) {
	querySort := r.URL.Query().Get("sort")
	books, err := h.repo.GetBooks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if querySort == "price" {
		sort.Slice(books, func(i, j int) bool {
			return books[i].Price > books[j].Price
		})
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

func (h *Handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var book domain.Book
	err = json.Unmarshal(bytes, &book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = h.repo.CreateBook(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
func (h *Handler) UpdateBookName(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var req struct {
		BookId  string `json:"book_id"`
		NewName string `json:"name"`
	}
	err = json.Unmarshal(bytes, &req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = h.repo.UpdateBookNameById(req.BookId, req.NewName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func (h *Handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var req struct {
		BookId string `json:"book_id"`
	}
	err = json.Unmarshal(bytes, &req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = h.repo.DeleteBookById(req.BookId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
