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
    cdn "github.com/jdcloud-api/jdcloud-sdk-go/services/cdn/models"
)

type QueryRefreshLimitRequest struct {

    core.JDCloudRequest

    /* 子账号，传哪些子账号就查哪些的，不传或传空默认返回所有子账号的额度 (Optional) */
    SubUsers []string `json:"subUsers"`

    /* 默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 默认为10，最大100 (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryRefreshLimitRequest(
) *QueryRefreshLimitRequest {

	return &QueryRefreshLimitRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/task:queryLimit",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param subUsers: 子账号，传哪些子账号就查哪些的，不传或传空默认返回所有子账号的额度 (Optional)
 * param pageNumber: 默认为1 (Optional)
 * param pageSize: 默认为10，最大100 (Optional)
 */
func NewQueryRefreshLimitRequestWithAllParams(
    subUsers []string,
    pageNumber *int,
    pageSize *int,
) *QueryRefreshLimitRequest {

    return &QueryRefreshLimitRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/task:queryLimit",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        SubUsers: subUsers,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryRefreshLimitRequestWithoutParam() *QueryRefreshLimitRequest {

    return &QueryRefreshLimitRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/task:queryLimit",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param subUsers: 子账号，传哪些子账号就查哪些的，不传或传空默认返回所有子账号的额度(Optional) */
func (r *QueryRefreshLimitRequest) SetSubUsers(subUsers []string) {
    r.SubUsers = subUsers
}

/* param pageNumber: 默认为1(Optional) */
func (r *QueryRefreshLimitRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 默认为10，最大100(Optional) */
func (r *QueryRefreshLimitRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryRefreshLimitRequest) GetRegionId() string {
    return ""
}

type QueryRefreshLimitResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryRefreshLimitResult `json:"result"`
}

type QueryRefreshLimitResult struct {
    User string `json:"user"`
    RefreshCount int64 `json:"refreshCount"`
    DirCount int64 `json:"dirCount"`
    PrefetchCount int64 `json:"prefetchCount"`
    PageNumber int `json:"pageNumber"`
    PageSize int `json:"pageSize"`
    Total int `json:"total"`
    SubUserQuota []cdn.SubUserRefreshLimit `json:"subUserQuota"`
}