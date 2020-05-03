package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"timeseries/commands/generate"
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
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, _ := config.Build()
	zap.ReplaceGlobals(logger)

}
