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

type SaveEvidenceRequest struct {

    core.JDCloudRequest

    /* 业务流水号  */
    BusinessId string `json:"businessId"`

    /* 存证数据json字符串的Base64  */
    File string `json:"file"`
}

/*
 * param businessId: 业务流水号 (Required)
 * param file: 存证数据json字符串的Base64 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSaveEvidenceRequest(
    businessId string,
    file string,
) *SaveEvidenceRequest {

	return &SaveEvidenceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/evidence:evidenceSave",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        BusinessId: businessId,
        File: file,
	}
}

/*
 * param businessId: 业务流水号 (Required)
 * param file: 存证数据json字符串的Base64 (Required)
 */
func NewSaveEvidenceRequestWithAllParams(
    businessId string,
    file string,
) *SaveEvidenceRequest {

    return &SaveEvidenceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/evidence:evidenceSave",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        BusinessId: businessId,
        File: file,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSaveEvidenceRequestWithoutParam() *SaveEvidenceRequest {

    return &SaveEvidenceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/evidence:evidenceSave",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param businessId: 业务流水号(Required) */
func (r *SaveEvidenceRequest) SetBusinessId(businessId string) {
    r.BusinessId = businessId
}
/* param file: 存证数据json字符串的Base64(Required) */
func (r *SaveEvidenceRequest) SetFile(file string) {
    r.File = file
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SaveEvidenceRequest) GetRegionId() string {
    return ""
}

type SaveEvidenceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SaveEvidenceResult `json:"result"`
}

type SaveEvidenceResult struct {
    Code string `json:"code"`
    Message string `json:"message"`
    Success bool `json:"success"`
    Data cloudsign.SaveEvidenceResp `json:"data"`
}