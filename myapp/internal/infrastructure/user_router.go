package infrastructure

import (
	"wow-ruby/internal/handlers"

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
		r.Post("/details", ur.userHandler.GetById)
		r.Post("/register", ur.userHandler.CreateUser)
		r.Post("/login", ur.userHandler.LoginUser)
	})
}
