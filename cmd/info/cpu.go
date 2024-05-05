package info

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/spf13/cobra"
)

// subCOMMAND
// cpuCmd represents the cpu command
var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "Get the CPU usage",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Get CPU usage
		cpuInfo, _ := cpu.Info()
		// Print CPU details
		for i, info := range cpuInfo {
			// model
			fmt.Printf("CPU #%d:\n", i+1)
			fmt.Printf(" - Model: %s\n", info.ModelName)
			fmt.Printf(" - CPU MHz: %.2f\n", info.Mhz)
			fmt.Printf(" - CPU Family: %s\n", info.Family)
			// cores
			physicalCores, _ := cpu.Counts(false)
			fmt.Printf(" - Physical Cores: %d\n", physicalCores)
			logicalCores, _ := cpu.Counts(true)
			fmt.Printf(" - Logical Cores: %d\n", logicalCores)
		}
	},
}

func init() {
	// Here you will define your flags and configuration settings.
	InfoCmd.AddCommand(cpuCmd)
}
