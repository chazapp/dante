package main

import (
	json2 "encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"tigbra/pkg"
)

var (
	Commands = []cli.Command{
		{
			Name:  "parse",
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
				docs, err := tigbra.ParseDataset(c.String("file"))
				if err != nil {
					return err
				} else {
					for _, item := range docs {
						json, _ := json2.Marshal(item)
						fmt.Println(string(json))
					}
				}
				return nil
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
