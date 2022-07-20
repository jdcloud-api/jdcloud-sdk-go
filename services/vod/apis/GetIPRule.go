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
    vod "github.com/lshuining/jdcloud-sdk-go/services/vod/models"
)

type GetIPRuleRequest struct {

    core.JDCloudRequest

    /* 域名ID  */
    DomainId int `json:"domainId"`
}

/*
 * param domainId: 域名ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetIPRuleRequest(
    domainId int,
) *GetIPRuleRequest {

	return &GetIPRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domains/{domainId}:getIPRule",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        DomainId: domainId,
	}
}

/*
 * param domainId: 域名ID (Required)
 */
func NewGetIPRuleRequestWithAllParams(
    domainId int,
) *GetIPRuleRequest {

    return &GetIPRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domains/{domainId}:getIPRule",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        DomainId: domainId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetIPRuleRequestWithoutParam() *GetIPRuleRequest {

    return &GetIPRuleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domains/{domainId}:getIPRule",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domainId: 域名ID(Required) */
func (r *GetIPRuleRequest) SetDomainId(domainId int) {
    r.DomainId = domainId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetIPRuleRequest) GetRegionId() string {
    return ""
}

type GetIPRuleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetIPRuleResult `json:"result"`
}

type GetIPRuleResult struct {
    RuleType string `json:"ruleType"`
    Config vod.IPRuleConfigObject `json:"config"`
    Enabled bool `json:"enabled"`
}