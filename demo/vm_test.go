package demo

import (
	"time"
	"fmt"
	"testing"
	. "github.com/jdcloud-api/jdcloud-sdk-go/services/vm/apis"
	. "github.com/jdcloud-api/jdcloud-sdk-go/services/vm/client"
	. "github.com/jdcloud-api/jdcloud-sdk-go/core"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
	vm "github.com/jdcloud-api/jdcloud-sdk-go/services/vm/models"
	vpc "github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/models"
)

func initVmClient() *VmClient {
	accessKey := "ak"
	secretKey := "sk"
	credentials := NewCredentials(accessKey, secretKey)

	return NewVmClient(credentials)
}

func TestVmGetInstances(t *testing.T) {
	client := initVmClient()
	req := NewDescribeInstancesRequest("cn-north-1")
	resp, err := client.DescribeInstances(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe instance failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.Instances)
	fmt.Println(resp.Result.TotalCount)
	fmt.Println(len(resp.Result.Instances))
}

func TestVmGetInstancesByPage(t *testing.T) {
	client := initVmClient()
	req := NewDescribeInstancesRequest("cn-north-1")
	req.SetPageNumber(2)
	req.SetPageSize(10)
	resp, err := client.DescribeInstances(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe instance failed. ", err, resp.Error.Code, resp.Error.Message)
	}

	if len(resp.Result.Instances) != 10 {
		t.Error("describe instances by paging failed.", len(resp.Result.Instances))
	}
	fmt.Println(resp.Result.Instances)
	fmt.Println(resp.Result.TotalCount)
	fmt.Println(len(resp.Result.Instances))
}

func TestVmGetInstanceTypes(t *testing.T) {
	client := initVmClient()
	req := NewDescribeInstanceTypesRequest("cn-north-1")
	resp, err := client.DescribeInstanceTypes(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe instance types failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.InstanceTypes)
}

func TestVmCreateInstance(t *testing.T) {
	instanceId := createInstance(t)
	testVmGetInstancesWithFilter(t, instanceId)
	try(10, 5*time.Second, stopInstance, t, instanceId)
	try(10, 5*time.Second, deleteInstance, t, instanceId)
}

func createInstance(t *testing.T) string {
	client := initVmClient()

	az := "cn-north-1a"
	instancetype := "g.s1.micro"
	imageid := "98d44a0f-88c1-451a-8971-f1f769073b6c"
	description := "golang-sdk"
	diskCategory := "local"
	spec := vm.InstanceSpec{
		Az:           &az,
		InstanceType: &instancetype,
		ImageId:      &imageid,
		Name:         "golang-sdk-test",
		PrimaryNetworkInterface: &vm.InstanceNetworkInterfaceAttachmentSpec{
			NetworkInterface: &vpc.NetworkInterfaceSpec{SubnetId: "subnet-3dm13k30gh", Az: &az},
		},
		SystemDisk:  &vm.InstanceDiskAttachmentSpec{DiskCategory: &diskCategory},
		Description: &description,
	}
	req := NewCreateInstancesRequest("cn-north-1")
	req.SetInstanceSpec(&spec)
	req.SetMaxCount(1)
	resp, err := client.CreateInstances(req)
	if err != nil {
		t.Error(err.Error())
		return ""
	}

	if resp.Error.Code != 0 {
		t.Error(fmt.Sprintf("create instance failed, code=%d, message=%s", resp.Error.Code, resp.Error.Message))
	}

	return resp.Result.InstanceIds[0]
}

func testVmGetInstancesWithFilter(t *testing.T, instanceId string) {
	client := initVmClient()
	op := "eq"
	filter := models.Filter{Name: "instanceId", Operator: &op, Values: []string{instanceId}}
	req := NewDescribeInstancesRequest("cn-north-1")
	req.SetFilters([]models.Filter{filter})
	resp, err := client.DescribeInstances(req)
	if err != nil || resp.Result.TotalCount != 1 {
		t.Error("describe instance with filter failed. ", err, resp.Error.Message)
	}
	fmt.Println(resp.Result.Instances)
}

func stopInstance(instanceId string) bool {
	req := NewStopInstanceRequest("cn-north-1", instanceId)
	client := initVmClient()

	config := NewConfig()
	config.SetEndpoint("vm.internal.cn-north-1.jdcloud-api.com") //指定非默认访问点，如供VPC专用调用的域名
	config.SetScheme(SchemeHttp) //设置使用HTTP而不是HTTPS，vpc专用域名不支持HTTPS
	config.SetTimeout(20 * time.Second) //设置请求超时
	client.SetConfig(config)

	client.DisableLogger() //关闭日志打印

	resp, err := client.StopInstance(req)
	return err == nil && resp.Error.Code == 0
}

func deleteInstance(instanceId string) bool {
	req := NewDeleteInstanceRequest("cn-north-1", instanceId)
	client := initVmClient()
	resp, err := client.DeleteInstance(req)
	return err == nil && resp.Error.Code == 0
}

func try(count int, wait time.Duration, op func(instanceId string) bool, t *testing.T, instanceId string) {
	for count > 0 {
		result := op(instanceId)
		if result {
			return
		}

		time.Sleep(wait)
		count--
		fmt.Println("Could try times: ", count)
	}
}
