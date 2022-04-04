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

	gkeCmd.AddCommand(GKEContainerCmd)
	GKEContainerCmd.AddCommand(GKEContainerImagesCmd)
	GKEContainerImagesCmd.AddCommand(GKEImagesAddTagCmd)
	GKEContainerImagesCmd.AddCommand(GKEImagesDeleteCmd)
	GKEContainerImagesCmd.AddCommand(GKEImagesDescribeCmd)
	GKEContainerImagesCmd.AddCommand(GKEImagesListCmd)
	GKEContainerImagesCmd.AddCommand(GKEImagesListTagsCmd)
	GKEContainerImagesCmd.AddCommand(GKEImagesUnTagCmd)

	GKEContainerCmd.AddCommand(GKEContainerOperationsCmd)
	GKEContainerOperationsCmd.AddCommand(GKEOperationDescribeCmd)
	GKEContainerOperationsCmd.AddCommand(GKEOperationsListCmd)
	GKEContainerOperationsCmd.AddCommand(GKEOperationsWaitCmd)

	gkeCmd.AddCommand(GKEAuthCmd)
	GKEAuthCmd.AddCommand(GKEAuthConfigureDockerCmd)
	GKEAuthCmd.AddCommand(GKEAuthListCmd)
	GKEAuthCmd.AddCommand(GKEAuthLoginCmd)
	GKEAuthCmd.AddCommand(GKEAuthRevokeCmd)

	gkeFlags()
}
