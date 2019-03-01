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
    ipanti "github.com/jdcloud-api/jdcloud-sdk-go/services/ipanti/models"
)

type ModifyInstanceNameRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 实例 ID  */
    InstanceId int `json:"instanceId"`

    /* 修改实例名称请求参数  */
    RenameInstanceSpec *ipanti.RenameInstanceSpec `json:"renameInstanceSpec"`
}

/*
 * param regionId: Region ID (Required)
 * param instanceId: 实例 ID (Required)
 * param renameInstanceSpec: 修改实例名称请求参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyInstanceNameRequest(
    regionId string,
    instanceId int,
    renameInstanceSpec *ipanti.RenameInstanceSpec,
) *ModifyInstanceNameRequest {

	return &ModifyInstanceNameRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:rename",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        RenameInstanceSpec: renameInstanceSpec,
	}
}

/*
 * param regionId: Region ID (Required)
 * param instanceId: 实例 ID (Required)
 * param renameInstanceSpec: 修改实例名称请求参数 (Required)
 */
func NewModifyInstanceNameRequestWithAllParams(
    regionId string,
    instanceId int,
    renameInstanceSpec *ipanti.RenameInstanceSpec,
) *ModifyInstanceNameRequest {

    return &ModifyInstanceNameRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:rename",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        RenameInstanceSpec: renameInstanceSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyInstanceNameRequestWithoutParam() *ModifyInstanceNameRequest {

    return &ModifyInstanceNameRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:rename",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *ModifyInstanceNameRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 实例 ID(Required) */
func (r *ModifyInstanceNameRequest) SetInstanceId(instanceId int) {
    r.InstanceId = instanceId
}

/* param renameInstanceSpec: 修改实例名称请求参数(Required) */
func (r *ModifyInstanceNameRequest) SetRenameInstanceSpec(renameInstanceSpec *ipanti.RenameInstanceSpec) {
    r.RenameInstanceSpec = renameInstanceSpec
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyInstanceNameRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyInstanceNameResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyInstanceNameResult `json:"result"`
}

type ModifyInstanceNameResult struct {
    Code int `json:"code"`
    Message string `json:"message"`
}