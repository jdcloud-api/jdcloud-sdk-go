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
    lb "github.com/lshuining/jdcloud-sdk-go/services/lb/models"
)

type DescribeListenerRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 监听器ID  */
    ListenerId string `json:"listenerId"`
}

/*
 * param regionId: Region ID (Required)
 * param listenerId: 监听器ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeListenerRequest(
    regionId string,
    listenerId string,
) *DescribeListenerRequest {

	return &DescribeListenerRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/listeners/{listenerId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ListenerId: listenerId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param listenerId: 监听器ID (Required)
 */
func NewDescribeListenerRequestWithAllParams(
    regionId string,
    listenerId string,
) *DescribeListenerRequest {

    return &DescribeListenerRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/listeners/{listenerId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ListenerId: listenerId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeListenerRequestWithoutParam() *DescribeListenerRequest {

    return &DescribeListenerRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/listeners/{listenerId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeListenerRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param listenerId: 监听器ID(Required) */
func (r *DescribeListenerRequest) SetListenerId(listenerId string) {
    r.ListenerId = listenerId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeListenerRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeListenerResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeListenerResult `json:"result"`
}

type DescribeListenerResult struct {
    Listener lb.Listener `json:"listener"`
}