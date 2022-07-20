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

type SetURLRuleRequest struct {

    core.JDCloudRequest

    /* 域名ID  */
    DomainId int `json:"domainId"`

    /* URL鉴权规则配置对象  */
    Config *vod.URLRuleConfigObject `json:"config"`

    /* 是否启用该规则  */
    Enabled bool `json:"enabled"`
}

/*
 * param domainId: 域名ID (Required)
 * param config: URL鉴权规则配置对象 (Required)
 * param enabled: 是否启用该规则 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSetURLRuleRequest(
    domainId int,
    config *vod.URLRuleConfigObject,
    enabled bool,
) *SetURLRuleRequest {

	return &SetURLRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domains/{domainId}:setURLRule",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        DomainId: domainId,
        Config: config,
        Enabled: enabled,
	}
}

/*
 * param domainId: 域名ID (Required)
 * param config: URL鉴权规则配置对象 (Required)
 * param enabled: 是否启用该规则 (Required)
 */
func NewSetURLRuleRequestWithAllParams(
    domainId int,
    config *vod.URLRuleConfigObject,
    enabled bool,
) *SetURLRuleRequest {

    return &SetURLRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domains/{domainId}:setURLRule",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        DomainId: domainId,
        Config: config,
        Enabled: enabled,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSetURLRuleRequestWithoutParam() *SetURLRuleRequest {

    return &SetURLRuleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domains/{domainId}:setURLRule",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domainId: 域名ID(Required) */
func (r *SetURLRuleRequest) SetDomainId(domainId int) {
    r.DomainId = domainId
}

/* param config: URL鉴权规则配置对象(Required) */
func (r *SetURLRuleRequest) SetConfig(config *vod.URLRuleConfigObject) {
    r.Config = config
}

/* param enabled: 是否启用该规则(Required) */
func (r *SetURLRuleRequest) SetEnabled(enabled bool) {
    r.Enabled = enabled
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SetURLRuleRequest) GetRegionId() string {
    return ""
}

type SetURLRuleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SetURLRuleResult `json:"result"`
}

type SetURLRuleResult struct {
}