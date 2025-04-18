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
    shenhaiplatform "github.com/jdcloud-api/jdcloud-sdk-go/services/shenhaiplatform/models"
)

type CopilotChatHistoryRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 话题搜索词 (Optional) */
    TopicLike *string `json:"topicLike"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCopilotChatHistoryRequest(
    regionId string,
    appName string,
) *CopilotChatHistoryRequest {

	return &CopilotChatHistoryRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/copilotChatHistory",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        AppName: appName,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param topicLike: 话题搜索词 (Optional)
 */
func NewCopilotChatHistoryRequestWithAllParams(
    regionId string,
    appName string,
    topicLike *string,
) *CopilotChatHistoryRequest {

    return &CopilotChatHistoryRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/copilotChatHistory",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
        TopicLike: topicLike,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCopilotChatHistoryRequestWithoutParam() *CopilotChatHistoryRequest {

    return &CopilotChatHistoryRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/copilotChatHistory",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *CopilotChatHistoryRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *CopilotChatHistoryRequest) SetAppName(appName string) {
    r.AppName = appName
}
/* param topicLike: 话题搜索词(Optional) */
func (r *CopilotChatHistoryRequest) SetTopicLike(topicLike string) {
    r.TopicLike = &topicLike
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CopilotChatHistoryRequest) GetRegionId() string {
    return r.RegionId
}

type CopilotChatHistoryResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CopilotChatHistoryResult `json:"result"`
}

type CopilotChatHistoryResult struct {
    Code string `json:"code"`
    Message string `json:"message"`
    Result []shenhaiplatform.ChatVo `json:"result"`
}