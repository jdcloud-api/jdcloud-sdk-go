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
    cdn "github.com/jdcloud-api/jdcloud-sdk-go/services/cdn/models"
)

type GetDomainDetailRequest struct {

    core.JDCloudRequest

    /* 用户域名  */
    Domain string `json:"domain"`
}

/*
 * param domain: 用户域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetDomainDetailRequest(
    domain string,
) *GetDomainDetailRequest {

	return &GetDomainDetailRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domain/{domain}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        Domain: domain,
	}
}

/*
 * param domain: 用户域名 (Required)
 */
func NewGetDomainDetailRequestWithAllParams(
    domain string,
) *GetDomainDetailRequest {

    return &GetDomainDetailRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        Domain: domain,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetDomainDetailRequestWithoutParam() *GetDomainDetailRequest {

    return &GetDomainDetailRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domain: 用户域名(Required) */
func (r *GetDomainDetailRequest) SetDomain(domain string) {
    r.Domain = domain
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetDomainDetailRequest) GetRegionId() string {
    return ""
}

type GetDomainDetailResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetDomainDetailResult `json:"result"`
}

type GetDomainDetailResult struct {
    AllStatus string `json:"allStatus"`
    AllowNoReferHeader string `json:"allowNoReferHeader"`
    AllowNullReferHeader string `json:"allowNullReferHeader"`
    DailyBandWidth int `json:"dailyBandWidth"`
    ForbiddenType string `json:"forbiddenType"`
    MaxFileSize int64 `json:"maxFileSize"`
    MinFileSize int64 `json:"minFileSize"`
    SumFileSize int64 `json:"sumFileSize"`
    AvgFileSize int64 `json:"avgFileSize"`
    ReferList []string `json:"referList"`
    ReferType string `json:"referType"`
    Domain string `json:"domain"`
    Cname string `json:"cname"`
    ArchiveNo string `json:"archiveNo"`
    Type string `json:"type"`
    Created string `json:"created"`
    Modified string `json:"modified"`
    Status string `json:"status"`
    AuditStatus string `json:"auditStatus"`
    Source cdn.BackSourceInfo `json:"source"`
    SourceType string `json:"sourceType"`
    DefaultSourceHost string `json:"defaultSourceHost"`
    BackSourceType string `json:"backSourceType"`
    HttpType string `json:"httpType"`
    Certificate string `json:"certificate"`
    RsaKey string `json:"rsaKey"`
    JumpType string `json:"jumpType"`
}