package cli

import (
    "fmt"
    route "github.com/itp-backend/backend-b-antar-jemput/routes"
    "os"
    "os/signal"

    "github.com/gofiber/fiber/v2"
    "github.com/itp-backend/backend-b-antar-jemput/app"
    "github.com/itp-backend/backend-b-antar-jemput/config"
    databaseconn "github.com/itp-backend/backend-b-antar-jemput/config/database"
    "github.com/itp-backend/backend-b-antar-jemput/controller"
    "github.com/itp-backend/backend-b-antar-jemput/repositories"
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
	appFiber := fiber.New(fiberConfig)

	// set up connection
	db := databaseconn.InitDb()
    //Repository
	userRepo := repositories.NewUserRepository(db)
	customerRepo := repositories.NewCustomerRepository(db)
    provinceRepo := repositories.NewProvinceRepository(db)
    regencyRepo := repositories.NewRegencyRepository(db)
    // Service
	customerService := service.NewCustomerService(customerRepo, userRepo, db)
    locationService := service.NewLocationService(provinceRepo, regencyRepo)
    loginService := service.NewLoginService(userRepo)
    // Controller
    customerController := controller.NewCustomerController(customerService)
    locationController := controller.NewLocationController(locationService)
    loginController := controller.NewLoginController(loginService)
    route.LoggerRoute(appFiber)
    appFiber.Post("/customer", customerController.RegisterCustomer)
	appFiber.Get("/location/provinces", locationController.GetAllProvinces)
	appFiber.Get("/location", locationController.GetAllRegenciesByProvinceId)
    appFiber.Post("/login", loginController.Login)
    route.NotFoundRoute(appFiber)
    StartServerWithGracefulShutdown(appFiber, application.Config.AppPort)

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
