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
    cloudsign "github.com/jdcloud-api/jdcloud-sdk-go/services/cloudsign/models"
)

type DescribeStampHistoryListRequest struct {

    core.JDCloudRequest

    /* 分页大小, 默认为10, 取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 页码, 默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 印章ID  */
    StampId string `json:"stampId"`
}

/*
 * param stampId: 印章ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeStampHistoryListRequest(
    stampId string,
) *DescribeStampHistoryListRequest {

	return &DescribeStampHistoryListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/smqStamphistory:describeStampHistoryList",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        StampId: stampId,
	}
}

/*
 * param pageSize: 分页大小, 默认为10, 取值范围[10, 100] (Optional)
 * param pageNumber: 页码, 默认为1 (Optional)
 * param stampId: 印章ID (Required)
 */
func NewDescribeStampHistoryListRequestWithAllParams(
    pageSize *int,
    pageNumber *int,
    stampId string,
) *DescribeStampHistoryListRequest {

    return &DescribeStampHistoryListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/smqStamphistory:describeStampHistoryList",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PageSize: pageSize,
        PageNumber: pageNumber,
        StampId: stampId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeStampHistoryListRequestWithoutParam() *DescribeStampHistoryListRequest {

    return &DescribeStampHistoryListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/smqStamphistory:describeStampHistoryList",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param pageSize: 分页大小, 默认为10, 取值范围[10, 100](Optional) */
func (r *DescribeStampHistoryListRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}
/* param pageNumber: 页码, 默认为1(Optional) */
func (r *DescribeStampHistoryListRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}
/* param stampId: 印章ID(Required) */
func (r *DescribeStampHistoryListRequest) SetStampId(stampId string) {
    r.StampId = stampId
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeStampHistoryListRequest) GetRegionId() string {
    return ""
}

type DescribeStampHistoryListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeStampHistoryListResult `json:"result"`
}

type DescribeStampHistoryListResult struct {
    Success bool `json:"success"`
    Message string `json:"message"`
    Code string `json:"code"`
    Data cloudsign.PageStampHistoryResp `json:"data"`
}