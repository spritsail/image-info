package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spritsail/image-info/badge"
	mb "github.com/spritsail/image-info/microbadger"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "image-info"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "bind",
			Value:  ":8080",
			Usage:  "host and port to bind to",
			EnvVar: "HOST_PORT",
		},
		cli.StringFlag{
			Name:   "baseURL, base",
			Value:  "/",
			Usage:  "base URL to serve API",
			EnvVar: "BASE_URL",
		},
	}
	app.Action = func(c *cli.Context) (err error) {
		router := gin.Default()

		// Define routes
		base := router.Group(c.String("baseURL"))
		badge.BuildRoutes(base.Group("/badge"), c)

		mb.Init(nil)

		return router.Run(c.String("bind"))
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
