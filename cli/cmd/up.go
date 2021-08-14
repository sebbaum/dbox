package cmd

import (
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Start up the environment",
	Long:  `Runs the docker-compose up command for you in order to start up the local development environment`,
	Run: func(cmd *cobra.Command, args []string) {
		call := exec.Command("docker-compose", "up", "-d")
		call.Stdout = os.Stdout
		call.Stderr = os.Stderr
		err := call.Run()
		if err != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(upCmd)
}
