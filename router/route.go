package router

import (
	swagger "github.com/arsmn/fiber-swagger/v2"

	// "afeilulu.com/example/middleware"
	"afeilulu.com/example/handler"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {

	app.Get("/swagger/*", swagger.Handler) // default

	// Middleware of basic auth, uncomment if you need
	// api := app.Group("/api", middleware.AuthReq())
	// api.Get("/groups-list", handler.GetGroups)

	app.Post("/api/groups", handler.CreateGroup)
	app.Get("/api/groups-list", handler.GetGroups)
	app.Get("/api/groups-paged", handler.GetPagedGroups)
	app.Get("/api/groups/:id", handler.GetGroup)
	app.Post("/api/groups/:id", handler.UpdateGroup)
	app.Delete("/api/groups/:id", handler.DeleteGroup)

	app.Post("/api/users", handler.CreateUser)
	app.Get("/api/users", handler.GetUsers)

}
