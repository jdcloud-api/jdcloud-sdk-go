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
    vm "github.com/lshuining/jdcloud-sdk-go/services/vm/models"
    common "github.com/lshuining/jdcloud-sdk-go/services/common/models"
)

type DescribeInstanceTypesRequest struct {

    core.JDCloudRequest

    /* 地域ID。  */
    RegionId string `json:"regionId"`

    /* 产品线类型，默认为 `vm`。支持范围：`vm` 云主机，`nc` 原生容器。 (Optional) */
    ServiceName *string `json:"serviceName"`

    /* <b>filters 中支持使用以下关键字进行过滤</b>
`instanceTypes`: 实例规格，精确匹配，支持多个
`az`: 可用区，精确匹配，支持多个
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: 地域ID。 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeInstanceTypesRequest(
    regionId string,
) *DescribeInstanceTypesRequest {

	return &DescribeInstanceTypesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instanceTypes",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID。 (Required)
 * param serviceName: 产品线类型，默认为 `vm`。支持范围：`vm` 云主机，`nc` 原生容器。 (Optional)
 * param filters: <b>filters 中支持使用以下关键字进行过滤</b>
`instanceTypes`: 实例规格，精确匹配，支持多个
`az`: 可用区，精确匹配，支持多个
 (Optional)
 */
func NewDescribeInstanceTypesRequestWithAllParams(
    regionId string,
    serviceName *string,
    filters []common.Filter,
) *DescribeInstanceTypesRequest {

    return &DescribeInstanceTypesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instanceTypes",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ServiceName: serviceName,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeInstanceTypesRequestWithoutParam() *DescribeInstanceTypesRequest {

    return &DescribeInstanceTypesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instanceTypes",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID。(Required) */
func (r *DescribeInstanceTypesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param serviceName: 产品线类型，默认为 `vm`。支持范围：`vm` 云主机，`nc` 原生容器。(Optional) */
func (r *DescribeInstanceTypesRequest) SetServiceName(serviceName string) {
    r.ServiceName = &serviceName
}

/* param filters: <b>filters 中支持使用以下关键字进行过滤</b>
`instanceTypes`: 实例规格，精确匹配，支持多个
`az`: 可用区，精确匹配，支持多个
(Optional) */
func (r *DescribeInstanceTypesRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeInstanceTypesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeInstanceTypesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeInstanceTypesResult `json:"result"`
}

type DescribeInstanceTypesResult struct {
    InstanceTypes []vm.InstanceType `json:"instanceTypes"`
    SpecificInstanceTypes []vm.InstanceType `json:"specificInstanceTypes"`
    TotalCount int `json:"totalCount"`
}