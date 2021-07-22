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

type RemoveRoomUserRequest struct {

    core.JDCloudRequest

    /* 应用ID  */
    AppId string `json:"appId"`

    /* 房间ID  */
    RoomId int `json:"roomId"`

    /* peerId列表,最多支持20个peerId (Optional) */
    PeerIds []int64 `json:"peerIds"`
}

/*
 * param appId: 应用ID (Required)
 * param roomId: 房间ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewRemoveRoomUserRequest(
    appId string,
    roomId int,
) *RemoveRoomUserRequest {

	return &RemoveRoomUserRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/roomUser/{appId}/removeUser/{roomId}",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        AppId: appId,
        RoomId: roomId,
	}
}

/*
 * param appId: 应用ID (Required)
 * param roomId: 房间ID (Required)
 * param peerIds: peerId列表,最多支持20个peerId (Optional)
 */
func NewRemoveRoomUserRequestWithAllParams(
    appId string,
    roomId int,
    peerIds []int64,
) *RemoveRoomUserRequest {

    return &RemoveRoomUserRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/roomUser/{appId}/removeUser/{roomId}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        AppId: appId,
        RoomId: roomId,
        PeerIds: peerIds,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewRemoveRoomUserRequestWithoutParam() *RemoveRoomUserRequest {

    return &RemoveRoomUserRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/roomUser/{appId}/removeUser/{roomId}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param appId: 应用ID(Required) */
func (r *RemoveRoomUserRequest) SetAppId(appId string) {
    r.AppId = appId
}

/* param roomId: 房间ID(Required) */
func (r *RemoveRoomUserRequest) SetRoomId(roomId int) {
    r.RoomId = roomId
}

/* param peerIds: peerId列表,最多支持20个peerId(Optional) */
func (r *RemoveRoomUserRequest) SetPeerIds(peerIds []int64) {
    r.PeerIds = peerIds
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r RemoveRoomUserRequest) GetRegionId() string {
    return ""
}

type RemoveRoomUserResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result RemoveRoomUserResult `json:"result"`
}

type RemoveRoomUserResult struct {
}