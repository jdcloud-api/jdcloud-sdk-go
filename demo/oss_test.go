package demo

import (
	"fmt"
	"testing"
	. "github.com/lshuining/jdcloud-sdk-go/core"
	"github.com/lshuining/jdcloud-sdk-go/services/oss/client"
	"github.com/lshuining/jdcloud-sdk-go/services/oss/apis"
)

func initOssClient() *client.OssClient {
	accessKey := "ak"
	secretKey := "sk"
	credentials := NewCredentials(accessKey, secretKey)

	client := client.NewOssClient(credentials)
	config := NewConfig()
	config.SetEndpoint("192.168.180.18")
	config.SetScheme(SchemeHttp)
	client.SetConfig(config)
	return client
}

func TestListBuckets(t *testing.T) {
	req := apis.NewListBucketsRequest("cn-north-1")
	client := initOssClient()
	resp, err := client.ListBuckets(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("oss list bucket failed.", err, resp.Error)
	}

	fmt.Println(resp.RequestID)
	fmt.Println(resp.Result.Buckets)
}

func TestHeadBucket(t *testing.T) {
	req := apis.NewHeadBucketRequest("cn-north-1", "bibo")
	client := initOssClient()
	resp, err := client.HeadBucket(req)
	if err != nil {
		t.Error("oss head bucket failed.", err)
	} else {
		fmt.Println(resp.RequestID)
	}
}
