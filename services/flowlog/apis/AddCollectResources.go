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
    flowlog "github.com/jdcloud-api/jdcloud-sdk-go/services/flowlog/models"
)

type AddCollectResourcesRequest struct {

    core.JDCloudRequest

    /* 地域 ID  */
    RegionId string `json:"regionId"`

    /* 流日志ID  */
    FlowLogId string `json:"flowLogId"`

    /* 采集资源列表  */
    CollectResources []flowlog.CollectResourceSpec `json:"collectResources"`
}

/*
 * param regionId: 地域 ID (Required)
 * param flowLogId: 流日志ID (Required)
 * param collectResources: 采集资源列表 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAddCollectResourcesRequest(
    regionId string,
    flowLogId string,
    collectResources []flowlog.CollectResourceSpec,
) *AddCollectResourcesRequest {

	return &AddCollectResourcesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/flowLogs/{flowLogId}:addCollectResources",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        FlowLogId: flowLogId,
        CollectResources: collectResources,
	}
}

/*
 * param regionId: 地域 ID (Required)
 * param flowLogId: 流日志ID (Required)
 * param collectResources: 采集资源列表 (Required)
 */
func NewAddCollectResourcesRequestWithAllParams(
    regionId string,
    flowLogId string,
    collectResources []flowlog.CollectResourceSpec,
) *AddCollectResourcesRequest {

    return &AddCollectResourcesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/flowLogs/{flowLogId}:addCollectResources",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        FlowLogId: flowLogId,
        CollectResources: collectResources,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAddCollectResourcesRequestWithoutParam() *AddCollectResourcesRequest {

    return &AddCollectResourcesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/flowLogs/{flowLogId}:addCollectResources",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 ID(Required) */
func (r *AddCollectResourcesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param flowLogId: 流日志ID(Required) */
func (r *AddCollectResourcesRequest) SetFlowLogId(flowLogId string) {
    r.FlowLogId = flowLogId
}
/* param collectResources: 采集资源列表(Required) */
func (r *AddCollectResourcesRequest) SetCollectResources(collectResources []flowlog.CollectResourceSpec) {
    r.CollectResources = collectResources
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AddCollectResourcesRequest) GetRegionId() string {
    return r.RegionId
}

type AddCollectResourcesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AddCollectResourcesResult `json:"result"`
}

type AddCollectResourcesResult struct {
}