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

type ConfigBackSourceIpVersionRequest struct {

    core.JDCloudRequest

    /* 用户域名  */
    Domain string `json:"domain"`

    /* 回源IP类型，取值ipv4/ipv6/ipv46。ipv4：回源IP为ipv4；ipv6：ipv6优先；ipv46：ipv4/ipv6负载均衡 (Optional) */
    BackSourceIpVersion *string `json:"backSourceIpVersion"`
}

/*
 * param domain: 用户域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewConfigBackSourceIpVersionRequest(
    domain string,
) *ConfigBackSourceIpVersionRequest {

	return &ConfigBackSourceIpVersionRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domain/{domain}/configBackSourceIpVersion",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Domain: domain,
	}
}

/*
 * param domain: 用户域名 (Required)
 * param backSourceIpVersion: 回源IP类型，取值ipv4/ipv6/ipv46。ipv4：回源IP为ipv4；ipv6：ipv6优先；ipv46：ipv4/ipv6负载均衡 (Optional)
 */
func NewConfigBackSourceIpVersionRequestWithAllParams(
    domain string,
    backSourceIpVersion *string,
) *ConfigBackSourceIpVersionRequest {

    return &ConfigBackSourceIpVersionRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/configBackSourceIpVersion",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Domain: domain,
        BackSourceIpVersion: backSourceIpVersion,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewConfigBackSourceIpVersionRequestWithoutParam() *ConfigBackSourceIpVersionRequest {

    return &ConfigBackSourceIpVersionRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/configBackSourceIpVersion",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domain: 用户域名(Required) */
func (r *ConfigBackSourceIpVersionRequest) SetDomain(domain string) {
    r.Domain = domain
}
/* param backSourceIpVersion: 回源IP类型，取值ipv4/ipv6/ipv46。ipv4：回源IP为ipv4；ipv6：ipv6优先；ipv46：ipv4/ipv6负载均衡(Optional) */
func (r *ConfigBackSourceIpVersionRequest) SetBackSourceIpVersion(backSourceIpVersion string) {
    r.BackSourceIpVersion = &backSourceIpVersion
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ConfigBackSourceIpVersionRequest) GetRegionId() string {
    return ""
}

type ConfigBackSourceIpVersionResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ConfigBackSourceIpVersionResult `json:"result"`
}

type ConfigBackSourceIpVersionResult struct {
}