package app

import "github.com/urfave/cli/v2"

// Generate a new app.
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line Interface"
	app.Usage = "A CLI for the Go programming language."

	return app
}
