package http

import (
	"github.com/Onelvay/halyklife-library/internal/repository"
	"github.com/gorilla/mux"
)

//Пути:
//- /authors
//- /books
//- /members
//
//Специальные пути:
//- /authors/{id}/books - список книг автора которые есть в наличии
//- /members/{id}/books - список книг взятых читателем

func InitRoutes(repo *repository.Repo) *mux.Router {
	handlers := newHandler(repo)

	router := mux.NewRouter().StrictSlash(true)

	authors := router.PathPrefix("/authors").Subrouter()
	{
		authors.HandleFunc("", handlers.GetAuthors).Methods("GET")
		authors.HandleFunc("/{id}/books", handlers.GetAuthorBooks).Methods("GET")
	}
	member := router.PathPrefix("/members").Subrouter()
	{
		member.HandleFunc("", handlers.GetMembers).Methods("GET")
		member.HandleFunc("/{id}/books", handlers.GetMemberBooks).Methods("GET")
	}
	router.HandleFunc("/books", handlers.GetBooks).Methods("GET")

	return router
}
