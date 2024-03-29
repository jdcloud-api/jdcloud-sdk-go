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

type StartFlowLogRequest struct {

    core.JDCloudRequest

    /* 地域 ID  */
    RegionId string `json:"regionId"`

    /* 流日志ID  */
    FlowLogId string `json:"flowLogId"`
}

/*
 * param regionId: 地域 ID (Required)
 * param flowLogId: 流日志ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewStartFlowLogRequest(
    regionId string,
    flowLogId string,
) *StartFlowLogRequest {

	return &StartFlowLogRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/flowLogs/{flowLogId}:startFlowLog",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        FlowLogId: flowLogId,
	}
}

/*
 * param regionId: 地域 ID (Required)
 * param flowLogId: 流日志ID (Required)
 */
func NewStartFlowLogRequestWithAllParams(
    regionId string,
    flowLogId string,
) *StartFlowLogRequest {

    return &StartFlowLogRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/flowLogs/{flowLogId}:startFlowLog",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        FlowLogId: flowLogId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewStartFlowLogRequestWithoutParam() *StartFlowLogRequest {

    return &StartFlowLogRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/flowLogs/{flowLogId}:startFlowLog",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 ID(Required) */
func (r *StartFlowLogRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param flowLogId: 流日志ID(Required) */
func (r *StartFlowLogRequest) SetFlowLogId(flowLogId string) {
    r.FlowLogId = flowLogId
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r StartFlowLogRequest) GetRegionId() string {
    return r.RegionId
}

type StartFlowLogResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result StartFlowLogResult `json:"result"`
}

type StartFlowLogResult struct {
}