package app

import (
	"wow-ruby/internal/infrastructure"

	"github.com/go-chi/chi"
)

func RegisterAllRoutes(r chi.Router, app *AppInstance) {
	userRouter := infrastructure.NewUserRouter(app.Handlers.UserHandler)
	itemRouter := infrastructure.NewItemRouter(app.Handlers.ItemHandler)
	questRouter := infrastructure.NewQuestRouter(app.Handlers.QuestHandler)
	npcRouter := infrastructure.NewNpcRouter(app.Handlers.NpcHandler)

	r.Route("/api", func(r chi.Router) {
		userRouter.RegisterRoutes(r)
		itemRouter.RegisterRoutes(r)
		questRouter.RegisterRoutes(r)
		npcRouter.RegisterRoutes(r)
	})
}
