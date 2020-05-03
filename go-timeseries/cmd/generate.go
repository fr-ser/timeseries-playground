package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"timeseries/tools"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var l *zap.SugaredLogger

// flag variables
var (
	start, end, outputFile, outputFolder string
	machines, readingInterval            int
	deleteDest                           bool
)

// generateCmd is the command to create random data
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate creates random data for testing databases",
	Long: `generate creates typical IOT data to be used for testing databases.

The generated data is mostly constant with a random deviation but tries to simulate existing
machine data. For example no usage during the night is simulated.
	`,
	Args: cobra.ExactArgs(0),
	Run: func(command *cobra.Command, args []string) {
		l = zap.S()

		startTime, err := time.Parse(time.RFC3339, start+"T00:00:00Z")
		tools.CheckError(l, "Start time parsing failed", err)

		endTime, err := time.Parse(time.RFC3339, end+"T00:00:00Z")
		tools.CheckError(l, "End time parsing failed", err)

		var destination string

		if outputFile != "" {
			destination = outputFile
		} else {
			destination = fmt.Sprintf(
				"%s/%d_machines_%d_days.csv",
				outputFolder, machines, int(endTime.Sub(startTime).Hours()/24),
			)
		}

		generate(startTime, endTime, destination)
	},
}

func initGenerate() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(&start, "start", "s", "2020-01-01", "Start date for the data")
	generateCmd.Flags().StringVarP(&end, "end", "e", "2020-02-01", "End date (exclusive) for the data")
	generateCmd.Flags().StringVarP(&outputFolder, "output-folder", "o", "./",
		"Folder to save the data in")
	generateCmd.Flags().StringVarP(&outputFile, "output-file", "f", "",
		"Full path of file to save the data in (output-folder is ignored and NOT prefixed)")
	generateCmd.Flags().IntVarP(&machines, "machines", "m", 5, "Number of machines")
	generateCmd.Flags().IntVarP(&readingInterval, "interval", "i", 5,
		"Interval of seconds between readings")
	generateCmd.Flags().BoolVarP(&deleteDest, "delete-out", "d", false,
		"Deletes the previous output file if it exists")
}

// generate is the entrypoint to generate the random machine data
// based on the arguments it is given
func generate(start, end time.Time, destination string) {
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
