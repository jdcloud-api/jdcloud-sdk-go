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
    cloudsign "github.com/jdcloud-api/jdcloud-sdk-go/services/cloudsign/models"
)

type PaddingTemplateRequest struct {

    core.JDCloudRequest

    /* 合同模板ID  */
    TemplateId string `json:"templateId"`

    /*   */
    PaddingSpec *cloudsign.PaddingSpec `json:"paddingSpec"`
}

/*
 * param templateId: 合同模板ID (Required)
 * param paddingSpec:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewPaddingTemplateRequest(
    templateId string,
    paddingSpec *cloudsign.PaddingSpec,
) *PaddingTemplateRequest {

	return &PaddingTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/smqTmplate/{templateId}:paddingTemplate",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        TemplateId: templateId,
        PaddingSpec: paddingSpec,
	}
}

/*
 * param templateId: 合同模板ID (Required)
 * param paddingSpec:  (Required)
 */
func NewPaddingTemplateRequestWithAllParams(
    templateId string,
    paddingSpec *cloudsign.PaddingSpec,
) *PaddingTemplateRequest {

    return &PaddingTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/smqTmplate/{templateId}:paddingTemplate",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        TemplateId: templateId,
        PaddingSpec: paddingSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewPaddingTemplateRequestWithoutParam() *PaddingTemplateRequest {

    return &PaddingTemplateRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/smqTmplate/{templateId}:paddingTemplate",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param templateId: 合同模板ID(Required) */
func (r *PaddingTemplateRequest) SetTemplateId(templateId string) {
    r.TemplateId = templateId
}
/* param paddingSpec: (Required) */
func (r *PaddingTemplateRequest) SetPaddingSpec(paddingSpec *cloudsign.PaddingSpec) {
    r.PaddingSpec = paddingSpec
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r PaddingTemplateRequest) GetRegionId() string {
    return ""
}

type PaddingTemplateResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result PaddingTemplateResult `json:"result"`
}

type PaddingTemplateResult struct {
    TemplateId string `json:"templateId"`
}