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
    antipro "github.com/lshuining/jdcloud-sdk-go/services/antipro/models"
)

type DescribeInstancesRequest struct {

    core.JDCloudRequest

    /* 地域 Id, DDoS 防护包目前支持华北-北京, 华东-宿迁, 华东-上海  */
    RegionId string `json:"regionId"`

    /* 页码 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小 (Optional) */
    PageSize *int `json:"pageSize"`

    /* 防护包名称, 支持模糊匹配 (Optional) */
    Name *string `json:"name"`
}

/*
 * param regionId: 地域 Id, DDoS 防护包目前支持华北-北京, 华东-宿迁, 华东-上海 (Required)
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
 * param regionId: 地域 Id, DDoS 防护包目前支持华北-北京, 华东-宿迁, 华东-上海 (Required)
 * param pageNumber: 页码 (Optional)
 * param pageSize: 分页大小 (Optional)
 * param name: 防护包名称, 支持模糊匹配 (Optional)
 */
func NewDescribeInstancesRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    name *string,
) *DescribeInstancesRequest {

    return &DescribeInstancesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Name: name,
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

/* param regionId: 地域 Id, DDoS 防护包目前支持华北-北京, 华东-宿迁, 华东-上海(Required) */
func (r *DescribeInstancesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 页码(Optional) */
func (r *DescribeInstancesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小(Optional) */
func (r *DescribeInstancesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param name: 防护包名称, 支持模糊匹配(Optional) */
func (r *DescribeInstancesRequest) SetName(name string) {
    r.Name = &name
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
    DataList []antipro.Instance `json:"dataList"`
    CurrentCount int `json:"currentCount"`
    TotalCount int `json:"totalCount"`
    TotalPage int `json:"totalPage"`
}