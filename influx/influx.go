package influx

import (
	"fmt"

	"github.com/influxdata/influxdb/client/v2"
)

type Influx struct {
	inClient client.Client
}

func NewInflux(INFLUX_IP, INFLUX_PORT, username, password string) *Influx {

	inf := &Influx{
		inClient: InfluxDBClient(INFLUX_IP, INFLUX_PORT, username, password),
	}
	return inf
}

func InfluxDBClient(INFLUX_IP, INFLUX_PORT, username, password string) client.Client {

	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://" + INFLUX_IP + ":" + INFLUX_PORT,
		Username: username,
		Password: password,
	})
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return c
}

func (in *Influx) GetPodData(podNum int, ns, clusterName string) []client.Result {

	pod_Num := fmt.Sprint(podNum)
	q := client.NewQuery("SELECT * FROM Pods WHERE cluster = '"+clusterName+"' ORDER BY DESC LIMIT "+pod_Num+" ", "Metrics", "")

	response, err := in.inClient.Query(q)

	if err == nil && response.Error() == nil {
		return response.Results
	}
	return nil
}

func (in *Influx) GetNodeData(clusterName string) []client.Result {

	q := client.NewQuery("SELECT * FROM Nodes WHERE cluster = '"+clusterName+"' ORDER BY DESC LIMIT 3", "Metrics", "")

	response, err := in.inClient.Query(q)

	if err == nil && response.Error() == nil {
		return response.Results
	}
	return nil

}
