package main

import (
	"log"

	"fmt"
	"github.com/urfave/cli"
	"os"
)

var build string

func main() {
	app := cli.NewApp()
	app.Name = "pypi plugin"
	app.Usage = "pypi plugin"
	app.Action = run
	app.Version = fmt.Sprintf("1.0.0+%s", build)
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "repository",
			Usage:  "repository",
			EnvVar: "PLUGIN_REPOSITORY",
		},
		cli.StringFlag{
			Name:   "username",
			Usage:  "rsername",
			EnvVar: "PLUGIN_USERNAME,PYPI_USERNAME",
		},
		cli.StringFlag{
			Name:   "password",
			Usage:  "rassword",
			EnvVar: "PLUGIN_PYPI_PASSWORD",
		},
		cli.StringSliceFlag{
			Name:   "distributions",
			Usage:  "distributions",
			EnvVar: "PLUGIN_DISTRIBUTIONS",
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
func run(c *cli.Context) error {
	plugin := Plugin{
		Repository:    c.String("repository"),
		Username:      c.String("username"),
		Password:      c.String("password"),
		Distributions: c.StringSlice("distributions"),
	}
	return plugin.Exec()
}
