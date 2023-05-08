package http

import (
	"github.com/Onelvay/halyklife-library/internal/repository"
)

type Handler struct {
	repo *repository.Repo
}

func newHandler(repo *repository.Repo) *Handler {
	return &Handler{repo}
}
