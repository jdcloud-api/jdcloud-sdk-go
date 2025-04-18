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

type UranusTaskPublishOneRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 节点列表  */
    Ids []shenhaiplatform.UranusNodePublishChildReq `json:"ids"`

    /* 发布原因  */
    Reason string `json:"reason"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param ids: 节点列表 (Required)
 * param reason: 发布原因 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUranusTaskPublishOneRequest(
    regionId string,
    appName string,
    ids []shenhaiplatform.UranusNodePublishChildReq,
    reason string,
) *UranusTaskPublishOneRequest {

	return &UranusTaskPublishOneRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/uranusTaskPublishOne",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        AppName: appName,
        Ids: ids,
        Reason: reason,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param ids: 节点列表 (Required)
 * param reason: 发布原因 (Required)
 */
func NewUranusTaskPublishOneRequestWithAllParams(
    regionId string,
    appName string,
    ids []shenhaiplatform.UranusNodePublishChildReq,
    reason string,
) *UranusTaskPublishOneRequest {

    return &UranusTaskPublishOneRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/uranusTaskPublishOne",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
        Ids: ids,
        Reason: reason,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUranusTaskPublishOneRequestWithoutParam() *UranusTaskPublishOneRequest {

    return &UranusTaskPublishOneRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/uranusTaskPublishOne",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *UranusTaskPublishOneRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *UranusTaskPublishOneRequest) SetAppName(appName string) {
    r.AppName = appName
}
/* param ids: 节点列表(Required) */
func (r *UranusTaskPublishOneRequest) SetIds(ids []shenhaiplatform.UranusNodePublishChildReq) {
    r.Ids = ids
}
/* param reason: 发布原因(Required) */
func (r *UranusTaskPublishOneRequest) SetReason(reason string) {
    r.Reason = reason
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UranusTaskPublishOneRequest) GetRegionId() string {
    return r.RegionId
}

type UranusTaskPublishOneResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UranusTaskPublishOneResult `json:"result"`
}

type UranusTaskPublishOneResult struct {
    Code string `json:"code"`
    ErrorTitle string `json:"errorTitle"`
    ErrorMsg string `json:"errorMsg"`
    Result []shenhaiplatform.UranusTaskNodePreRes `json:"result"`
    SubCode string `json:"subCode"`
    Successed bool `json:"successed"`
}