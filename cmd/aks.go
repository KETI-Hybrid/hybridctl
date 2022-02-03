// Copyright © 2021 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
)

// aksCmd represents the aks command
var aksCmd = &cobra.Command{
	Use:   "aks",
	Short: "A brief description of your command",
	Long: ` 

	`,
}

func init() {
	RootCmd.AddCommand(aksCmd)
	aksCmd.AddCommand(AddonCmd)
	AddonCmd.AddCommand(AKSDisableAddonsCmd)
	AddonCmd.AddCommand(AKSEnableAddonsCmd)
	AddonCmd.AddCommand(AKSListAddonsCmd)
	AddonCmd.AddCommand(AKSListAddonsAvailableCmd)
	AddonCmd.AddCommand(AKSShowAddonsCmd)
	AddonCmd.AddCommand(AKSUpdateAddonsCmd)

	aksCmd.AddCommand(AKSPodIdentityCmd)
	AKSPodIdentityCmd.AddCommand(AKSPIAddCmd)
	AKSPodIdentityCmd.AddCommand(AKSPIDeleteCmd)
	AKSPodIdentityCmd.AddCommand(AKSPIListCmd)
	AKSPodIdentityCmd.AddCommand(AKSPIExceptionCmd)
	AKSPIExceptionCmd.AddCommand(AKSPIExceptionAddCmd)
	AKSPIExceptionCmd.AddCommand(AKSPIExceptionDeleteCmd)
	AKSPIExceptionCmd.AddCommand(AKSPIExceptionListCmd)
	AKSPIExceptionCmd.AddCommand(AKSPIExceptionUpdateCmd)

	aksCmd.AddCommand(StartCmd)
	aksCmd.AddCommand(StopCmd)
	aksCmd.AddCommand(RotateCertsCmd)
	aksCmd.AddCommand(GetOSoptionsCmd)
	aksCmd.AddCommand(MaintenanceconfigurationCmd)
	MaintenanceconfigurationCmd.AddCommand(MCAddCmd)
	MaintenanceconfigurationCmd.AddCommand(MCDeleteCmd)
	MaintenanceconfigurationCmd.AddCommand(MCUpdateCmd)
	MaintenanceconfigurationCmd.AddCommand(MCListCmd)
	MaintenanceconfigurationCmd.AddCommand(MCShowCmd)

	aksCmd.AddCommand(AKSAppUpCmd)
	aksCmd.AddCommand(AKSBrowseCmd)
	aksCmd.AddCommand(AKSCheckAcrCmd)
	aksCmd.AddCommand(AKSGetUpgradesCmd)
	aksCmd.AddCommand(AKSGetVersionsCmd)
	aksCmd.AddCommand(AKSKanalyzeCmd)
	aksCmd.AddCommand(AKSKollectCmd)
	aksCmd.AddCommand(Nodepool)
	Nodepool.AddCommand(AKSNodepoolGetUpgradesCmd)
	aksCmd.AddCommand(AKSInstallCLICmd)

	aksCmd.AddCommand(AKSConnectedK8sCmd)
	AKSConnectedK8sCmd.AddCommand(AKSConnectedk8sConnectCmd)
	AKSConnectedK8sCmd.AddCommand(AKSConnectedk8sDeleteCmd)
	AKSConnectedK8sCmd.AddCommand(AKSConnectedk8sDisableFeaturesCmd)
	AKSConnectedK8sCmd.AddCommand(AKSConnectedk8sEnableFeaturesCmd)
	AKSConnectedK8sCmd.AddCommand(AKSConnectedk8sListCmd)
	AKSConnectedK8sCmd.AddCommand(AKSConnectedk8sProxyCmd)
	AKSConnectedK8sCmd.AddCommand(AKSConnectedk8sShowCmd)
	AKSConnectedK8sCmd.AddCommand(AKSConnectedk8sUpdateCmd)
	AKSConnectedK8sCmd.AddCommand(AKSConnectedk8sUpgradeCmd)

	aksCmd.AddCommand(AKSk8sConfiguration)
	AKSk8sConfiguration.AddCommand(AKSConfigurationCreate)
	AKSk8sConfiguration.AddCommand(AKSConfigurationDelete)
	AKSk8sConfiguration.AddCommand(AKSConfigurationList)
	AKSk8sConfiguration.AddCommand(AKSConfigurationShow)
	aksFlags()
}
