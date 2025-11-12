package service

import (
	"wow-ruby/internal/repository"
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
