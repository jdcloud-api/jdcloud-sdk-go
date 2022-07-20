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

type DescribeStreamInfoByStreamIdRequest struct {

    core.JDCloudRequest

    /* 应用ID  */
    AppId string `json:"appId"`

    /* 流ID  */
    StreamId string `json:"streamId"`
}

/*
 * param appId: 应用ID (Required)
 * param streamId: 流ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeStreamInfoByStreamIdRequest(
    appId string,
    streamId string,
) *DescribeStreamInfoByStreamIdRequest {

	return &DescribeStreamInfoByStreamIdRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/describeStreamInfoByStreamId/{appId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        AppId: appId,
        StreamId: streamId,
	}
}

/*
 * param appId: 应用ID (Required)
 * param streamId: 流ID (Required)
 */
func NewDescribeStreamInfoByStreamIdRequestWithAllParams(
    appId string,
    streamId string,
) *DescribeStreamInfoByStreamIdRequest {

    return &DescribeStreamInfoByStreamIdRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/describeStreamInfoByStreamId/{appId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        AppId: appId,
        StreamId: streamId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeStreamInfoByStreamIdRequestWithoutParam() *DescribeStreamInfoByStreamIdRequest {

    return &DescribeStreamInfoByStreamIdRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/describeStreamInfoByStreamId/{appId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param appId: 应用ID(Required) */
func (r *DescribeStreamInfoByStreamIdRequest) SetAppId(appId string) {
    r.AppId = appId
}

/* param streamId: 流ID(Required) */
func (r *DescribeStreamInfoByStreamIdRequest) SetStreamId(streamId string) {
    r.StreamId = streamId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeStreamInfoByStreamIdRequest) GetRegionId() string {
    return ""
}

type DescribeStreamInfoByStreamIdResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeStreamInfoByStreamIdResult `json:"result"`
}

type DescribeStreamInfoByStreamIdResult struct {
    AppId string `json:"appId"`
    UserRoomId string `json:"userRoomId"`
    RoomName string `json:"roomName"`
    UserId string `json:"userId"`
    UserName string `json:"userName"`
    NickName string `json:"nickName"`
    StreamId string `json:"streamId"`
    StreamName string `json:"streamName"`
    Kind int `json:"kind"`
    Status int `json:"status"`
    DeviceType int `json:"deviceType"`
    PublishCount int `json:"publishCount"`
    PublishTime string `json:"publishTime"`
    CreateTime string `json:"createTime"`
}