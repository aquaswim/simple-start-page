package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"net/http"
	"simple-start-page/web"
)

type Server struct {
	cfg *Config
	App *fiber.App
}

func NewServer(cfg *Config) Server {
	app := fiber.New()
	app.Use(recover.New(recover.Config{EnableStackTrace: cfg.Dev}))
	// routing

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
