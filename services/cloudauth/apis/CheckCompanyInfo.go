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
    cloudauth "github.com/jdcloud-api/jdcloud-sdk-go/services/cloudauth/models"
)

type CheckCompanyInfoRequest struct {

    core.JDCloudRequest

    /*   */
    CompanyInfo *cloudauth.CompanyInfo `json:"companyInfo"`
}

/*
 * param companyInfo:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCheckCompanyInfoRequest(
    companyInfo *cloudauth.CompanyInfo,
) *CheckCompanyInfoRequest {

	return &CheckCompanyInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/company:info",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        CompanyInfo: companyInfo,
	}
}

/*
 * param companyInfo:  (Required)
 */
func NewCheckCompanyInfoRequestWithAllParams(
    companyInfo *cloudauth.CompanyInfo,
) *CheckCompanyInfoRequest {

    return &CheckCompanyInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/company:info",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        CompanyInfo: companyInfo,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCheckCompanyInfoRequestWithoutParam() *CheckCompanyInfoRequest {

    return &CheckCompanyInfoRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/company:info",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param companyInfo: (Required) */
func (r *CheckCompanyInfoRequest) SetCompanyInfo(companyInfo *cloudauth.CompanyInfo) {
    r.CompanyInfo = companyInfo
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CheckCompanyInfoRequest) GetRegionId() string {
    return ""
}

type CheckCompanyInfoResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CheckCompanyInfoResult `json:"result"`
}

type CheckCompanyInfoResult struct {
    AuthInfo cloudauth.AuthInfo `json:"authInfo"`
}