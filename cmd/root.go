package cmd

import (
	"fmt"
	"os"

	"github.com/mrkouhadi/go-cli/cmd/info"
	"github.com/mrkouhadi/go-cli/cmd/net"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-cli",
	Short: "Display the Usage of CPU, RAM, DISK, and NETWORK.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		toggle, _ := cmd.Flags().GetBool("toggle")
		if toggle {
			fmt.Println("Toggling...")
		} else {
			fmt.Println("Toggle is off")
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func AddSubCommandsPalletes() {
	rootCmd.AddCommand(net.NetCmd)
	rootCmd.AddCommand(info.InfoCmd)
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	_ = rootCmd.PersistentFlags().Parse(os.Args[1:]) // Parse flags before adding subcommands

	AddSubCommandsPalletes()
}
