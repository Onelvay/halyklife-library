package http

import (
	"encoding/json"
	"github.com/Onelvay/halyklife-library/pkg/domain"
	"io/ioutil"
	"net/http"
)

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
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
func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
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
	err = h.repo.SignIn(author.Username, author.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
