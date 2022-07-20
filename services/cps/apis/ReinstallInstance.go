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
    cps "github.com/lshuining/jdcloud-sdk-go/services/cps/models"
)

type ReinstallInstanceRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* 云物理服务器ID  */
    InstanceId string `json:"instanceId"`

    /* 由客户端生成，用于保证请求的幂等性，长度不能超过36个字符；<br/>
如果多个请求使用了相同的clientToken，只会执行第一个请求，之后的请求直接返回第一个请求的结果<br/>
 (Optional) */
    ClientToken *string `json:"clientToken"`

    /* 云物理服务器配置  */
    InstanceSpec *cps.ReinstallInstanceSpec `json:"instanceSpec"`
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
 * param instanceId: 云物理服务器ID (Required)
 * param instanceSpec: 云物理服务器配置 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewReinstallInstanceRequest(
    regionId string,
    instanceId string,
    instanceSpec *cps.ReinstallInstanceSpec,
) *ReinstallInstanceRequest {

	return &ReinstallInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:reinstallInstance",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        InstanceSpec: instanceSpec,
	}
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
 * param instanceId: 云物理服务器ID (Required)
 * param clientToken: 由客户端生成，用于保证请求的幂等性，长度不能超过36个字符；<br/>
如果多个请求使用了相同的clientToken，只会执行第一个请求，之后的请求直接返回第一个请求的结果<br/>
 (Optional)
 * param instanceSpec: 云物理服务器配置 (Required)
 */
func NewReinstallInstanceRequestWithAllParams(
    regionId string,
    instanceId string,
    clientToken *string,
    instanceSpec *cps.ReinstallInstanceSpec,
) *ReinstallInstanceRequest {

    return &ReinstallInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:reinstallInstance",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        ClientToken: clientToken,
        InstanceSpec: instanceSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewReinstallInstanceRequestWithoutParam() *ReinstallInstanceRequest {

    return &ReinstallInstanceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:reinstallInstance",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域(Required) */
func (r *ReinstallInstanceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 云物理服务器ID(Required) */
func (r *ReinstallInstanceRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param clientToken: 由客户端生成，用于保证请求的幂等性，长度不能超过36个字符；<br/>
如果多个请求使用了相同的clientToken，只会执行第一个请求，之后的请求直接返回第一个请求的结果<br/>
(Optional) */
func (r *ReinstallInstanceRequest) SetClientToken(clientToken string) {
    r.ClientToken = &clientToken
}

/* param instanceSpec: 云物理服务器配置(Required) */
func (r *ReinstallInstanceRequest) SetInstanceSpec(instanceSpec *cps.ReinstallInstanceSpec) {
    r.InstanceSpec = instanceSpec
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ReinstallInstanceRequest) GetRegionId() string {
    return r.RegionId
}

type ReinstallInstanceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ReinstallInstanceResult `json:"result"`
}

type ReinstallInstanceResult struct {
    Success bool `json:"success"`
}