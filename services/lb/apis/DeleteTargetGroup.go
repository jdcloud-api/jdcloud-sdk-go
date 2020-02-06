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

type DeleteTargetGroupRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* TargetGroup Id  */
    TargetGroupId string `json:"targetGroupId"`
}

/*
 * param regionId: Region ID (Required)
 * param targetGroupId: TargetGroup Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteTargetGroupRequest(
    regionId string,
    targetGroupId string,
) *DeleteTargetGroupRequest {

	return &DeleteTargetGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/targetGroups/{targetGroupId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        TargetGroupId: targetGroupId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param targetGroupId: TargetGroup Id (Required)
 */
func NewDeleteTargetGroupRequestWithAllParams(
    regionId string,
    targetGroupId string,
) *DeleteTargetGroupRequest {

    return &DeleteTargetGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/targetGroups/{targetGroupId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        TargetGroupId: targetGroupId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteTargetGroupRequestWithoutParam() *DeleteTargetGroupRequest {

    return &DeleteTargetGroupRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/targetGroups/{targetGroupId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DeleteTargetGroupRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param targetGroupId: TargetGroup Id(Required) */
func (r *DeleteTargetGroupRequest) SetTargetGroupId(targetGroupId string) {
    r.TargetGroupId = targetGroupId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteTargetGroupRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteTargetGroupResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteTargetGroupResult `json:"result"`
}

type DeleteTargetGroupResult struct {
}