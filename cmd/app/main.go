package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/guneyin/bookstore/api"
	"github.com/guneyin/bookstore/common"
	"github.com/guneyin/bookstore/config"
	"log"
	"time"
)

const (
	defaultReadTimeout  = 5 * time.Second
	defaultWriteTimeout = 5 * time.Second
)

type Application struct {
	Name       string
	Version    string
	Config     *config.Config
	HttpServer *fiber.App
	Api        *api.Api
}

func NewApplication(name string) (*Application, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	httpServer := fiber.New(fiber.Config{
		ServerHeader:      fmt.Sprintf("%s HTTP Server", name),
		AppName:           name,
		EnablePrintRoutes: true,
		ReadTimeout:       defaultReadTimeout,
		WriteTimeout:      defaultWriteTimeout,
	})
	httpServer.Use(recover.New())
	apiGroup := httpServer.Group("/api")

	return &Application{
		Name:       name,
		Version:    common.GetVersion(),
		Config:     cfg,
		HttpServer: httpServer,
		Api:        api.New(cfg, apiGroup),
	}, nil
}

func (app *Application) Run() error {
	return app.HttpServer.Listen(fmt.Sprintf(":%d", app.Config.Port))
}

func main() {
	app, err := NewApplication("The Online Book Store")
	if err != nil {
		log.Fatal(err)
	}

	common.SetLastRun(time.Now())

	log.Fatal(app.Run())
}
