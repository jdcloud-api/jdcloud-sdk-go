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
    ipanti "github.com/jdcloud-api/jdcloud-sdk-go/services/ipanti/models"
)

type DescribeIpSetRequest struct {

    core.JDCloudRequest

    /* 区域 ID, 高防不区分区域, 传 cn-north-1 即可  */
    RegionId string `json:"regionId"`

    /* 高防实例 Id  */
    InstanceId string `json:"instanceId"`

    /* IP 黑白名单 Id  */
    IpSetId string `json:"ipSetId"`
}

/*
 * param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可 (Required)
 * param instanceId: 高防实例 Id (Required)
 * param ipSetId: IP 黑白名单 Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeIpSetRequest(
    regionId string,
    instanceId string,
    ipSetId string,
) *DescribeIpSetRequest {

	return &DescribeIpSetRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/ipSets/{ipSetId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        IpSetId: ipSetId,
	}
}

/*
 * param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可 (Required)
 * param instanceId: 高防实例 Id (Required)
 * param ipSetId: IP 黑白名单 Id (Required)
 */
func NewDescribeIpSetRequestWithAllParams(
    regionId string,
    instanceId string,
    ipSetId string,
) *DescribeIpSetRequest {

    return &DescribeIpSetRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/ipSets/{ipSetId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        IpSetId: ipSetId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeIpSetRequestWithoutParam() *DescribeIpSetRequest {

    return &DescribeIpSetRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/ipSets/{ipSetId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可(Required) */
func (r *DescribeIpSetRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 高防实例 Id(Required) */
func (r *DescribeIpSetRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param ipSetId: IP 黑白名单 Id(Required) */
func (r *DescribeIpSetRequest) SetIpSetId(ipSetId string) {
    r.IpSetId = ipSetId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeIpSetRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeIpSetResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeIpSetResult `json:"result"`
}

type DescribeIpSetResult struct {
    Data ipanti.IpSet `json:"data"`
}