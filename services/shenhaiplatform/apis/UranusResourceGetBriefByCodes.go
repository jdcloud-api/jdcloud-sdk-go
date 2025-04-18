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

type UranusResourceGetBriefByCodesRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 资源codes  */
    ResourceCodes []string `json:"resourceCodes"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param resourceCodes: 资源codes (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUranusResourceGetBriefByCodesRequest(
    regionId string,
    appName string,
    resourceCodes []string,
) *UranusResourceGetBriefByCodesRequest {

	return &UranusResourceGetBriefByCodesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/uranusResourceGetBriefByCodes",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        AppName: appName,
        ResourceCodes: resourceCodes,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param resourceCodes: 资源codes (Required)
 */
func NewUranusResourceGetBriefByCodesRequestWithAllParams(
    regionId string,
    appName string,
    resourceCodes []string,
) *UranusResourceGetBriefByCodesRequest {

    return &UranusResourceGetBriefByCodesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/uranusResourceGetBriefByCodes",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
        ResourceCodes: resourceCodes,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUranusResourceGetBriefByCodesRequestWithoutParam() *UranusResourceGetBriefByCodesRequest {

    return &UranusResourceGetBriefByCodesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/uranusResourceGetBriefByCodes",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *UranusResourceGetBriefByCodesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *UranusResourceGetBriefByCodesRequest) SetAppName(appName string) {
    r.AppName = appName
}
/* param resourceCodes: 资源codes(Required) */
func (r *UranusResourceGetBriefByCodesRequest) SetResourceCodes(resourceCodes []string) {
    r.ResourceCodes = resourceCodes
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UranusResourceGetBriefByCodesRequest) GetRegionId() string {
    return r.RegionId
}

type UranusResourceGetBriefByCodesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UranusResourceGetBriefByCodesResult `json:"result"`
}

type UranusResourceGetBriefByCodesResult struct {
    Code string `json:"code"`
    ErrorTitle string `json:"errorTitle"`
    ErrorMsg string `json:"errorMsg"`
    Result shenhaiplatform.ResourceBriefInfoVo `json:"result"`
    SubCode string `json:"subCode"`
    Successed bool `json:"successed"`
}