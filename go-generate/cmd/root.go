package cmd

import (
	"fmt"
	"generate/tools"
	"os"
	"time"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// flag variables
var (
	start, end, outputFile, outputFolder string
	machines, readingInterval            int
	deleteDest                           bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
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

		generate(startTime, endTime, destination, machines, readingInterval, deleteDest)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&start, "start", "s", "2020-01-01", "Start date for the data")
	rootCmd.Flags().StringVarP(&end, "end", "e", "2020-02-01", "End date (exclusive) for the data")
	rootCmd.Flags().StringVarP(&outputFolder, "output-folder", "o", "./",
		"Folder to save the data in")
	rootCmd.Flags().StringVarP(&outputFile, "output-file", "f", "",
		"Full path of file to save the data in (output-folder is ignored and NOT prefixed)")
	rootCmd.Flags().IntVarP(&machines, "machines", "m", 5, "Number of machines")
	rootCmd.Flags().IntVarP(&readingInterval, "interval", "i", 5,
		"Interval of seconds between readings")
	rootCmd.Flags().BoolVarP(&deleteDest, "delete-out", "d", false,
		"Deletes the previous output file if it exists")
}
