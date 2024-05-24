package net

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var netversion string

// IPCmd represents the ip command
var IPCmd = &cobra.Command{
	Use:   "ip",
	Short: "Get IP Address",
	Long:  `ip command retrieves and displays the IP addresses of the host, filtered by the specified network version (IPv4 or IPv6).`,
	Run: func(cmd *cobra.Command, args []string) {
		// Start spinner
		color.Cyan("Processing your request...")
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		s.Start()

		// Get hostname
		host, err := os.Hostname()
		if err != nil {
			s.Stop()
			color.Red("Error getting hostname: %v", err)
			return
		}
		fmt.Println("Host:", host)

		// Get local interface addresses
		addrs, err := net.InterfaceAddrs()
		if err != nil {
			s.Stop()
			color.Red("Error getting network interfaces: %v", err)
			return
		}

		// Display IP addresses based on the specified network version
		found := false
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if netversion == "ipv4" && ipnet.IP.To4() != nil {
					color.Green("IPv4: %s", ipnet.IP.String())
					found = true
				} else if netversion == "ipv6" && ipnet.IP.To4() == nil {
					color.Green("IPv6: %s", ipnet.IP.String())
					found = true
				}
			}
		}

		if !found {
			if netversion == "ipv4" {
				color.Red("No IPv4 address found")
			} else if netversion == "ipv6" {
				color.Red("No IPv6 address found")
			} else {
				color.Red("Invalid network version specified. Use 'ipv4' or 'ipv6'.")
			}
		}

		// Stop spinner
		s.Stop()
		color.Green("DONE!")
	},
}

func init() {
	IPCmd.Flags().StringVarP(&netversion, "version", "v", "", "Version of network: 'ipv4' or 'ipv6'")
	NetCmd.AddCommand(IPCmd)
}
