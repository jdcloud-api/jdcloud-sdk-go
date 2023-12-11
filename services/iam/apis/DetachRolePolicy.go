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

type DetachRolePolicyRequest struct {

    core.JDCloudRequest

    /* 角色名称  */
    RoleName string `json:"roleName"`

    /* 策略名称  */
    PolicyName string `json:"policyName"`

    /* 资源组id (Optional) */
    ScopeId *string `json:"scopeId"`

    /* 允许解除策略："Deny" 不允许，Allow 允许，空情况默认允许，兼容历史数据 (Optional) */
    AllowDetachAddPolicy *string `json:"allowDetachAddPolicy"`
}

/*
 * param roleName: 角色名称 (Required)
 * param policyName: 策略名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDetachRolePolicyRequest(
    roleName string,
    policyName string,
) *DetachRolePolicyRequest {

	return &DetachRolePolicyRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/role/{roleName}:detachRolePolicy",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RoleName: roleName,
        PolicyName: policyName,
	}
}

/*
 * param roleName: 角色名称 (Required)
 * param policyName: 策略名称 (Required)
 * param scopeId: 资源组id (Optional)
 * param allowDetachAddPolicy: 允许解除策略："Deny" 不允许，Allow 允许，空情况默认允许，兼容历史数据 (Optional)
 */
func NewDetachRolePolicyRequestWithAllParams(
    roleName string,
    policyName string,
    scopeId *string,
    allowDetachAddPolicy *string,
) *DetachRolePolicyRequest {

    return &DetachRolePolicyRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/role/{roleName}:detachRolePolicy",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RoleName: roleName,
        PolicyName: policyName,
        ScopeId: scopeId,
        AllowDetachAddPolicy: allowDetachAddPolicy,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDetachRolePolicyRequestWithoutParam() *DetachRolePolicyRequest {

    return &DetachRolePolicyRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/role/{roleName}:detachRolePolicy",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param roleName: 角色名称(Required) */
func (r *DetachRolePolicyRequest) SetRoleName(roleName string) {
    r.RoleName = roleName
}
/* param policyName: 策略名称(Required) */
func (r *DetachRolePolicyRequest) SetPolicyName(policyName string) {
    r.PolicyName = policyName
}
/* param scopeId: 资源组id(Optional) */
func (r *DetachRolePolicyRequest) SetScopeId(scopeId string) {
    r.ScopeId = &scopeId
}
/* param allowDetachAddPolicy: 允许解除策略："Deny" 不允许，Allow 允许，空情况默认允许，兼容历史数据(Optional) */
func (r *DetachRolePolicyRequest) SetAllowDetachAddPolicy(allowDetachAddPolicy string) {
    r.AllowDetachAddPolicy = &allowDetachAddPolicy
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DetachRolePolicyRequest) GetRegionId() string {
    return ""
}

type DetachRolePolicyResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DetachRolePolicyResult `json:"result"`
}

type DetachRolePolicyResult struct {
}