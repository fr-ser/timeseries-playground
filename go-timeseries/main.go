package main

import (
	"math/rand"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"timeseries/cmd"
)

func main() {
	initLogger()
	cmd.Execute()
}

func initLogger() {
	rand.Seed(time.Now().UnixNano())

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, _ := config.Build()
	zap.ReplaceGlobals(logger)

}
