package info

import (
	"fmt"

	"github.com/shirou/gopsutil/mem"
	"github.com/spf13/cobra"
)

// subCOMMAND
// ramCmd represents the ram command
var ramCmd = &cobra.Command{
	Use:   "ram",
	Short: "Get the RAM usage",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Get RAM usage
		// Get memory info
		memInfo, _ := mem.VirtualMemory()
		// Print memory details
		fmt.Printf("Total RAM: %.2f GB\n", float64(memInfo.Total)/1024/1024/1024)
		fmt.Printf("Available RAM: %.2f GB\n", float64(memInfo.Available)/1024/1024/1024)
		fmt.Printf("Used RAM: %.2f GB\n", float64(memInfo.Used)/1024/1024/1024)
		fmt.Printf("Free RAM: %.2f GB\n", float64(memInfo.Free)/1024/1024/1024)
		fmt.Printf("RAM Usage: %.2f%%\n", memInfo.UsedPercent)
	},
}

func init() {
	// Here you will define your flags and configuration settings.
	InfoCmd.AddCommand(ramCmd)
}
