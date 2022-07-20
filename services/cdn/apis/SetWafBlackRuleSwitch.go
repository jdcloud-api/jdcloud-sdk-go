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

type SetWafBlackRuleSwitchRequest struct {

    core.JDCloudRequest

    /* 用户域名  */
    Domain string `json:"domain"`

    /* 开关状态, on:开启,off:关闭 (Optional) */
    SwitchStatus *string `json:"switchStatus"`
}

/*
 * param domain: 用户域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSetWafBlackRuleSwitchRequest(
    domain string,
) *SetWafBlackRuleSwitchRequest {

	return &SetWafBlackRuleSwitchRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domain/{domain}/wafBlackRuleSwitch",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Domain: domain,
	}
}

/*
 * param domain: 用户域名 (Required)
 * param switchStatus: 开关状态, on:开启,off:关闭 (Optional)
 */
func NewSetWafBlackRuleSwitchRequestWithAllParams(
    domain string,
    switchStatus *string,
) *SetWafBlackRuleSwitchRequest {

    return &SetWafBlackRuleSwitchRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/wafBlackRuleSwitch",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Domain: domain,
        SwitchStatus: switchStatus,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSetWafBlackRuleSwitchRequestWithoutParam() *SetWafBlackRuleSwitchRequest {

    return &SetWafBlackRuleSwitchRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/wafBlackRuleSwitch",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domain: 用户域名(Required) */
func (r *SetWafBlackRuleSwitchRequest) SetDomain(domain string) {
    r.Domain = domain
}

/* param switchStatus: 开关状态, on:开启,off:关闭(Optional) */
func (r *SetWafBlackRuleSwitchRequest) SetSwitchStatus(switchStatus string) {
    r.SwitchStatus = &switchStatus
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SetWafBlackRuleSwitchRequest) GetRegionId() string {
    return ""
}

type SetWafBlackRuleSwitchResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SetWafBlackRuleSwitchResult `json:"result"`
}

type SetWafBlackRuleSwitchResult struct {
}