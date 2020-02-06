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

type DeleteRulesRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 转发规则组Id  */
    UrlMapId string `json:"urlMapId"`

    /* rule Id列表  */
    RuleIds []string `json:"ruleIds"`
}

/*
 * param regionId: Region ID (Required)
 * param urlMapId: 转发规则组Id (Required)
 * param ruleIds: rule Id列表 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteRulesRequest(
    regionId string,
    urlMapId string,
    ruleIds []string,
) *DeleteRulesRequest {

	return &DeleteRulesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/urlMaps/{urlMapId}:deleteRules",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        UrlMapId: urlMapId,
        RuleIds: ruleIds,
	}
}

/*
 * param regionId: Region ID (Required)
 * param urlMapId: 转发规则组Id (Required)
 * param ruleIds: rule Id列表 (Required)
 */
func NewDeleteRulesRequestWithAllParams(
    regionId string,
    urlMapId string,
    ruleIds []string,
) *DeleteRulesRequest {

    return &DeleteRulesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/urlMaps/{urlMapId}:deleteRules",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        UrlMapId: urlMapId,
        RuleIds: ruleIds,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteRulesRequestWithoutParam() *DeleteRulesRequest {

    return &DeleteRulesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/urlMaps/{urlMapId}:deleteRules",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DeleteRulesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param urlMapId: 转发规则组Id(Required) */
func (r *DeleteRulesRequest) SetUrlMapId(urlMapId string) {
    r.UrlMapId = urlMapId
}

/* param ruleIds: rule Id列表(Required) */
func (r *DeleteRulesRequest) SetRuleIds(ruleIds []string) {
    r.RuleIds = ruleIds
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteRulesRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteRulesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteRulesResult `json:"result"`
}

type DeleteRulesResult struct {
}