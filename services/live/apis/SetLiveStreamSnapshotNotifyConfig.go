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

type SetLiveStreamSnapshotNotifyConfigRequest struct {

    core.JDCloudRequest

    /* 您的推流加速域名  */
    PublishDomain string `json:"publishDomain"`

    /* 设置直播流信息推送到的 URL 地址:
  - 以 http:// 开头
  - 正则校验
  */
    NotifyUrl string `json:"notifyUrl"`
}

/*
 * param publishDomain: 您的推流加速域名 (Required)
 * param notifyUrl: 设置直播流信息推送到的 URL 地址:
  - 以 http:// 开头
  - 正则校验
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSetLiveStreamSnapshotNotifyConfigRequest(
    publishDomain string,
    notifyUrl string,
) *SetLiveStreamSnapshotNotifyConfigRequest {

	return &SetLiveStreamSnapshotNotifyConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/snapshotNotifys:config",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        PublishDomain: publishDomain,
        NotifyUrl: notifyUrl,
	}
}

/*
 * param publishDomain: 您的推流加速域名 (Required)
 * param notifyUrl: 设置直播流信息推送到的 URL 地址:
  - 以 http:// 开头
  - 正则校验
 (Required)
 */
func NewSetLiveStreamSnapshotNotifyConfigRequestWithAllParams(
    publishDomain string,
    notifyUrl string,
) *SetLiveStreamSnapshotNotifyConfigRequest {

    return &SetLiveStreamSnapshotNotifyConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/snapshotNotifys:config",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        PublishDomain: publishDomain,
        NotifyUrl: notifyUrl,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSetLiveStreamSnapshotNotifyConfigRequestWithoutParam() *SetLiveStreamSnapshotNotifyConfigRequest {

    return &SetLiveStreamSnapshotNotifyConfigRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/snapshotNotifys:config",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param publishDomain: 您的推流加速域名(Required) */
func (r *SetLiveStreamSnapshotNotifyConfigRequest) SetPublishDomain(publishDomain string) {
    r.PublishDomain = publishDomain
}

/* param notifyUrl: 设置直播流信息推送到的 URL 地址:
  - 以 http:// 开头
  - 正则校验
(Required) */
func (r *SetLiveStreamSnapshotNotifyConfigRequest) SetNotifyUrl(notifyUrl string) {
    r.NotifyUrl = notifyUrl
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SetLiveStreamSnapshotNotifyConfigRequest) GetRegionId() string {
    return ""
}

type SetLiveStreamSnapshotNotifyConfigResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SetLiveStreamSnapshotNotifyConfigResult `json:"result"`
}

type SetLiveStreamSnapshotNotifyConfigResult struct {
}