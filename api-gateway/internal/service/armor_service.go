package service

import ()

type ArmorProvider interface {
}

type ArmorService struct {
	auth ArmorProvider
}

func NewArmorService() *ArmorService {
	return &ArmorService{}
}
