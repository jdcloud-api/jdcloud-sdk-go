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
    dbs "github.com/lshuining/jdcloud-sdk-go/services/dbs/models"
)

type AgentTasksRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》]  */
    RegionId string `json:"regionId"`

    /* Agent ID  */
    AgentId string `json:"agentId"`

    /* agent的状态 (Optional) */
    Stat *string `json:"stat"`

    /* agent的MAC地址  */
    Mac string `json:"mac"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》] (Required)
 * param agentId: Agent ID (Required)
 * param mac: agent的MAC地址 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAgentTasksRequest(
    regionId string,
    agentId string,
    mac string,
) *AgentTasksRequest {

	return &AgentTasksRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/manage/{agentId}/tasks",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        AgentId: agentId,
        Mac: mac,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》] (Required)
 * param agentId: Agent ID (Required)
 * param stat: agent的状态 (Optional)
 * param mac: agent的MAC地址 (Required)
 */
func NewAgentTasksRequestWithAllParams(
    regionId string,
    agentId string,
    stat *string,
    mac string,
) *AgentTasksRequest {

    return &AgentTasksRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/manage/{agentId}/tasks",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AgentId: agentId,
        Stat: stat,
        Mac: mac,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAgentTasksRequestWithoutParam() *AgentTasksRequest {

    return &AgentTasksRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/manage/{agentId}/tasks",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](Required) */
func (r *AgentTasksRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param agentId: Agent ID(Required) */
func (r *AgentTasksRequest) SetAgentId(agentId string) {
    r.AgentId = agentId
}

/* param stat: agent的状态(Optional) */
func (r *AgentTasksRequest) SetStat(stat string) {
    r.Stat = &stat
}

/* param mac: agent的MAC地址(Required) */
func (r *AgentTasksRequest) SetMac(mac string) {
    r.Mac = mac
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AgentTasksRequest) GetRegionId() string {
    return r.RegionId
}

type AgentTasksResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AgentTasksResult `json:"result"`
}

type AgentTasksResult struct {
    Tasks []dbs.AgentTasks `json:"tasks"`
}