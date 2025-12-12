package http

import (
	"github.com/solluzumo/wow-ruby/gateway/internal/handlers"
	"github.com/solluzumo/wow-ruby/gateway/internal/middleware"

	"github.com/go-chi/chi"
)

type QuestRouter struct {
	questHandler *handlers.QuestHandler
}

func NewQuestRouter(handler *handlers.QuestHandler) *QuestRouter {
	return &QuestRouter{
		questHandler: handler,
	}
}

func (qr *QuestRouter) RegisterRoutes(router chi.Router) {
	router.Route("/quest", func(r chi.Router) {
		r.With(middleware.AuthMiddleware).Get("/{id}", qr.questHandler.GetById)
		r.With(middleware.AuthMiddleware).Post("/list", qr.questHandler.GetQuestList)
	})
}
