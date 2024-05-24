package info

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/spf13/cobra"
)

// subCOMMAND
// cpuCmd represents the cpu command
var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "Get the CPU usage",
	Long:  `cpu command provides details about the CPU, including model, speed, and core count.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()

		// Start spinner to indicate processing
		color.Cyan("Processing your request...")
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		s.Start()

		// Get CPU information
		cpuInfo, err := cpu.Info()
		if err != nil {
			s.Stop()
			color.Red("Error retrieving CPU information: %v", err)
			return
		}

		// Print CPU details
		for i, info := range cpuInfo {
			color.Cyan("CPU #%d:", i+1)
			color.Green(" - Model: %s", info.ModelName)
			color.Green(" - CPU MHz: %.2f", info.Mhz)
			color.Green(" - CPU Family: %s", info.Family)

			// Get core counts
			physicalCores, err := cpu.Counts(false)
			if err != nil {
				color.Red("Error retrieving physical core count: %v", err)
			} else {
				color.Green(" - Physical Cores: %d", physicalCores)
			}

			logicalCores, err := cpu.Counts(true)
			if err != nil {
				color.Red("Error retrieving logical core count: %v", err)
			} else {
				color.Green(" - Logical Cores: %d", logicalCores)
			}
		}

		// Stop spinner and print completion message
		s.Stop()
		fmt.Println()
		color.Green("DONE!")

	},
}

func init() {
	// Here you will define your flags and configuration settings.
	InfoCmd.AddCommand(cpuCmd)
}
