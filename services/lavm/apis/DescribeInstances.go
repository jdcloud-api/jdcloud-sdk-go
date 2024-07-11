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
    lavm "github.com/jdcloud-api/jdcloud-sdk-go/services/lavm/models"
)

type DescribeInstancesRequest struct {

    core.JDCloudRequest

    /* regionId
  */
    RegionId string `json:"regionId"`

    /* 轻量应用云主机的实例ID, `[\"lavm-xxx\", \"lavm-yyy\"]`, json array 字串
 (Optional) */
    InstanceIds *string `json:"instanceIds"`

    /* 实例的计费方式, 目前只支持且默认值prepaid_by_duration, 包年包月, 
 (Optional) */
    ChargeType *string `json:"chargeType"`

    /* 轻量应用云主机的公网IP, 例如: `[\"114.1.x.y\", \"114.2.x.z\"]`, json array 字串
 (Optional) */
    PublicIpAddresses *string `json:"publicIpAddresses"`

    /* 轻量应用云主机的实例名称, 支持模糊搜索, 例如: `[\"instanceName-1\", \"instanceName-2\"]`, json array 字串
 (Optional) */
    Names *string `json:"names"`

    /* 页码；默认为1。 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；<br>默认为20；取值范围[10, 100]。 (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: regionId
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeInstancesRequest(
    regionId string,
) *DescribeInstancesRequest {

	return &DescribeInstancesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: regionId
 (Required)
 * param instanceIds: 轻量应用云主机的实例ID, `[\"lavm-xxx\", \"lavm-yyy\"]`, json array 字串
 (Optional)
 * param chargeType: 实例的计费方式, 目前只支持且默认值prepaid_by_duration, 包年包月, 
 (Optional)
 * param publicIpAddresses: 轻量应用云主机的公网IP, 例如: `[\"114.1.x.y\", \"114.2.x.z\"]`, json array 字串
 (Optional)
 * param names: 轻量应用云主机的实例名称, 支持模糊搜索, 例如: `[\"instanceName-1\", \"instanceName-2\"]`, json array 字串
 (Optional)
 * param pageNumber: 页码；默认为1。 (Optional)
 * param pageSize: 分页大小；<br>默认为20；取值范围[10, 100]。 (Optional)
 */
func NewDescribeInstancesRequestWithAllParams(
    regionId string,
    instanceIds *string,
    chargeType *string,
    publicIpAddresses *string,
    names *string,
    pageNumber *int,
    pageSize *int,
) *DescribeInstancesRequest {

    return &DescribeInstancesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceIds: instanceIds,
        ChargeType: chargeType,
        PublicIpAddresses: publicIpAddresses,
        Names: names,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeInstancesRequestWithoutParam() *DescribeInstancesRequest {

    return &DescribeInstancesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: regionId
(Required) */
func (r *DescribeInstancesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param instanceIds: 轻量应用云主机的实例ID, `[\"lavm-xxx\", \"lavm-yyy\"]`, json array 字串
(Optional) */
func (r *DescribeInstancesRequest) SetInstanceIds(instanceIds string) {
    r.InstanceIds = &instanceIds
}
/* param chargeType: 实例的计费方式, 目前只支持且默认值prepaid_by_duration, 包年包月, 
(Optional) */
func (r *DescribeInstancesRequest) SetChargeType(chargeType string) {
    r.ChargeType = &chargeType
}
/* param publicIpAddresses: 轻量应用云主机的公网IP, 例如: `[\"114.1.x.y\", \"114.2.x.z\"]`, json array 字串
(Optional) */
func (r *DescribeInstancesRequest) SetPublicIpAddresses(publicIpAddresses string) {
    r.PublicIpAddresses = &publicIpAddresses
}
/* param names: 轻量应用云主机的实例名称, 支持模糊搜索, 例如: `[\"instanceName-1\", \"instanceName-2\"]`, json array 字串
(Optional) */
func (r *DescribeInstancesRequest) SetNames(names string) {
    r.Names = &names
}
/* param pageNumber: 页码；默认为1。(Optional) */
func (r *DescribeInstancesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}
/* param pageSize: 分页大小；<br>默认为20；取值范围[10, 100]。(Optional) */
func (r *DescribeInstancesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeInstancesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeInstancesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeInstancesResult `json:"result"`
}

type DescribeInstancesResult struct {
    Instances []lavm.Instance `json:"instances"`
    TotalCount int `json:"totalCount"`
}