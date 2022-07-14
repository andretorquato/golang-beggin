package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli/v2"
)

// Generate a new app.
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line Interface"
	app.Usage = "A CLI for the Go programming language."

	app.Commands = []*cli.Command{
		{
			Name:  "ip",
			Usage: "Get the IP address from the internet",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "host",
					Value: "google.com.br",
				},
			},
			Action: getIps,
		},
	}
	return app
}

func getIps(c *cli.Context) error {
	host := c.String("host")
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}

	return nil
}
