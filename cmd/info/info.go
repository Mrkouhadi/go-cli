package info

import (
	"fmt"
	"github.com/spf13/cobra"
)
// COMMAND
// infoCmd represents the info command
var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "info is a pallete that contains commands to get info of CPU, RAM, and DISK usage",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please specify what command do you want. ")
		fmt.Println(" - CPU usage: go-cli info cpu")
		fmt.Println(" - RAM usage: go-cli info ram")
		fmt.Println(" - DISK usage: go-cli info disk")
	},
}
func init() {}
