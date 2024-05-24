package net

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// NetCmd represents the net command
var NetCmd = &cobra.Command{
	Use:   "net",
	Short: "net is a palette that contains network-based commands: network details and IP addresses",
	Long:  `net command provides a collection of subcommands to retrieve network details, IP addresses, and perform network operations such as pinging a URL.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()
		color.Cyan("Please specify what command you want to run.")
		fmt.Println()
		color.Yellow("Available Commands:")
		fmt.Println(" - IP Addresses:")
		color.Green("   go-cli net ip -v ipv4")
		color.Green("   go-cli net ip -v ipv6")
		fmt.Println(" - Network Details:")
		color.Green("   go-cli net info")
		fmt.Println(" - Ping a URL:")
		color.Green("   go-cli net ping -u brifel.com")
		fmt.Println()
	},
}

func init() {}
