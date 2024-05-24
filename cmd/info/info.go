package info

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// COMMAND
// infoCmd represents the info command
var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "info is a pallete that contains commands to get info of CPU, RAM, and DISK usage",
	Long:  `info command provides a collection of subcommands to retrieve info of CPU, RAM, and DISK usage`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()
		color.Cyan("Please specify which command you want to run.")
		fmt.Println()
		color.Yellow("Available Commands:")
		fmt.Println(" - CPU usage:")
		color.Green("   go-cli info cpu")
		fmt.Println(" - RAM usage:")
		color.Green("   go-cli info ram")
		fmt.Println(" - Disk usage:")
		color.Green("   go-cli info disk")
		fmt.Println(" - GPU usage:")
		color.Green("   go-cli info gpu")
		fmt.Println()
	},
}

func init() {}
