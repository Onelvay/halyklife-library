package http

import (
	"encoding/json"
	"github.com/Onelvay/halyklife-library/internal/repository"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	repo *repository.Repo
}

func newHandler(repo *repository.Repo) *Handler {
	return &Handler{repo}
}

// - /authors/{id}/books - список книг автора которые есть в наличии
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

// - /members/{id}/books - список книг взятых читателем
func (h *Handler) GetMemberBooks(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	books, err := h.repo.GetMemberBooks(id)
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

// - /authors
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

// - /books
func (h *Handler) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.repo.GetBooks()
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

// - /members
func (h *Handler) GetMembers(w http.ResponseWriter, r *http.Request) {
	members, err := h.repo.GetMembers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	js, err := json.Marshal(members)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	w.WriteHeader(http.StatusOK)
}
