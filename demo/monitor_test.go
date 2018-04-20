package demo

import (
	"fmt"
	"testing"
	. "github.com/jdcloud-api/jdcloud-sdk-go/core"
	. "github.com/jdcloud-api/jdcloud-sdk-go/services/monitor/client"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/monitor/apis"
)

func initMonitorClient() *MonitorClient {
	accessKey := "ak"
	secretKey := "sk"
	credentials := NewCredentials(accessKey, secretKey)

	return NewMonitorClient(credentials)
}

func TestDescribeMetrics(t *testing.T) {
	req := apis.NewDescribeMetricsRequest("vm")
	client := initMonitorClient()
	resp, err := client.DescribeMetrics(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("monitor describe metrics failed.", err, resp.Error.Message)
	}
	fmt.Println(resp.Result.Metrics)
}
