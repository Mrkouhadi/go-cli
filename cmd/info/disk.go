package info

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/spf13/cobra"
)

// subCOMMAND
// diskCmd represents the disk command
var diskCmd = &cobra.Command{
	Use:   "disk",
	Short: "Get the disk usage",
	Long:  `disk command provides details about disk usage including total, used, and free space.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()

		// Start spinner to indicate processing
		color.Cyan("Processing your request...")
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		s.Start()

		// Get disk usage
		usage, err := disk.Usage("/")
		if err != nil {
			s.Stop()
			color.Red("Error retrieving disk usage: %v", err)
			return
		}

		// Print disk usage
		color.Cyan("Disk Usage Information:")
		color.Green(" - Total Disk Space: %.2f GB", float64(usage.Total)/1024/1024/1024)
		color.Green(" - Used Disk Space: %.2f GB", float64(usage.Used)/1024/1024/1024)
		color.Green(" - Free Disk Space: %.2f GB", float64(usage.Free)/1024/1024/1024)

		// Stop spinner and print completion message
		s.Stop()
		color.Green("DONE!")
		fmt.Println()

	},
}

func init() {
	// Here you will define your flags and configuration settings.
	InfoCmd.AddCommand(diskCmd)
}
