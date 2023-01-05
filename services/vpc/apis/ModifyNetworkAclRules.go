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
    vpc "github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/models"
)

type ModifyNetworkAclRulesRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* networkAclId ID  */
    NetworkAclId string `json:"networkAclId"`

    /* networkAcl规则列表  */
    ModifyNetworkAclRuleSpecs []vpc.ModifyNetworkAclRuleSpec `json:"modifyNetworkAclRuleSpecs"`
}

/*
 * param regionId: Region ID (Required)
 * param networkAclId: networkAclId ID (Required)
 * param modifyNetworkAclRuleSpecs: networkAcl规则列表 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyNetworkAclRulesRequest(
    regionId string,
    networkAclId string,
    modifyNetworkAclRuleSpecs []vpc.ModifyNetworkAclRuleSpec,
) *ModifyNetworkAclRulesRequest {

	return &ModifyNetworkAclRulesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/networkAcls/{networkAclId}:modifyNetworkAclRules",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        NetworkAclId: networkAclId,
        ModifyNetworkAclRuleSpecs: modifyNetworkAclRuleSpecs,
	}
}

/*
 * param regionId: Region ID (Required)
 * param networkAclId: networkAclId ID (Required)
 * param modifyNetworkAclRuleSpecs: networkAcl规则列表 (Required)
 */
func NewModifyNetworkAclRulesRequestWithAllParams(
    regionId string,
    networkAclId string,
    modifyNetworkAclRuleSpecs []vpc.ModifyNetworkAclRuleSpec,
) *ModifyNetworkAclRulesRequest {

    return &ModifyNetworkAclRulesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/networkAcls/{networkAclId}:modifyNetworkAclRules",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        NetworkAclId: networkAclId,
        ModifyNetworkAclRuleSpecs: modifyNetworkAclRuleSpecs,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyNetworkAclRulesRequestWithoutParam() *ModifyNetworkAclRulesRequest {

    return &ModifyNetworkAclRulesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/networkAcls/{networkAclId}:modifyNetworkAclRules",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *ModifyNetworkAclRulesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param networkAclId: networkAclId ID(Required) */
func (r *ModifyNetworkAclRulesRequest) SetNetworkAclId(networkAclId string) {
    r.NetworkAclId = networkAclId
}
/* param modifyNetworkAclRuleSpecs: networkAcl规则列表(Required) */
func (r *ModifyNetworkAclRulesRequest) SetModifyNetworkAclRuleSpecs(modifyNetworkAclRuleSpecs []vpc.ModifyNetworkAclRuleSpec) {
    r.ModifyNetworkAclRuleSpecs = modifyNetworkAclRuleSpecs
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyNetworkAclRulesRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyNetworkAclRulesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyNetworkAclRulesResult `json:"result"`
}

type ModifyNetworkAclRulesResult struct {
}