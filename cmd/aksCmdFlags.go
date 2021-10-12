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
}
