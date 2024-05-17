package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/guneyin/bookstore/api"
	"github.com/guneyin/bookstore/api/middleware"
	"github.com/guneyin/bookstore/common"
	"github.com/guneyin/bookstore/config"
	"github.com/guneyin/bookstore/database"
	"github.com/guneyin/bookstore/util"
	"github.com/spf13/cobra"
	"log"
	"time"
)

const (
	defaultReadTimeout  = 30 * time.Second
	defaultWriteTimeout = 30 * time.Second
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
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return middleware.Error(ctx, err)
		},
	})

	httpServer.Use(recover.New())
	httpServer.Use(favicon.New())

	apiGroup := httpServer.Group("/api")

	return &Application{
		Name:       name,
		Version:    common.GetVersion().Version,
		Config:     cfg,
		HttpServer: httpServer,
		Api:        api.New(cfg, apiGroup),
	}, nil
}

func (app *Application) Run() error {
	common.SetLastRun(time.Now())

	return app.HttpServer.Listen(fmt.Sprintf(":%d", app.Config.Port))
}

var Cmd = &cobra.Command{
	Use: "run",
	Run: func(cmd *cobra.Command, args []string) {
		runApp()
	},
}

func runApp() {
	app, err := NewApplication("The Online Book Store")
	if err != nil {
		log.Fatal(err)
	}

	err = database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	err = util.MigrateDB(database.DB)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(app.Run())
}
