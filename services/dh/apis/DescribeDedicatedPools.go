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
    dh "github.com/lshuining/jdcloud-sdk-go/services/dh/models"
    common "github.com/lshuining/jdcloud-sdk-go/services/common/models"
)

type DescribeDedicatedPoolsRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 页码；默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认为20；取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* dedicatedPoolId - 专有宿主机池ID，精确匹配，支持多个
name - 专有宿主机名称，模糊匹配，支持单个
dedicatedHostType - 专有宿主机机型，精确匹配，支持多个
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: 地域ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeDedicatedPoolsRequest(
    regionId string,
) *DescribeDedicatedPoolsRequest {

	return &DescribeDedicatedPoolsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/dedicatedPools",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param pageNumber: 页码；默认为1 (Optional)
 * param pageSize: 分页大小；默认为20；取值范围[10, 100] (Optional)
 * param filters: dedicatedPoolId - 专有宿主机池ID，精确匹配，支持多个
name - 专有宿主机名称，模糊匹配，支持单个
dedicatedHostType - 专有宿主机机型，精确匹配，支持多个
 (Optional)
 */
func NewDescribeDedicatedPoolsRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    filters []common.Filter,
) *DescribeDedicatedPoolsRequest {

    return &DescribeDedicatedPoolsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dedicatedPools",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeDedicatedPoolsRequestWithoutParam() *DescribeDedicatedPoolsRequest {

    return &DescribeDedicatedPoolsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dedicatedPools",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DescribeDedicatedPoolsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 页码；默认为1(Optional) */
func (r *DescribeDedicatedPoolsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小；默认为20；取值范围[10, 100](Optional) */
func (r *DescribeDedicatedPoolsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param filters: dedicatedPoolId - 专有宿主机池ID，精确匹配，支持多个
name - 专有宿主机名称，模糊匹配，支持单个
dedicatedHostType - 专有宿主机机型，精确匹配，支持多个
(Optional) */
func (r *DescribeDedicatedPoolsRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeDedicatedPoolsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeDedicatedPoolsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeDedicatedPoolsResult `json:"result"`
}

type DescribeDedicatedPoolsResult struct {
    DedicatedPools []dh.DedicatedPool `json:"dedicatedPools"`
    TotalCount int `json:"totalCount"`
}