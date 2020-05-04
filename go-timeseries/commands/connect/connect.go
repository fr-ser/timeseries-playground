package connect

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"timeseries/commands"
)

const (
	host     = "0.0.0.0"
	user     = "db_admin"
	pass     = "pass"
	database = "zlab"
)

// ConnectCommand is the command to connect do the databases
var ConnectCommand = &cli.Command{
	Name:  "connect",
	Usage: "opens a shell with a connection to the database",
	Description: `Returns a command to be executed. Use like this:
>
> $(timeseries connect the_command)`,
	Subcommands: []*cli.Command{
		{
			Name:   "postgres",
			Usage:  "connect to the plain postgres",
			Before: commands.NoArguments,
			Action: func(c *cli.Context) error {
				port := "35434"
				fmt.Printf("psql -x postgres://%s:%s@%s:%s/%s", user, pass, host, port, database)
				return nil
			},
		},
		{
			Name:   "timescale",
			Usage:  "connect to the timescale db",
			Before: commands.NoArguments,
			Action: func(c *cli.Context) error {
				port := "35433"
				fmt.Printf("psql -x postgres://%s:%s@%s:%s/%s", user, pass, host, port, database)
				return nil
			},
		},
		{
			Name:   "influx",
			Usage:  "connect to the influx db through docker compose",
			Before: commands.NoArguments,
			Action: func(c *cli.Context) error {
				fmt.Print("docker-compose exec influx influx -precision rfc3339")
				return nil
			},
		},
	},
}
