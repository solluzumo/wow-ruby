package http

import (
	"github.com/solluzumo/wow-ruby/gateway/internal/handlers"

	"github.com/go-chi/chi"
)

type UserRouter struct {
	userHandler *handlers.UserHandler
}

func NewUserRouter(handler *handlers.UserHandler) *UserRouter {
	return &UserRouter{
		userHandler: handler,
	}
}

func (ur *UserRouter) RegisterRoutes(router chi.Router) {
	router.Route("/user", func(r chi.Router) {
		r.Post("/details", ur.userHandler.GetByEmail)
		r.Post("/register", ur.userHandler.Register)
		r.Post("/login", ur.userHandler.Login)
	})
}
