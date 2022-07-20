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

type RemoveNetworkSecurityGroupRulesRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* NetworkSecurityGroup ID  */
    NetworkSecurityGroupId string `json:"networkSecurityGroupId"`

    /* 安全组规则Id列表  */
    RuleIds []string `json:"ruleIds"`
}

/*
 * param regionId: Region ID (Required)
 * param networkSecurityGroupId: NetworkSecurityGroup ID (Required)
 * param ruleIds: 安全组规则Id列表 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewRemoveNetworkSecurityGroupRulesRequest(
    regionId string,
    networkSecurityGroupId string,
    ruleIds []string,
) *RemoveNetworkSecurityGroupRulesRequest {

	return &RemoveNetworkSecurityGroupRulesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/networkSecurityGroups/{networkSecurityGroupId}:removeNetworkSecurityGroupRules",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        NetworkSecurityGroupId: networkSecurityGroupId,
        RuleIds: ruleIds,
	}
}

/*
 * param regionId: Region ID (Required)
 * param networkSecurityGroupId: NetworkSecurityGroup ID (Required)
 * param ruleIds: 安全组规则Id列表 (Required)
 */
func NewRemoveNetworkSecurityGroupRulesRequestWithAllParams(
    regionId string,
    networkSecurityGroupId string,
    ruleIds []string,
) *RemoveNetworkSecurityGroupRulesRequest {

    return &RemoveNetworkSecurityGroupRulesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/networkSecurityGroups/{networkSecurityGroupId}:removeNetworkSecurityGroupRules",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        NetworkSecurityGroupId: networkSecurityGroupId,
        RuleIds: ruleIds,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewRemoveNetworkSecurityGroupRulesRequestWithoutParam() *RemoveNetworkSecurityGroupRulesRequest {

    return &RemoveNetworkSecurityGroupRulesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/networkSecurityGroups/{networkSecurityGroupId}:removeNetworkSecurityGroupRules",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *RemoveNetworkSecurityGroupRulesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param networkSecurityGroupId: NetworkSecurityGroup ID(Required) */
func (r *RemoveNetworkSecurityGroupRulesRequest) SetNetworkSecurityGroupId(networkSecurityGroupId string) {
    r.NetworkSecurityGroupId = networkSecurityGroupId
}

/* param ruleIds: 安全组规则Id列表(Required) */
func (r *RemoveNetworkSecurityGroupRulesRequest) SetRuleIds(ruleIds []string) {
    r.RuleIds = ruleIds
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r RemoveNetworkSecurityGroupRulesRequest) GetRegionId() string {
    return r.RegionId
}

type RemoveNetworkSecurityGroupRulesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result RemoveNetworkSecurityGroupRulesResult `json:"result"`
}

type RemoveNetworkSecurityGroupRulesResult struct {
}