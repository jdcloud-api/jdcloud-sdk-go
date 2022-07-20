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
    vod "github.com/lshuining/jdcloud-sdk-go/services/vod/models"
)

type GetTranscodeTemplateGroupRequest struct {

    core.JDCloudRequest

    /* 模板组ID  */
    GroupId string `json:"groupId"`
}

/*
 * param groupId: 模板组ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetTranscodeTemplateGroupRequest(
    groupId string,
) *GetTranscodeTemplateGroupRequest {

	return &GetTranscodeTemplateGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/transcodeTemplateGroups/{groupId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        GroupId: groupId,
	}
}

/*
 * param groupId: 模板组ID (Required)
 */
func NewGetTranscodeTemplateGroupRequestWithAllParams(
    groupId string,
) *GetTranscodeTemplateGroupRequest {

    return &GetTranscodeTemplateGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/transcodeTemplateGroups/{groupId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        GroupId: groupId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetTranscodeTemplateGroupRequestWithoutParam() *GetTranscodeTemplateGroupRequest {

    return &GetTranscodeTemplateGroupRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/transcodeTemplateGroups/{groupId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param groupId: 模板组ID(Required) */
func (r *GetTranscodeTemplateGroupRequest) SetGroupId(groupId string) {
    r.GroupId = groupId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetTranscodeTemplateGroupRequest) GetRegionId() string {
    return ""
}

type GetTranscodeTemplateGroupResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetTranscodeTemplateGroupResult `json:"result"`
}

type GetTranscodeTemplateGroupResult struct {
    GroupId string `json:"groupId"`
    GroupName string `json:"groupName"`
    Templates []vod.GroupedTranscodeTemplateData `json:"templates"`
    CreateTime string `json:"createTime"`
    UpdateTime string `json:"updateTime"`
}