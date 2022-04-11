package cmd

import (
	cobrautil "Hybrid_Cloud/hybridctl/util"
)

func gkeFlags() {
	GKEInitCmd.Flags().StringVarP(&cobrautil.CONFIGURATION, "configuration", "", "", "CONFIGURATION")
	GKEInitCmd.MarkFlagRequired("configuration")
	GKEInitCmd.Flags().StringVarP(&cobrautil.PROJECT_ID, "project-id", "", "", "PROJECT_ID")
	GKEInitCmd.MarkFlagRequired("project-id")
	GKEInitCmd.Flags().StringVarP(&cobrautil.ZONE, "zone", "", "", "Compute zone (e.g. us-central1-a) for the cluster. Overrides the default compute/zone property value for this command invocation.")
	GKEInitCmd.Flags().StringVarP(&cobrautil.REGION, "region", "", "", "REGION")

	GKEImagesListCmd.Flags().StringP("repository", "", "", "The name of the repository. Format: *.gcr.io/repository. Defaults to gcr.io/<project>, for the active project.")
	GKEImagesListCmd.Flags().StringP("filter", "", "", "Apply a Boolean filter EXPRESSION to each resource item to be listed. If the expression evaluates True, then that item is listed. For more details and examples of filter expressions, run $ gcloud topic filters. This flag interacts with other flags that are applied in this order: --flatten, --sort-by, --filter, --limit.")
	GKEImagesListCmd.Flags().StringP("limit", "", "", "Maximum number of resources to list. The default is unlimited. This flag interacts with other flags that are applied in this order: --flatten, --sort-by, --filter, --limit.")
	GKEImagesListCmd.Flags().StringP("page-size", "", "", "Some services group resource list output into pages. This flag specifies the maximum number of resources per page. The default is determined by the service if it supports paging, otherwise it is unlimited (no paging). Paging may be applied before or after --filter and --limit depending on the service.")
	GKEImagesListCmd.Flags().StringP("sort-by", "", "", "Comma-separated list of resource field key names to sort by. The default order is ascending. Prefix a field with ``~´´ for descending order on that field. This flag interacts with other flags that are applied in this order: --flatten, --sort-by, --filter, --limit.")
	GKEImagesListCmd.Flags().BoolP("uri", "", false, "Print a list of resource URIs instead of the default output, and change the command output to a list of URIs. If this flag is used with --format, the formatting is applied on this URI list. To display URIs alongside other keys instead, use the uri() transform.")

	GKEImagesDeleteCmd.Flags().BoolP("force-delete-tags", "", false, "If there are tags pointing to an image to be deleted then they must all be specified explicitly, or this flag must be specified, for the command to succeed.")

	GKEImagesListTagsCmd.Flags().StringP("filter", "", "", "Apply a Boolean filter EXPRESSION to each resource item to be listed. If the expression evaluates True, then that item is listed. For more details and examples of filter expressions, run $ gcloud topic filters. This flag interacts with other flags that are applied in this order: --flatten, --sort-by, --filter, --limit.")
	GKEImagesListTagsCmd.Flags().StringP("limit", "", "", "Maximum number of resources to list. The default is unlimited. This flag interacts with other flags that are applied in this order: --flatten, --sort-by, --filter, --limit.")
	GKEImagesListTagsCmd.Flags().StringP("page-size", "", "", "Some services group resource list output into pages. This flag specifies the maximum number of resources per page. The default is determined by the service if it supports paging, otherwise it is unlimited (no paging). Paging may be applied before or after --filter and --limit depending on the service.")
	GKEImagesListTagsCmd.Flags().StringP("sort-by", "", "", "Comma-separated list of resource field key names to sort by. The default order is ascending. Prefix a field with ``~´´ for descending order on that field. This flag interacts with other flags that are applied in this order: --flatten, --sort-by, --filter, --limit.")

	GKEAuthListCmd.Flags().StringP("filter-account", "", "", "List only credentials for one account. Use --filter=\"account~PATTERN\" to select accounts that match PATTERN.")
	GKEAuthListCmd.Flags().StringP("filter", "", "", "Apply a Boolean filter EXPRESSION to each resource item to be listed. If the expression evaluates True, then that item is listed. For more details and examples of filter expressions, run $ gcloud topic filters. This flag interacts with other flags that are applied in this order: --flatten, --sort-by, --filter, --limit.")
	GKEAuthListCmd.Flags().StringP("limit", "", "", "Maximum number of resources to list. The default is unlimited. This flag interacts with other flags that are applied in this order: --flatten, --sort-by, --filter, --limit.")
	GKEAuthListCmd.Flags().StringP("page-size", "", "", "Some services group resource list output into pages. This flag specifies the maximum number of resources per page. The default is determined by the service if it supports paging, otherwise it is unlimited (no paging). Paging may be applied before or after --filter and --limit depending on the service.")
	GKEAuthListCmd.Flags().StringP("sort-by", "", "", "Comma-separated list of resource field key names to sort by. The default order is ascending. Prefix a field with ``~´´ for descending order on that field. This flag interacts with other flags that are applied in this order: --flatten, --sort-by, --filter, --limit.")

	GKEAuthRevokeCmd.Flags().BoolP("all", "", false, "Revoke credentials for all accounts.")

	GKEAuthLoginCmd.Flags().StringP("cred-file", "", "", "Path to the external account configuration file (workload identity pool, generated by the Cloud Console or gcloud iam workload-identity-pools create-cred-config) or service account credential key file (JSON).")
	GKEAuthLoginCmd.MarkFlagRequired("cred-file")

	GKEDockerCmd.Flags().BoolP("authorize-only", "a", false, "Configure Docker authorization only; do not launch the Docker command-line.")
	GKEDockerCmd.Flags().StringP("docker-host", "", "", "URL to connect to Docker Daemon. Format: tcp://host:port or unix:///path/to/socket.")
	GKEDockerCmd.Flags().StringP("server", "", "", "Address of the Google Cloud Registry.")
	GKEDockerCmd.Flags().StringP("args", "", "", "Arguments to pass to Docker. The '--' argument must be specified between gcloud specific args on the left and DOCKER_ARGS on the right.")

	GKENodePoolsRollbackCmd.Flags().StringP("cluster", "", "", "The cluster from which to rollback the node pool. Overrides the default container/cluster property value for this command invocation.")
	GKENodePoolsRollbackCmd.Flags().StringP("zone", "z", "", "Compute zone (e.g. us-central1-a) for the cluster. Overrides the default compute/zone property value for this command invocation.")

	GKEOperationDescribeCmd.Flags().StringVarP(&cobrautil.ZONE, "zone", "z", "", "Compute zone (e.g. us-central1-a) for the cluster. Overrides the default compute/zone property value for this command invocation.")
	GKEOperationsListCmd.Flags().StringVarP(&cobrautil.ZONE, "zone", "z", "", "Compute zone (e.g. us-central1-a) for the cluster. Overrides the default compute/zone property value for this command invocation.")
	GKEOperationsWaitCmd.Flags().StringVarP(&cobrautil.ZONE, "zone", "z", "", "Compute zone (e.g. us-central1-a) for the cluster. Overrides the default compute/zone property value for this command invocation.")

	GKEProjectConfigsUpdateCmd.Flags().BoolP("disable-pushblock", "", false, "Disable PushBlock for all repositories under current project. PushBlock allows repository owners to block git push transactions containing private key data.")
	GKEProjectConfigsUpdateCmd.Flags().BoolP("enable-pushblock", "", false, "Enable PushBlock for all repositories under current project. PushBlock allows repository owners to block git push transactions containing private key data.")
	GKEProjectConfigsUpdateCmd.Flags().StringP("message-format", "", "", "The format of the message to publish to the topic. MESSAGE_FORMAT must be one of: json, protobuf.")
	GKEProjectConfigsUpdateCmd.Flags().StringP("service-account", "", "", "Email address of the service account used for publishing Cloud Pub/Sub messages. This service account needs to be in the same project as the project. \nWhen added, the caller needs to have iam.serviceAccounts.actAs permission on this service account.\n If unspecified, it defaults to the Compute Engine default service account.")
	GKEProjectConfigsUpdateCmd.Flags().StringP("topic-project", "", "", "Cloud project for the topic. If not set, the currently set project will be used.")
	GKEProjectConfigsUpdateCmd.Flags().StringP("add-topic", "", "", "ID of the topic or fully qualified identifier for the topic.")
	GKEProjectConfigsUpdateCmd.Flags().StringP("remove-topic", "", "", "ID of the topic or fully qualified identifier for the topic.")
	GKEProjectConfigsUpdateCmd.Flags().StringP("update-topic", "", "", "ID of the topic or fully qualified identifier for the topic.")

	GKEConfigSetCmd.Flags().BoolP("installation", "", false, "If set, the property is updated for the entire Google Cloud CLI installation. Otherwise, by default, the property is updated only in the currently active configuration.")
}
