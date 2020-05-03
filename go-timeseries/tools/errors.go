package tools

import (
	log "github.com/sirupsen/logrus"
)

// CheckError checks the error variable and prints a fatal log
// if it is set
func CheckError(message string, err error) {
	if err != nil {
		log.Fatalf("%s - Err: %v", message, err)
	}
}
