package cmd

import (
	"encoding/csv"
	"os"
	"time"

	"generate/tools"

	"go.uber.org/zap"
)

var l *zap.SugaredLogger

// generate is the entrypoint to generate the random machine data
// based on the arguments it is given
func generate(
	start, end time.Time,
	destination string,
	machines, readingInterval int,
	deleteDest bool,
) {
	l = zap.S()

	file, csvWriter := setupOutfile(destination, deleteDest)
	defer file.Close()
	defer csvWriter.Flush()

	l.Info("Starting to generate data...")
	for machineID := 0; machineID < machines; machineID++ {
		readingTime := start
		for readingTime.Before(end) {
			createAndSaveReading(csvWriter, readingTime, machineID)

			readingTime = readingTime.Add(time.Second * time.Duration(readingInterval))
		}
		csvWriter.Flush()
		l.Debugf("Finished machine %d of %d", machineID+1, machines)
	}
	readingsPerMachine := int(end.Unix()-start.Unix()) / readingInterval
	numberOfMetrics := 5
	l.Infof("Finished. Wrote %d readings to the file", machines*readingsPerMachine*numberOfMetrics)

}

func setupOutfile(destination string, deleteDest bool) (*os.File, *csv.Writer) {
	if tools.FileExists(destination) {
		if deleteDest {
			err := os.Remove(destination)
			tools.CheckError(l, "Cannot remove output file", err)
		} else {
			l.Fatalf("ERR: The destination file (%s) already exists \n", destination)
		}
	}

	file, err := os.Create(destination)
	tools.CheckError(l, "Cannot create file", err)

	csvWriter := csv.NewWriter(file)

	// write headers
	err = csvWriter.Write([]string{"timestamp", "metric", "machine", "value"})
	tools.CheckError(l, "Cannot write to file", err)

	return file, csvWriter
}
