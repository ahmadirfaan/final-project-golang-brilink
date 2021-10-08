package cli

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/itp-backend/backend-b-antar-jemput/app"
	"github.com/itp-backend/backend-b-antar-jemput/config"
	databaseconn "github.com/itp-backend/backend-b-antar-jemput/config/database"
	"github.com/itp-backend/backend-b-antar-jemput/controller"
	"github.com/itp-backend/backend-b-antar-jemput/repositories"
	route "github.com/itp-backend/backend-b-antar-jemput/routes"
	"github.com/itp-backend/backend-b-antar-jemput/service"
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
	db := databaseconn.InitDb()
	userRepo := repositories.NewUserRepository(db)
	customerRepo := repositories.NewCustomerRepository(db)
	customerService := service.NewCustomerService(customerRepo, userRepo, db)
	customerController := controller.NewCustomerController(customerService)

	agentRepo := repositories.NewAgentRepository(db)
	agentService := service.NewAgentService(agentRepo, userRepo, db)
	agentController := controller.NewAgentController(agentService)

	//location controller
	provinceRepo := repositories.NewProvinceRepository(db)
    regencyRepo := repositories.NewRegencyRepository(db)
	locationService := service.NewLocationService(provinceRepo, regencyRepo)
	locationController := controller.NewLocationController(locationService)
	app.Post("/customer", customerController.RegisterCustomer)
	app.Get("/location/provinces", locationController.GetAllProvinces)
	app.Post("/agent", agentController.RegisterAgent)
	route.NotFoundRoute(app)
	log.Println(app.Server())
	app.Get("/location", locationController.GetAllRegenciesByProvinceId)
	route.NotFoundRoute(app)
    log.Println("Server jalan: ")
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
