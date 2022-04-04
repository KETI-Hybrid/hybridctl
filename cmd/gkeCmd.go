package cmd

import (
	cobrautil "Hybrid_Cloud/hybridctl/util"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var GKE_HELP = "Use \"hybridctl gke container images [command] --help\" for more information about a command."

var GKEContainer = &cobra.Command{
	Use:   "container",
	Short: "deploy and manage clusters of machines for running containers",
}

// images
type Images struct {
	SRC_IMAGE  string
	DEST_IMAGE string
	IMAGE_NAME string
}

var GKEContainerImages = &cobra.Command{
	Use:   "images",
	Short: "list and manipulate Google Container Registry images",
}

var GKEImagesAddTag = &cobra.Command{
	Use:   "add-tag",
	Short: "adds tags to existing image",
	Run: func(cmd *cobra.Command, args []string) {

		// gcloud container images add-tag SRC_IMAGE DEST_IMAGE
		if len(args) < 2 {
			fmt.Println(GKE_HELP)
		} else {
			/*
				i = &Images{
					SRC_IMAGE:  "gcr.io/keti-container/busybox",
					DEST_IMAGE: "gcr.io/keti-container/busybox:mytag3",
				}
			*/
			input := &Images{
				SRC_IMAGE:  args[0],
				DEST_IMAGE: args[1],
			}
			input.AddTag()
		}
	},
}

var GKEImagesDelete = &cobra.Command{
	Use:   "delete",
	Short: "delete existing images",
	Run: func(cmd *cobra.Command, args []string) {

		// gcloud container images delete IMAGE_NAME
		if len(args) < 1 {
			fmt.Println(GKE_HELP)
		} else {
			/*
				i = &Images{
					IMAGE_NAME: "gcr.io/keti-container/busybox",
				}
			*/
			input := &Images{
				IMAGE_NAME: args[0],
			}
			input.Delete()
		}
	},
}

var GKEImagesDescribe = &cobra.Command{
	Use:   "describe",
	Short: "lists information about the specified image",
	Run: func(cmd *cobra.Command, args []string) {

		// gcloud container images describe IMAGE_NAME
		if len(args) < 1 {
			fmt.Println(GKE_HELP)
		} else {
			/*
				i = &Images{
					IMAGE_NAME: "gcr.io/keti-container/busybox",
				}
			*/
			input := &Images{
				IMAGE_NAME: args[0],
			}
			input.Describe()
		}
	},
}

var GKEImagesList = &cobra.Command{
	Use:   "list",
	Short: "list existing images",
	Run: func(cmd *cobra.Command, args []string) {

		// gcloud container images list [--repository=REPOSITORY] [--filter=EXPRESSION] [--limit=LIMIT] [--page-size=PAGE_SIZE] [--sort-by=[FIELD,…]] [--uri]

		/*
			i = &Images{
				IMAGE_NAME: "gcr.io/keti-container/busybox",
			}
		*/
		input := &Images{}
		input.List()

	},
}

var GKEImagesListTags = &cobra.Command{
	Use:   "list-tags",
	Short: "list tags and digests for the specified image",
	Run: func(cmd *cobra.Command, args []string) {

		// gcloud container images list-tags IMAGE_NAME [--filter=EXPRESSION] [--limit=LIMIT] [--page-size=PAGE_SIZE] [--sort-by=[FIELD,…]; default="~timestamp"]
		if len(args) < 1 {
			fmt.Println(GKE_HELP)
		} else {
			/*
				i = &Images{
					IMAGE_NAME: "gcr.io/keti-container/busybox",
				}
			*/
			input := &Images{
				IMAGE_NAME: args[0],
			}
			input.ListTags()
		}
	},
}

var GKEImagesUnTag = &cobra.Command{
	Use:   "untag",
	Short: "remove existing image tags",
	Run: func(cmd *cobra.Command, args []string) {

		// gcloud container images untag IMAGE_NAME
		if len(args) < 1 {
			fmt.Println(GKE_HELP)
		} else {
			/*
				i = &Images{
					IMAGE_NAME: "gcr.io/keti-container/busybox:mytag3",
				}
			*/
			input := &Images{
				IMAGE_NAME: args[0],
			}
			input.UnTags()
		}
	},
}

var GKEInitCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize or reinitialize gcloud",
	Long: `hybridctl gke init
	
	[REQUIRED]    --configuration : configuration name
					1) if exist, just activate configuration
					2) if no exist, create a new configuration

	[REQUIRED]    --project-id : projectID

	[NO REQUIRED] --zone : default zone

	[NO REQUIRED] --region : default region
	`,
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
