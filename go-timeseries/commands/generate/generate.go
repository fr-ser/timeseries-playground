package generate

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"timeseries/commands"
	"timeseries/tools"
)

// pass through flag variables
var (
	machines, readingInterval int
	deleteDest                bool
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name: "start", Aliases: []string{"s"}, Value: "2020-01-01",
		Usage: "Start date for the data",
	},
	&cli.StringFlag{
		Name: "end", Aliases: []string{"e"}, Value: "2020-02-01",
		Usage: "End date for the data",
	},
	&cli.IntFlag{
		Name: "machines", Aliases: []string{"m"}, Value: 5, Destination: &machines,
		Usage: "Number of machines",
	},
	&cli.IntFlag{
		Name: "interval", Aliases: []string{"i"}, Value: 5, Destination: &readingInterval,
		Usage: "Interval of seconds between readings",
	},
	&cli.StringFlag{
		Name: "output-folder", Aliases: []string{"o"}, Value: "./",
		Usage: "Folder to save the data in",
	},
	&cli.StringFlag{
		Name: "output-file", Aliases: []string{"f"},
		Usage: "Full path of file to save the data in (output-folder is ignored and NOT prefixed)",
	},
	&cli.BoolFlag{
		Name: "delete-out", Value: false, Destination: &deleteDest,
		Usage: "Deletes the previous output file if it exists",
	},
}

// GenerateCommand is the command to create random data
var GenerateCommand = &cli.Command{
	Name:   "generate",
	Usage:  "creates random data for testing databases",
	Flags:  flags,
	Before: commands.NoArguments,
	Action: func(c *cli.Context) error {

		startTime, err := time.Parse(time.RFC3339, c.String("start")+"T00:00:00Z")
		tools.CheckError("Start time parsing failed", err)

		endTime, err := time.Parse(time.RFC3339, c.String("end")+"T00:00:00Z")
		tools.CheckError("End time parsing failed", err)

		var destination string

		if c.String("output-file") != "" {
			destination = c.String("output-file")
		} else {
			destination = fmt.Sprintf(
				"%s/%d_machines_%d_days.csv",
				c.String("output-folder"), machines, int(endTime.Sub(startTime).Hours()/24),
			)
		}

		generate(startTime, endTime, destination)
		return nil
	},
}

// generate is the entrypoint to generate the random machine data
// based on the arguments it is given
func generate(start, end time.Time, destination string) {
	file, csvWriter := setupOutfile(destination, deleteDest)
	defer file.Close()
	defer csvWriter.Flush()

	log.Info("Starting to generate data...")
	for machineID := 0; machineID < machines; machineID++ {
		readingTime := start
		for readingTime.Before(end) {
			createAndSaveReading(csvWriter, readingTime, machineID)

			readingTime = readingTime.Add(time.Second * time.Duration(readingInterval))
		}
		csvWriter.Flush()
		log.Debugf("Finished machine %d of %d", machineID+1, machines)
	}
	readingsPerMachine := int(end.Unix()-start.Unix()) / readingInterval
	numberOfMetrics := 5
	log.Infof("Finished. Wrote %d readings to the file", machines*readingsPerMachine*numberOfMetrics)

}

func setupOutfile(destination string, deleteDest bool) (*os.File, *csv.Writer) {
	if tools.FileExists(destination) {
		if deleteDest {
			err := os.Remove(destination)
			tools.CheckError("Cannot remove output file", err)
		} else {
			log.Fatalf("ERR: The destination file (%s) already exists \n", destination)
		}
	}

	file, err := os.Create(destination)
	tools.CheckError("Cannot create file", err)

	csvWriter := csv.NewWriter(file)

	// write headers
	err = csvWriter.Write([]string{"timestamp", "metric", "machine", "value"})
	tools.CheckError("Cannot write to file", err)

	return file, csvWriter
}
