package app

import "github.com/urfave/cli/v2"

// Generate a new app.
func Generate() *cli.App {
	return cli.NewApp()
}
