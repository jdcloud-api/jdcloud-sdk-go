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
    shenhaiplatform "github.com/jdcloud-api/jdcloud-sdk-go/services/shenhaiplatform/models"
)

type UranusTaskExportParamListRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 任务code  */
    TaskCode string `json:"taskCode"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param taskCode: 任务code (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUranusTaskExportParamListRequest(
    regionId string,
    appName string,
    taskCode string,
) *UranusTaskExportParamListRequest {

	return &UranusTaskExportParamListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/uranusTaskExportParamList",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        AppName: appName,
        TaskCode: taskCode,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param taskCode: 任务code (Required)
 */
func NewUranusTaskExportParamListRequestWithAllParams(
    regionId string,
    appName string,
    taskCode string,
) *UranusTaskExportParamListRequest {

    return &UranusTaskExportParamListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/uranusTaskExportParamList",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
        TaskCode: taskCode,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUranusTaskExportParamListRequestWithoutParam() *UranusTaskExportParamListRequest {

    return &UranusTaskExportParamListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/uranusTaskExportParamList",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *UranusTaskExportParamListRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *UranusTaskExportParamListRequest) SetAppName(appName string) {
    r.AppName = appName
}
/* param taskCode: 任务code(Required) */
func (r *UranusTaskExportParamListRequest) SetTaskCode(taskCode string) {
    r.TaskCode = taskCode
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UranusTaskExportParamListRequest) GetRegionId() string {
    return r.RegionId
}

type UranusTaskExportParamListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UranusTaskExportParamListResult `json:"result"`
}

type UranusTaskExportParamListResult struct {
    Code string `json:"code"`
    ErrorTitle string `json:"errorTitle"`
    ErrorMsg string `json:"errorMsg"`
    Result []shenhaiplatform.UranusTaskParamReq `json:"result"`
    SubCode string `json:"subCode"`
    Successed bool `json:"successed"`
}