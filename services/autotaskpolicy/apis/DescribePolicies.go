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
    autotaskpolicy "github.com/jdcloud-api/jdcloud-sdk-go/services/autotaskpolicy/models"
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type DescribePoliciesRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 页码；默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小，默认为20，最大100。 (Optional) */
    PageSize *int `json:"pageSize"`

    /* policyId - 策略ID，精确匹配，支持多个
policyState - 策略状态，精确匹配，支持多个
policyType - 策略类型，精确匹配，支持多个
policyName - 策略名称，模糊匹配，支持一个
resourceId - 已关联的资源ID，精确匹配，支持多个
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: 地域ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribePoliciesRequest(
    regionId string,
) *DescribePoliciesRequest {

	return &DescribePoliciesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/policies",
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
 * param pageSize: 分页大小，默认为20，最大100。 (Optional)
 * param filters: policyId - 策略ID，精确匹配，支持多个
policyState - 策略状态，精确匹配，支持多个
policyType - 策略类型，精确匹配，支持多个
policyName - 策略名称，模糊匹配，支持一个
resourceId - 已关联的资源ID，精确匹配，支持多个
 (Optional)
 */
func NewDescribePoliciesRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    filters []common.Filter,
) *DescribePoliciesRequest {

    return &DescribePoliciesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/policies",
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
func NewDescribePoliciesRequestWithoutParam() *DescribePoliciesRequest {

    return &DescribePoliciesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/policies",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DescribePoliciesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param pageNumber: 页码；默认为1(Optional) */
func (r *DescribePoliciesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}
/* param pageSize: 分页大小，默认为20，最大100。(Optional) */
func (r *DescribePoliciesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}
/* param filters: policyId - 策略ID，精确匹配，支持多个
policyState - 策略状态，精确匹配，支持多个
policyType - 策略类型，精确匹配，支持多个
policyName - 策略名称，模糊匹配，支持一个
resourceId - 已关联的资源ID，精确匹配，支持多个
(Optional) */
func (r *DescribePoliciesRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribePoliciesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribePoliciesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribePoliciesResult `json:"result"`
}

type DescribePoliciesResult struct {
    Policies []autotaskpolicy.Policy `json:"policies"`
    TotalCount int `json:"totalCount"`
}