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

type SetIpBlackListRequest struct {

    core.JDCloudRequest

    /* 用户域名  */
    Domain string `json:"domain"`

    /* ip黑名单,ips中url不能超过50条 (Optional) */
    Ips []string `json:"ips"`
}

/*
 * param domain: 用户域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSetIpBlackListRequest(
    domain string,
) *SetIpBlackListRequest {

	return &SetIpBlackListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domain/{domain}/ipBlackList",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Domain: domain,
	}
}

/*
 * param domain: 用户域名 (Required)
 * param ips: ip黑名单,ips中url不能超过50条 (Optional)
 */
func NewSetIpBlackListRequestWithAllParams(
    domain string,
    ips []string,
) *SetIpBlackListRequest {

    return &SetIpBlackListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/ipBlackList",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Domain: domain,
        Ips: ips,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSetIpBlackListRequestWithoutParam() *SetIpBlackListRequest {

    return &SetIpBlackListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/ipBlackList",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domain: 用户域名(Required) */
func (r *SetIpBlackListRequest) SetDomain(domain string) {
    r.Domain = domain
}

/* param ips: ip黑名单,ips中url不能超过50条(Optional) */
func (r *SetIpBlackListRequest) SetIps(ips []string) {
    r.Ips = ips
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SetIpBlackListRequest) GetRegionId() string {
    return ""
}

type SetIpBlackListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SetIpBlackListResult `json:"result"`
}

type SetIpBlackListResult struct {
}