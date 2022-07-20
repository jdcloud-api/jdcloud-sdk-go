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
    apigateway "github.com/lshuining/jdcloud-sdk-go/services/apigateway/models"
    common "github.com/lshuining/jdcloud-sdk-go/services/common/models"
)

type QueryApisRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 分组ID  */
    ApiGroupId string `json:"apiGroupId"`

    /* 版本号  */
    Revision string `json:"revision"`

    /* 页码, 默认为1, 取值范围：[1,∞) (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小，默认为20，取值范围：[10,100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* apiName - API名称，模糊匹配，支持单个
action - 动作，精确匹配，支持多个
backServiceType- 后端服务类型，精确匹配，支持多个
path - 路径，模糊匹配，支持单个
description - 描述，模糊匹配，支持单个
isApiProduct - 是否API产品，精确匹配，1为是
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: 地域ID (Required)
 * param apiGroupId: 分组ID (Required)
 * param revision: 版本号 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryApisRequest(
    regionId string,
    apiGroupId string,
    revision string,
) *QueryApisRequest {

	return &QueryApisRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apiGroups/{apiGroupId}/revision/{revision}/apis",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ApiGroupId: apiGroupId,
        Revision: revision,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param apiGroupId: 分组ID (Required)
 * param revision: 版本号 (Required)
 * param pageNumber: 页码, 默认为1, 取值范围：[1,∞) (Optional)
 * param pageSize: 分页大小，默认为20，取值范围：[10,100] (Optional)
 * param filters: apiName - API名称，模糊匹配，支持单个
action - 动作，精确匹配，支持多个
backServiceType- 后端服务类型，精确匹配，支持多个
path - 路径，模糊匹配，支持单个
description - 描述，模糊匹配，支持单个
isApiProduct - 是否API产品，精确匹配，1为是
 (Optional)
 */
func NewQueryApisRequestWithAllParams(
    regionId string,
    apiGroupId string,
    revision string,
    pageNumber *int,
    pageSize *int,
    filters []common.Filter,
) *QueryApisRequest {

    return &QueryApisRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apiGroups/{apiGroupId}/revision/{revision}/apis",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ApiGroupId: apiGroupId,
        Revision: revision,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryApisRequestWithoutParam() *QueryApisRequest {

    return &QueryApisRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apiGroups/{apiGroupId}/revision/{revision}/apis",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *QueryApisRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param apiGroupId: 分组ID(Required) */
func (r *QueryApisRequest) SetApiGroupId(apiGroupId string) {
    r.ApiGroupId = apiGroupId
}

/* param revision: 版本号(Required) */
func (r *QueryApisRequest) SetRevision(revision string) {
    r.Revision = revision
}

/* param pageNumber: 页码, 默认为1, 取值范围：[1,∞)(Optional) */
func (r *QueryApisRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小，默认为20，取值范围：[10,100](Optional) */
func (r *QueryApisRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param filters: apiName - API名称，模糊匹配，支持单个
action - 动作，精确匹配，支持多个
backServiceType- 后端服务类型，精确匹配，支持多个
path - 路径，模糊匹配，支持单个
description - 描述，模糊匹配，支持单个
isApiProduct - 是否API产品，精确匹配，1为是
(Optional) */
func (r *QueryApisRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryApisRequest) GetRegionId() string {
    return r.RegionId
}

type QueryApisResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryApisResult `json:"result"`
}

type QueryApisResult struct {
    Apis []apigateway.Api `json:"apis"`
    TotalCount int `json:"totalCount"`
}