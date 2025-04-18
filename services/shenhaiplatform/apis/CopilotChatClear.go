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

type CopilotChatClearRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 对话id  */
    ChatId string `json:"chatId"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param chatId: 对话id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCopilotChatClearRequest(
    regionId string,
    appName string,
    chatId string,
) *CopilotChatClearRequest {

	return &CopilotChatClearRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/copilotChatClear",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        AppName: appName,
        ChatId: chatId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param chatId: 对话id (Required)
 */
func NewCopilotChatClearRequestWithAllParams(
    regionId string,
    appName string,
    chatId string,
) *CopilotChatClearRequest {

    return &CopilotChatClearRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/copilotChatClear",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
        ChatId: chatId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCopilotChatClearRequestWithoutParam() *CopilotChatClearRequest {

    return &CopilotChatClearRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/copilotChatClear",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *CopilotChatClearRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *CopilotChatClearRequest) SetAppName(appName string) {
    r.AppName = appName
}
/* param chatId: 对话id(Required) */
func (r *CopilotChatClearRequest) SetChatId(chatId string) {
    r.ChatId = chatId
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CopilotChatClearRequest) GetRegionId() string {
    return r.RegionId
}

type CopilotChatClearResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CopilotChatClearResult `json:"result"`
}

type CopilotChatClearResult struct {
    Code string `json:"code"`
    Message string `json:"message"`
}