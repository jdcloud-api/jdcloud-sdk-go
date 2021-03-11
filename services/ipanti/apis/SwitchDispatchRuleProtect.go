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

type SwitchDispatchRuleProtectRequest struct {

    core.JDCloudRequest

    /* 区域 ID, 高防不区分区域, 传 cn-north-1 即可  */
    RegionId string `json:"regionId"`

    /* 高防实例 Id  */
    InstanceId string `json:"instanceId"`

    /* 防护调度规则 Id  */
    DispatchRuleId string `json:"dispatchRuleId"`
}

/*
 * param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可 (Required)
 * param instanceId: 高防实例 Id (Required)
 * param dispatchRuleId: 防护调度规则 Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSwitchDispatchRuleProtectRequest(
    regionId string,
    instanceId string,
    dispatchRuleId string,
) *SwitchDispatchRuleProtectRequest {

	return &SwitchDispatchRuleProtectRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/dispatchRules/{dispatchRuleId}:protect",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        DispatchRuleId: dispatchRuleId,
	}
}

/*
 * param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可 (Required)
 * param instanceId: 高防实例 Id (Required)
 * param dispatchRuleId: 防护调度规则 Id (Required)
 */
func NewSwitchDispatchRuleProtectRequestWithAllParams(
    regionId string,
    instanceId string,
    dispatchRuleId string,
) *SwitchDispatchRuleProtectRequest {

    return &SwitchDispatchRuleProtectRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/dispatchRules/{dispatchRuleId}:protect",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        DispatchRuleId: dispatchRuleId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSwitchDispatchRuleProtectRequestWithoutParam() *SwitchDispatchRuleProtectRequest {

    return &SwitchDispatchRuleProtectRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/dispatchRules/{dispatchRuleId}:protect",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可(Required) */
func (r *SwitchDispatchRuleProtectRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 高防实例 Id(Required) */
func (r *SwitchDispatchRuleProtectRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param dispatchRuleId: 防护调度规则 Id(Required) */
func (r *SwitchDispatchRuleProtectRequest) SetDispatchRuleId(dispatchRuleId string) {
    r.DispatchRuleId = dispatchRuleId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SwitchDispatchRuleProtectRequest) GetRegionId() string {
    return r.RegionId
}

type SwitchDispatchRuleProtectResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SwitchDispatchRuleProtectResult `json:"result"`
}

type SwitchDispatchRuleProtectResult struct {
    Code int `json:"code"`
    Message string `json:"message"`
}