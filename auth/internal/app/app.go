package app

import (
	"sync"

	"github.com/solluzumo/wow-ruby/auth/internal/domain"
)

type UserChannels struct {
	HashChannel *chan domain.HashTaskDomain
}

type App struct {
	WG              *sync.WaitGroup
	HashTaskChannel chan domain.HashTaskDomain
	UserChannels    *UserChannels
}

func NewAppInstance(wg *sync.WaitGroup, hsChan *chan domain.HashTaskDomain) *App {
	userChannels := &UserChannels{
		HashChannel: hsChan,
	}
	return &App{
		WG:              wg,
		HashTaskChannel: *hsChan,
		UserChannels:    userChannels,
	}
}
