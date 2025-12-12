package service

import (
	"github.com/solluzumo/wow-ruby/ruby-api/internal/repository"
)

type ArmorService struct {
	repo repository.ArmorRepository
}

func NewArmorService(repo repository.ArmorRepository) *ArmorService {
	return &ArmorService{
		repo: repo,
	}
}
