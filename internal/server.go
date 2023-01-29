package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"net/http"
	"simple-start-page/internal/handlers"
	"simple-start-page/web"
)

type Server struct {
	cfg *Config
	App *fiber.App
}

func NewServer(cfg *Config, apiHandler handlers.ApiHandler) Server {
	app := fiber.New(fiber.Config{
		EnablePrintRoutes: cfg.Dev,
	})
	app.Use(recover.New(recover.Config{EnableStackTrace: cfg.Dev}))
	if cfg.Dev {
		app.Use(logger.New())
	}
	// routing
	api := app.Group("/api")
	api.Get("/links", apiHandler.GetListLink)
	api.Get("/setting", apiHandler.GetSetting)
	api.Post("/auth/login", apiHandler.Login)
	api.Put("/auth/login", apiHandler.CekAuth, apiHandler.UpdateAuth)
	api.Put("/setting", apiHandler.CekAuth, apiHandler.UpdateSetting)
	api.Put("/links", apiHandler.CekAuth, apiHandler.UpdateLinks)

	// static files
	if cfg.Dev {
		// serve for dist directory
		app.Use(filesystem.New(filesystem.Config{
			Root: http.Dir("./web/dist"),
		}))
	} else {
		// serve from embed
		app.Use(filesystem.New(filesystem.Config{
			Root:       http.FS(web.Dist),
			PathPrefix: "dist",
		}))
	}

	return Server{
		cfg: cfg,
		App: app,
	}
}

func (s *Server) Start() error {
	return s.App.Listen(s.cfg.ListenAddr)
}
