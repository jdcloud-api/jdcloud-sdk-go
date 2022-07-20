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

type CreatePortSetRequest struct {

    core.JDCloudRequest

    /* 地域 Id, DDoS 防护包目前支持华北-北京, 华东-宿迁, 华东-上海  */
    RegionId string `json:"regionId"`

    /* 防护包实例 Id  */
    InstanceId string `json:"instanceId"`

    /* 创建实例的端口库请求参数  */
    PortSetSpec *antipro.PortSetSpec `json:"portSetSpec"`
}

/*
 * param regionId: 地域 Id, DDoS 防护包目前支持华北-北京, 华东-宿迁, 华东-上海 (Required)
 * param instanceId: 防护包实例 Id (Required)
 * param portSetSpec: 创建实例的端口库请求参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreatePortSetRequest(
    regionId string,
    instanceId string,
    portSetSpec *antipro.PortSetSpec,
) *CreatePortSetRequest {

	return &CreatePortSetRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/portSets",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        PortSetSpec: portSetSpec,
	}
}

/*
 * param regionId: 地域 Id, DDoS 防护包目前支持华北-北京, 华东-宿迁, 华东-上海 (Required)
 * param instanceId: 防护包实例 Id (Required)
 * param portSetSpec: 创建实例的端口库请求参数 (Required)
 */
func NewCreatePortSetRequestWithAllParams(
    regionId string,
    instanceId string,
    portSetSpec *antipro.PortSetSpec,
) *CreatePortSetRequest {

    return &CreatePortSetRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/portSets",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        PortSetSpec: portSetSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreatePortSetRequestWithoutParam() *CreatePortSetRequest {

    return &CreatePortSetRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/portSets",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id, DDoS 防护包目前支持华北-北京, 华东-宿迁, 华东-上海(Required) */
func (r *CreatePortSetRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 防护包实例 Id(Required) */
func (r *CreatePortSetRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param portSetSpec: 创建实例的端口库请求参数(Required) */
func (r *CreatePortSetRequest) SetPortSetSpec(portSetSpec *antipro.PortSetSpec) {
    r.PortSetSpec = portSetSpec
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreatePortSetRequest) GetRegionId() string {
    return r.RegionId
}

type CreatePortSetResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreatePortSetResult `json:"result"`
}

type CreatePortSetResult struct {
    PortSetId string `json:"portSetId"`
}