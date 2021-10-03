package cli

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/itp-backend/backend-b-antar-jemput/app"
	"github.com/itp-backend/backend-b-antar-jemput/config"
	databaseconn "github.com/itp-backend/backend-b-antar-jemput/config/database"
	route "github.com/itp-backend/backend-b-antar-jemput/routes"
	log "github.com/sirupsen/logrus"
)

type Cli struct {
	Args []string
}

func NewCli(args []string) *Cli {
	return &Cli{
		Args: args,
	}
}

func (cli *Cli) Run(application *app.Application) {
	fiberConfig := config.FiberConfig()
	app := fiber.New(fiberConfig)

	// set up connection
	databaseconn.InitDb()

	route.NotFoundRoute(app)

	StartServerWithGracefulShutdown(app, application.Config.AppPort)

}

func StartServerWithGracefulShutdown(app *fiber.App, port string) {
	appPort := fmt.Sprintf(`:%s`, port)
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := app.Shutdown(); err != nil {
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	// Run server.
	if err := app.Listen(appPort); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}
