package cmd

import (
	apiserverutil "Hybrid_Cloud/hcp-apiserver/pkg/util"
	"Hybrid_Cloud/hybridctl/util"
	cobrautil "Hybrid_Cloud/hybridctl/util"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	containerpb "google.golang.org/genproto/googleapis/container/v1"
)

const (
	GKE_CONTAINER_PATH = "/gke/container"
	GKE_AUTH_PATH      = "/gke/auth"
	GKE_CONFIG_PATH    = "/gke/config"
	GKE_HELP           = "Use \"hybridctl gke [command] --help\" for more information about a command."
)

var GKEContainerCmd = &cobra.Command{
	Use:   "container",
	Short: "deploy and manage clusters of machines for running containers",
}

var GKEContainerImagesCmd = &cobra.Command{
	Use:   "images",
	Short: "list and manipulate Google Container Registry images",
}

var GKEContainerOperationsCmd = &cobra.Command{
	Use:   "operations",
	Short: "get and list operations for Google Kubernetes Engine clusters",
}

var GKEContainerGetServerConfigCmd = &cobra.Command{
	Use:   "get-server-config",
	Short: "list and manipulate Google Container Registry images",
	Run: func(cmd *cobra.Command, args []string) {

		/*
			input := &containerpb.GetServerConfigRequest{
				ProjectId: "keti-container",
				Zone:      "us-central1-a",
			}
		*/
		ReloadGKEConfigValue()
		input := &containerpb.GetServerConfigRequest{
			ProjectId: os.Getenv("GKE_PROJECT_ID"),
		}
		zone, _ := cmd.Flags().GetString("zone")
		if zone == "" {
			input.Zone = os.Getenv("GKE_DEFAULT_ZONE")
		} else {
			input.Zone = zone
		}

		httpPostUrl := "http://localhost:8080" + GKE_CONTAINER_PATH + "/server-config/get"
		bytes := util.HTTPPostRequest(input, httpPostUrl)

		var output apiserverutil.Output
		json.Unmarshal(bytes, &output)
		if output.Stderr != nil {
			fmt.Println(string(output.Stderr))
		}

		if output.Stdout != nil {
			stdout := output.Stdout
			var resp containerpb.ServerConfig
			json.Unmarshal(stdout, &resp)
			fmt.Printf("Fetching server config for %s\n", input.Zone)
			PrintServerConfig(resp)
		}
	},
}

func PrintServerConfig(resp containerpb.ServerConfig) {
	//	var field string
	fmt.Println("channels:")
	for _, c := range resp.Channels {
		fmt.Println("- channel:", c.GetChannel())
		fmt.Printf("  defaultVersion: %s\n", c.GetDefaultVersion())
		fmt.Println("  validVersions:")
		for _, j := range c.GetValidVersions() {
			fmt.Println("  - ", j)
		}
	}

	fmt.Println("defaultClusterVersion: ", resp.DefaultClusterVersion)
	fmt.Println("defaultImageType: ", resp.DefaultImageType)

	fmt.Println("validImageTypes:")
	for _, c := range resp.ValidImageTypes {
		fmt.Println("- ", c)
	}

	fmt.Println("validMasterVersions:")
	for _, c := range resp.ValidMasterVersions {
		fmt.Println("- ", c)
	}

	fmt.Println("validNodeVersions:")
	for _, c := range resp.ValidNodeVersions {
		fmt.Println("- ", c)
	}
}

var GKEImagesAddTagCmd = &cobra.Command{
	Use:   "add-tag",
	Short: "adds tags to existing image",
	Run: func(cmd *cobra.Command, args []string) {

		// gcloud container images add-tag SRC_IMAGE DEST_IMAGE
		if len(args) < 2 {
			cmd.Help()
		} else {
			/*
				i = &Images{
					SRC_IMAGE:  "gcr.io/keti-container/busybox",
					DEST_IMAGE: "gcr.io/keti-container/busybox:mytag3",
				}
			*/

			temp := []string{}
			for i := 1; i < len(args); i++ {
				temp = append(temp, args[i])
			}

			input := &util.GKEImages{
				SRC_IMAGE:  args[0],
				DEST_IMAGE: temp,
			}

			httpPostUrl := "http://localhost:8080" + GKE_CONTAINER_PATH + "/images/tag/add"
			bytes := util.HTTPPostRequest(input, httpPostUrl)
			util.PrintOutput(bytes)
		}
	},
}

var GKEImagesDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete existing images",
	Run: func(cmd *cobra.Command, args []string) {

		// gcloud container images delete IMAGE_NAME
		if len(args) < 1 {
			cmd.Help()
		} else {
			/*
				i = &Images{
					IMAGE_NAME: "gcr.io/keti-container/busybox",
				}
			*/

			temp := []string{args[0]}
			input := &util.GKEImages{
				IMAGE_NAME: temp,
			}

			bol, _ := cmd.Flags().GetBool("force-delete-tags")
			input.FORCE_DELETE_TAGS = bol

			httpPostUrl := "http://localhost:8080" + GKE_CONTAINER_PATH + "/images/delete"
			bytes := util.HTTPPostRequest(input, httpPostUrl)
			util.PrintOutput(bytes)
		}
	},
}

var GKEImagesDescribeCmd = &cobra.Command{
	Use:   "describe",
	Short: "lists information about the specified image",
	Run: func(cmd *cobra.Command, args []string) {

		// gcloud container images describe IMAGE_NAME
		if len(args) < 1 {
			cmd.Help()
		} else {
			/*
				i = &Images{
					IMAGE_NAME: "gcr.io/keti-container/busybox",
				}
			*/
			temp := []string{args[0]}
			input := &util.GKEImages{
				IMAGE_NAME: temp,
			}
			httpPostUrl := "http://localhost:8080" + GKE_CONTAINER_PATH + "/images/describe"
			bytes := util.HTTPPostRequest(input, httpPostUrl)
			util.PrintOutput(bytes)
		}
	},
}

var GKEImagesListCmd = &cobra.Command{
	Use:   "list",
	Short: "list existing images",
	Run: func(cmd *cobra.Command, args []string) {

		// gcloud container images list [--repository=REPOSITORY] [--filter=EXPRESSION] [--limit=LIMIT] [--page-size=PAGE_SIZE] [--sort-by=[FIELD,…]] [--uri]

		/*
			i = &Images{
				IMAGE_NAME: "gcr.io/keti-container/busybox",
			}
		*/
		var input util.GKEImages

		str, _ := cmd.Flags().GetString("repository")
		if str != "" {
			input.REPOSITORY = str
		}

		str, _ = cmd.Flags().GetString("filter")
		if str != "" {
			input.FILTER = str
		}

		str, _ = cmd.Flags().GetString("limit")
		if str != "" {
			input.LIMIT = str
		}

		str, _ = cmd.Flags().GetString("page-size")
		if str != "" {
			input.PAGE_SIZE = str
		}

		str, _ = cmd.Flags().GetString("sort-by")
		if str != "" {
			input.SORT_BY = str
		}

		bol, _ := cmd.Flags().GetBool("uri")
		input.URI = bol

		httpPostUrl := "http://localhost:8080" + GKE_CONTAINER_PATH + "/images/list"
		bytes := util.HTTPPostRequest(input, httpPostUrl)
		util.PrintOutput(bytes)
	},
}

var GKEImagesListTagsCmd = &cobra.Command{
	Use:   "list-tags",
	Short: "list tags and digests for the specified image",
	Run: func(cmd *cobra.Command, args []string) {

		// gcloud container images list-tags IMAGE_NAME [--filter=EXPRESSION] [--limit=LIMIT] [--page-size=PAGE_SIZE] [--sort-by=[FIELD,…]; default="~timestamp"]
		if len(args) < 1 {
			cmd.Help()
		} else {
			/*
				i = &Images{
					IMAGE_NAME: "gcr.io/keti-container/busybox",
				}
			*/
			temp := []string{args[0]}
			input := &util.GKEImages{
				IMAGE_NAME: temp,
			}

			str, _ := cmd.Flags().GetString("filter")
			if str != "" {
				input.FILTER = str
			}

			str, _ = cmd.Flags().GetString("limit")
			if str != "" {
				input.LIMIT = str
			}

			str, _ = cmd.Flags().GetString("page-size")
			if str != "" {
				input.PAGE_SIZE = str
			}

			str, _ = cmd.Flags().GetString("sort-by")
			if str != "" {
				input.SORT_BY = str
			}

			httpPostUrl := "http://localhost:8080" + GKE_CONTAINER_PATH + "/images/tag/list"
			bytes := util.HTTPPostRequest(input, httpPostUrl)
			util.PrintOutput(bytes)
		}
	},
}

var GKEImagesUnTagCmd = &cobra.Command{
	Use:   "untag",
	Short: "remove existing image tags",
	Run: func(cmd *cobra.Command, args []string) {

		// gcloud container images untag IMAGE_NAME
		if len(args) < 1 {
			cmd.Help()
		} else {
			/*
				i = &Images{
					IMAGE_NAME: "gcr.io/keti-container/busybox:mytag3",
				}
			*/
			var input util.GKEImages

			temp := []string{}
			for i := 0; i < len(args); i++ {
				temp = append(temp, args[i])
			}

			input.IMAGE_NAME = temp

			httpPostUrl := "http://localhost:8080" + GKE_CONTAINER_PATH + "/images/untags"
			bytes := util.HTTPPostRequest(input, httpPostUrl)
			util.PrintOutput(bytes)
		}
	},
}

var GKEOperationDescribeCmd = &cobra.Command{
	Use:   "describe",
	Short: "describe an operation",
	Long:  `gcloud container operations describe OPERATION_ID [--zone=ZONE, -z ZONE]`,
	Run: func(cmd *cobra.Command, args []string) {

		// gcloud container operations describe OPERATION_ID
		if len(args) != 1 {
			cmd.Help()
		} else {
			/*
				input := &containerpb.GetOperationRequest{
					ProjectId:   "keti-container",
					Zone:        "us-central1-a",
					OperationId: "operation-1648309236003-34160983",
					Name:        "operation-1648309236003-34160983",
				}
			*/
			ReloadGKEConfigValue()
			input := &containerpb.GetOperationRequest{
				ProjectId:   os.Getenv("GKE_PROJECT_ID"),
				OperationId: args[0],
			}
			zone, _ := cmd.Flags().GetString("zone")
			if zone == "" {
				input.Zone = os.Getenv("GKE_DEFAULT_ZONE")
			} else {
				input.Zone = zone
			}

			httpPostUrl := "http://localhost:8080" + GKE_CONTAINER_PATH + "/operations/describe"
			bytes := util.HTTPPostRequest(input, httpPostUrl)

			var output apiserverutil.Output
			json.Unmarshal(bytes, &output)
			if output.Stderr != nil {
				fmt.Println(string(output.Stderr))
			}

			if output.Stdout != nil {
				stdout := output.Stdout
				var resp *containerpb.Operation
				json.Unmarshal(stdout, &resp)
				header := []string{"EndTime", "Name", "OperationType", "SelfLink", "StartTime", "Status", "TargetLink", "Zone"}
				for _, i := range header {
					fmt.Printf("%s: %s\n", i, reflect.ValueOf(resp).Elem().FieldByName(i))
				}
			}
		}

	},
}

var GKEOperationsListCmd = &cobra.Command{
	Use:   "list",
	Short: "list operations for container clusters",
	Long:  `gcloud container operations list [--zone=ZONE, -z ZONE]`,
	Run: func(cmd *cobra.Command, args []string) {

		// gcloud container operations list [--zone=ZONE, -z ZONE]
		/*
			op := &containerpb.ListOperationsRequest{
				ProjectId: "keti-container",
				Zone:      "-",
			}
		*/
		if len(args) > 0 {
			cmd.Help()
		} else {
			ReloadGKEConfigValue()
			input := &containerpb.ListOperationsRequest{
				ProjectId: os.Getenv("GKE_PROJECT_ID"),
			}

			zone, _ := cmd.Flags().GetString("zone")
			if zone == "" {
				input.Zone = "-"
			} else {
				input.Zone = zone
			}

			httpPostUrl := "http://localhost:8080" + GKE_CONTAINER_PATH + "/operations/list"
			bytes := util.HTTPPostRequest(input, httpPostUrl)

			var output apiserverutil.Output
			json.Unmarshal(bytes, &output)
			if output.Stderr != nil {
				fmt.Println(string(output.Stderr))
			}

			if output.Stdout != nil {
				stdout := output.Stdout
				var resp *containerpb.ListOperationsResponse
				json.Unmarshal(stdout, &resp)
				table := tablewriter.NewWriter(os.Stdout)
				header := []string{"NAME", "TYPE", "LOCATION", "TARGET", "STATUS_MESSAGE", "STATUS", "START_TIME", "END_TIME"}
				table.SetHeader(header)
				for _, v := range resp.Operations {
					targetLink := v.GetTargetLink()
					target := targetLink[strings.LastIndex(targetLink, "/")+1:]
					temp := []string{v.Name, v.OperationType.String(), v.Location, target, v.StatusMessage, v.Status.String(), v.StartTime, v.EndTime}
					table.Append(temp)
				}
				table.Render()
			}
		}
	},
}

var GKEOperationsWaitCmd = &cobra.Command{
	Use:   "wait",
	Short: "poll an operation for completion",
	Long:  `gcloud container operations wait OPERATION_ID [--zone=ZONE, -z ZONE]`,
	Run: func(cmd *cobra.Command, args []string) {

		// gcloud container operations wait OPERATION_ID [--zone=ZONE, -z ZONE]
		if len(args) != 1 {
			cmd.Help()
		} else {

			var input = &util.GKEOperations{
				OPERATION_ID: args[0],
			}

			zone, _ := cmd.Flags().GetString("zone")
			if zone != "" {
				input.ZONE = zone
			}

			httpPostUrl := "http://localhost:8080" + GKE_CONTAINER_PATH + "/operations/wait"
			bytes := util.HTTPPostRequest(input, httpPostUrl)
			util.PrintOutput(bytes)
		}

	},
}

var GKEContainerNodePoolsCmd = &cobra.Command{
	Use:   "node-pools",
	Short: "rollback a node-pool upgrade",
}

var GKENodePoolsRollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "rollback a node-pool upgrade",
	Run: func(cmd *cobra.Command, args []string) {

		// gcloud container node-pools rollback NAME [--async] [--cluster=CLUSTER] [--region=REGION     | --zone=ZONE, -z ZONE]
		if len(args) < 1 {
			cmd.Help()
		} else {
			/*
				input := &containerpb.RollbackNodePoolUpgradeRequest{
					ProjectId: "keti-container",
					Zone:      "us-central1-a",
					ClusterId: "hcp-cluster",
					Name:      "pool-1",
				}
			*/
			ReloadGKEConfigValue()
			input := &containerpb.RollbackNodePoolUpgradeRequest{
				ProjectId: os.Getenv("GKE_PROJECT_ID"),
				Name:      args[0],
				ClusterId: os.Getenv("GKE_DEFAULT_CLUSTER"),
			}

			cluster, _ := cmd.Flags().GetString("cluster")
			if cluster == "" {
				input.ClusterId = os.Getenv("GKE_DEFAULT_CLUSTSER")
			} else {
				input.ClusterId = cluster
			}

			zone, _ := cmd.Flags().GetString("zone")
			if zone == "" {
				input.Zone = os.Getenv("GKE_DEFAULT_ZONE")
			} else {
				input.Zone = zone
			}

			httpPostUrl := "http://localhost:8080" + GKE_CONTAINER_PATH + "/nodepool-upgrade/rollback"
			bytes := util.HTTPPostRequest(input, httpPostUrl)

			var output apiserverutil.Output
			json.Unmarshal(bytes, &output)
			if output.Stderr != nil {
				fmt.Println(string(output.Stderr))
			}

			if output.Stdout != nil {
				stdout := output.Stdout
				var resp containerpb.Operation
				json.Unmarshal(stdout, &resp)
				fmt.Printf("Updated [%s]\n", resp.TargetLink)
				fmt.Printf("operationId: %s\nprojectId: %s\nzone: %s\n", resp.GetName(), resp.GetZone(), input.GetProjectId())
			}
		}
	},
}

var GKEAuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "manage oauth2 credentials for the Google Cloud CLI",
}

var GKEAuthConfigureDockerCmd = &cobra.Command{
	Use:   "configure-docker",
	Short: "register gcloud as a Docker credential helper",
	Long:  `hybridctl auth configure-docker [REGISTRIES]`,
	Run: func(cmd *cobra.Command, args []string) {

		// gcloud auth configure-docker [REGISTRIES]

		if len(args) > 1 {
			cmd.Help()
		} else {
			var input util.GKEAuth
			if len(args) == 1 {
				input.REGISTRIES = args[0]
			}

			/*
				var temp []string
				if len(args) == 1 {
					temp = strings.Split(args[0], ",")
				}
				input.REGISTRIES = append(input.REGISTRIES, temp...)
			*/

			httpPostUrl := "http://localhost:8080" + GKE_AUTH_PATH + "/configure-docker"
			bytes := util.HTTPPostRequest(input, httpPostUrl)
			util.PrintOutput(bytes)
		}
	},
}

var GKEAuthListCmd = &cobra.Command{
	Use:   "list",
	Short: "lists credentialed accounts",
	Long:  `hybridctl auth list [--filter-account=FILTER_ACCOUNT] [--filter=EXPRESSION] [--limit=LIMIT] [--page-size=PAGE_SIZE] [--sort-by=[FIELD,…]]`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 0 {
			cmd.Help()
		} else {
			// gcloud auth list
			var input util.GKEAuth

			str, _ := cmd.Flags().GetString("filter-account")
			if str != "" {
				input.FILTER_ACCOUNT = str
			}

			str, _ = cmd.Flags().GetString("filter")
			if str != "" {
				input.FILTER = str
			}

			str, _ = cmd.Flags().GetString("limit")
			if str != "" {
				input.LIMIT = str
			}

			str, _ = cmd.Flags().GetString("page-size")
			if str != "" {
				input.PAGE_SIZE = str
			}

			str, _ = cmd.Flags().GetString("sort-by")
			if str != "" {
				input.SORT_BY = str
			}

			httpPostUrl := "http://localhost:8080" + GKE_AUTH_PATH + "/list"
			bytes := util.HTTPPostRequest(input, httpPostUrl)
			util.PrintOutput(bytes)
		}
	},
}

var GKEAuthRevokeCmd = &cobra.Command{
	Use:   "revoke",
	Short: "revoke access credentials for an account",
	Long:  "hybridctl auth revoke [ACCOUNTS …] [--all]",
	Run: func(cmd *cobra.Command, args []string) {

		// gcloud auth revoke [ACCOUNTS …] [--all]

		if len(args) > 1 {
			cmd.Help()
		} else {

			var input util.GKEAuth
			if len(args) == 1 {
				input.ACCOUNTS = args[0]
			}

			// set flags
			bol, _ := cmd.Flags().GetBool("all")
			input.ALL = bol

			httpPostUrl := "http://localhost:8080" + GKE_AUTH_PATH + "/revoke"
			bytes := util.HTTPPostRequest(input, httpPostUrl)
			util.PrintOutput(bytes)
		}
	},
}

var GKEAuthLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "authorize gcloud to access the Cloud Platform with Google user credentials",
	Long:  `hybridcl auth login [ACCOUNT] [--cred-file=CRED_FILE]`,
	Run: func(cmd *cobra.Command, args []string) {

		// gcloud auth login [ACCOUNT] [--cred-file=CRED_FILE]

		if len(args) > 1 {
			cmd.Help()
		} else {
			/*
				a = &Auth{
					CRED_FILE: "/root/hcp-key.json",
				}
			*/

			var input util.GKEAuth
			if len(args) == 1 {
				input.ACCOUNTS = args[0]
			}

			str, _ := cmd.Flags().GetString("cred-file")
			input.CRED_FILE = str

			httpPostUrl := "http://localhost:8080" + GKE_AUTH_PATH + "/login"
			bytes := util.HTTPPostRequest(input, httpPostUrl)
			util.PrintOutput(bytes)
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

		if len(args) > 0 {
			cmd.Help()
		} else {
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
		}
	},
}

var GKEConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "authorize gcloud to access the Cloud Platform with Google user credentials",
}

var GKEConfigSetCmd = &cobra.Command{
	Use:   "set",
	Short: "set a Google Cloud CLI property",
	Long:  `hybridctl gke config set SECTION/PROPERTY VALUE [--installation]`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 2 {
			cmd.Help()
		} else {
			// gcloud config set SECTION/PROPERTY VALUE [--installation] [GCLOUD_WIDE_FLAG …]

			/*
				input := SetProperty{
					SECTION:  "compute",
					PROPERTY: "zone",
					VALUE:    "us-central1-a",
				}
			*/

			var input util.GKESetProperty

			// set SECTION/PROPERTY
			if strings.Contains(args[0], "/") {
				cnt := strings.Count(args[0], "/")
				if cnt != 1 {
					fmt.Println("ERROR: Invalid Input. Enter in the correct command format.\n Usage: hybridctl gke config set SECTION/PROPERTY VALUE")
					return
				}
				temp := strings.Split(args[0], "/")
				input.SECTION = temp[0]
				input.PROPERTY = temp[1]
			} else {
				input.PROPERTY = args[0]
			}

			// set VALUE
			input.VALUE = args[1]

			// set flags
			bol, _ := cmd.Flags().GetBool("installation")
			input.INSTALLATION = bol

			httpPostUrl := "http://localhost:8080/gke/config/set"
			bytes := util.HTTPPostRequest(input, httpPostUrl)
			util.PrintOutput(bytes)
		}

	},
}

func ReloadGKEConfigValue() {
	cmd := exec.Command("bash", "-c", "gcloud config get-value project")
	_, stdout := apiserverutil.CombinedOutput2(cmd)
	GKE_project_id := strings.ReplaceAll(string(stdout), "\n", "")
	if GKE_project_id == "" {
		os.Setenv("GKE_PROJECT_ID", "")
	} else {
		os.Setenv("GKE_PROJECT_ID", GKE_project_id)
	}

	cmd = exec.Command("bash", "-c", "gcloud config get-value compute/zone")
	_, stdout = apiserverutil.CombinedOutput2(cmd)
	GKE_default_zone := strings.ReplaceAll(string(stdout), "\n", "")
	if GKE_default_zone == "" {
		os.Setenv("GKE_DEFAULT_ZONE", "")
	} else {
		os.Setenv("GKE_DEFAULT_ZONE", GKE_default_zone)
	}

	cmd = exec.Command("bash", "-c", "gcloud config get-value container/cluster")
	_, stdout = apiserverutil.CombinedOutput2(cmd)
	GKE_default_cluster := strings.ReplaceAll(string(stdout), "\n", "")
	if GKE_default_cluster == "" {
		os.Setenv("GKE_DEFAULT_CLUSTER", "")
	} else {
		os.Setenv("GKE_DEFAULT_CLUSTER", GKE_default_cluster)
	}
}
