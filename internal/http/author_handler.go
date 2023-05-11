package http

import (
	"encoding/json"
	"github.com/gorilla/mux"
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
