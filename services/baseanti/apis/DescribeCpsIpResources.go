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
    baseanti "github.com/jdcloud-api/jdcloud-sdk-go/services/baseanti/models"
)

type DescribeCpsIpResourcesRequest struct {

    core.JDCloudRequest

    /* 地域编码. 基础防护已支持华北-北京, 华东-宿迁, 华东-上海, 华南-广州  */
    RegionId string `json:"regionId"`

    /* 页码 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小 (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: 地域编码. 基础防护已支持华北-北京, 华东-宿迁, 华东-上海, 华南-广州 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeCpsIpResourcesRequest(
    regionId string,
) *DescribeCpsIpResourcesRequest {

	return &DescribeCpsIpResourcesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/cpsIpResources",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域编码. 基础防护已支持华北-北京, 华东-宿迁, 华东-上海, 华南-广州 (Required)
 * param pageNumber: 页码 (Optional)
 * param pageSize: 分页大小 (Optional)
 */
func NewDescribeCpsIpResourcesRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
) *DescribeCpsIpResourcesRequest {

    return &DescribeCpsIpResourcesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cpsIpResources",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeCpsIpResourcesRequestWithoutParam() *DescribeCpsIpResourcesRequest {

    return &DescribeCpsIpResourcesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cpsIpResources",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域编码. 基础防护已支持华北-北京, 华东-宿迁, 华东-上海, 华南-广州(Required) */
func (r *DescribeCpsIpResourcesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 页码(Optional) */
func (r *DescribeCpsIpResourcesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小(Optional) */
func (r *DescribeCpsIpResourcesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeCpsIpResourcesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeCpsIpResourcesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeCpsIpResourcesResult `json:"result"`
}

type DescribeCpsIpResourcesResult struct {
    DataList []baseanti.IpResource `json:"dataList"`
    CurrentCount int `json:"currentCount"`
    TotalCount int `json:"totalCount"`
    TotalPage int `json:"totalPage"`
}