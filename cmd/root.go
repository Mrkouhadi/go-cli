package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/mrkouhadi/go-cli/cmd/info"
	"github.com/mrkouhadi/go-cli/cmd/net"
	"github.com/spf13/cobra"
)

// Define the version of the application
var version = "1.0.0"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-cli",
	Short: "Display the Usage of CPU, RAM, DISK, and NETWORK.",
	Long:  `go-cli is a command line interface to display the usage of CPU, RAM, DISK, and NETWORK.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Display usage information if no command is specified
		cmd.Usage()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		color.Red("Error: %v", err)
		os.Exit(1)
	}
}

func AddSubCommandsPalletes() {
	rootCmd.AddCommand(net.NetCmd)
	rootCmd.AddCommand(info.InfoCmd)
}

func init() {
	// getting the version
	rootCmd.PersistentFlags().BoolP("version", "v", false, "Print the version number of go-cli")
	cobra.OnInitialize(func() {
		// Check if the version flag is set
		versionFlag, err := rootCmd.Flags().GetBool("version")
		if err != nil {
			color.Red("Error: %v", err)
			os.Exit(1)
		}
		if versionFlag {
			color.Green("go-cli version: %s", version)
			os.Exit(0)
		}
	})
	// adding subcommands
	AddSubCommandsPalletes()
}
