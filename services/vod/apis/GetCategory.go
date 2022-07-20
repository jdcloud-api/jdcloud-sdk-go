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
)

type GetCategoryRequest struct {

    core.JDCloudRequest

    /* 分类ID  */
    CategoryId int `json:"categoryId"`
}

/*
 * param categoryId: 分类ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetCategoryRequest(
    categoryId int,
) *GetCategoryRequest {

	return &GetCategoryRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/categories/{categoryId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        CategoryId: categoryId,
	}
}

/*
 * param categoryId: 分类ID (Required)
 */
func NewGetCategoryRequestWithAllParams(
    categoryId int,
) *GetCategoryRequest {

    return &GetCategoryRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/categories/{categoryId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        CategoryId: categoryId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetCategoryRequestWithoutParam() *GetCategoryRequest {

    return &GetCategoryRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/categories/{categoryId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param categoryId: 分类ID(Required) */
func (r *GetCategoryRequest) SetCategoryId(categoryId int) {
    r.CategoryId = categoryId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetCategoryRequest) GetRegionId() string {
    return ""
}

type GetCategoryResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetCategoryResult `json:"result"`
}

type GetCategoryResult struct {
    Id int64 `json:"id"`
    Name string `json:"name"`
    Level int `json:"level"`
    ParentId int64 `json:"parentId"`
    Description string `json:"description"`
    CreateTime string `json:"createTime"`
    UpdateTime string `json:"updateTime"`
}