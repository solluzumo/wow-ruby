package service

import (
	"github.com/solluzumo/wow-ruby/ruby-api/internal/repository"
)

type WeaponService struct {
	repo repository.WeaponRepository
	jwt  string
}

func NewWeaponService(repo repository.WeaponRepository, jwt string) *WeaponService {
	return &WeaponService{
		repo: repo,
		jwt:  jwt,
	}
}
