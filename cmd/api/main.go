package main

import (
	"fmt"
	"github.com/chazapp/dante/pkg/api"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
	"log"
	"os"
)

var (
	Commands = []cli.Command{
		{
			Name: "start",
			Usage: "Start the API at given port",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name: "port, p",
					Usage: "Port the API will listen to.",
					Value: 4242,
				},
				cli.StringFlag{
					Name: "db",
					Usage: "The MongoDB instance to be used.",
					Value: "mongodb://localhost:27017",
				},
				cli.StringFlag{
					Name: "env",
					Usage: "The environment file to be used.",
					Value: ".env",
				},
			},
			Action: func(c *cli.Context) error {
				_ = godotenv.Load()
				fmt.Println("DanteAPI starts at port ", c.Int("port"))
				err := api.StartRESTAPI(c.Int("port"), c.String("db"))
				return err
			},
		},
	}
)

func main() {
	app := cli.NewApp()
	app.Commands = Commands
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
