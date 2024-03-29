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

type DisableAsAlarmRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 告警任务ID  */
    AsAlarmId string `json:"asAlarmId"`
}

/*
 * param regionId: 地域ID (Required)
 * param asAlarmId: 告警任务ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDisableAsAlarmRequest(
    regionId string,
    asAlarmId string,
) *DisableAsAlarmRequest {

	return &DisableAsAlarmRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/asAlarms/{asAlarmId}:disable",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        AsAlarmId: asAlarmId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param asAlarmId: 告警任务ID (Required)
 */
func NewDisableAsAlarmRequestWithAllParams(
    regionId string,
    asAlarmId string,
) *DisableAsAlarmRequest {

    return &DisableAsAlarmRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/asAlarms/{asAlarmId}:disable",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        AsAlarmId: asAlarmId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDisableAsAlarmRequestWithoutParam() *DisableAsAlarmRequest {

    return &DisableAsAlarmRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/asAlarms/{asAlarmId}:disable",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DisableAsAlarmRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param asAlarmId: 告警任务ID(Required) */
func (r *DisableAsAlarmRequest) SetAsAlarmId(asAlarmId string) {
    r.AsAlarmId = asAlarmId
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DisableAsAlarmRequest) GetRegionId() string {
    return r.RegionId
}

type DisableAsAlarmResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DisableAsAlarmResult `json:"result"`
}

type DisableAsAlarmResult struct {
}