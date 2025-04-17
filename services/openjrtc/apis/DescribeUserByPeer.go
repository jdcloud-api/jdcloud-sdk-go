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

type DescribeUserByPeerRequest struct {

    core.JDCloudRequest

    /* peerId  */
    PeerId int `json:"peerId"`
}

/*
 * param peerId: peerId (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeUserByPeerRequest(
    peerId int,
) *DescribeUserByPeerRequest {

	return &DescribeUserByPeerRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/describeUserByPeer",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        PeerId: peerId,
	}
}

/*
 * param peerId: peerId (Required)
 */
func NewDescribeUserByPeerRequestWithAllParams(
    peerId int,
) *DescribeUserByPeerRequest {

    return &DescribeUserByPeerRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/describeUserByPeer",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PeerId: peerId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeUserByPeerRequestWithoutParam() *DescribeUserByPeerRequest {

    return &DescribeUserByPeerRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/describeUserByPeer",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param peerId: peerId(Required) */
func (r *DescribeUserByPeerRequest) SetPeerId(peerId int) {
    r.PeerId = peerId
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeUserByPeerRequest) GetRegionId() string {
    return ""
}

type DescribeUserByPeerResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeUserByPeerResult `json:"result"`
}

type DescribeUserByPeerResult struct {
    PeerId int64 `json:"peerId"`
    AppId string `json:"appId"`
    UserId string `json:"userId"`
    Temporary bool `json:"temporary"`
    CreateTime string `json:"createTime"`
}