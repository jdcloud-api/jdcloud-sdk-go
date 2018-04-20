package demo

import (
	"fmt"
	"testing"

	. "github.com/jdcloud-api/jdcloud-sdk-go/core"
	. "github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/apis"
	. "github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/client"
	vpc "github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/models"
)

const (
	ELASTIC_IP_ID        = "fip-6j7ssgqr0o"
	NETWORK_INTERFACE_ID = "port-m2bbrmmaxz"
	SECONDARY_IP1        = "10.0.1.10"
	SECONDARY_IP2        = "10.0.1.11"
	SUBNET_ID            = "subnet-63ndensu4g"
	SECURITYGROUP_ID     = "sg-7lpiabk6lo"

	VPCPEERING_NAME       = "wwc_test_peering"
	VPCPEERING_NAME2      = "wwc_test_peering2"
	PEERING_LOCAL_VPC_ID  = "vpc-4utqaaxvm1"
	PEERING_REMOTE_VPC_ID = "vpc-9z7fy2dskf"
	VPC_ID                = "vpc-yzqu131my1"
	VPEPEERING_MOD_DEL_ID = "vpcpr-9otx9as779"
)

var vpcPeeringId = ""

func initVpcClient() *VpcClient {
	accessKey := "ak"
	secretKey := "sk"
	credentials := NewCredentials(accessKey, secretKey)

	return NewVpcClient(credentials)
}

func TestVpcGetElasticIps(t *testing.T) {
	client := initVpcClient()
	req := NewDescribeElasticIpsRequest("cn-north-1")
	req.SetPageNumber(1)
	req.SetPageSize(100)
	resp, err := client.DescribeElasticIps(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe elasticips failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.ElasticIps)
	fmt.Println(resp.Result.TotalCount)
	fmt.Println(len(resp.Result.ElasticIps))
}

func TestVpcGetElasticIp(t *testing.T) {
	client := initVpcClient()
	req := NewDescribeElasticIpRequest("cn-north-1", ELASTIC_IP_ID)
	resp, err := client.DescribeElasticIp(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe elasticip failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.ElasticIp)
}

func TestVpcCreateElasticIps(t *testing.T) {
	client := initVpcClient()
	bandWidth := 2
	provider := "bgp"
	spec := vpc.ElasticIpSpec{
		BandwidthMbps: bandWidth,
		Provider:      provider,
	}
	req := NewCreateElasticIpsRequest("cn-north-1", 1, &spec)
	resp, err := client.CreateElasticIps(req)
	if err != nil {
		t.Error(err.Error())
		fmt.Println("error")
	}

	if resp.Error.Code != 0 {
		t.Error(fmt.Sprintf("create elasticIp failed, code=%d, message=%s", resp.Error.Code, resp.Error.Message))
	}

	//fmt.Println(resp.Result.ElasticIpIds[0])
}
func TestVpcDeleteElasticIp(t *testing.T) {
	client := initVpcClient()
	req := NewDeleteElasticIpRequest("cn-north-1", "fip-i7wpd87kzu")
	resp, err := client.DeleteElasticIp(req)
	if err != nil {
		t.Error(err.Error())
		fmt.Println("error")
	}
	if resp.Error.Code != 0 {
		t.Error(fmt.Sprintf("delete elasticIp failed, code=%d, message=%s", resp.Error.Code, resp.Error.Message))
	}

}

func TestVpcAssociateElasticIp(t *testing.T) {
	client := initVpcClient()
	req := NewAssociateElasticIpRequest("cn-north-1", NETWORK_INTERFACE_ID)
	eipId := ELASTIC_IP_ID
	req.ElasticIpId = &eipId
	resp, err := client.AssociateElasticIp(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("associate elastic ip failed. ", err, resp.Error.Code, resp.Error.Message)
	} else {
		fmt.Println("AssociateElasticIp success")
	}

	fmt.Println(resp.RequestID)

}

func TestVpcDisassociateElasticIp(t *testing.T) {
	client := initVpcClient()
	req := NewDisassociateElasticIpRequest("cn-north-1", NETWORK_INTERFACE_ID)
	resp, err := client.DisassociateElasticIp(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("disassociate elastic ip failed. ", err, resp.Error.Code, resp.Error.Message)
	} else {
		fmt.Println("DisassociateElasticIp success")
	}
}

func TestVpcAssignSecondaryIps(t *testing.T) {
	client := initVpcClient()
	req := NewAssignSecondaryIpsRequest("cn-north-1", NETWORK_INTERFACE_ID)
	req.SecondaryIps = []string{SECONDARY_IP1, SECONDARY_IP2}
	resp, err := client.AssignSecondaryIps(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("assign secondary ips failed. ", err, resp.Error.Code, resp.Error.Message)
	} else {
		fmt.Println("AssignSecondaryIps success")
	}

}

func TestVpcUnassignSecondaryIps(t *testing.T) {
	client := initVpcClient()
	req := NewUnassignSecondaryIpsRequest("cn-north-1", NETWORK_INTERFACE_ID)
	req.SecondaryIps = []string{SECONDARY_IP1, SECONDARY_IP2}
	resp, err := client.UnassignSecondaryIps(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("unassign secondary ips failed. ", err, resp.Error.Code, resp.Error.Message)
	} else {
		fmt.Println("UnassignSecondaryIps success")
	}
}

func TestVpcDescribeSubnets(t *testing.T) {
	client := initVpcClient()
	req := NewDescribeSubnetsRequest("cn-north-1")
	req.SetPageNumber(1)
	req.SetPageSize(100)
	resp, err := client.DescribeSubnets(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe subnets failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.Subnets)
	fmt.Println(resp.Result.TotalCount)
	fmt.Println(len(resp.Result.Subnets))
}

func TestVpcDescribeSubnet(t *testing.T) {
	client := initVpcClient()
	req := NewDescribeSubnetRequest("cn-north-1", SUBNET_ID)
	resp, err := client.DescribeSubnet(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe subnet failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.Subnet)
}

func TestVpcDescribeNetworkSecurityGroups(t *testing.T) {
	client := initVpcClient()
	req := NewDescribeNetworkSecurityGroupsRequest("cn-north-1")
	req.SetPageNumber(1)
	req.SetPageSize(100)
	resp, err := client.DescribeNetworkSecurityGroups(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe network securitygroups failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.NetworkSecurityGroups)
	fmt.Println(resp.Result.TotalCount)
	fmt.Println(len(resp.Result.NetworkSecurityGroups))
}

func TestVpcDescribeNetworkSecurityGroup(t *testing.T) {
	client := initVpcClient()
	req := NewDescribeNetworkSecurityGroupRequest("cn-north-1", SECURITYGROUP_ID)
	resp, err := client.DescribeNetworkSecurityGroup(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe network securitygroup failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.NetworkSecurityGroup)
}

func TestVpcDescribeVpcPeerings(t *testing.T) {
	client := initVpcClient()
	req := NewDescribeVpcPeeringsRequest("cn-north-1")
	req.SetPageNumber(1)
	req.SetPageSize(100)
	resp, err := client.DescribeVpcPeerings(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe vpcpeerings failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.VpcPeerings)
	fmt.Println(resp.Result.TotalCount)
	fmt.Println(len(resp.Result.VpcPeerings))
}

func TestVpcCreateVpcPeering(t *testing.T) {
	client := initVpcClient()
	req := NewCreateVpcPeeringRequest("cn-north-1", VPCPEERING_NAME, PEERING_LOCAL_VPC_ID, PEERING_REMOTE_VPC_ID)
	resp, err := client.CreateVpcPeering(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("create vpcpeering failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.VpcPeering)
	vpcPeeringId = resp.Result.VpcPeering.VpcPeeringId
}

func TestVpcModifyVpcPeering(t *testing.T) {
	client := initVpcClient()
	req := NewModifyVpcPeeringRequest("cn-north-1", vpcPeeringId)
	req.SetVpcPeeringName(VPCPEERING_NAME2)
	resp, err := client.ModifyVpcPeering(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("modify vpcpeering failed. ", err, resp.Error.Code, resp.Error.Message, resp.RequestID)
	} else {
		fmt.Println("ModifyVpcPeering success")
	}
}

func TestVpcDescribeVpcPeering(t *testing.T) {
	client := initVpcClient()
	req := NewDescribeVpcPeeringRequest("cn-north-1", vpcPeeringId)
	resp, err := client.DescribeVpcPeering(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe vpcpeering failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.VpcPeering)
}

func TestVpcDeleteVpcPeering(t *testing.T) {
	client := initVpcClient()
	req := NewDeleteVpcPeeringRequest("cn-north-1", vpcPeeringId)
	resp, err := client.DeleteVpcPeering(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("delete vpcpeering failed. ", err, resp.Error.Code, resp.Error.Message)
	} else {
		fmt.Println("DeleteVpcPeering success")
	}
}

func TestVpcDescribeVpcs(t *testing.T) {
	client := initVpcClient()
	req := NewDescribeVpcsRequest("cn-north-1")
	req.SetPageNumber(1)
	req.SetPageSize(100)
	resp, err := client.DescribeVpcs(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe vpcs failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.Vpcs)
	fmt.Println(resp.Result.TotalCount)
	fmt.Println(len(resp.Result.Vpcs))
}

func TestVpcDescribeVpc(t *testing.T) {
	client := initVpcClient()
	req := NewDescribeVpcRequest("cn-north-1", VPC_ID)
	resp, err := client.DescribeVpc(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe vpc failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.Vpc)
}
