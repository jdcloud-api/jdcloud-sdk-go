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

type RebootPodRequest struct {

    core.JDCloudRequest

    /* 地域代码  */
    RegionId string `json:"regionId"`

    /* 实例ID  */
    InstanceId string `json:"instanceId"`

    /* 重启指定类型的pod,支持Tikv,TiDB,PD,TiFlash  */
    NodeType []string `json:"nodeType"`
}

/*
 * param regionId: 地域代码 (Required)
 * param instanceId: 实例ID (Required)
 * param nodeType: 重启指定类型的pod,支持Tikv,TiDB,PD,TiFlash (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewRebootPodRequest(
    regionId string,
    instanceId string,
    nodeType []string,
) *RebootPodRequest {

	return &RebootPodRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:rebootpod",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        NodeType: nodeType,
	}
}

/*
 * param regionId: 地域代码 (Required)
 * param instanceId: 实例ID (Required)
 * param nodeType: 重启指定类型的pod,支持Tikv,TiDB,PD,TiFlash (Required)
 */
func NewRebootPodRequestWithAllParams(
    regionId string,
    instanceId string,
    nodeType []string,
) *RebootPodRequest {

    return &RebootPodRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:rebootpod",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        NodeType: nodeType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewRebootPodRequestWithoutParam() *RebootPodRequest {

    return &RebootPodRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:rebootpod",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码(Required) */
func (r *RebootPodRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 实例ID(Required) */
func (r *RebootPodRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param nodeType: 重启指定类型的pod,支持Tikv,TiDB,PD,TiFlash(Required) */
func (r *RebootPodRequest) SetNodeType(nodeType []string) {
    r.NodeType = nodeType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r RebootPodRequest) GetRegionId() string {
    return r.RegionId
}

type RebootPodResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result RebootPodResult `json:"result"`
}

type RebootPodResult struct {
}