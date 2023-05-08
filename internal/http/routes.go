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
		authors.HandleFunc("", handlers.CreateAuthor).Methods("POST")
		authors.HandleFunc("/{id}/books", handlers.GetAuthorBooks).Methods("GET")
	}
	member := router.PathPrefix("/members").Subrouter()
	{
		member.HandleFunc("", handlers.GetMembers).Methods("GET")
		member.HandleFunc("", handlers.CreateMember).Methods("POST")
		member.HandleFunc("/{id}/books", handlers.AddBookToMember).Methods("POST")
		member.HandleFunc("/{id}/books", handlers.GetMemberBooks).Methods("GET")
	}
	book := router.PathPrefix("/books").Subrouter()
	{
		book.HandleFunc("", handlers.GetBooks).Methods("GET")
		book.HandleFunc("", handlers.UpdateBookName).Methods("PUT")
		book.HandleFunc("", handlers.DeleteBook).Methods("DELETE")
		book.HandleFunc("", handlers.CreateBook).Methods("POST")
	}

	return router
}
