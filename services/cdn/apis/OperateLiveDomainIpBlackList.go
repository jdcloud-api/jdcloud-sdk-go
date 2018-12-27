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

type OperateLiveDomainIpBlackListRequest struct {

    core.JDCloudRequest

    /* 用户域名  */
    Domain string `json:"domain"`

    /*   */
    BlackIpsEnable string `json:"blackIpsEnable"`
}

/*
 * param domain: 用户域名 (Required)
 * param blackIpsEnable:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewOperateLiveDomainIpBlackListRequest(
    domain string,
    blackIpsEnable string,
) *OperateLiveDomainIpBlackListRequest {

	return &OperateLiveDomainIpBlackListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/liveDomain/{domain}/ipBlackList:{blackIpsEnable}",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Domain: domain,
        BlackIpsEnable: blackIpsEnable,
	}
}

/*
 * param domain: 用户域名 (Required)
 * param blackIpsEnable:  (Required)
 */
func NewOperateLiveDomainIpBlackListRequestWithAllParams(
    domain string,
    blackIpsEnable string,
) *OperateLiveDomainIpBlackListRequest {

    return &OperateLiveDomainIpBlackListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveDomain/{domain}/ipBlackList:{blackIpsEnable}",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Domain: domain,
        BlackIpsEnable: blackIpsEnable,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewOperateLiveDomainIpBlackListRequestWithoutParam() *OperateLiveDomainIpBlackListRequest {

    return &OperateLiveDomainIpBlackListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveDomain/{domain}/ipBlackList:{blackIpsEnable}",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domain: 用户域名(Required) */
func (r *OperateLiveDomainIpBlackListRequest) SetDomain(domain string) {
    r.Domain = domain
}

/* param blackIpsEnable: (Required) */
func (r *OperateLiveDomainIpBlackListRequest) SetBlackIpsEnable(blackIpsEnable string) {
    r.BlackIpsEnable = blackIpsEnable
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r OperateLiveDomainIpBlackListRequest) GetRegionId() string {
    return ""
}

type OperateLiveDomainIpBlackListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result OperateLiveDomainIpBlackListResult `json:"result"`
}

type OperateLiveDomainIpBlackListResult struct {
}