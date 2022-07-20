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
    antipro "github.com/lshuining/jdcloud-sdk-go/services/antipro/models"
)

type ModifyInstanceRequest struct {

    core.JDCloudRequest

    /* 地域 Id, DDoS 防护包目前支持华北-北京, 华东-宿迁, 华东-上海  */
    RegionId string `json:"regionId"`

    /* 防护包实例 Id  */
    InstanceId string `json:"instanceId"`

    /* 升级防护包实例请求参数  */
    ModifyInstanceSpec *antipro.ModifyInstanceSpec `json:"modifyInstanceSpec"`
}

/*
 * param regionId: 地域 Id, DDoS 防护包目前支持华北-北京, 华东-宿迁, 华东-上海 (Required)
 * param instanceId: 防护包实例 Id (Required)
 * param modifyInstanceSpec: 升级防护包实例请求参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyInstanceRequest(
    regionId string,
    instanceId string,
    modifyInstanceSpec *antipro.ModifyInstanceSpec,
) *ModifyInstanceRequest {

	return &ModifyInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        ModifyInstanceSpec: modifyInstanceSpec,
	}
}

/*
 * param regionId: 地域 Id, DDoS 防护包目前支持华北-北京, 华东-宿迁, 华东-上海 (Required)
 * param instanceId: 防护包实例 Id (Required)
 * param modifyInstanceSpec: 升级防护包实例请求参数 (Required)
 */
func NewModifyInstanceRequestWithAllParams(
    regionId string,
    instanceId string,
    modifyInstanceSpec *antipro.ModifyInstanceSpec,
) *ModifyInstanceRequest {

    return &ModifyInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        ModifyInstanceSpec: modifyInstanceSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyInstanceRequestWithoutParam() *ModifyInstanceRequest {

    return &ModifyInstanceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id, DDoS 防护包目前支持华北-北京, 华东-宿迁, 华东-上海(Required) */
func (r *ModifyInstanceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 防护包实例 Id(Required) */
func (r *ModifyInstanceRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param modifyInstanceSpec: 升级防护包实例请求参数(Required) */
func (r *ModifyInstanceRequest) SetModifyInstanceSpec(modifyInstanceSpec *antipro.ModifyInstanceSpec) {
    r.ModifyInstanceSpec = modifyInstanceSpec
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyInstanceRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyInstanceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyInstanceResult `json:"result"`
}

type ModifyInstanceResult struct {
    InstanceId string `json:"instanceId"`
}