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
)

type ModyfyAgentLimitsRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* agentId  */
    AgentId string `json:"agentId"`

    /* 是否限制 0 不限制 1 限制(cpuLimit/memLimit必填) (Optional) */
    LimitStatus *int `json:"limitStatus"`

    /* cpu使用限制（1%-50%） (Optional) */
    CpuLimit *int `json:"cpuLimit"`

    /* 内存占用限额（1%-50%） (Optional) */
    MemLimit *int `json:"memLimit"`
}

/*
 * param regionId: 地域 Id (Required)
 * param agentId: agentId (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModyfyAgentLimitsRequest(
    regionId string,
    agentId string,
) *ModyfyAgentLimitsRequest {

	return &ModyfyAgentLimitsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/agents/{agentId}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        AgentId: agentId,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param agentId: agentId (Required)
 * param limitStatus: 是否限制 0 不限制 1 限制(cpuLimit/memLimit必填) (Optional)
 * param cpuLimit: cpu使用限制（1%-50%） (Optional)
 * param memLimit: 内存占用限额（1%-50%） (Optional)
 */
func NewModyfyAgentLimitsRequestWithAllParams(
    regionId string,
    agentId string,
    limitStatus *int,
    cpuLimit *int,
    memLimit *int,
) *ModyfyAgentLimitsRequest {

    return &ModyfyAgentLimitsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/agents/{agentId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        AgentId: agentId,
        LimitStatus: limitStatus,
        CpuLimit: cpuLimit,
        MemLimit: memLimit,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModyfyAgentLimitsRequestWithoutParam() *ModyfyAgentLimitsRequest {

    return &ModyfyAgentLimitsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/agents/{agentId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *ModyfyAgentLimitsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param agentId: agentId(Required) */
func (r *ModyfyAgentLimitsRequest) SetAgentId(agentId string) {
    r.AgentId = agentId
}

/* param limitStatus: 是否限制 0 不限制 1 限制(cpuLimit/memLimit必填)(Optional) */
func (r *ModyfyAgentLimitsRequest) SetLimitStatus(limitStatus int) {
    r.LimitStatus = &limitStatus
}

/* param cpuLimit: cpu使用限制（1%-50%）(Optional) */
func (r *ModyfyAgentLimitsRequest) SetCpuLimit(cpuLimit int) {
    r.CpuLimit = &cpuLimit
}

/* param memLimit: 内存占用限额（1%-50%）(Optional) */
func (r *ModyfyAgentLimitsRequest) SetMemLimit(memLimit int) {
    r.MemLimit = &memLimit
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModyfyAgentLimitsRequest) GetRegionId() string {
    return r.RegionId
}

type ModyfyAgentLimitsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModyfyAgentLimitsResult `json:"result"`
}

type ModyfyAgentLimitsResult struct {
}