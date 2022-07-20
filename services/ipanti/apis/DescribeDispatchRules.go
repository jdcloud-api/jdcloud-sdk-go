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
    "github.com/lshuining/jdcloud-sdk-go/core"
    ipanti "github.com/lshuining/jdcloud-sdk-go/services/ipanti/models"
)

type DescribeDispatchRulesRequest struct {

    core.JDCloudRequest

    /* 区域 ID, 高防不区分区域, 传 cn-north-1 即可  */
    RegionId string `json:"regionId"`

    /* 高防实例 Id  */
    InstanceId string `json:"instanceId"`

    /* 页码, 默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小, 默认为10, 取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 实例名称, 可模糊匹配 (Optional) */
    Name *string `json:"name"`

    /* 云内IP, 可模糊匹配 (Optional) */
    InnerIp *string `json:"innerIp"`

    /* 高防IP, 可模糊匹配 (Optional) */
    ServiceIp *string `json:"serviceIp"`
}

/*
 * param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可 (Required)
 * param instanceId: 高防实例 Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeDispatchRulesRequest(
    regionId string,
    instanceId string,
) *DescribeDispatchRulesRequest {

	return &DescribeDispatchRulesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/dispatchRules",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可 (Required)
 * param instanceId: 高防实例 Id (Required)
 * param pageNumber: 页码, 默认为1 (Optional)
 * param pageSize: 分页大小, 默认为10, 取值范围[10, 100] (Optional)
 * param name: 实例名称, 可模糊匹配 (Optional)
 * param innerIp: 云内IP, 可模糊匹配 (Optional)
 * param serviceIp: 高防IP, 可模糊匹配 (Optional)
 */
func NewDescribeDispatchRulesRequestWithAllParams(
    regionId string,
    instanceId string,
    pageNumber *int,
    pageSize *int,
    name *string,
    innerIp *string,
    serviceIp *string,
) *DescribeDispatchRulesRequest {

    return &DescribeDispatchRulesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/dispatchRules",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Name: name,
        InnerIp: innerIp,
        ServiceIp: serviceIp,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeDispatchRulesRequestWithoutParam() *DescribeDispatchRulesRequest {

    return &DescribeDispatchRulesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/dispatchRules",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可(Required) */
func (r *DescribeDispatchRulesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 高防实例 Id(Required) */
func (r *DescribeDispatchRulesRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param pageNumber: 页码, 默认为1(Optional) */
func (r *DescribeDispatchRulesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小, 默认为10, 取值范围[10, 100](Optional) */
func (r *DescribeDispatchRulesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param name: 实例名称, 可模糊匹配(Optional) */
func (r *DescribeDispatchRulesRequest) SetName(name string) {
    r.Name = &name
}

/* param innerIp: 云内IP, 可模糊匹配(Optional) */
func (r *DescribeDispatchRulesRequest) SetInnerIp(innerIp string) {
    r.InnerIp = &innerIp
}

/* param serviceIp: 高防IP, 可模糊匹配(Optional) */
func (r *DescribeDispatchRulesRequest) SetServiceIp(serviceIp string) {
    r.ServiceIp = &serviceIp
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeDispatchRulesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeDispatchRulesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeDispatchRulesResult `json:"result"`
}

type DescribeDispatchRulesResult struct {
    DataList []ipanti.DispatchRule `json:"dataList"`
    CurrentCount int `json:"currentCount"`
    TotalCount int `json:"totalCount"`
    TotalPage int `json:"totalPage"`
}