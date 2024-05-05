package net

import (
	"fmt"
	"net"
	"os"

	"github.com/spf13/cobra"
)

var netversion string

// subCOMMAND
// IPCmd represents the ip command
var IPCmd = &cobra.Command{
	Use:   "ip",
	Short: "Get IP Address",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := os.Hostname()
		fmt.Println("Host:", host)
		// Get local interface addresses
		addrs, _ := net.InterfaceAddrs()
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if netversion == "ipv4" && ipnet.IP.To4() != nil {
					fmt.Println("IPv4: ", ipnet.IP.String())
				} else if netversion != "ipv4" {
					fmt.Println("IPv6: ", ipnet.IP.String())
				}
			}
		}
	},
}

func init() {
	IPCmd.Flags().StringVarP(&netversion, "version", "v", "", "version of network: IPv4 or IPv6") // --url or -u
	NetCmd.AddCommand(IPCmd)
}
