package app

import (
	"github.com/solluzumo/wow-ruby/gateway/internal/infrastructure/http"

	"github.com/go-chi/chi"
)

func RegisterAllRoutes(r chi.Router, app *AppInstance) {
	itemRouter := http.NewItemRouter(app.Handlers.ItemHandler)
	questRouter := http.NewQuestRouter(app.Handlers.QuestHandler)
	npcRouter := http.NewNpcRouter(app.Handlers.NpcHandler)
	userRouter := http.NewUserRouter(app.Handlers.UserHandler)

	r.Route("/api", func(r chi.Router) {
		itemRouter.RegisterRoutes(r)
		questRouter.RegisterRoutes(r)
		npcRouter.RegisterRoutes(r)
		userRouter.RegisterRoutes(r)
	})
}
