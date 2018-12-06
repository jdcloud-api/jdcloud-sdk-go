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

type EnableAlarmRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 规则id  */
    AlarmId string `json:"alarmId"`
}

/*
 * param regionId: 地域 Id (Required)
 * param alarmId: 规则id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewEnableAlarmRequest(
    regionId string,
    alarmId string,
) *EnableAlarmRequest {

	return &EnableAlarmRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/alarms/{alarmId}/enable",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        AlarmId: alarmId,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param alarmId: 规则id (Required)
 */
func NewEnableAlarmRequestWithAllParams(
    regionId string,
    alarmId string,
) *EnableAlarmRequest {

    return &EnableAlarmRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/alarms/{alarmId}/enable",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        AlarmId: alarmId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewEnableAlarmRequestWithoutParam() *EnableAlarmRequest {

    return &EnableAlarmRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/alarms/{alarmId}/enable",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *EnableAlarmRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param alarmId: 规则id(Required) */
func (r *EnableAlarmRequest) SetAlarmId(alarmId string) {
    r.AlarmId = alarmId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r EnableAlarmRequest) GetRegionId() string {
    return r.RegionId
}

type EnableAlarmResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result EnableAlarmResult `json:"result"`
}

type EnableAlarmResult struct {
}