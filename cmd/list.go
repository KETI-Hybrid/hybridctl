// Copyright © 2022 NAME HERE <EMAIL ADDRESS>
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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"

	"github.com/KETI-Hybrid/hybridctl-v1/pkg/nks"

	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"cloud.google.com/go/pubsub"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"

	"k8s.io/klog"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print Kubernetes engine resource list",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("list called")
	},
}

// func gke_getBearer() string {
// 	ctx := context.Background()

// 	conf := &oauth2.Config{
// 		ClientID:     "235318948778-2hmm3cvglu5jrdgpab017m4f710l2iq9.apps.googleusercontent.com",
// 		ClientSecret: "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",
// 		Scopes:       []string{"SCOPE1", "SCOPE2"},
// 		Endpoint: oauth2.Endpoint{
// 			TokenURL: "https://provider.com/o/oauth2/token",
// 			AuthURL:  "https://provider.com/o/oauth2/auth",
// 		},
// 	}

// 	// Redirect user to consent page to ask for permission
// 	// for the scopes specified above.
// 	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
// 	fmt.Printf("Visit the URL for the auth dialog: %v", url)

// 	// Use the authorization code that is pushed to the redirect
// 	// URL. Exchange will do the handshake to retrieve the
// 	// initial access token. The HTTP Client returned by
// 	// conf.Client will refresh the token as necessary.
// 	var code string
// 	if _, err := fmt.Scan(&code); err != nil {
// 		log.Fatal(err)
// 	}

// 	// Use the custom HTTP client when requesting a token.
// 	httpClient := &http.Client{Timeout: 2 * time.Second}
// 	ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)

// 	tok, err := conf.Exchange(ctx, code)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	client := conf.Client(ctx, tok)
// 	_ = client

// 	return tok.AccessToken
// }

// oauthClient shows how to use an OAuth client ID to authenticate as an end-user.
func oauthClient() string {
	ctx := context.Background()

	// Please make sure the redirect URL is the same as the one you specified when you
	// created the client ID.
	redirectURL := os.Getenv("OAUTH2_CALLBACK")
	if redirectURL == "" {
		redirectURL = "https://example.com"
	}
	config := &oauth2.Config{
		ClientID:     "235318948778-2hmm3cvglu5jrdgpab017m4f710l2iq9.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  redirectURL,
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}

	// Dummy authorization flow to read auth code from stdin.
	authURL := config.AuthCodeURL("your state")
	fmt.Printf("Follow the link in your browser to obtain auth code: %s", authURL)

	//testing

	//
	// Read the authentication code from the command line
	var code string

	if _, err := fmt.Scan(&code); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}
	// code = fs.ReadLine()
	// code = "4/0AX4XfWi-Ubk09zqzB5uUTPJjDN-V1MafN3I_uNjb-nfz2FpjvlyZ_4PXl1PgwmLRC8POuw"

	// Exchange auth code for OAuth token.
	token, err := config.Exchange(ctx, code)
	if err != nil {
		fmt.Printf("config.Exchange: %v", err)
	}

	client, err := pubsub.NewClient(ctx, "keti-container", option.WithTokenSource(config.TokenSource(ctx, token)))
	if err != nil {
		fmt.Printf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	// Use the authenticated client.
	_ = client

	return token.AccessToken
}

var listClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("list cluster called")
		platform_name, err := cmd.Flags().GetString("platform")
		if err != nil {
			panic(err)
		}
		var cli = Cli{}

		cli.PlatformName = platform_name
		fmt.Println(cli.PlatformName)
		if cli.PlatformName == "gke" {

			cmd := exec.Command("gcloud", "container", "clusters", "list", "--region", "us-central1-a", "--uri")
			output, err := cmd.Output()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(string(output))
			}

			// fmt.Println("kubernetes engine Name :", cli.PlatformName)

			// URL := os.Getenv("GKE_URL")

			// request, _ := http.NewRequest("GET", URL, nil)

			// request.Header.Add("Authorization", "Bearer"+oauthClient())

			// client := &http.Client{}
			// response, err := client.Do(request)
			// if err != nil {
			// 	panic(err)
			// }

			// // fmt.Println(resp)
			// // 결과 출력

			// defer response.Body.Close()
			// bytes, _ := ioutil.ReadAll(response.Body)
			// str := string(bytes) //바이트를 문자열로
			// // string to struck 과정 필요
			// fmt.Println("str: ", str)

			// // data := AKS_Cluster_API{} //json to struck
			// // if err := json.Unmarshal([]byte(str), &data); err != nil {
			// // 	panic(err)
			// // }

		} else if cli.PlatformName == "eks" {

			fmt.Println("kubernetes engine Name : ", cli.PlatformName)

			sess, err := session.NewSession(&aws.Config{
				Region: aws.String("us-east-2")},
			)
			if err != nil {
				fmt.Print(err)
			}

			svc := eks.New(sess)
			input := &eks.ListClustersInput{}

			result, err := svc.ListClusters(input)
			if err != nil {
				if aerr, ok := err.(awserr.Error); ok {
					switch aerr.Code() {
					case eks.ErrCodeInvalidParameterException:
						fmt.Println(eks.ErrCodeInvalidParameterException, aerr.Error())
					case eks.ErrCodeClientException:
						fmt.Println(eks.ErrCodeClientException, aerr.Error())
					case eks.ErrCodeServerException:
						fmt.Println(eks.ErrCodeServerException, aerr.Error())
					case eks.ErrCodeServiceUnavailableException:
						fmt.Println(eks.ErrCodeServiceUnavailableException, aerr.Error())
					default:
						fmt.Println(aerr.Error())
					}
				} else {
					// Print the error, cast err to awserr.Error to get the Code and
					// Message from an error.
					fmt.Println(err.Error())
				}
				return
			}

			fmt.Println("---Show eks Cluster list---")
			for i := 0; i < len(result.Clusters); i++ {
				fmt.Println("[", i+1, "]", *result.Clusters[i]) // *string 형태로 eks cluster list 출력
			}

		} else if cli.PlatformName == "aks" {

			// var bearer = "Bearer " + "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsIng1dCI6ImpTMVhvMU9XRGpfNTJ2YndHTmd2UU8yVnpNYyIsImtpZCI6ImpTMVhvMU9XRGpfNTJ2YndHTmd2UU8yVnpNYyJ9.eyJhdWQiOiJodHRwczovL21hbmFnZW1lbnQuY29yZS53aW5kb3dzLm5ldCIsImlzcyI6Imh0dHBzOi8vc3RzLndpbmRvd3MubmV0L2M4ZWE5MWI1LTZhYWMtNGM1Yy1hZTM0LTk3MTdhODcyMTU5Zi8iLCJpYXQiOjE2NDk3Mzk0OTgsIm5iZiI6MTY0OTczOTQ5OCwiZXhwIjoxNjQ5NzQzNzUyLCJhY3IiOiIxIiwiYWlvIjoiQVRRQXkvOFRBQUFBc0djaThYRThkT1hFZDJwQUljVUk4V3JtaWJUMkJMdHZRcTB6ZEVlUWNHQ0pwY05wbFYxUkk3MERyR1FtZXBYdCIsImFtciI6WyJwd2QiXSwiYXBwaWQiOiIxOGZiY2ExNi0yMjI0LTQ1ZjYtODViMC1mN2JmMmIzOWIzZjMiLCJhcHBpZGFjciI6IjAiLCJncm91cHMiOlsiZDljMDdmMTYtMDQ1My00MzIzLWJmNjMtZjAyYzVhOGI3YzQ2IiwiM2FiZDliMGQtY2M1Ni00MGYxLThjMjItODZiNTNhMWZmMzJkIiwiZTAxOWI2NGItY2ZkZS00ZDkzLTg2MTAtNThmM2MzZGJkNmE1IiwiNzFjODE0ZDMtYjBiNy00NzcyLTg0MjktMzI1NmRkYWNiZjUyIiwiMjc4NGFhNTAtODlhOS00OTY5LWFlMzMtNWM1NzI2Mzc2MGUwIiwiMzllYTI4MzUtYmYwZi00YWQyLThiOTMtMWY3ODJhZDU5ZGUxIl0sImlwYWRkciI6IjExNS45NC4xNDEuNjIiLCJuYW1lIjoiY29yZWh1biIsIm9pZCI6ImZiODdkNDJiLTljMDItNGU5MS1hYjVkLTIwMjAxNDQzNzQ2NiIsInB1aWQiOiIxMDAzMjAwMTVCMkYyMEEwIiwicmgiOiIwLkFVb0F0WkhxeUt4cVhFeXVOSmNYcUhJVm4wWklmM2tBdXRkUHVrUGF3ZmoyTUJOS0FPQS4iLCJzY3AiOiJ1c2VyX2ltcGVyc29uYXRpb24iLCJzdWIiOiJnUXRjRkNfSDkxQWI1clYxVWF0Mnc0TFJyS0VKQlZXdktoRHJ6U1pKMTRvIiwidGlkIjoiYzhlYTkxYjUtNmFhYy00YzVjLWFlMzQtOTcxN2E4NzIxNTlmIiwidW5pcXVlX25hbWUiOiJjb3JlaHVuQHNwdGVrY2xvdWQub25taWNyb3NvZnQuY29tIiwidXBuIjoiY29yZWh1bkBzcHRla2Nsb3VkLm9ubWljcm9zb2Z0LmNvbSIsInV0aSI6IlFiVVItYm5SSUV5bmtoQ2x3cjBkQUEiLCJ2ZXIiOiIxLjAiLCJ3aWRzIjpbIjc2OThhNzcyLTc4N2ItNGFjOC05MDFmLTYwZDZiMDhhZmZkMiIsIjliODk1ZDkyLTJjZDMtNDRjNy05ZDAyLWE2YWMyZDVlYTVjMyIsIjkzNjBmZWI1LWY0MTgtNGJhYS04MTc1LWUyYTAwYmFjNDMwMSIsIjg4ZDhlM2UzLThmNTUtNGExZS05NTNhLTliOTg5OGI4ODc2YiIsIjE1OGMwNDdhLWM5MDctNDU1Ni1iN2VmLTQ0NjU1MWE2YjVmNyIsImNmMWMzOGU1LTM2MjEtNDAwNC1hN2NiLTg3OTYyNGRjZWQ3YyIsImI3OWZiZjRkLTNlZjktNDY4OS04MTQzLTc2YjE5NGU4NTUwOSJdLCJ4bXNfdGNkdCI6MTU0MzU0MTQ2NH0.qC11gd1hHgnGczpUhu4aDCsqYDKwgaYiUZ8FLB2L0yDze5Z_GeMTu2dZAZo93YSUV5cC2SM9uWpmlHNtpswa-LpzjuHYPM34i25myJqusP8Jj6iImGh67NIstNcYgc_5uvhF6xsY82Gi-_PHiWYSmc5_ZWDIUhcclGqdkxQTbHAOjwrQGUiVW7Wnk68HhkuwZNreW7T1rtUL_KkGNjnNn3S0oO4NtH6fpy4NbkOTQoSDhRhwlrt9KMDf411ZW0W_10Tn5PzvJlSUMhcaP_zWdpWYVLC_Yby9ITjyo6_r6Q5u0EwYIVfsuIoFhhQjyiXOmYmfY34ASYp3OlD0P2sSsw"
			fmt.Println("kubernetes engine Name : ", cli.PlatformName)
			// resp, err := http.NewRequest("GET", "https://management.azure.com/subscriptions/ccfc0c6c-d3c6-4de2-9a6c-c09ca498ff73/providers/Microsoft.ContainerService/managedClusters?api-version=2022-02-01", nil)
			// if err != nil {
			// 	panic(err)
			// }

			URL := os.Getenv("AKS_URL")

			resp, err := AuthorizationAndHTTP("GET", URL, nil)
			if err != nil {
				fmt.Println(err)
			}

			// fmt.Println("!")
			// fmt.Println(resp)
			// 결과 출력

			// defer resp.Body.Close()
			bytes, _ := ioutil.ReadAll(resp.Body)
			str := string(bytes) //바이트를 문자열로
			// string to struck 과정 필요
			fmt.Println("str: ", str)

			data := AKS_Cluster_API{} //json to struck
			if err := json.Unmarshal([]byte(str), &data); err != nil {
				panic(err)
			}

			fmt.Println("--- Show aks Cluster list ---")

			// 출력

			// fmt.Println(data.Value)
			for i := 0; i < len(data.Value); i++ {
				fmt.Println("[", i+1, "]", data.Value[i].Name)
			}

		} else if cli.PlatformName == "nks" {
			nks.NksgetClusterlist()
		} else {
			fmt.Println("Error : please enter your flag")
		}

	},
}

var listNodeCmd = &cobra.Command{
	Use:   "node",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("list node called")
		platform_name, err := cmd.Flags().GetString("platform")
		if err != nil {
			panic(err)
		}
		cluster_name, err := cmd.Flags().GetString("cluster-name")
		if err != nil {
			panic(err)
		}
		var cli = Cli{}

		cli.PlatformName = platform_name
		fmt.Println(cli.PlatformName)
		cli.ClusterName = cluster_name
		if cli.PlatformName == "eks" {

			sess, err := session.NewSession(&aws.Config{
				Region: aws.String("us-east-2")},
			)
			if err != nil {
				fmt.Print(err)
			}

			svc := eks.New(sess)

			input := &eks.ListNodegroupsInput{
				ClusterName: &args[1],
			}

			result, err := svc.ListNodegroups(input)
			if err != nil {
				if aerr, ok := err.(awserr.Error); ok {
					switch aerr.Code() {
					case eks.ErrCodeInvalidParameterException:
						fmt.Println(eks.ErrCodeInvalidParameterException, aerr.Error())
					case eks.ErrCodeClientException:
						fmt.Println(eks.ErrCodeClientException, aerr.Error())
					case eks.ErrCodeServerException:
						fmt.Println(eks.ErrCodeServerException, aerr.Error())
					case eks.ErrCodeServiceUnavailableException:
						fmt.Println(eks.ErrCodeServiceUnavailableException, aerr.Error())
					default:
						fmt.Println(aerr.Error())
					}
				} else {
					// Print the error, cast err to awserr.Error to get the Code and
					// Message from an error.
					fmt.Println(err.Error())
				}
				return
			}
			fmt.Println("--- Show eks Nodegroups list ---")

			fmt.Println(result)
			for i := 0; i < len(result.Nodegroups); i++ {
				fmt.Println("[", i+1, "]", result.Nodegroups[i])
			}

		} else if cli.PlatformName == "gke" {
			cmd := exec.Command("gcloud", "container", "node-pools", "list", "--region", "us-central1-a", "--cluster", cli.ClusterName, "--uri")
			output, err := cmd.Output()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(string(output))
			}
		} else if cli.PlatformName == "aks" {

			fmt.Println("kubernetes engine Name : ", cli.PlatformName)

			URL := os.Getenv("AKS_URL")

			resp, _ := AuthorizationAndHTTP("GET", URL, nil)

			defer resp.Body.Close()

			// 결과 출력
			bytes, _ := ioutil.ReadAll(resp.Body)
			str := string(bytes) //바이트를 문자열로
			// string to struck 과정 필요

			data := AKS_Cluster_API{}
			if err := json.Unmarshal([]byte(str), &data); err != nil {
				panic(err)
			}

			for i := 0; i < len(data.Value); i++ {
				if data.Value[i].Name == cli.ClusterName {
					fmt.Println("---", data.Value[i].Name, "Cluster's nodepool list ---")
					for j := 0; j < len(data.Value[i].Properties.AgentPoolProfiles); j++ {
						fmt.Println("[", j+1, "]", data.Value[i].Properties.AgentPoolProfiles[j].Name)
					}
				}
			}

		}
	},
}

func AuthorizationAndHTTP(method string, hosturl string, input interface{}) (*http.Response, error) {

	var request *http.Request
	var err error
	fmt.Println(GetBearer().Access_token)
	switch method {
	case "POST":
		params := url.Values{}
		params.Add("resource", `https://management.azure.com/`)
		body := strings.NewReader(params.Encode())
		request, _ = http.NewRequest(method, hosturl, body)
		break
	case "GET":
		request, _ = http.NewRequest(method, hosturl, nil)
		break
	case "DELETE":
		request, _ = http.NewRequest(method, hosturl, nil)
		break
	case "PUT":
		jsonData, _ := json.Marshal(input)
		request, _ = http.NewRequest(method, hosturl, bytes.NewBuffer(jsonData))
		break
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Add("Authorization", "Bearer "+GetBearer().Access_token)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("err2 : ", err)
	} else {
		fmt.Println(response.Status)
	}
	return response, err
}

func GetBearer() bearerToken {

	params := url.Values{}
	params.Add("client_id", os.Getenv("ClientId"))
	params.Add("grant_type", `client_credentials`)
	params.Add("resource", `https://management.azure.com/`)
	params.Add("client_secret", os.Getenv("ClientSecret"))
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("POST", "https://login.microsoftonline.com/"+os.Getenv("TenantId")+"/oauth2/token", body)
	if err != nil {
		klog.Error(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		klog.Error(err)
	}
	defer response.Body.Close()
	bytes, _ := ioutil.ReadAll(response.Body)
	token := bearerToken{}
	json.Unmarshal(bytes, &token)

	return token
}

func init() {
	RootCmd.AddCommand(listCmd)
	listCmd.AddCommand(listClusterCmd)
	listCmd.AddCommand(listNodeCmd)

	listClusterCmd.Flags().String("platform", "", "input your platform name")

	listNodeCmd.Flags().String("platform", "", "input your platform name")
	listNodeCmd.Flags().String("cluster-name", "", "input your cluster name")

	os.Setenv("AKS_URL", "https://management.azure.com/subscriptions/ccfc0c6c-d3c6-4de2-9a6c-c09ca498ff73/providers/Microsoft.ContainerService/managedClusters?api-version=2021-05-01")
	os.Setenv("ClientId", "9bd853b8-04d5-4c24-8d07-f0e7cd0a680e")
	os.Setenv("ClientSecret", "Q9r7Q~9ypO-BmymfM8vHOq22tjW6kJaRHz~4y")
	os.Setenv("SubscriptionId", "ccfc0c6c-d3c6-4de2-9a6c-c09ca498ff73")
	os.Setenv("TenantId", "c8ea91b5-6aac-4c5c-ae34-9717a872159f")

	os.Setenv("GKE_URL", "https://container.googleapis.com/v1beta1/projects/keti-container/locations/us-central1-a/clusters")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
