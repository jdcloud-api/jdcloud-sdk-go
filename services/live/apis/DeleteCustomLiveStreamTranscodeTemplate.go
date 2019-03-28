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

type DeleteCustomLiveStreamTranscodeTemplateRequest struct {

    core.JDCloudRequest

    /* 转码模板
  */
    Template string `json:"template"`
}

/*
 * param template: 转码模板
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteCustomLiveStreamTranscodeTemplateRequest(
    template string,
) *DeleteCustomLiveStreamTranscodeTemplateRequest {

	return &DeleteCustomLiveStreamTranscodeTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/transcodeCustoms/{template}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        Template: template,
	}
}

/*
 * param template: 转码模板
 (Required)
 */
func NewDeleteCustomLiveStreamTranscodeTemplateRequestWithAllParams(
    template string,
) *DeleteCustomLiveStreamTranscodeTemplateRequest {

    return &DeleteCustomLiveStreamTranscodeTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/transcodeCustoms/{template}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        Template: template,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteCustomLiveStreamTranscodeTemplateRequestWithoutParam() *DeleteCustomLiveStreamTranscodeTemplateRequest {

    return &DeleteCustomLiveStreamTranscodeTemplateRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/transcodeCustoms/{template}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param template: 转码模板
(Required) */
func (r *DeleteCustomLiveStreamTranscodeTemplateRequest) SetTemplate(template string) {
    r.Template = template
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteCustomLiveStreamTranscodeTemplateRequest) GetRegionId() string {
    return ""
}

type DeleteCustomLiveStreamTranscodeTemplateResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteCustomLiveStreamTranscodeTemplateResult `json:"result"`
}

type DeleteCustomLiveStreamTranscodeTemplateResult struct {
}