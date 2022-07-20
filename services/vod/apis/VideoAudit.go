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

type VideoAuditRequest struct {

    core.JDCloudRequest

    /* 视频ID  */
    VideoId string `json:"videoId"`

    /* 审核结果，取值范围:
 block(封禁)
 unblock(解封)
  */
    AuditResult string `json:"auditResult"`
}

/*
 * param videoId: 视频ID (Required)
 * param auditResult: 审核结果，取值范围:
 block(封禁)
 unblock(解封)
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewVideoAuditRequest(
    videoId string,
    auditResult string,
) *VideoAuditRequest {

	return &VideoAuditRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/videos/{videoId}:audit",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        VideoId: videoId,
        AuditResult: auditResult,
	}
}

/*
 * param videoId: 视频ID (Required)
 * param auditResult: 审核结果，取值范围:
 block(封禁)
 unblock(解封)
 (Required)
 */
func NewVideoAuditRequestWithAllParams(
    videoId string,
    auditResult string,
) *VideoAuditRequest {

    return &VideoAuditRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/videos/{videoId}:audit",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        VideoId: videoId,
        AuditResult: auditResult,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewVideoAuditRequestWithoutParam() *VideoAuditRequest {

    return &VideoAuditRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/videos/{videoId}:audit",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param videoId: 视频ID(Required) */
func (r *VideoAuditRequest) SetVideoId(videoId string) {
    r.VideoId = videoId
}

/* param auditResult: 审核结果，取值范围:
 block(封禁)
 unblock(解封)
(Required) */
func (r *VideoAuditRequest) SetAuditResult(auditResult string) {
    r.AuditResult = auditResult
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r VideoAuditRequest) GetRegionId() string {
    return ""
}

type VideoAuditResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result VideoAuditResult `json:"result"`
}

type VideoAuditResult struct {
}