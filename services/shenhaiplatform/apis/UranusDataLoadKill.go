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
)

type UranusDataLoadKillRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 任务id  */
    ApplicationId string `json:"applicationId"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param applicationId: 任务id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUranusDataLoadKillRequest(
    regionId string,
    appName string,
    applicationId string,
) *UranusDataLoadKillRequest {

	return &UranusDataLoadKillRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/uranusDataLoadKill",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        AppName: appName,
        ApplicationId: applicationId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param applicationId: 任务id (Required)
 */
func NewUranusDataLoadKillRequestWithAllParams(
    regionId string,
    appName string,
    applicationId string,
) *UranusDataLoadKillRequest {

    return &UranusDataLoadKillRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/uranusDataLoadKill",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
        ApplicationId: applicationId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUranusDataLoadKillRequestWithoutParam() *UranusDataLoadKillRequest {

    return &UranusDataLoadKillRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/uranusDataLoadKill",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *UranusDataLoadKillRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *UranusDataLoadKillRequest) SetAppName(appName string) {
    r.AppName = appName
}
/* param applicationId: 任务id(Required) */
func (r *UranusDataLoadKillRequest) SetApplicationId(applicationId string) {
    r.ApplicationId = applicationId
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UranusDataLoadKillRequest) GetRegionId() string {
    return r.RegionId
}

type UranusDataLoadKillResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UranusDataLoadKillResult `json:"result"`
}

type UranusDataLoadKillResult struct {
    Code string `json:"code"`
    Message string `json:"message"`
}