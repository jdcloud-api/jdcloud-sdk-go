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
    iam "github.com/lshuining/jdcloud-sdk-go/services/iam/models"
)

type DescribePoliciesRequest struct {

    core.JDCloudRequest

    /* 页码，默认1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小，默认50，取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 关键字 (Optional) */
    Keyword *string `json:"keyword"`

    /* 策略类型：0-全部（默认），1-系统策略，2-自定义策略 (Optional) */
    QueryType *int `json:"queryType"`

    /* 排序规则：0-创建时间顺序排序，1-创建时间倒序排序 (Optional) */
    Sort *int `json:"sort"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribePoliciesRequest(
) *DescribePoliciesRequest {

	return &DescribePoliciesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/policies",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param pageNumber: 页码，默认1 (Optional)
 * param pageSize: 分页大小，默认50，取值范围[10, 100] (Optional)
 * param keyword: 关键字 (Optional)
 * param queryType: 策略类型：0-全部（默认），1-系统策略，2-自定义策略 (Optional)
 * param sort: 排序规则：0-创建时间顺序排序，1-创建时间倒序排序 (Optional)
 */
func NewDescribePoliciesRequestWithAllParams(
    pageNumber *int,
    pageSize *int,
    keyword *string,
    queryType *int,
    sort *int,
) *DescribePoliciesRequest {

    return &DescribePoliciesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/policies",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PageNumber: pageNumber,
        PageSize: pageSize,
        Keyword: keyword,
        QueryType: queryType,
        Sort: sort,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribePoliciesRequestWithoutParam() *DescribePoliciesRequest {

    return &DescribePoliciesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/policies",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param pageNumber: 页码，默认1(Optional) */
func (r *DescribePoliciesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小，默认50，取值范围[10, 100](Optional) */
func (r *DescribePoliciesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param keyword: 关键字(Optional) */
func (r *DescribePoliciesRequest) SetKeyword(keyword string) {
    r.Keyword = &keyword
}

/* param queryType: 策略类型：0-全部（默认），1-系统策略，2-自定义策略(Optional) */
func (r *DescribePoliciesRequest) SetQueryType(queryType int) {
    r.QueryType = &queryType
}

/* param sort: 排序规则：0-创建时间顺序排序，1-创建时间倒序排序(Optional) */
func (r *DescribePoliciesRequest) SetSort(sort int) {
    r.Sort = &sort
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribePoliciesRequest) GetRegionId() string {
    return ""
}

type DescribePoliciesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribePoliciesResult `json:"result"`
}

type DescribePoliciesResult struct {
    Total int `json:"total"`
    Policies []iam.Policy `json:"policies"`
}