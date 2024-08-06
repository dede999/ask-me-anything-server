package main

import (
	"github.com/dede999/ask-me-anything-server/cmd/tools/terndotenv/workers"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
	"os"
)

func main() {
	defer workers.ErrorRecovery()

	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	app := cli.NewApp()
	app.Name = "Migration Workers"
	app.Usage = "Run migration workers"
	app.Commands = []cli.Command{
		{
			Name:     "migrate",
			HelpName: "migrate",
			Action:   workers.Migration,
			Usage:    "Migrate current database",
			Description: `If a version is not set, it will migrate to the latest version.
						If a negative version is set, it will rollback the database.`,
			Flags: []cli.Flag{
				&cli.IntFlag{
					Name:  "version",
					Usage: "Version to migrate. If not set, it will migrate to the latest version",
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
