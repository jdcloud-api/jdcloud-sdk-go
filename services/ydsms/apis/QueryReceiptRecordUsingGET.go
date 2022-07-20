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

type QueryReceiptRecordUsingGETRequest struct {

    core.JDCloudRequest

    /* 应用id  */
    AppId string `json:"appId"`

    /* 结束时间,pattern ="yyyy-MM-dd HH:mm:ss",timezone="GMT+8"  */
    EndTime string `json:"endTime"`

    /* 开始时间,pattern ="yyyy-MM-dd HH:mm:ss",timezone="GMT+8"  */
    StartTime string `json:"startTime"`
}

/*
 * param appId: 应用id (Required)
 * param endTime: 结束时间,pattern ="yyyy-MM-dd HH:mm:ss",timezone="GMT+8" (Required)
 * param startTime: 开始时间,pattern ="yyyy-MM-dd HH:mm:ss",timezone="GMT+8" (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryReceiptRecordUsingGETRequest(
    appId string,
    endTime string,
    startTime string,
) *QueryReceiptRecordUsingGETRequest {

	return &QueryReceiptRecordUsingGETRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/smsApps/{appId}:queryReceiptRecord",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        AppId: appId,
        EndTime: endTime,
        StartTime: startTime,
	}
}

/*
 * param appId: 应用id (Required)
 * param endTime: 结束时间,pattern ="yyyy-MM-dd HH:mm:ss",timezone="GMT+8" (Required)
 * param startTime: 开始时间,pattern ="yyyy-MM-dd HH:mm:ss",timezone="GMT+8" (Required)
 */
func NewQueryReceiptRecordUsingGETRequestWithAllParams(
    appId string,
    endTime string,
    startTime string,
) *QueryReceiptRecordUsingGETRequest {

    return &QueryReceiptRecordUsingGETRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/smsApps/{appId}:queryReceiptRecord",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        AppId: appId,
        EndTime: endTime,
        StartTime: startTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryReceiptRecordUsingGETRequestWithoutParam() *QueryReceiptRecordUsingGETRequest {

    return &QueryReceiptRecordUsingGETRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/smsApps/{appId}:queryReceiptRecord",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param appId: 应用id(Required) */
func (r *QueryReceiptRecordUsingGETRequest) SetAppId(appId string) {
    r.AppId = appId
}

/* param endTime: 结束时间,pattern ="yyyy-MM-dd HH:mm:ss",timezone="GMT+8"(Required) */
func (r *QueryReceiptRecordUsingGETRequest) SetEndTime(endTime string) {
    r.EndTime = endTime
}

/* param startTime: 开始时间,pattern ="yyyy-MM-dd HH:mm:ss",timezone="GMT+8"(Required) */
func (r *QueryReceiptRecordUsingGETRequest) SetStartTime(startTime string) {
    r.StartTime = startTime
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryReceiptRecordUsingGETRequest) GetRegionId() string {
    return ""
}

type QueryReceiptRecordUsingGETResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryReceiptRecordUsingGETResult `json:"result"`
}

type QueryReceiptRecordUsingGETResult struct {
    Pin string `json:"pin"`
    AppId string `json:"appId"`
    ReceiptRecords []ydsms.ReceiptRecord `json:"receiptRecords"`
}