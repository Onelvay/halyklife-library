package http

import (
	"encoding/json"
	"github.com/Onelvay/halyklife-library/pkg/domain"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

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

func (h *Handler) AddBookToMember(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var req struct {
		MemberId string `json:"member_id"`
		BookId   string `json:"book_id"`
	}
	err = json.Unmarshal(bytes, &req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = h.repo.AddBookToMember(req.BookId, req.MemberId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func (h *Handler) CreateMember(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var mem domain.Member
	err = json.Unmarshal(bytes, &mem)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = h.repo.CreateMember(mem)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
