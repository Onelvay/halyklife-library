package main

import (
	"fmt"
	"github.com/Onelvay/halyklife-library/internal/repository"
)

func main() {
	repo := repository.NewRepo()
	fmt.Println(repo)
}
