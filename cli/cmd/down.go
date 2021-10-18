package cmd

import (
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// downCmd represents the down command
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Shut down the environment",
	Long:  `Runs the docker-compose down command for you in order to shut down the local development environment.`,
	Run: func(cmd *cobra.Command, args []string) {
		call := exec.Command("docker-compose", "-f", "../.config/docker-compose.yaml", "down")
		call.Stdout = os.Stdout
		call.Stderr = os.Stderr
		err := call.Run()
		if err != nil {
			log.Fatalf("down command failed with %s\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(downCmd)
}
