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

type ConfigBackSourceRuleRequest struct {

    core.JDCloudRequest

    /* 用户域名  */
    Domain string `json:"domain"`

    /* 回源改写之前的正则表达式 (Optional) */
    BeforeRegex *string `json:"beforeRegex"`

    /* 回源改写之后的正则表达式 (Optional) */
    AfterRegex *string `json:"afterRegex"`
}

/*
 * param domain: 用户域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewConfigBackSourceRuleRequest(
    domain string,
) *ConfigBackSourceRuleRequest {

	return &ConfigBackSourceRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domain/{domain}/configBackSourceRule",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Domain: domain,
	}
}

/*
 * param domain: 用户域名 (Required)
 * param beforeRegex: 回源改写之前的正则表达式 (Optional)
 * param afterRegex: 回源改写之后的正则表达式 (Optional)
 */
func NewConfigBackSourceRuleRequestWithAllParams(
    domain string,
    beforeRegex *string,
    afterRegex *string,
) *ConfigBackSourceRuleRequest {

    return &ConfigBackSourceRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/configBackSourceRule",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Domain: domain,
        BeforeRegex: beforeRegex,
        AfterRegex: afterRegex,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewConfigBackSourceRuleRequestWithoutParam() *ConfigBackSourceRuleRequest {

    return &ConfigBackSourceRuleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/configBackSourceRule",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domain: 用户域名(Required) */
func (r *ConfigBackSourceRuleRequest) SetDomain(domain string) {
    r.Domain = domain
}

/* param beforeRegex: 回源改写之前的正则表达式(Optional) */
func (r *ConfigBackSourceRuleRequest) SetBeforeRegex(beforeRegex string) {
    r.BeforeRegex = &beforeRegex
}

/* param afterRegex: 回源改写之后的正则表达式(Optional) */
func (r *ConfigBackSourceRuleRequest) SetAfterRegex(afterRegex string) {
    r.AfterRegex = &afterRegex
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ConfigBackSourceRuleRequest) GetRegionId() string {
    return ""
}

type ConfigBackSourceRuleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ConfigBackSourceRuleResult `json:"result"`
}

type ConfigBackSourceRuleResult struct {
}