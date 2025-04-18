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

type GravityParticleJobMonitorJobTypeCountRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 是否统计本人作业 (Optional) */
    IsMine *bool `json:"isMine"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGravityParticleJobMonitorJobTypeCountRequest(
    regionId string,
    appName string,
) *GravityParticleJobMonitorJobTypeCountRequest {

	return &GravityParticleJobMonitorJobTypeCountRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/gravityParticleJobMonitorJobTypeCount",
			Method:  "GET",
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
 * param isMine: 是否统计本人作业 (Optional)
 */
func NewGravityParticleJobMonitorJobTypeCountRequestWithAllParams(
    regionId string,
    appName string,
    isMine *bool,
) *GravityParticleJobMonitorJobTypeCountRequest {

    return &GravityParticleJobMonitorJobTypeCountRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/gravityParticleJobMonitorJobTypeCount",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
        IsMine: isMine,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGravityParticleJobMonitorJobTypeCountRequestWithoutParam() *GravityParticleJobMonitorJobTypeCountRequest {

    return &GravityParticleJobMonitorJobTypeCountRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/gravityParticleJobMonitorJobTypeCount",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *GravityParticleJobMonitorJobTypeCountRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *GravityParticleJobMonitorJobTypeCountRequest) SetAppName(appName string) {
    r.AppName = appName
}
/* param isMine: 是否统计本人作业(Optional) */
func (r *GravityParticleJobMonitorJobTypeCountRequest) SetIsMine(isMine bool) {
    r.IsMine = &isMine
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GravityParticleJobMonitorJobTypeCountRequest) GetRegionId() string {
    return r.RegionId
}

type GravityParticleJobMonitorJobTypeCountResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GravityParticleJobMonitorJobTypeCountResult `json:"result"`
}

type GravityParticleJobMonitorJobTypeCountResult struct {
    Success int `json:"success"`
    Result []shenhaiplatform.GpmnPieData `json:"result"`
    Code string `json:"code"`
    Msg string `json:"msg"`
    _REQ_SEQ_NO_ string `json:"_REQ_SEQ_NO_"`
}