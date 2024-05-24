package info

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/shirou/gopsutil/mem"
	"github.com/spf13/cobra"
)

// ramCmd represents the ram command
var ramCmd = &cobra.Command{
	Use:   "ram",
	Short: "Get the RAM usage",
	Long:  `ram command provides all details about your RAM including total RAM, used RAM, Free RAM, and RAM usage in %`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()
		// Start spinner to indicate processing
		color.Cyan("Processing your request...")
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		s.Start()
		// Get RAM usage
		memInfo, _ := mem.VirtualMemory()
		// Print memory details
		color.Cyan("RAM Usage:")
		color.Green(" - Total RAM: %.2f GB", float64(memInfo.Total)/1024/1024/1024)
		color.Green(" - Available RAM: %.2f GB", float64(memInfo.Available)/1024/1024/1024)
		color.Green(" - Used RAM: %.2f GB", float64(memInfo.Used)/1024/1024/1024)
		color.Green(" - Free RAM: %.2f GB", float64(memInfo.Free)/1024/1024/1024)
		color.Green(" - RAM Usage: %.2f%%", memInfo.UsedPercent)
		// Stop spinner and print completion message
		s.Stop()
		color.Green("DONE!")
		fmt.Println()
	},
}

func init() {
	InfoCmd.AddCommand(ramCmd)
}
