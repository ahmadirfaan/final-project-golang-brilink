package main

import (
	"os"

	"github.com/itp-backend/backend-b-antar-jemput/app"
	"github.com/itp-backend/backend-b-antar-jemput/cli"
)


func main() {
	c := cli.NewCli(os.Args)
	c.Run(app.Init())

}
