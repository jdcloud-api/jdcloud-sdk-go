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
    logs "github.com/lshuining/jdcloud-sdk-go/services/logs/models"
)

type UpdateCollectResourcesRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 采集配置 UID  */
    CollectInfoUID string `json:"collectInfoUID"`

    /* action  */
    Action string `json:"action"`

    /* 采集实例列表（系统日志存在上限限制） (Optional) */
    Resources []logs.Resource `json:"resources"`
}

/*
 * param regionId: 地域 Id (Required)
 * param collectInfoUID: 采集配置 UID (Required)
 * param action: action (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateCollectResourcesRequest(
    regionId string,
    collectInfoUID string,
    action string,
) *UpdateCollectResourcesRequest {

	return &UpdateCollectResourcesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/collectinfos/{collectInfoUID}:updateResources",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        CollectInfoUID: collectInfoUID,
        Action: action,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param collectInfoUID: 采集配置 UID (Required)
 * param action: action (Required)
 * param resources: 采集实例列表（系统日志存在上限限制） (Optional)
 */
func NewUpdateCollectResourcesRequestWithAllParams(
    regionId string,
    collectInfoUID string,
    action string,
    resources []logs.Resource,
) *UpdateCollectResourcesRequest {

    return &UpdateCollectResourcesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/collectinfos/{collectInfoUID}:updateResources",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        CollectInfoUID: collectInfoUID,
        Action: action,
        Resources: resources,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateCollectResourcesRequestWithoutParam() *UpdateCollectResourcesRequest {

    return &UpdateCollectResourcesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/collectinfos/{collectInfoUID}:updateResources",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *UpdateCollectResourcesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param collectInfoUID: 采集配置 UID(Required) */
func (r *UpdateCollectResourcesRequest) SetCollectInfoUID(collectInfoUID string) {
    r.CollectInfoUID = collectInfoUID
}

/* param action: action(Required) */
func (r *UpdateCollectResourcesRequest) SetAction(action string) {
    r.Action = action
}

/* param resources: 采集实例列表（系统日志存在上限限制）(Optional) */
func (r *UpdateCollectResourcesRequest) SetResources(resources []logs.Resource) {
    r.Resources = resources
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateCollectResourcesRequest) GetRegionId() string {
    return r.RegionId
}

type UpdateCollectResourcesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateCollectResourcesResult `json:"result"`
}

type UpdateCollectResourcesResult struct {
}