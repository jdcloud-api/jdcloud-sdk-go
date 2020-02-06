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

type UpdateTargetGroupRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* TargetGroup Id  */
    TargetGroupId string `json:"targetGroupId"`

    /* 虚拟服务器组描述,允许输入UTF-8编码下的全部字符，不超过256字符 (Optional) */
    Description *string `json:"description"`

    /* 虚拟服务器组名称,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符 (Optional) */
    TargetGroupName *string `json:"targetGroupName"`
}

/*
 * param regionId: Region ID (Required)
 * param targetGroupId: TargetGroup Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateTargetGroupRequest(
    regionId string,
    targetGroupId string,
) *UpdateTargetGroupRequest {

	return &UpdateTargetGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/targetGroups/{targetGroupId}",
			Method:  "PATCH",
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
 * param description: 虚拟服务器组描述,允许输入UTF-8编码下的全部字符，不超过256字符 (Optional)
 * param targetGroupName: 虚拟服务器组名称,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符 (Optional)
 */
func NewUpdateTargetGroupRequestWithAllParams(
    regionId string,
    targetGroupId string,
    description *string,
    targetGroupName *string,
) *UpdateTargetGroupRequest {

    return &UpdateTargetGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/targetGroups/{targetGroupId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        TargetGroupId: targetGroupId,
        Description: description,
        TargetGroupName: targetGroupName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateTargetGroupRequestWithoutParam() *UpdateTargetGroupRequest {

    return &UpdateTargetGroupRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/targetGroups/{targetGroupId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *UpdateTargetGroupRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param targetGroupId: TargetGroup Id(Required) */
func (r *UpdateTargetGroupRequest) SetTargetGroupId(targetGroupId string) {
    r.TargetGroupId = targetGroupId
}

/* param description: 虚拟服务器组描述,允许输入UTF-8编码下的全部字符，不超过256字符(Optional) */
func (r *UpdateTargetGroupRequest) SetDescription(description string) {
    r.Description = &description
}

/* param targetGroupName: 虚拟服务器组名称,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符(Optional) */
func (r *UpdateTargetGroupRequest) SetTargetGroupName(targetGroupName string) {
    r.TargetGroupName = &targetGroupName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateTargetGroupRequest) GetRegionId() string {
    return r.RegionId
}

type UpdateTargetGroupResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateTargetGroupResult `json:"result"`
}

type UpdateTargetGroupResult struct {
}