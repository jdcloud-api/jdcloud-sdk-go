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

type DeleteCommandsRequest struct {

    core.JDCloudRequest

    /* 地域ID。  */
    RegionId string `json:"regionId"`

    /* 用户创建的命令Id
  */
    CommandIds []string `json:"commandIds"`
}

/*
 * param regionId: 地域ID。 (Required)
 * param commandIds: 用户创建的命令Id
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteCommandsRequest(
    regionId string,
    commandIds []string,
) *DeleteCommandsRequest {

	return &DeleteCommandsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/deleteCommands",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        CommandIds: commandIds,
	}
}

/*
 * param regionId: 地域ID。 (Required)
 * param commandIds: 用户创建的命令Id
 (Required)
 */
func NewDeleteCommandsRequestWithAllParams(
    regionId string,
    commandIds []string,
) *DeleteCommandsRequest {

    return &DeleteCommandsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/deleteCommands",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        CommandIds: commandIds,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteCommandsRequestWithoutParam() *DeleteCommandsRequest {

    return &DeleteCommandsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/deleteCommands",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID。(Required) */
func (r *DeleteCommandsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param commandIds: 用户创建的命令Id
(Required) */
func (r *DeleteCommandsRequest) SetCommandIds(commandIds []string) {
    r.CommandIds = commandIds
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteCommandsRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteCommandsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteCommandsResult `json:"result"`
}

type DeleteCommandsResult struct {
    CommandId string `json:"commandId"`
}