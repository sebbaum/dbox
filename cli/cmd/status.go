package cmd

import (
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Print a status of the environment",
	Long: `Prints the docker-composer ps command's output to get status information about the containers in the
environment`,
	Run: func(cmd *cobra.Command, args []string) {
		call := exec.Command("docker-compose", "ps")
		call.Stdout = os.Stdout
		call.Stderr = os.Stderr
		err := call.Run()
		if err != nil {
			log.Fatalf("status command failed with %s\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
