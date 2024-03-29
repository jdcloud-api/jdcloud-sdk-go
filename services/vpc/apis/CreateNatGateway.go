// Copyright 2018 JDCLOUD.COM
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
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    vpc "github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/models"
    charge "github.com/jdcloud-api/jdcloud-sdk-go/services/charge/models"
)

type CreateNatGatewayRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* NAT网关名称  */
    NatGatewayName string `json:"natGatewayName"`

    /* NAT网关规格，取值small（100万并发连接数），medium（300万并发连接数），large（1000万并发连接数），默认small (Optional) */
    NatGatewaySpec *string `json:"natGatewaySpec"`

    /* 私有网络ID  */
    VpcId string `json:"vpcId"`

    /* 子网ID  */
    SubnetId string `json:"subnetId"`

    /* NAT网关的可用区属性，即将废弃 (Optional) */
    AzIpSpecs []vpc.AzIpSpec `json:"azIpSpecs"`

    /* NAT网关可用区 (Optional) */
    Azs []string `json:"azs"`

    /* 选择已有公网IP列表。选择已有和新购公网IP可以同时配置，也可以配置其一 (Optional) */
    ElasticIpIds []string `json:"elasticIpIds"`

    /* 新购公网IP数量 (Optional) */
    ElasticIpCount *int `json:"elasticIpCount"`

    /* 新购公网IP配置。NAT网关仅支持打包创建标准公网IP，不支持边缘公网IP。且标准公网IP仅支持按配置、按用量两种计费模式。 (Optional) */
    ElasticIpSpec *vpc.ElasticIpSpec `json:"elasticIpSpec"`

    /* 计费配置，仅支持按配置，默认按配置 (Optional) */
    NatGatewayCharge *charge.ChargeSpec `json:"natGatewayCharge"`

    /* 描述,​ 允许输入UTF-8编码下的全部字符，不超过256字符 (Optional) */
    Description *string `json:"description"`
}

/*
 * param regionId: Region ID (Required)
 * param natGatewayName: NAT网关名称 (Required)
 * param vpcId: 私有网络ID (Required)
 * param subnetId: 子网ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateNatGatewayRequest(
    regionId string,
    natGatewayName string,
    vpcId string,
    subnetId string,
) *CreateNatGatewayRequest {

	return &CreateNatGatewayRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/natGateways/",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        NatGatewayName: natGatewayName,
        VpcId: vpcId,
        SubnetId: subnetId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param natGatewayName: NAT网关名称 (Required)
 * param natGatewaySpec: NAT网关规格，取值small（100万并发连接数），medium（300万并发连接数），large（1000万并发连接数），默认small (Optional)
 * param vpcId: 私有网络ID (Required)
 * param subnetId: 子网ID (Required)
 * param azIpSpecs: NAT网关的可用区属性，即将废弃 (Optional)
 * param azs: NAT网关可用区 (Optional)
 * param elasticIpIds: 选择已有公网IP列表。选择已有和新购公网IP可以同时配置，也可以配置其一 (Optional)
 * param elasticIpCount: 新购公网IP数量 (Optional)
 * param elasticIpSpec: 新购公网IP配置。NAT网关仅支持打包创建标准公网IP，不支持边缘公网IP。且标准公网IP仅支持按配置、按用量两种计费模式。 (Optional)
 * param natGatewayCharge: 计费配置，仅支持按配置，默认按配置 (Optional)
 * param description: 描述,​ 允许输入UTF-8编码下的全部字符，不超过256字符 (Optional)
 */
func NewCreateNatGatewayRequestWithAllParams(
    regionId string,
    natGatewayName string,
    natGatewaySpec *string,
    vpcId string,
    subnetId string,
    azIpSpecs []vpc.AzIpSpec,
    azs []string,
    elasticIpIds []string,
    elasticIpCount *int,
    elasticIpSpec *vpc.ElasticIpSpec,
    natGatewayCharge *charge.ChargeSpec,
    description *string,
) *CreateNatGatewayRequest {

    return &CreateNatGatewayRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/natGateways/",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        NatGatewayName: natGatewayName,
        NatGatewaySpec: natGatewaySpec,
        VpcId: vpcId,
        SubnetId: subnetId,
        AzIpSpecs: azIpSpecs,
        Azs: azs,
        ElasticIpIds: elasticIpIds,
        ElasticIpCount: elasticIpCount,
        ElasticIpSpec: elasticIpSpec,
        NatGatewayCharge: natGatewayCharge,
        Description: description,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateNatGatewayRequestWithoutParam() *CreateNatGatewayRequest {

    return &CreateNatGatewayRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/natGateways/",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CreateNatGatewayRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param natGatewayName: NAT网关名称(Required) */
func (r *CreateNatGatewayRequest) SetNatGatewayName(natGatewayName string) {
    r.NatGatewayName = natGatewayName
}
/* param natGatewaySpec: NAT网关规格，取值small（100万并发连接数），medium（300万并发连接数），large（1000万并发连接数），默认small(Optional) */
func (r *CreateNatGatewayRequest) SetNatGatewaySpec(natGatewaySpec string) {
    r.NatGatewaySpec = &natGatewaySpec
}
/* param vpcId: 私有网络ID(Required) */
func (r *CreateNatGatewayRequest) SetVpcId(vpcId string) {
    r.VpcId = vpcId
}
/* param subnetId: 子网ID(Required) */
func (r *CreateNatGatewayRequest) SetSubnetId(subnetId string) {
    r.SubnetId = subnetId
}
/* param azIpSpecs: NAT网关的可用区属性，即将废弃(Optional) */
func (r *CreateNatGatewayRequest) SetAzIpSpecs(azIpSpecs []vpc.AzIpSpec) {
    r.AzIpSpecs = azIpSpecs
}
/* param azs: NAT网关可用区(Optional) */
func (r *CreateNatGatewayRequest) SetAzs(azs []string) {
    r.Azs = azs
}
/* param elasticIpIds: 选择已有公网IP列表。选择已有和新购公网IP可以同时配置，也可以配置其一(Optional) */
func (r *CreateNatGatewayRequest) SetElasticIpIds(elasticIpIds []string) {
    r.ElasticIpIds = elasticIpIds
}
/* param elasticIpCount: 新购公网IP数量(Optional) */
func (r *CreateNatGatewayRequest) SetElasticIpCount(elasticIpCount int) {
    r.ElasticIpCount = &elasticIpCount
}
/* param elasticIpSpec: 新购公网IP配置。NAT网关仅支持打包创建标准公网IP，不支持边缘公网IP。且标准公网IP仅支持按配置、按用量两种计费模式。(Optional) */
func (r *CreateNatGatewayRequest) SetElasticIpSpec(elasticIpSpec *vpc.ElasticIpSpec) {
    r.ElasticIpSpec = elasticIpSpec
}
/* param natGatewayCharge: 计费配置，仅支持按配置，默认按配置(Optional) */
func (r *CreateNatGatewayRequest) SetNatGatewayCharge(natGatewayCharge *charge.ChargeSpec) {
    r.NatGatewayCharge = natGatewayCharge
}
/* param description: 描述,​ 允许输入UTF-8编码下的全部字符，不超过256字符(Optional) */
func (r *CreateNatGatewayRequest) SetDescription(description string) {
    r.Description = &description
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateNatGatewayRequest) GetRegionId() string {
    return r.RegionId
}

type CreateNatGatewayResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateNatGatewayResult `json:"result"`
}

type CreateNatGatewayResult struct {
    NatGatewayId string `json:"natGatewayId"`
}