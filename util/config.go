package util

import (
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type KubeConfig struct {
	APIVersion string `yaml:"apiVersion"`
	Clusters   []struct {
		Cluster struct {
			CertificateAuthorityData string `yaml:"certificate-authority-data"`
			Server                   string `yaml:"server"`
		} `yaml:"cluster"`
		Name string `yaml:"name"`
	} `yaml:"clusters"`
	Contexts []struct {
		Context struct {
			Cluster string `yaml:"cluster"`
			User    string `yaml:"user"`
		} `yaml:"context"`
		Name string `yaml:"name"`
	} `yaml:"contexts"`
	CurrentContext string `yaml:"current-context"`
	Kind           string `yaml:"kind"`
	Preferences    struct {
	} `yaml:"preferences"`
	Users []struct {
		Name string `yaml:"name"`
		User struct {
			ClientCertificateData string `yaml:"client-certificate-data,omitempty"`
			ClientKeyData         string `yaml:"client-key-data,omitempty"`
			Token                 string `yaml:"token,omitempty"`
			AuthProvider          struct {
				Config struct {
					AccessToken string `yaml:"access-token,omitempty"`
					CmdArgs     string `yaml:"cmd-args,omitempty"`
					CmdPath     string `yaml:"cmd-path,omitempty"`
					Expiry      string `yaml:"expiry,omitempty"`
					ExpiryKey   string `yaml:"expiry-key,omitempty"`
					TokenKey    string `yaml:"token-key,omitempty"`
				} `yaml:"config,omitempty"`
				Name string `yaml:"name,omitempty"`
			} `yaml:"auth-provider,omitempty"`
			Exec struct {
				APIVersion string      `yaml:"apiVersion,omitempty"`
				Args       []string    `yaml:"args,omitempty"`
				Command    string      `yaml:"command,omitempty"`
				Env        interface{} `yaml:"env,omitempty"`
			} `yaml:"exec,omitempty"`
		} `yaml:"user"`
	} `yaml:"users"`
}

func BuildConfigFromFlags(context, kubeconfigPath string) (*rest.Config, error) {
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfigPath},
		&clientcmd.ConfigOverrides{
			CurrentContext: context,
		}).ClientConfig()
}
