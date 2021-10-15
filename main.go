package main

import (
    "github.com/itp-backend/backend-b-antar-jemput/app"
    "github.com/itp-backend/backend-b-antar-jemput/cli"
    "os"
)

func main() {
	c := cli.NewCli(os.Args)
	c.Run(app.Init())
}
