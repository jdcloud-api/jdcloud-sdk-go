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

type DescribeWebRuleBlackListUsageRequest struct {

    core.JDCloudRequest

    /* 区域 ID, 高防不区分区域, 传 cn-north-1 即可  */
    RegionId string `json:"regionId"`

    /* 高防实例 Id  */
    InstanceId string `json:"instanceId"`

    /* 网站规则 Id  */
    WebRuleId string `json:"webRuleId"`
}

/*
 * param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可 (Required)
 * param instanceId: 高防实例 Id (Required)
 * param webRuleId: 网站规则 Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeWebRuleBlackListUsageRequest(
    regionId string,
    instanceId string,
    webRuleId string,
) *DescribeWebRuleBlackListUsageRequest {

	return &DescribeWebRuleBlackListUsageRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/webRules/{webRuleId}:describeWebRuleBlackListUsage",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        WebRuleId: webRuleId,
	}
}

/*
 * param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可 (Required)
 * param instanceId: 高防实例 Id (Required)
 * param webRuleId: 网站规则 Id (Required)
 */
func NewDescribeWebRuleBlackListUsageRequestWithAllParams(
    regionId string,
    instanceId string,
    webRuleId string,
) *DescribeWebRuleBlackListUsageRequest {

    return &DescribeWebRuleBlackListUsageRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/webRules/{webRuleId}:describeWebRuleBlackListUsage",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        WebRuleId: webRuleId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeWebRuleBlackListUsageRequestWithoutParam() *DescribeWebRuleBlackListUsageRequest {

    return &DescribeWebRuleBlackListUsageRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/webRules/{webRuleId}:describeWebRuleBlackListUsage",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可(Required) */
func (r *DescribeWebRuleBlackListUsageRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 高防实例 Id(Required) */
func (r *DescribeWebRuleBlackListUsageRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param webRuleId: 网站规则 Id(Required) */
func (r *DescribeWebRuleBlackListUsageRequest) SetWebRuleId(webRuleId string) {
    r.WebRuleId = webRuleId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeWebRuleBlackListUsageRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeWebRuleBlackListUsageResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeWebRuleBlackListUsageResult `json:"result"`
}

type DescribeWebRuleBlackListUsageResult struct {
    AllocatedNum int `json:"allocatedNum"`
    ActiveNum int `json:"activeNum"`
    SurplusAllocateNum int `json:"surplusAllocateNum"`
    MaxAllocateNum int `json:"maxAllocateNum"`
}