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

type ModifyInstanceRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* 云物理服务器ID  */
    InstanceId string `json:"instanceId"`

    /* 云物理服务器名称 (Optional) */
    Name *string `json:"name"`

    /* 云物理服务器描述 (Optional) */
    Description *string `json:"description"`
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
 * param instanceId: 云物理服务器ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyInstanceRequest(
    regionId string,
    instanceId string,
) *ModifyInstanceRequest {

	return &ModifyInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:modifyInstance",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
 * param instanceId: 云物理服务器ID (Required)
 * param name: 云物理服务器名称 (Optional)
 * param description: 云物理服务器描述 (Optional)
 */
func NewModifyInstanceRequestWithAllParams(
    regionId string,
    instanceId string,
    name *string,
    description *string,
) *ModifyInstanceRequest {

    return &ModifyInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:modifyInstance",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        Name: name,
        Description: description,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyInstanceRequestWithoutParam() *ModifyInstanceRequest {

    return &ModifyInstanceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:modifyInstance",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域(Required) */
func (r *ModifyInstanceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 云物理服务器ID(Required) */
func (r *ModifyInstanceRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param name: 云物理服务器名称(Optional) */
func (r *ModifyInstanceRequest) SetName(name string) {
    r.Name = &name
}

/* param description: 云物理服务器描述(Optional) */
func (r *ModifyInstanceRequest) SetDescription(description string) {
    r.Description = &description
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
    Name string `json:"name"`
    Description string `json:"description"`
}