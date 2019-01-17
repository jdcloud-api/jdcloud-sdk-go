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
    jcq "github.com/jdcloud-api/jdcloud-sdk-go/services/jcq/models"
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type DescribeTopicsRequest struct {

    core.JDCloudRequest

    /* 所在区域的Region ID  */
    RegionId string `json:"regionId"`

    /* 分页之中的每页大小 (Optional) */
    PageSize *int `json:"pageSize"`

    /* 分页之中的页码 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* topic名称的过滤条件，大小写不敏感 (Optional) */
    TopicFilter *string `json:"topicFilter"`

    /* 标签过滤条件 (Optional) */
    TagFilters []common.TagFilter `json:"tagFilters"`
}

/*
 * param regionId: 所在区域的Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeTopicsRequest(
    regionId string,
) *DescribeTopicsRequest {

	return &DescribeTopicsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/topics",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 所在区域的Region ID (Required)
 * param pageSize: 分页之中的每页大小 (Optional)
 * param pageNumber: 分页之中的页码 (Optional)
 * param topicFilter: topic名称的过滤条件，大小写不敏感 (Optional)
 * param tagFilters: 标签过滤条件 (Optional)
 */
func NewDescribeTopicsRequestWithAllParams(
    regionId string,
    pageSize *int,
    pageNumber *int,
    topicFilter *string,
    tagFilters []common.TagFilter,
) *DescribeTopicsRequest {

    return &DescribeTopicsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/topics",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageSize: pageSize,
        PageNumber: pageNumber,
        TopicFilter: topicFilter,
        TagFilters: tagFilters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeTopicsRequestWithoutParam() *DescribeTopicsRequest {

    return &DescribeTopicsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/topics",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 所在区域的Region ID(Required) */
func (r *DescribeTopicsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageSize: 分页之中的每页大小(Optional) */
func (r *DescribeTopicsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param pageNumber: 分页之中的页码(Optional) */
func (r *DescribeTopicsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param topicFilter: topic名称的过滤条件，大小写不敏感(Optional) */
func (r *DescribeTopicsRequest) SetTopicFilter(topicFilter string) {
    r.TopicFilter = &topicFilter
}

/* param tagFilters: 标签过滤条件(Optional) */
func (r *DescribeTopicsRequest) SetTagFilters(tagFilters []common.TagFilter) {
    r.TagFilters = tagFilters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeTopicsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeTopicsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeTopicsResult `json:"result"`
}

type DescribeTopicsResult struct {
    Topics []jcq.Topic `json:"topics"`
    TotalCount int `json:"totalCount"`
}