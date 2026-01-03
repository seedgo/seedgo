package main

import (
	"context"
	"log"
	"os"

	"github.com/seedgo/seedgo/action"
	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		EnableShellCompletion: true,
		Name:                  "seedgo",
		Version:               "v0.0.5",
		Usage:                 "command line tool for seedgo golang framework",
		UsageText:             "seedgo create project [projectName]",
		Commands: []*cli.Command{
			{
				Name:  "create",
				Usage: "create project",
				Commands: []*cli.Command{
					{
						Name:   "project",
						Usage:  "create a project",
						Action: action.CreateProject,
					},
				},
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
