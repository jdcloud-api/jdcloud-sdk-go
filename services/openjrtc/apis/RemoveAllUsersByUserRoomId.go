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

type RemoveAllUsersByUserRoomIdRequest struct {

    core.JDCloudRequest

    /* 应用ID  */
    AppId string `json:"appId"`

    /* 业务接入方定义的且在JRTC系统内注册过的房间号 (Optional) */
    UserRoomId *string `json:"userRoomId"`
}

/*
 * param appId: 应用ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewRemoveAllUsersByUserRoomIdRequest(
    appId string,
) *RemoveAllUsersByUserRoomIdRequest {

	return &RemoveAllUsersByUserRoomIdRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/removeAllUsersByUserRoomId/{appId}",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        AppId: appId,
	}
}

/*
 * param appId: 应用ID (Required)
 * param userRoomId: 业务接入方定义的且在JRTC系统内注册过的房间号 (Optional)
 */
func NewRemoveAllUsersByUserRoomIdRequestWithAllParams(
    appId string,
    userRoomId *string,
) *RemoveAllUsersByUserRoomIdRequest {

    return &RemoveAllUsersByUserRoomIdRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/removeAllUsersByUserRoomId/{appId}",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        AppId: appId,
        UserRoomId: userRoomId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewRemoveAllUsersByUserRoomIdRequestWithoutParam() *RemoveAllUsersByUserRoomIdRequest {

    return &RemoveAllUsersByUserRoomIdRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/removeAllUsersByUserRoomId/{appId}",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param appId: 应用ID(Required) */
func (r *RemoveAllUsersByUserRoomIdRequest) SetAppId(appId string) {
    r.AppId = appId
}

/* param userRoomId: 业务接入方定义的且在JRTC系统内注册过的房间号(Optional) */
func (r *RemoveAllUsersByUserRoomIdRequest) SetUserRoomId(userRoomId string) {
    r.UserRoomId = &userRoomId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r RemoveAllUsersByUserRoomIdRequest) GetRegionId() string {
    return ""
}

type RemoveAllUsersByUserRoomIdResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result RemoveAllUsersByUserRoomIdResult `json:"result"`
}

type RemoveAllUsersByUserRoomIdResult struct {
}