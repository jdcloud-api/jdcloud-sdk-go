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
    smartdba "github.com/jdcloud-api/jdcloud-sdk-go/services/smartdba/models"
)

type DescribeSystemInfoRequest struct {

    core.JDCloudRequest

    /* 地域代码  */
    RegionId string `json:"regionId"`

    /* RDS 实例ID，唯一标识一个RDS实例  */
    InstanceGid string `json:"instanceGid"`
}

/*
 * param regionId: 地域代码 (Required)
 * param instanceGid: RDS 实例ID，唯一标识一个RDS实例 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeSystemInfoRequest(
    regionId string,
    instanceGid string,
) *DescribeSystemInfoRequest {

	return &DescribeSystemInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instance/{instanceGid}/systemInfo",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        InstanceGid: instanceGid,
	}
}

/*
 * param regionId: 地域代码 (Required)
 * param instanceGid: RDS 实例ID，唯一标识一个RDS实例 (Required)
 */
func NewDescribeSystemInfoRequestWithAllParams(
    regionId string,
    instanceGid string,
) *DescribeSystemInfoRequest {

    return &DescribeSystemInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instance/{instanceGid}/systemInfo",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        InstanceGid: instanceGid,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeSystemInfoRequestWithoutParam() *DescribeSystemInfoRequest {

    return &DescribeSystemInfoRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instance/{instanceGid}/systemInfo",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域代码(Required) */
func (r *DescribeSystemInfoRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceGid: RDS 实例ID，唯一标识一个RDS实例(Required) */
func (r *DescribeSystemInfoRequest) SetInstanceGid(instanceGid string) {
    r.InstanceGid = instanceGid
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeSystemInfoRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeSystemInfoResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeSystemInfoResult `json:"result"`
}

type DescribeSystemInfoResult struct {
    BufferPoolInfo smartdba.BufferPoolInfo `json:"bufferPoolInfo"`
    ConnectionInfo smartdba.ConnectionInfo `json:"connectionInfo"`
    ServerInfo smartdba.ServerInfo `json:"serverInfo"`
}