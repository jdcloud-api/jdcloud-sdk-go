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

type UranusFunctionDeleteRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 函数id (Optional) */
    FunctionId *int `json:"functionId"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUranusFunctionDeleteRequest(
    regionId string,
    appName string,
) *UranusFunctionDeleteRequest {

	return &UranusFunctionDeleteRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/uranusFunctionDelete",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        AppName: appName,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param functionId: 函数id (Optional)
 */
func NewUranusFunctionDeleteRequestWithAllParams(
    regionId string,
    appName string,
    functionId *int,
) *UranusFunctionDeleteRequest {

    return &UranusFunctionDeleteRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/uranusFunctionDelete",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
        FunctionId: functionId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUranusFunctionDeleteRequestWithoutParam() *UranusFunctionDeleteRequest {

    return &UranusFunctionDeleteRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/uranusFunctionDelete",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *UranusFunctionDeleteRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *UranusFunctionDeleteRequest) SetAppName(appName string) {
    r.AppName = appName
}
/* param functionId: 函数id(Optional) */
func (r *UranusFunctionDeleteRequest) SetFunctionId(functionId int) {
    r.FunctionId = &functionId
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UranusFunctionDeleteRequest) GetRegionId() string {
    return r.RegionId
}

type UranusFunctionDeleteResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UranusFunctionDeleteResult `json:"result"`
}

type UranusFunctionDeleteResult struct {
    Code string `json:"code"`
    ErrorTitle string `json:"errorTitle"`
    ErrorMsg string `json:"errorMsg"`
    Result bool `json:"result"`
    SubCode string `json:"subCode"`
    Successed bool `json:"successed"`
}