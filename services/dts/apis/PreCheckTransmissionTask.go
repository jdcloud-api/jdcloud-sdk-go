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

type PreCheckTransmissionTaskRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》]  */
    RegionId string `json:"regionId"`

    /* DTS数据传输任务ID  */
    TaskId string `json:"taskId"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》] (Required)
 * param taskId: DTS数据传输任务ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewPreCheckTransmissionTaskRequest(
    regionId string,
    taskId string,
) *PreCheckTransmissionTaskRequest {

	return &PreCheckTransmissionTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/dts/{taskId}/precheck",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        TaskId: taskId,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》] (Required)
 * param taskId: DTS数据传输任务ID (Required)
 */
func NewPreCheckTransmissionTaskRequestWithAllParams(
    regionId string,
    taskId string,
) *PreCheckTransmissionTaskRequest {

    return &PreCheckTransmissionTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dts/{taskId}/precheck",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        TaskId: taskId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewPreCheckTransmissionTaskRequestWithoutParam() *PreCheckTransmissionTaskRequest {

    return &PreCheckTransmissionTaskRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dts/{taskId}/precheck",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](Required) */
func (r *PreCheckTransmissionTaskRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param taskId: DTS数据传输任务ID(Required) */
func (r *PreCheckTransmissionTaskRequest) SetTaskId(taskId string) {
    r.TaskId = taskId
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r PreCheckTransmissionTaskRequest) GetRegionId() string {
    return r.RegionId
}

type PreCheckTransmissionTaskResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result PreCheckTransmissionTaskResult `json:"result"`
}

type PreCheckTransmissionTaskResult struct {
}