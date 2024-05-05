package net

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
)

// subCOMMAND

// INFOCmd represents the info command
var INFOCmd = &cobra.Command{
	Use:   "info",
	Short: "Get All network details",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Get network interfaces
		interfaces, _ := net.Interfaces()
		for _, intf := range interfaces {
			addrs, _ := intf.Addrs()
			fmt.Printf("Interface Name: %s\n", intf.Name)
			fmt.Printf("  Hardware Address (MAC): %s\n", intf.HardwareAddr)
			fmt.Printf("  Flags: %v\n", intf.Flags)
			fmt.Printf("  MTU: %d\n", intf.MTU)
			fmt.Printf("  IP Addresses:\n")
			for _, addr := range addrs {
				fmt.Printf("    - %s\n", addr.String())
			}
			fmt.Println()
		}
	},
}

func init() {
	NetCmd.AddCommand(INFOCmd)
}
