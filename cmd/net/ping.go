package net

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

// subCOMMAND

var (
	urlPath string
)

// helper function
func ping(domain string) (int, error) {
	url := "http://" + domain
	client := http.Client{
		Timeout: 2 * time.Second,
	}
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return 0, err
	}
	res, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	res.Body.Close()
	return res.StatusCode, nil
}

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Pings a remote URL",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		//  do the logic here
		if resp, err := ping(urlPath); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp)
		}
	},
}

func init() {
	pingCmd.Flags().StringVarP(&urlPath, "url", "u", "", "The url to ping")
	// if we run `net ping`  we will get this error
	if err := pingCmd.MarkFlagRequired("url"); err != nil {
		// fmt.Println(err)
		fmt.Println("please provide a url")
	}
	NetCmd.AddCommand(pingCmd)
}
