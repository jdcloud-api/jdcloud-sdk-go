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
    lb "github.com/jdcloud-api/jdcloud-sdk-go/services/lb/models"
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type DescribeCustomizedConfigurationsRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 页码, 默认为1, 取值范围：[1,∞), 页码超过总页数时, 显示最后一页 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小，默认为20，取值范围：[10,100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* customizedConfigurationIds - 个性化配置ID，支持多个
customizedConfigurationNames - 个性化配置名称，支持多个
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeCustomizedConfigurationsRequest(
    regionId string,
) *DescribeCustomizedConfigurationsRequest {

	return &DescribeCustomizedConfigurationsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/customizedConfigurations/",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param pageNumber: 页码, 默认为1, 取值范围：[1,∞), 页码超过总页数时, 显示最后一页 (Optional)
 * param pageSize: 分页大小，默认为20，取值范围：[10,100] (Optional)
 * param filters: customizedConfigurationIds - 个性化配置ID，支持多个
customizedConfigurationNames - 个性化配置名称，支持多个
 (Optional)
 */
func NewDescribeCustomizedConfigurationsRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    filters []common.Filter,
) *DescribeCustomizedConfigurationsRequest {

    return &DescribeCustomizedConfigurationsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/customizedConfigurations/",
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
func NewDescribeCustomizedConfigurationsRequestWithoutParam() *DescribeCustomizedConfigurationsRequest {

    return &DescribeCustomizedConfigurationsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/customizedConfigurations/",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeCustomizedConfigurationsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param pageNumber: 页码, 默认为1, 取值范围：[1,∞), 页码超过总页数时, 显示最后一页(Optional) */
func (r *DescribeCustomizedConfigurationsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}
/* param pageSize: 分页大小，默认为20，取值范围：[10,100](Optional) */
func (r *DescribeCustomizedConfigurationsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}
/* param filters: customizedConfigurationIds - 个性化配置ID，支持多个
customizedConfigurationNames - 个性化配置名称，支持多个
(Optional) */
func (r *DescribeCustomizedConfigurationsRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeCustomizedConfigurationsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeCustomizedConfigurationsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeCustomizedConfigurationsResult `json:"result"`
}

type DescribeCustomizedConfigurationsResult struct {
    CustomizedConfigurations []lb.CustomizedConfiguration `json:"customizedConfigurations"`
    TotalCount int `json:"totalCount"`
}