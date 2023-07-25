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
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type DescribeNatGatewaysRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 页码, 默认为1, 取值范围：[1,∞), 页码超过总页数时, 显示最后一页 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小，默认为20，取值范围：[10,100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* natGatewayIds - NAT网关ID列表，支持多个
natGatewayNames - NAT网关名称列表，支持多个
natGatewayPublicIp - NAT网关公网IP，支持单个，即将废弃，请使用elasticIpAddress
elasticIpAddress - 公网IP，支持单个
natGatewayPrivateIp - NAT网关私网IP，支持单个
vpcId - NAT网关所属VPC ID，支持单个
subnetId - NAT网关所属子网ID，支持单个
 (Optional) */
    Filters []common.Filter `json:"filters"`

    /* Tag筛选条件 (Optional) */
    Tags []vpc.TagFilter `json:"tags"`
}

/*
 * param regionId: Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeNatGatewaysRequest(
    regionId string,
) *DescribeNatGatewaysRequest {

	return &DescribeNatGatewaysRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/natGateways/",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param pageNumber: 页码, 默认为1, 取值范围：[1,∞), 页码超过总页数时, 显示最后一页 (Optional)
 * param pageSize: 分页大小，默认为20，取值范围：[10,100] (Optional)
 * param filters: natGatewayIds - NAT网关ID列表，支持多个
natGatewayNames - NAT网关名称列表，支持多个
natGatewayPublicIp - NAT网关公网IP，支持单个，即将废弃，请使用elasticIpAddress
elasticIpAddress - 公网IP，支持单个
natGatewayPrivateIp - NAT网关私网IP，支持单个
vpcId - NAT网关所属VPC ID，支持单个
subnetId - NAT网关所属子网ID，支持单个
 (Optional)
 * param tags: Tag筛选条件 (Optional)
 */
func NewDescribeNatGatewaysRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    filters []common.Filter,
    tags []vpc.TagFilter,
) *DescribeNatGatewaysRequest {

    return &DescribeNatGatewaysRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/natGateways/",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Filters: filters,
        Tags: tags,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeNatGatewaysRequestWithoutParam() *DescribeNatGatewaysRequest {

    return &DescribeNatGatewaysRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/natGateways/",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeNatGatewaysRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param pageNumber: 页码, 默认为1, 取值范围：[1,∞), 页码超过总页数时, 显示最后一页(Optional) */
func (r *DescribeNatGatewaysRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}
/* param pageSize: 分页大小，默认为20，取值范围：[10,100](Optional) */
func (r *DescribeNatGatewaysRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}
/* param filters: natGatewayIds - NAT网关ID列表，支持多个
natGatewayNames - NAT网关名称列表，支持多个
natGatewayPublicIp - NAT网关公网IP，支持单个，即将废弃，请使用elasticIpAddress
elasticIpAddress - 公网IP，支持单个
natGatewayPrivateIp - NAT网关私网IP，支持单个
vpcId - NAT网关所属VPC ID，支持单个
subnetId - NAT网关所属子网ID，支持单个
(Optional) */
func (r *DescribeNatGatewaysRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}
/* param tags: Tag筛选条件(Optional) */
func (r *DescribeNatGatewaysRequest) SetTags(tags []vpc.TagFilter) {
    r.Tags = tags
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeNatGatewaysRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeNatGatewaysResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeNatGatewaysResult `json:"result"`
}

type DescribeNatGatewaysResult struct {
    NatGateways []vpc.NatGateway `json:"natGateways"`
    TotalCount int `json:"totalCount"`
}