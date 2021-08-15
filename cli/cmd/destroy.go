package cmd

import (
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// destroyCmd represents the destroy command
var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroy the whole environment",
	Long:  `Destroys the whole environment by shutting down all services and removing the docker images`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: optimize
		// TODO: add a prompt for confirmation
		// TODO: later also remove data volumes of databases for example
		//call := exec.Command("docker-compose", "down")
		call := exec.Command("docker-compose", "down")
		call.Stdout = os.Stdout
		call.Stderr = os.Stderr
		err := call.Run()
		if err != nil {
			log.Fatalf("destroy command failed with %s\n", err)
		}

		// TODO: docker rmi is not working
		call = exec.Command("docker rmi", "cbox_nginx", "cbox_php")
		call.Stdout = os.Stdout
		call.Stderr = os.Stderr
		err = call.Run()
		if err != nil {
			log.Fatalf("destroy command failed with %s\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(destroyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// destroyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// destroyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
