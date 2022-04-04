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
	gkeCmd.AddCommand(GKEContainer)
	GKEContainer.AddCommand(GKEContainerImages)
	GKEContainerImages.AddCommand(GKEImagesAddTag)
	GKEContainerImages.AddCommand(GKEImagesDelete)
	GKEContainerImages.AddCommand(GKEImagesDescribe)
	GKEContainerImages.AddCommand(GKEImagesList)
	GKEContainerImages.AddCommand(GKEImagesListTags)
	GKEContainerImages.AddCommand(GKEImagesUnTag)

	gkeFlags()
}
