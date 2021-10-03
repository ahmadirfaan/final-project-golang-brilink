package main

import (
	"os"

	"github.com/itp-backend/backend-b-antar-jemput/app"
)

func main() {
	c := cli.
		c.Run(app.Init())
}
