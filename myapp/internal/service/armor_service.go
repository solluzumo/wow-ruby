package service

import (
	"wow-ruby/internal/repository"
)

type ArmorService struct {
	repo repository.ArmorRepository
	jwt  string
}

func NewArmorService(repo repository.ArmorRepository, jwt string) *ArmorService {
	return &ArmorService{
		repo: repo,
		jwt:  jwt,
	}
}
