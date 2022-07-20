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
    iothub "github.com/lshuining/jdcloud-sdk-go/services/iothub/models"
    common "github.com/lshuining/jdcloud-sdk-go/services/common/models"
)

type ListProductAbilitiesRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 产品Key  */
    ProductKey string `json:"productKey"`

    /* 页码, 默认为1, 取值范围：[1,∞) (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小，默认为10，取值范围：[10,100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* abilityName-功能名称，精确匹配
abilityType-功能类型，精确匹配
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: 地域ID (Required)
 * param productKey: 产品Key (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewListProductAbilitiesRequest(
    regionId string,
    productKey string,
) *ListProductAbilitiesRequest {

	return &ListProductAbilitiesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/products/{productKey}/abilities",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        ProductKey: productKey,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param productKey: 产品Key (Required)
 * param pageNumber: 页码, 默认为1, 取值范围：[1,∞) (Optional)
 * param pageSize: 分页大小，默认为10，取值范围：[10,100] (Optional)
 * param filters: abilityName-功能名称，精确匹配
abilityType-功能类型，精确匹配
 (Optional)
 */
func NewListProductAbilitiesRequestWithAllParams(
    regionId string,
    productKey string,
    pageNumber *int,
    pageSize *int,
    filters []common.Filter,
) *ListProductAbilitiesRequest {

    return &ListProductAbilitiesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/products/{productKey}/abilities",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        ProductKey: productKey,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewListProductAbilitiesRequestWithoutParam() *ListProductAbilitiesRequest {

    return &ListProductAbilitiesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/products/{productKey}/abilities",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *ListProductAbilitiesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param productKey: 产品Key(Required) */
func (r *ListProductAbilitiesRequest) SetProductKey(productKey string) {
    r.ProductKey = productKey
}

/* param pageNumber: 页码, 默认为1, 取值范围：[1,∞)(Optional) */
func (r *ListProductAbilitiesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小，默认为10，取值范围：[10,100](Optional) */
func (r *ListProductAbilitiesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param filters: abilityName-功能名称，精确匹配
abilityType-功能类型，精确匹配
(Optional) */
func (r *ListProductAbilitiesRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ListProductAbilitiesRequest) GetRegionId() string {
    return r.RegionId
}

type ListProductAbilitiesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ListProductAbilitiesResult `json:"result"`
}

type ListProductAbilitiesResult struct {
    Page iothub.PageinfoVO `json:"page"`
    Abilities []iothub.ProductAbility `json:"abilities"`
}