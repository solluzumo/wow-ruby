package http

import (
	"github.com/solluzumo/wow-ruby/gateway/internal/handlers"
	"github.com/solluzumo/wow-ruby/gateway/internal/middleware"

	"github.com/go-chi/chi"
)

type ItemRouter struct {
	itemHandler *handlers.ItemHandler
}

func NewItemRouter(handler *handlers.ItemHandler) *ItemRouter {
	return &ItemRouter{
		itemHandler: handler,
	}
}

func (ir *ItemRouter) RegisterRoutes(router chi.Router) {
	router.Route("/item", func(r chi.Router) {
		r.With(middleware.AuthMiddleware).Get("/{id}", ir.itemHandler.GetById)
		r.With(middleware.AuthMiddleware).Post("/list", ir.itemHandler.GetItemList)
	})
}
