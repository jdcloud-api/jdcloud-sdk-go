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
    detection "github.com/lshuining/jdcloud-sdk-go/services/detection/models"
)

type DescribeProbeTaskRequest struct {

    core.JDCloudRequest

    /* 探测任务的task_id  */
    ProbeTaskID string `json:"probeTaskID"`
}

/*
 * param probeTaskID: 探测任务的task_id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeProbeTaskRequest(
    probeTaskID string,
) *DescribeProbeTaskRequest {

	return &DescribeProbeTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/probeTask/{probeTaskID}",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        ProbeTaskID: probeTaskID,
	}
}

/*
 * param probeTaskID: 探测任务的task_id (Required)
 */
func NewDescribeProbeTaskRequestWithAllParams(
    probeTaskID string,
) *DescribeProbeTaskRequest {

    return &DescribeProbeTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/probeTask/{probeTaskID}",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        ProbeTaskID: probeTaskID,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeProbeTaskRequestWithoutParam() *DescribeProbeTaskRequest {

    return &DescribeProbeTaskRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/probeTask/{probeTaskID}",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param probeTaskID: 探测任务的task_id(Required) */
func (r *DescribeProbeTaskRequest) SetProbeTaskID(probeTaskID string) {
    r.ProbeTaskID = probeTaskID
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeProbeTaskRequest) GetRegionId() string {
    return ""
}

type DescribeProbeTaskResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeProbeTaskResult `json:"result"`
}

type DescribeProbeTaskResult struct {
    Probes []detection.Probe `json:"probes"`
    Address string `json:"address"`
    CreateTime string `json:"createTime"`
    Deleted bool `json:"deleted"`
    Enable bool `json:"enable"`
    Frequency int64 `json:"frequency"`
    HttpBody string `json:"httpBody"`
    HttpCookie []detection.KeyValue `json:"httpCookie"`
    HttpHeader []detection.KeyValue `json:"httpHeader"`
    HttpType int64 `json:"httpType"`
    Name string `json:"name"`
    Pin string `json:"pin"`
    Port int64 `json:"port"`
    ProbeType int64 `json:"probeType"`
    TargetId string `json:"targetId"`
    TargetRegion string `json:"targetRegion"`
    TaskId string `json:"taskId"`
    TaskType int64 `json:"taskType"`
    Timeout int64 `json:"timeout"`
    UpdateTime string `json:"updateTime"`
}