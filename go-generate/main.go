package main

import (
	"generate/cmd"
	"math/rand"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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
