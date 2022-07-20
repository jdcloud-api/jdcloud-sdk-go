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
    cps "github.com/lshuining/jdcloud-sdk-go/services/cps/models"
    common "github.com/lshuining/jdcloud-sdk-go/services/common/models"
)

type DescribeCertsRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* 页码；默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认为20；取值范围[20, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 证书名称 (Optional) */
    Name *string `json:"name"`

    /* certId - 证书ID，精确匹配，支持多个<br/>
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeCertsRequest(
    regionId string,
) *DescribeCertsRequest {

	return &DescribeCertsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/certs",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
 * param pageNumber: 页码；默认为1 (Optional)
 * param pageSize: 分页大小；默认为20；取值范围[20, 100] (Optional)
 * param name: 证书名称 (Optional)
 * param filters: certId - 证书ID，精确匹配，支持多个<br/>
 (Optional)
 */
func NewDescribeCertsRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    name *string,
    filters []common.Filter,
) *DescribeCertsRequest {

    return &DescribeCertsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/certs",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Name: name,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeCertsRequestWithoutParam() *DescribeCertsRequest {

    return &DescribeCertsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/certs",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域(Required) */
func (r *DescribeCertsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 页码；默认为1(Optional) */
func (r *DescribeCertsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小；默认为20；取值范围[20, 100](Optional) */
func (r *DescribeCertsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param name: 证书名称(Optional) */
func (r *DescribeCertsRequest) SetName(name string) {
    r.Name = &name
}

/* param filters: certId - 证书ID，精确匹配，支持多个<br/>
(Optional) */
func (r *DescribeCertsRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeCertsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeCertsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeCertsResult `json:"result"`
}

type DescribeCertsResult struct {
    Certs []cps.Cert `json:"certs"`
    PageNumber int `json:"pageNumber"`
    PageSize int `json:"pageSize"`
    TotalCount int `json:"totalCount"`
}