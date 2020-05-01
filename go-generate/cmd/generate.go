package cmd

import (
	"fmt"

	"generate/tools"

	"go.uber.org/zap"
)

// METRICS is a map of the names of metrics and their IDs
var METRICS = map[string]int{
	"engine_temperature": 1,
	"oil_temperature":    2,
	"oil_pressure":       3,
	"running_hours":      4,
	"engine_load":        5,
}

// generate generates the data based on the arguments
// it is supposed to be called by the CLI
func generate(start, end string, machines int) {
	var logger = zap.S()

	var destination = "../out.csv"
	// var readingInterval = 5

	if tools.FileExists(destination) {
		logger.Fatalf("ERR: The destination file (%s) already exists \n", destination)
	}

	for idx := 0; idx < machines; idx++ {
		fmt.Printf("cmd %s for ship %d \n", destination, idx)

	}

}
