package info

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/disk"
	"github.com/spf13/cobra"
)
// subCOMMAND
// diskCmd represents the disk command
var diskCmd = &cobra.Command{
	Use:   "disk",
	Short: "Get the disk usage",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Get disk usage
		usage, _ := disk.Usage("/")
		// Print disk usage
		fmt.Printf("Total Disk Space: %.2f GB\n", float64(usage.Total)/1024/1024/1024)
		fmt.Printf("Used Disk Space: %.2f GB\n", float64(usage.Used)/1024/1024/1024)
		fmt.Printf("Free Disk Space: %.2f GB\n", float64(usage.Free)/1024/1024/1024)
	},
}
func init() {
	// Here you will define your flags and configuration settings.
	InfoCmd.AddCommand(diskCmd)
}