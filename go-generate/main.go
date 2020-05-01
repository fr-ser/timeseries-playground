package main

import (
	"generate/cmd"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	initLogger()
	cmd.Execute()
}

func initLogger() {

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, _ := config.Build()
	zap.ReplaceGlobals(logger)

}
