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
    live "github.com/jdcloud-api/jdcloud-sdk-go/services/live/models"
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type DescribeCustomLiveStreamRecordConfigRequest struct {

    core.JDCloudRequest

    /* 页码；默认为1 (Optional) */
    PageNum *int `json:"pageNum"`

    /* 分页大小；默认为10；取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 转码模板查询过滤条件, 不传递分页参数时默认返回10条 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeCustomLiveStreamRecordConfigRequest(
) *DescribeCustomLiveStreamRecordConfigRequest {

	return &DescribeCustomLiveStreamRecordConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/records:config",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param pageNum: 页码；默认为1 (Optional)
 * param pageSize: 分页大小；默认为10；取值范围[10, 100] (Optional)
 * param filters: 转码模板查询过滤条件, 不传递分页参数时默认返回10条 (Optional)
 */
func NewDescribeCustomLiveStreamRecordConfigRequestWithAllParams(
    pageNum *int,
    pageSize *int,
    filters []common.Filter,
) *DescribeCustomLiveStreamRecordConfigRequest {

    return &DescribeCustomLiveStreamRecordConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/records:config",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PageNum: pageNum,
        PageSize: pageSize,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeCustomLiveStreamRecordConfigRequestWithoutParam() *DescribeCustomLiveStreamRecordConfigRequest {

    return &DescribeCustomLiveStreamRecordConfigRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/records:config",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param pageNum: 页码；默认为1(Optional) */
func (r *DescribeCustomLiveStreamRecordConfigRequest) SetPageNum(pageNum int) {
    r.PageNum = &pageNum
}

/* param pageSize: 分页大小；默认为10；取值范围[10, 100](Optional) */
func (r *DescribeCustomLiveStreamRecordConfigRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param filters: 转码模板查询过滤条件, 不传递分页参数时默认返回10条(Optional) */
func (r *DescribeCustomLiveStreamRecordConfigRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeCustomLiveStreamRecordConfigRequest) GetRegionId() string {
    return ""
}

type DescribeCustomLiveStreamRecordConfigResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeCustomLiveStreamRecordConfigResult `json:"result"`
}

type DescribeCustomLiveStreamRecordConfigResult struct {
    PageNumber int `json:"pageNumber"`
    PageSize int `json:"pageSize"`
    TotalCount int `json:"totalCount"`
    RecordConfigs []live.LiveRecordConfig `json:"recordConfigs"`
}