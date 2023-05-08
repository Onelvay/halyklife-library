package crud_app

import (
	httpRepo "github.com/Onelvay/halyklife-library/internal/http"
	"github.com/Onelvay/halyklife-library/internal/repository"
	"net/http"
)

func Run() {
	repo := repository.NewRepo()
	router := httpRepo.InitRoutes(repo)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
