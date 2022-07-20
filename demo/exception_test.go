package demo

import (
	"fmt"
	"testing"
	. "github.com/lshuining/jdcloud-sdk-go/core"
	. "github.com/lshuining/jdcloud-sdk-go/services/monitor/client"
	"github.com/lshuining/jdcloud-sdk-go/services/monitor/apis"
)

func TestCredentialNotCorrect(t *testing.T) {
	credentials := NewCredentials("", "")
	client := NewMonitorClient(credentials)
	req := apis.NewDescribeMetricsRequest("vm")
	resp, err := client.DescribeMetrics(req)
	if err != nil {
		t.Error("monitor describe vm metrics, err should nil. ", err.Error())
	}
	if resp.Error.Code != 403 {
		t.Error("when credential is nil, should return 403 error.")
	}
	fmt.Println(resp.Result.Metrics)
}

func TestCredentialNil(t *testing.T) {
	client := NewMonitorClient(nil)
	if client != nil {
		t.Error("when credential is nil, client is nil")
	}
}

func TestLackOfScheme(t *testing.T) {
	credentials := NewCredentials("", "")
	config := NewConfig()
	config.SetScheme("")

	client := NewMonitorClient(credentials)
	client.SetConfig(config)

	req := apis.NewDescribeMetricsRequest("vm")
	_, err := client.DescribeMetrics(req)
	if err == nil {
		t.Error("when scheme not correct, client is nil")
	}
}

func TestLackOfEndpoint(t *testing.T) {
	credentials := NewCredentials("", "")
	config := NewConfig()
	config.SetEndpoint("")

	client := NewMonitorClient(credentials)
	client.SetConfig(config)

	req := apis.NewDescribeMetricsRequest("vm")
	_, err := client.DescribeMetrics(req)
	if err == nil {
		t.Error("if endpoint is empty string, should return error")
	}
}

func TestSendNilRequest(t *testing.T) {
	credentials := NewCredentials("", "")
	client := NewMonitorClient(credentials)
	_, err := client.DescribeMetrics(nil)
	if err == nil {
		t.Error("request is nil, should return error")
	}
}
