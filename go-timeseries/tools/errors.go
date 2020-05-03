package tools

import "go.uber.org/zap"

// CheckError checks the error variable and prints a fatal log
// if it is set
func CheckError(l *zap.SugaredLogger, message string, err error) {
	if err != nil {
		l.Fatalf("%s - Err: %v", message, err)
	}
}
