package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// flag variables
var (
	start, end string
	machines   int
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
		generate(start, end, machines)
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
	rootCmd.Flags().StringVarP(&end, "end", "e", "2021-01-01", "End date for the data")
	rootCmd.Flags().IntVarP(&machines, "machines", "m", 100, "Number of machines")
}
