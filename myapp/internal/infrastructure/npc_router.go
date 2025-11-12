package infrastructure

import (
	"wow-ruby/internal/handlers"
	"wow-ruby/internal/middleware"

	"github.com/go-chi/chi"
)

type NpcRouter struct {
	npcHandler *handlers.NpcHandler
}

func NewNpcRouter(handler *handlers.NpcHandler) *NpcRouter {
	return &NpcRouter{
		npcHandler: handler,
	}
}

func (nr *NpcRouter) RegisterRoutes(router chi.Router) {
	router.Route("/npc", func(r chi.Router) {
		r.With(middleware.AuthMiddleware).Get("/{id}", nr.npcHandler.GetById)
		r.With(middleware.AuthMiddleware).Post("/list", nr.npcHandler.GetNpcList)
	})
}
