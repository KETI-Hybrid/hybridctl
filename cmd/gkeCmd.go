package cmd

import (
	cobrautil "Hybrid_Cloud/hybridctl/util"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var GKEInitCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize or reinitialize gcloud",
	Long:  `hybridctl gke init`,
	Run: func(cmd *cobra.Command, args []string) {

		var arguments []string
		arguments = append(arguments, "gcloud", cobrautil.CONFIGURATION, cobrautil.PROJECT_ID)

		if cobrautil.ZONE != "" {
			arguments = append(arguments, cobrautil.ZONE)
		}

		if cobrautil.REGION != "" {
			arguments = append(arguments, cobrautil.REGION)
		}

		command := &exec.Cmd{
			Path:   "./gcloud-init.sh",
			Args:   arguments,
			Stdout: os.Stdout,
			Stderr: os.Stderr,
		}
		err := command.Start()
		if err != nil {
			fmt.Println(err)
		}
		err = command.Wait()
		if err != nil {
			fmt.Println(err)
		}
	},
}
