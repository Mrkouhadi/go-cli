package net

import (
	"fmt"
	"net"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// INFOCmd represents the info command
var INFOCmd = &cobra.Command{
	Use:   "info",
	Short: "Get all network details",
	Long:  `info command retrieves and displays detailed information about all network interfaces on the system.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()
		// Start spinner
		color.Cyan("Processing your request...")
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		s.Start()

		// Get network interfaces
		interfaces, err := net.Interfaces()
		if err != nil {
			s.Stop()
			color.Red("Error fetching network interfaces: %v", err)
			return
		}

		// Display network interfaces
		for _, intf := range interfaces {
			color.Yellow("Interface Name: %s", intf.Name)
			fmt.Printf("  Hardware Address (MAC): %s\n", intf.HardwareAddr)
			fmt.Printf("  Flags: %v\n", intf.Flags)
			fmt.Printf("  MTU: %d\n", intf.MTU)
			fmt.Printf("  IP Addresses:\n")

			addrs, err := intf.Addrs()
			if err != nil {
				color.Red("  Error fetching addresses for interface %s: %v", intf.Name, err)
				continue
			}

			for _, addr := range addrs {
				fmt.Printf("    - %s\n", addr.String())
			}
			fmt.Println()
		}

		// Stop spinner
		s.Stop()
		fmt.Println()
		color.Green("DONE!")
	},
}

func init() {
	NetCmd.AddCommand(INFOCmd)
}
