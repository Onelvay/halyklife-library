package main

import (
	"fmt"
	"github.com/Onelvay/halyklife-library/internal/repository"
)

func main() {
	repo := repository.NewRepo()
	//repo.CreateBook("zxczxc", domain.Book{"qweqweqwe", "qweasdzxc", "strong", "234-34123-435-123", "qweqwe", domain.Author{}})
	res, err := repo.GetAuthorBooks("qweqwe")
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
