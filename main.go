package main

import (
	"github.com/seedgo/seedgo/action"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		HelpName:  "seedgo - command line tool for seedgo golang framework",
		UsageText: "seedgo create project [projectName]",
		Commands: []*cli.Command{
			{
				Name:  "create",
				Usage: "create project",
				Subcommands: []*cli.Command{
					{
						Name:   "project",
						Usage:  "create a project",
						Action: action.CreateProject,
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
