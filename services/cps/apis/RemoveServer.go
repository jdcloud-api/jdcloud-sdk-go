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

type RemoveServerRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeCPSLBRegions）获取云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* 服务器组ID  */
    ServerGroupId string `json:"serverGroupId"`

    /* 后端服务器ID  */
    ServerId string `json:"serverId"`
}

/*
 * param regionId: 地域ID，可调用接口（describeCPSLBRegions）获取云物理服务器支持的地域 (Required)
 * param serverGroupId: 服务器组ID (Required)
 * param serverId: 后端服务器ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewRemoveServerRequest(
    regionId string,
    serverGroupId string,
    serverId string,
) *RemoveServerRequest {

	return &RemoveServerRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/serverGroups/{serverGroupId}/servers/{serverId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ServerGroupId: serverGroupId,
        ServerId: serverId,
	}
}

/*
 * param regionId: 地域ID，可调用接口（describeCPSLBRegions）获取云物理服务器支持的地域 (Required)
 * param serverGroupId: 服务器组ID (Required)
 * param serverId: 后端服务器ID (Required)
 */
func NewRemoveServerRequestWithAllParams(
    regionId string,
    serverGroupId string,
    serverId string,
) *RemoveServerRequest {

    return &RemoveServerRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/serverGroups/{serverGroupId}/servers/{serverId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ServerGroupId: serverGroupId,
        ServerId: serverId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewRemoveServerRequestWithoutParam() *RemoveServerRequest {

    return &RemoveServerRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/serverGroups/{serverGroupId}/servers/{serverId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeCPSLBRegions）获取云物理服务器支持的地域(Required) */
func (r *RemoveServerRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param serverGroupId: 服务器组ID(Required) */
func (r *RemoveServerRequest) SetServerGroupId(serverGroupId string) {
    r.ServerGroupId = serverGroupId
}

/* param serverId: 后端服务器ID(Required) */
func (r *RemoveServerRequest) SetServerId(serverId string) {
    r.ServerId = serverId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r RemoveServerRequest) GetRegionId() string {
    return r.RegionId
}

type RemoveServerResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result RemoveServerResult `json:"result"`
}

type RemoveServerResult struct {
    Success bool `json:"success"`
}