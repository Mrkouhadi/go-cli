package info

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

// subCOMMAND
// gpuCmd represents the gpu command
var gpuCmd = &cobra.Command{
	Use:   "gpu",
	Short: "Get the GPU Deatils",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// check what OS is this cli running on :
		switch os := runtime.GOOS; os {
		case "darwin":
			getGPUInfoDarwin()
		case "windows":
			getGPUInfoNvidia()
		case "linux":
			getGPUInfoNvidia()
		default:
			fmt.Printf("Your Operating System -- %s --  is Not supported !\n", os)
		}
	},
}

func init() {
	// Here you will define your flags and configuration settings.
	InfoCmd.AddCommand(gpuCmd)
}

// Darwin system
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
		fmt.Println("GPU Information:")
		for key, value := range gpuInfo {
			fmt.Printf("%s: %s\n", key, value)
		}
	case err := <-errChan:
		fmt.Println("Error:", err)
	}
}

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

// ////////////////// NVIDIA:
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
		fmt.Println("GPU Information:")
		for key, value := range gpuInfo {
			fmt.Printf("%s: %s\n", key, value)
		}
	case <-errChan:
		fmt.Println("Your system is not using Nvidia. `Go-cli info gpu` works only on machines using Nvidia and Darwin system.")
	}
}

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
