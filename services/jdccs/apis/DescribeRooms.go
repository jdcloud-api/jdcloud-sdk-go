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
    jdccs "github.com/jdcloud-api/jdcloud-sdk-go/services/jdccs/models"
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type DescribeRoomsRequest struct {

    core.JDCloudRequest

    /* IDC机房ID  */
    Idc string `json:"idc"`

    /* 页码, 默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小，默认为20 (Optional) */
    PageSize *int `json:"pageSize"`

    /* roomNo - 房间号，精确匹配，支持多个
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param idc: IDC机房ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeRoomsRequest(
    idc string,
) *DescribeRoomsRequest {

	return &DescribeRoomsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/idcs/{idc}/rooms",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        Idc: idc,
	}
}

/*
 * param idc: IDC机房ID (Required)
 * param pageNumber: 页码, 默认为1 (Optional)
 * param pageSize: 分页大小，默认为20 (Optional)
 * param filters: roomNo - 房间号，精确匹配，支持多个
 (Optional)
 */
func NewDescribeRoomsRequestWithAllParams(
    idc string,
    pageNumber *int,
    pageSize *int,
    filters []common.Filter,
) *DescribeRoomsRequest {

    return &DescribeRoomsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/idcs/{idc}/rooms",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        Idc: idc,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeRoomsRequestWithoutParam() *DescribeRoomsRequest {

    return &DescribeRoomsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/idcs/{idc}/rooms",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param idc: IDC机房ID(Required) */
func (r *DescribeRoomsRequest) SetIdc(idc string) {
    r.Idc = idc
}

/* param pageNumber: 页码, 默认为1(Optional) */
func (r *DescribeRoomsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小，默认为20(Optional) */
func (r *DescribeRoomsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param filters: roomNo - 房间号，精确匹配，支持多个
(Optional) */
func (r *DescribeRoomsRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeRoomsRequest) GetRegionId() string {
    return ""
}

type DescribeRoomsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeRoomsResult `json:"result"`
}

type DescribeRoomsResult struct {
    Rooms []jdccs.Room `json:"rooms"`
    PageNumber int `json:"pageNumber"`
    PageSize int `json:"pageSize"`
    TotalCount int `json:"totalCount"`
}