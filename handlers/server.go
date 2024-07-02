package handlers

import (
	"log/slog"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rangzen/gotth-top/services"
	"github.com/rangzen/gotth-top/views"
	slogecho "github.com/samber/slog-echo"
)

type Server struct {
	log *slog.Logger
}

func NewServer() *Server {
	return &Server{
		log: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}
}

func (s *Server) Run() error {
	s.log.Info("starting GoTTH top")

	info, ok := debug.ReadBuildInfo()
	if ok {
		s.log.Info("build info",
			"go", info.GoVersion,
			"version", info.Main.Version,
			"path", info.Main.Path,
			"checksum", info.Main.Sum,
		)
	} else {
		s.log.Warn("build info not available")
	}

	app := echo.New()
	app.HideBanner = true
	app.HidePort = true

	// Middleware
	app.Use(slogecho.New(s.log))
	app.Use(middleware.Recover())

	// Handle static assets:
	assetHandler := http.FileServer(views.GetPublicAssetsFileSystem())
	app.GET("/public/*", echo.WrapHandler(http.StripPrefix("/public/", assetHandler)))

	// Routes
	app.GET("/", Index)

	app.GET("/home", Home)

	memService := services.NewMemService()
	memHandler := NewMemHandler(memService)
	app.GET("/mem", memHandler.ListMem)

	cpuService := services.NewCpuService()
	cpuHandler := NewCpuHandler(cpuService)
	app.GET("/cpu/percents", cpuHandler.ListCpuPercents)

	procService := services.NewProcService()
	procHandler := NewProcHandler(procService)
	app.GET("/proc", procHandler.ListProc)

	generalService := services.NewGeneralService()
	generalHandler := NewGeneralHandler(generalService)
	app.GET("/general", generalHandler.Base)

	// Start application
	s.log.Info("starting server on :8080")
	return app.Start(":8080")

}
