package handlers

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/rangzen/gotth-top/services"
	"github.com/rangzen/gotth-top/views"
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
	s.log.Info("starting top")

	app := echo.New()
	app.HideBanner = true


	// Handle static assets:
	assetHandler := http.FileServer(views.GetPublicAssetsFileSystem())
	app.GET("/public/*", echo.WrapHandler(http.StripPrefix("/public/", assetHandler)))

	app.GET("/", Index)

	app.GET("/home", Home)

	memService := services.NewMemService()
	memHandler := NewMemHandler(memService)
	app.GET("/mem", memHandler.ListMem)

	cpuService := services.NewCpuService()
	cpuHandler := NewCpuHandler(cpuService)
	app.GET("/cpu/percents", cpuHandler.ListCpuPercents)

	generalService := services.NewGeneralService()
	generalHandler := NewGeneralHandler(generalService)
	app.GET("/general", generalHandler.Base)

	return app.Start(":8080")

}
