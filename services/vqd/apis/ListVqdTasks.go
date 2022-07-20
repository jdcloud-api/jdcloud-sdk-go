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
    vqd "github.com/lshuining/jdcloud-sdk-go/services/vqd/models"
)

type ListVqdTasksRequest struct {

    core.JDCloudRequest

    /* 页码；默认值为 1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认值为 10；取值范围 [10, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /*  (Optional) */
    Filters []vqd.Filter `json:"filters"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewListVqdTasksRequest(
) *ListVqdTasksRequest {

	return &ListVqdTasksRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/vqdTasks",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param pageNumber: 页码；默认值为 1 (Optional)
 * param pageSize: 分页大小；默认值为 10；取值范围 [10, 100] (Optional)
 * param filters:  (Optional)
 */
func NewListVqdTasksRequestWithAllParams(
    pageNumber *int,
    pageSize *int,
    filters []vqd.Filter,
) *ListVqdTasksRequest {

    return &ListVqdTasksRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/vqdTasks",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PageNumber: pageNumber,
        PageSize: pageSize,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewListVqdTasksRequestWithoutParam() *ListVqdTasksRequest {

    return &ListVqdTasksRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/vqdTasks",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param pageNumber: 页码；默认值为 1(Optional) */
func (r *ListVqdTasksRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小；默认值为 10；取值范围 [10, 100](Optional) */
func (r *ListVqdTasksRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param filters: (Optional) */
func (r *ListVqdTasksRequest) SetFilters(filters []vqd.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ListVqdTasksRequest) GetRegionId() string {
    return ""
}

type ListVqdTasksResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ListVqdTasksResult `json:"result"`
}

type ListVqdTasksResult struct {
    PageNumber int `json:"pageNumber"`
    PageSize int `json:"pageSize"`
    TotalElements int `json:"totalElements"`
    TotalPages int `json:"totalPages"`
    Content []vqd.VqdTaskObject `json:"content"`
}