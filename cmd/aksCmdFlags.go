package cmd

func aksFlags() {
	aksCmd.PersistentFlags().StringP("resource-group", "g", "", "resourceGroup name")
	aksCmd.PersistentFlags().StringP("name", "n", "", "clustername")

	StartCmd.MarkPersistentFlagRequired("resource-group")
	StartCmd.MarkPersistentFlagRequired("name")

	StopCmd.MarkPersistentFlagRequired("resource-group")
	StopCmd.MarkPersistentFlagRequired("name")

	GetOSoptionsCmd.PersistentFlags().StringP("location", "l", "", "location")
	GetOSoptionsCmd.MarkPersistentFlagRequired("location")

	MaintenanceconfigurationCmd.MarkPersistentFlagRequired("resource-group")
	MaintenanceconfigurationCmd.MarkPersistentFlagRequired("name")

	MCAddCmd.Flags().StringP("config-name", "c", "", "configname")
	MCAddCmd.MarkFlagRequired("config-name")
	MCAddCmd.Flags().StringP("config-file", "", "", "configfile")
	MCAddCmd.MarkFlagRequired("config-file")

	MCDeleteCmd.Flags().StringP("config-name", "c", "", "configname")
	MCDeleteCmd.MarkFlagRequired("config-name")

	MCUpdateCmd.Flags().StringP("config-name", "c", "", "configname")
	MCUpdateCmd.MarkFlagRequired("config-name")
	MCUpdateCmd.Flags().StringP("config-file", "", "", "configfile")
	MCUpdateCmd.MarkFlagRequired("config-file")

	MCShowCmd.Flags().StringP("config-name", "c", "", "configname")
	MCShowCmd.MarkFlagRequired("config-name")
}
