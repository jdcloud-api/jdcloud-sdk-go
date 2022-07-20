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

type UpdateCCProtectRuleRequest struct {

    core.JDCloudRequest

    /* 用户域名  */
    Domain string `json:"domain"`

    /* 需要修改的规则ID  */
    Id string `json:"id"`

    /* null (Optional) */
    Uri *string `json:"uri"`

    /* null (Optional) */
    DetectPeriod *int `json:"detectPeriod"`

    /* null (Optional) */
    SingleIpLimit *int `json:"singleIpLimit"`

    /* null (Optional) */
    BlockType *int `json:"blockType"`

    /* null (Optional) */
    BlockTime *int `json:"blockTime"`
}

/*
 * param domain: 用户域名 (Required)
 * param id: 需要修改的规则ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateCCProtectRuleRequest(
    domain string,
    id string,
) *UpdateCCProtectRuleRequest {

	return &UpdateCCProtectRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domain/{domain}/ccProtectRule/{id}",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Domain: domain,
        Id: id,
	}
}

/*
 * param domain: 用户域名 (Required)
 * param id: 需要修改的规则ID (Required)
 * param uri: null (Optional)
 * param detectPeriod: null (Optional)
 * param singleIpLimit: null (Optional)
 * param blockType: null (Optional)
 * param blockTime: null (Optional)
 */
func NewUpdateCCProtectRuleRequestWithAllParams(
    domain string,
    id string,
    uri *string,
    detectPeriod *int,
    singleIpLimit *int,
    blockType *int,
    blockTime *int,
) *UpdateCCProtectRuleRequest {

    return &UpdateCCProtectRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/ccProtectRule/{id}",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Domain: domain,
        Id: id,
        Uri: uri,
        DetectPeriod: detectPeriod,
        SingleIpLimit: singleIpLimit,
        BlockType: blockType,
        BlockTime: blockTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateCCProtectRuleRequestWithoutParam() *UpdateCCProtectRuleRequest {

    return &UpdateCCProtectRuleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/ccProtectRule/{id}",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domain: 用户域名(Required) */
func (r *UpdateCCProtectRuleRequest) SetDomain(domain string) {
    r.Domain = domain
}

/* param id: 需要修改的规则ID(Required) */
func (r *UpdateCCProtectRuleRequest) SetId(id string) {
    r.Id = id
}

/* param uri: null(Optional) */
func (r *UpdateCCProtectRuleRequest) SetUri(uri string) {
    r.Uri = &uri
}

/* param detectPeriod: null(Optional) */
func (r *UpdateCCProtectRuleRequest) SetDetectPeriod(detectPeriod int) {
    r.DetectPeriod = &detectPeriod
}

/* param singleIpLimit: null(Optional) */
func (r *UpdateCCProtectRuleRequest) SetSingleIpLimit(singleIpLimit int) {
    r.SingleIpLimit = &singleIpLimit
}

/* param blockType: null(Optional) */
func (r *UpdateCCProtectRuleRequest) SetBlockType(blockType int) {
    r.BlockType = &blockType
}

/* param blockTime: null(Optional) */
func (r *UpdateCCProtectRuleRequest) SetBlockTime(blockTime int) {
    r.BlockTime = &blockTime
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateCCProtectRuleRequest) GetRegionId() string {
    return ""
}

type UpdateCCProtectRuleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateCCProtectRuleResult `json:"result"`
}

type UpdateCCProtectRuleResult struct {
}