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

type SetLiveStreamQualityDetectionNotifyConfigRequest struct {

    core.JDCloudRequest

    /* 推流域名  */
    PublishDomain string `json:"publishDomain"`

    /* 质量检测回调通知的URL地址:
- 以 http:// 开头,外网可访问的地址
  */
    NotifyUrl string `json:"notifyUrl"`
}

/*
 * param publishDomain: 推流域名 (Required)
 * param notifyUrl: 质量检测回调通知的URL地址:
- 以 http:// 开头,外网可访问的地址
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSetLiveStreamQualityDetectionNotifyConfigRequest(
    publishDomain string,
    notifyUrl string,
) *SetLiveStreamQualityDetectionNotifyConfigRequest {

	return &SetLiveStreamQualityDetectionNotifyConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/qualityDetectionNotifys:config",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        PublishDomain: publishDomain,
        NotifyUrl: notifyUrl,
	}
}

/*
 * param publishDomain: 推流域名 (Required)
 * param notifyUrl: 质量检测回调通知的URL地址:
- 以 http:// 开头,外网可访问的地址
 (Required)
 */
func NewSetLiveStreamQualityDetectionNotifyConfigRequestWithAllParams(
    publishDomain string,
    notifyUrl string,
) *SetLiveStreamQualityDetectionNotifyConfigRequest {

    return &SetLiveStreamQualityDetectionNotifyConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/qualityDetectionNotifys:config",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        PublishDomain: publishDomain,
        NotifyUrl: notifyUrl,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSetLiveStreamQualityDetectionNotifyConfigRequestWithoutParam() *SetLiveStreamQualityDetectionNotifyConfigRequest {

    return &SetLiveStreamQualityDetectionNotifyConfigRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/qualityDetectionNotifys:config",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param publishDomain: 推流域名(Required) */
func (r *SetLiveStreamQualityDetectionNotifyConfigRequest) SetPublishDomain(publishDomain string) {
    r.PublishDomain = publishDomain
}

/* param notifyUrl: 质量检测回调通知的URL地址:
- 以 http:// 开头,外网可访问的地址
(Required) */
func (r *SetLiveStreamQualityDetectionNotifyConfigRequest) SetNotifyUrl(notifyUrl string) {
    r.NotifyUrl = notifyUrl
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SetLiveStreamQualityDetectionNotifyConfigRequest) GetRegionId() string {
    return ""
}

type SetLiveStreamQualityDetectionNotifyConfigResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SetLiveStreamQualityDetectionNotifyConfigResult `json:"result"`
}

type SetLiveStreamQualityDetectionNotifyConfigResult struct {
}