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

type ModifyInstancesProtectedRequest struct {

    core.JDCloudRequest

    /* 地域  */
    RegionId string `json:"regionId"`

    /* 高可用组 ID  */
    AgId string `json:"agId"`

    /* 实例在弹性伸缩组是否为保护状态，保护状态的实例缩容时不可以删除。  */
    IsProtected bool `json:"isProtected"`

    /* 资源ID数组。  */
    InstanceIds []string `json:"instanceIds"`
}

/*
 * param regionId: 地域 (Required)
 * param agId: 高可用组 ID (Required)
 * param isProtected: 实例在弹性伸缩组是否为保护状态，保护状态的实例缩容时不可以删除。 (Required)
 * param instanceIds: 资源ID数组。 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyInstancesProtectedRequest(
    regionId string,
    agId string,
    isProtected bool,
    instanceIds []string,
) *ModifyInstancesProtectedRequest {

	return &ModifyInstancesProtectedRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/autoScaling/{agId}:modifyInstancesProtected",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        AgId: agId,
        IsProtected: isProtected,
        InstanceIds: instanceIds,
	}
}

/*
 * param regionId: 地域 (Required)
 * param agId: 高可用组 ID (Required)
 * param isProtected: 实例在弹性伸缩组是否为保护状态，保护状态的实例缩容时不可以删除。 (Required)
 * param instanceIds: 资源ID数组。 (Required)
 */
func NewModifyInstancesProtectedRequestWithAllParams(
    regionId string,
    agId string,
    isProtected bool,
    instanceIds []string,
) *ModifyInstancesProtectedRequest {

    return &ModifyInstancesProtectedRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/autoScaling/{agId}:modifyInstancesProtected",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        AgId: agId,
        IsProtected: isProtected,
        InstanceIds: instanceIds,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyInstancesProtectedRequestWithoutParam() *ModifyInstancesProtectedRequest {

    return &ModifyInstancesProtectedRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/autoScaling/{agId}:modifyInstancesProtected",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域(Required) */
func (r *ModifyInstancesProtectedRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param agId: 高可用组 ID(Required) */
func (r *ModifyInstancesProtectedRequest) SetAgId(agId string) {
    r.AgId = agId
}
/* param isProtected: 实例在弹性伸缩组是否为保护状态，保护状态的实例缩容时不可以删除。(Required) */
func (r *ModifyInstancesProtectedRequest) SetIsProtected(isProtected bool) {
    r.IsProtected = isProtected
}
/* param instanceIds: 资源ID数组。(Required) */
func (r *ModifyInstancesProtectedRequest) SetInstanceIds(instanceIds []string) {
    r.InstanceIds = instanceIds
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyInstancesProtectedRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyInstancesProtectedResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyInstancesProtectedResult `json:"result"`
}

type ModifyInstancesProtectedResult struct {
}