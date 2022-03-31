package cmd

import (
	cobrautil "Hybrid_Cloud/hybridctl/util"
)

func gkeFlags() {
	GKEInitCmd.Flags().StringVarP(&cobrautil.CONFIGURATION, "configuration", "", "", "CONFIGURATION")
	GKEInitCmd.MarkFlagRequired("configuration")
	GKEInitCmd.Flags().StringVarP(&cobrautil.PROJECT_ID, "project-id", "", "", "PROJECT_ID")
	GKEInitCmd.MarkFlagRequired("project-id")
	GKEInitCmd.Flags().StringVarP(&cobrautil.ZONE, "zone", "", "", "ZONE")
	GKEInitCmd.Flags().StringVarP(&cobrautil.REGION, "region", "", "", "REGION")
}
