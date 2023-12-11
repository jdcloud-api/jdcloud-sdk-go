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
    iam "github.com/jdcloud-api/jdcloud-sdk-go/services/iam/models"
)

type DescribeSubUsersRequest struct {

    core.JDCloudRequest

    /* 页码，默认1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小，默认50，取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 关键字 (Optional) */
    Keyword *string `json:"keyword"`

    /* 排序规则：0-创建时间顺序排序，1-创建时间倒序排序 (Optional) */
    Sort *int `json:"sort"`

    /* 手机号，和keyword互斥，不要都传 (Optional) */
    Phone *string `json:"phone"`

    /* 邮箱，和keyword互斥，不要都传 (Optional) */
    Email *string `json:"email"`

    /* 用户姓名，和keyword互斥，不要都传 (Optional) */
    Description *string `json:"description"`

    /* 子用户别名 (Optional) */
    NickName *string `json:"nickName"`

    /* 子用户姓名 (Optional) */
    Name *string `json:"name"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeSubUsersRequest(
) *DescribeSubUsersRequest {

	return &DescribeSubUsersRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/subUsers",
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
 * param sort: 排序规则：0-创建时间顺序排序，1-创建时间倒序排序 (Optional)
 * param phone: 手机号，和keyword互斥，不要都传 (Optional)
 * param email: 邮箱，和keyword互斥，不要都传 (Optional)
 * param description: 用户姓名，和keyword互斥，不要都传 (Optional)
 * param nickName: 子用户别名 (Optional)
 * param name: 子用户姓名 (Optional)
 */
func NewDescribeSubUsersRequestWithAllParams(
    pageNumber *int,
    pageSize *int,
    keyword *string,
    sort *int,
    phone *string,
    email *string,
    description *string,
    nickName *string,
    name *string,
) *DescribeSubUsersRequest {

    return &DescribeSubUsersRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/subUsers",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PageNumber: pageNumber,
        PageSize: pageSize,
        Keyword: keyword,
        Sort: sort,
        Phone: phone,
        Email: email,
        Description: description,
        NickName: nickName,
        Name: name,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeSubUsersRequestWithoutParam() *DescribeSubUsersRequest {

    return &DescribeSubUsersRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/subUsers",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param pageNumber: 页码，默认1(Optional) */
func (r *DescribeSubUsersRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}
/* param pageSize: 分页大小，默认50，取值范围[10, 100](Optional) */
func (r *DescribeSubUsersRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}
/* param keyword: 关键字(Optional) */
func (r *DescribeSubUsersRequest) SetKeyword(keyword string) {
    r.Keyword = &keyword
}
/* param sort: 排序规则：0-创建时间顺序排序，1-创建时间倒序排序(Optional) */
func (r *DescribeSubUsersRequest) SetSort(sort int) {
    r.Sort = &sort
}
/* param phone: 手机号，和keyword互斥，不要都传(Optional) */
func (r *DescribeSubUsersRequest) SetPhone(phone string) {
    r.Phone = &phone
}
/* param email: 邮箱，和keyword互斥，不要都传(Optional) */
func (r *DescribeSubUsersRequest) SetEmail(email string) {
    r.Email = &email
}
/* param description: 用户姓名，和keyword互斥，不要都传(Optional) */
func (r *DescribeSubUsersRequest) SetDescription(description string) {
    r.Description = &description
}
/* param nickName: 子用户别名(Optional) */
func (r *DescribeSubUsersRequest) SetNickName(nickName string) {
    r.NickName = &nickName
}
/* param name: 子用户姓名(Optional) */
func (r *DescribeSubUsersRequest) SetName(name string) {
    r.Name = &name
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeSubUsersRequest) GetRegionId() string {
    return ""
}

type DescribeSubUsersResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeSubUsersResult `json:"result"`
}

type DescribeSubUsersResult struct {
    SubUsers []iam.SubUser `json:"subUsers"`
    Total int `json:"total"`
}