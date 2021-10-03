package app

import "github.com/itp-backend/backend-b-antar-jemput/config"

type Application struct {
	Config *config.Config
}

func Init() *Application {
	application := &Application{
		Config: config.Init(),
	}

	return application
}
