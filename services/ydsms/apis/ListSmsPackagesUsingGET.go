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
    ydsms "github.com/lshuining/jdcloud-sdk-go/services/ydsms/models"
)

type ListSmsPackagesUsingGETRequest struct {

    core.JDCloudRequest

    /* 套餐包类型:1通道短信，2官方短信 (Optional) */
    PackageType *int `json:"packageType"`

    /* pageNumber (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* pageSize (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewListSmsPackagesUsingGETRequest(
) *ListSmsPackagesUsingGETRequest {

	return &ListSmsPackagesUsingGETRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/smsPackages",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param packageType: 套餐包类型:1通道短信，2官方短信 (Optional)
 * param pageNumber: pageNumber (Optional)
 * param pageSize: pageSize (Optional)
 */
func NewListSmsPackagesUsingGETRequestWithAllParams(
    packageType *int,
    pageNumber *int,
    pageSize *int,
) *ListSmsPackagesUsingGETRequest {

    return &ListSmsPackagesUsingGETRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/smsPackages",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PackageType: packageType,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewListSmsPackagesUsingGETRequestWithoutParam() *ListSmsPackagesUsingGETRequest {

    return &ListSmsPackagesUsingGETRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/smsPackages",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param packageType: 套餐包类型:1通道短信，2官方短信(Optional) */
func (r *ListSmsPackagesUsingGETRequest) SetPackageType(packageType int) {
    r.PackageType = &packageType
}

/* param pageNumber: pageNumber(Optional) */
func (r *ListSmsPackagesUsingGETRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: pageSize(Optional) */
func (r *ListSmsPackagesUsingGETRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ListSmsPackagesUsingGETRequest) GetRegionId() string {
    return ""
}

type ListSmsPackagesUsingGETResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ListSmsPackagesUsingGETResult `json:"result"`
}

type ListSmsPackagesUsingGETResult struct {
    SmsPackages []ydsms.SmsPackageVO `json:"smsPackages"`
    TotalCount int64 `json:"totalCount"`
}