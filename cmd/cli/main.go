package main

import (
	json2 "encoding/json"
	"fmt"
	"github.com/chazapp/dante/pkg"
	"github.com/urfave/cli"
	"log"
	"os"
)

var (
	Commands = []cli.Command{
		{
			Name:  "print",
			Usage: "Parses and output json format to stdin",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file, f",
					Usage: "Path to the input file.",
				},
				cli.StringFlag{
					Name:  "output, o",
					Usage: "Path to the output file",
					Value: "",
				},
			},
			Action: func(c *cli.Context) error {
				docs, err := dante.ParseDataset(c.String("file"))
				if err != nil {
					return err
				}
				for _, item := range docs {
					json, _ := json2.Marshal(item)
					fmt.Println(string(json))
				}
				return nil
			},
		},
		{
			Name:  "mongo",
			Usage: "Parses  and stores json format to given MongoDB",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "db",
					Usage: "The target Mongodb URI.",
				},
				cli.StringFlag{
					Name:  "file, f",
					Usage: "Path to the input file.",
				},
				cli.StringFlag{
					Name:  "name, n",
					Usage: "Name of the mongodb collection.",
				},
			},
			Action: func(c *cli.Context) error {
				docs, err := dante.ParseDataset(c.String("file"))
				if err != nil {
					return err
				}
				err = dante.MongoDB(docs, c.String("db"), c.String("name"))
				return err
			},
		},
		{
			Name:  "elastic",
			Usage: "Parses and stores json format to given ElasticSearch instance",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "db",
					Usage: "The target ElasticSearch URI.",
				},
				cli.StringFlag{
					Name:  "file, f",
					Usage: "Path to the input file.",
				},
				cli.StringFlag{
					Name:  "name, n",
					Usage: "Name of the index",
				},
			},
			Action: func(c *cli.Context) error {
				docs, err := dante.ParseDataset(c.String("file"))
				if err != nil {
					return err
				}
				err = dante.ElasticSearch(docs, c.String("db"), c.String("name"))
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
