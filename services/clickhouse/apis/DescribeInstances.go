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
    clickhouse "github.com/lshuining/jdcloud-sdk-go/services/clickhouse/models"
    common "github.com/lshuining/jdcloud-sdk-go/services/common/models"
)

type DescribeInstancesRequest struct {

    core.JDCloudRequest

    /* 地域代码  */
    RegionId string `json:"regionId"`

    /* 显示数据的页码，默认为1，取值范围：[-1,∞)。pageNumber为-1时，返回所有数据页码；超过总页数时，显示最后一页; (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 每页显示的数据条数，默认为10，取值范围：[10,100]，用于查询列表的接口 (Optional) */
    PageSize *int `json:"pageSize"`

    /* 过滤参数，多个过滤参数之间的关系为“与”(and)
支持以下属性的过滤：
instanceId, 支持operator选项：eq,ne
resourceId, 支持operator选项：eq,ne
instanceName, 支持operator选项：eq,ne,like
instanceStatus, 支持operator选项：eq,ne
 (Optional) */
    Filters []common.Filter `json:"filters"`

    /* 资源标签 (Optional) */
    TagFilters []common.TagFilter `json:"tagFilters"`
}

/*
 * param regionId: 地域代码 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeInstancesRequest(
    regionId string,
) *DescribeInstancesRequest {

	return &DescribeInstancesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances:describeInstances",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域代码 (Required)
 * param pageNumber: 显示数据的页码，默认为1，取值范围：[-1,∞)。pageNumber为-1时，返回所有数据页码；超过总页数时，显示最后一页; (Optional)
 * param pageSize: 每页显示的数据条数，默认为10，取值范围：[10,100]，用于查询列表的接口 (Optional)
 * param filters: 过滤参数，多个过滤参数之间的关系为“与”(and)
支持以下属性的过滤：
instanceId, 支持operator选项：eq,ne
resourceId, 支持operator选项：eq,ne
instanceName, 支持operator选项：eq,ne,like
instanceStatus, 支持operator选项：eq,ne
 (Optional)
 * param tagFilters: 资源标签 (Optional)
 */
func NewDescribeInstancesRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    filters []common.Filter,
    tagFilters []common.TagFilter,
) *DescribeInstancesRequest {

    return &DescribeInstancesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances:describeInstances",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Filters: filters,
        TagFilters: tagFilters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeInstancesRequestWithoutParam() *DescribeInstancesRequest {

    return &DescribeInstancesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances:describeInstances",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码(Required) */
func (r *DescribeInstancesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 显示数据的页码，默认为1，取值范围：[-1,∞)。pageNumber为-1时，返回所有数据页码；超过总页数时，显示最后一页;(Optional) */
func (r *DescribeInstancesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 每页显示的数据条数，默认为10，取值范围：[10,100]，用于查询列表的接口(Optional) */
func (r *DescribeInstancesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param filters: 过滤参数，多个过滤参数之间的关系为“与”(and)
支持以下属性的过滤：
instanceId, 支持operator选项：eq,ne
resourceId, 支持operator选项：eq,ne
instanceName, 支持operator选项：eq,ne,like
instanceStatus, 支持operator选项：eq,ne
(Optional) */
func (r *DescribeInstancesRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

/* param tagFilters: 资源标签(Optional) */
func (r *DescribeInstancesRequest) SetTagFilters(tagFilters []common.TagFilter) {
    r.TagFilters = tagFilters
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
    Instances []clickhouse.Instance `json:"instances"`
    TotalCount int `json:"totalCount"`
}