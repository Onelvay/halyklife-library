package repository

import (
	"fmt"
	"github.com/Onelvay/halyklife-library/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func NewRepo() *Repo {
	cfg := getConfig()
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.HOST, cfg.PORT, cfg.USER, cfg.DBNAME, cfg.PASS)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err = db.AutoMigrate(&domain.Author{},
		&domain.Book{},
		&domain.Member{},
		&domain.Purchase{}); err != nil {
		panic(err)
	}
	return &Repo{db}
}
