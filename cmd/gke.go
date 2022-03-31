package cmd

import (
	"github.com/spf13/cobra"
)

// gkeCmd represents the gke command
var gkeCmd = &cobra.Command{
	Use:   "gke",
	Short: "A brief description of your command",
	Long: ` 

	`,
}

func init() {
	RootCmd.AddCommand(gkeCmd)
	gkeCmd.AddCommand(GKEInitCmd)
	gkeFlags()
}
