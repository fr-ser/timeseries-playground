package main

import (
	"math/rand"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"timeseries/commands/generate"
	"timeseries/tools"
)

// func main() {
// rand.Seed(time.Now().UnixNano())
// 	initLogger()
// 	cmd.Execute()
// }

func main() {
	rand.Seed(time.Now().UnixNano())
	initLogger()

	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Commands = []*cli.Command{
		generate.GenerateCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func initLogger() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	switch logLevel := tools.GetEnvDefault("LOG_LEVEL", "DEBUG"); logLevel {
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	case "WARN":
		log.SetLevel(log.WarnLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
}
