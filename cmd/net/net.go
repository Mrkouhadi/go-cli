package net

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NetCmd represents the net command
var NetCmd = &cobra.Command{
	Use:   "net",
	Short: "Net is a pallete that contains network based commands:netwrok detaisl and IP Adresses",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please specify what command do you want. ")
		fmt.Println(" - IP Addresses: go-cli net ip -v ipv4 / go-cli net ip -v ipv6")
		fmt.Println(" - Network Details: go-cli net info")
		fmt.Println(" - Ping a url: go-cli net ping -u google.com")
	},
}

func init() {}
