package net

import (
	"fmt"
	"net/http"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
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
		Timeout: 10 * time.Second,
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
	Short: "ping pings a remote URL",
	Long:  `ping command pings the specified URL and returns the HTTP status code along with a message.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()

		// Start spinner
		color.Cyan("Processing your request...")
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		s.Start()

		// Perform the ping
		statusCode, err := ping(urlPath)

		// Stop spinner
		s.Stop()

		// Handle the response
		if err != nil {
			color.Red("Error: %v", err)
		} else {
			switch statusCode {
			case http.StatusOK:
				color.Green("The URL is reachable! Status Code: 200 OK")
			case http.StatusNotFound:
				color.Yellow("The URL was not found. Status Code: 404 Not Found")
			case http.StatusInternalServerError:
				color.Red("The server encountered an error. Status Code: 500 Internal Server Error")
			default:
				color.Blue("Received HTTP Status Code: %d", statusCode)
			}
		}
		fmt.Println()

		color.Green("DONE!")
	},
}

func init() {
	pingCmd.Flags().StringVarP(&urlPath, "url", "u", "", "The URL to ping")
	// Ensure the url flag is required
	if err := pingCmd.MarkFlagRequired("url"); err != nil {
		fmt.Println("please provide a url")
	}
	NetCmd.AddCommand(pingCmd)
}
