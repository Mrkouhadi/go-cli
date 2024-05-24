package info

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// gpuCmd represents the gpu command
var gpuCmd = &cobra.Command{
	Use:   "gpu",
	Short: "Get the GPU details",
	Long:  `gpu command provides details about the GPU, including model, memory usage, and other relevant information.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()
		// Start spinner to indicate processing
		color.Cyan("Processing your request...")
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		s.Start()
		// Check what OS is running the CLI
		switch os := runtime.GOOS; os {
		case "darwin":
			getGPUInfoDarwin()
		case "windows", "linux":
			getGPUInfoNvidia()
		default:
			s.Stop()
			color.Red("Your Operating System -- %s -- is not supported!", os)
			return
		}
		// Stop spinner and print completion message
		s.Stop()
		color.Green("DONE!")
		fmt.Println()
	},
}

func init() {
	InfoCmd.AddCommand(gpuCmd)
}

// getGPUInfoDarwin retrieves GPU information on Darwin systems
func getGPUInfoDarwin() {
	cmd := exec.Command("system_profiler", "SPDisplaysDataType")

	outputChan := make(chan string)
	errChan := make(chan error)

	go func() {
		output, err := cmd.Output()
		if err != nil {
			errChan <- err
			return
		}
		outputChan <- string(output)
	}()

	select {
	case output := <-outputChan:
		gpuInfo := parseSystemProfilerOutput(output)
		color.Blue("GPU Information:")
		for key, value := range gpuInfo {
			color.Green("%s: %s", key, value)
		}
	case err := <-errChan:
		color.Red("Error: %v", err)
	}
}

// parseSystemProfilerOutput parses the output of the system_profiler command on Darwin systems
func parseSystemProfilerOutput(output string) map[string]string {
	gpuInfo := make(map[string]string)
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			gpuInfo[key] = value
		}
	}
	return gpuInfo
}

// getGPUInfoNvidia retrieves GPU information on systems using Nvidia GPUs
func getGPUInfoNvidia() {
	cmd := exec.Command("nvidia-smi")

	outputChan := make(chan string)
	errChan := make(chan error)

	go func() {
		output, err := cmd.Output()
		if err != nil {
			errChan <- err
			return
		}
		outputChan <- string(output)
	}()

	select {
	case output := <-outputChan:
		gpuInfo := parseNvidiaSMIOutput(output)
		color.Cyan("GPU Information:")
		for key, value := range gpuInfo {
			color.Green("%s: %s", key, value)
		}
	case err := <-errChan:
		color.Red("Error: %v", err)
		color.Red("Your system is not using Nvidia. `go-cli info gpu` works only on machines using Nvidia and Darwin systems.")
	}
}

// parseNvidiaSMIOutput parses the output of the nvidia-smi command on Nvidia systems
func parseNvidiaSMIOutput(output string) map[string]string {
	gpuInfo := make(map[string]string)
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			gpuInfo[key] = value
		}
	}
	return gpuInfo
}
