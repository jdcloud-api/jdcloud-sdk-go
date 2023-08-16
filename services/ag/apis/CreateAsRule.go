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
    ag "github.com/jdcloud-api/jdcloud-sdk-go/services/ag/models"
)

type CreateAsRuleRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 高可用组ID  */
    AgId string `json:"agId"`

    /* 伸缩规则名称，长度为1~32个字符，只允许中文、数字、大小写字母、英文下划线（_）、连字符（-）  */
    Name string `json:"name"`

    /* 伸缩规则描述，最大长度为256个字符 (Optional) */
    Description *string `json:"description"`

    /* 伸缩规则类型，取值范围：[`Simple`,`Target`,`Step`]
- `Simple`：简单规则，直接设置调整方式和调整值
- `Target`：目标跟踪规则，根据监控项和目标值计算需要扩缩容的实例数量，尽量将监控项的指标值维持在目标值和目标值的90%之间
- `Step`：步进规则，根据阈值和指标值提供分步扩展方式
  */
    AsRuleType string `json:"asRuleType"`

    /* 简单规则相关参数，当`asRuleType`为`Simple`时必填 (Optional) */
    SimpleAsRuleSpec *ag.CreateSimpleAsRuleSpec `json:"simpleAsRuleSpec"`

    /* 目标跟踪规则相关参数，当`asRuleType`为`Target`时必填 (Optional) */
    TargetAsRuleSpec *ag.CreateTargetAsRuleSpec `json:"targetAsRuleSpec"`

    /* 步进规则相关参数，当`asRuleType`为`Step`时必填 (Optional) */
    StepAsRuleSpec *ag.CreateStepAsRuleSpec `json:"stepAsRuleSpec"`
}

/*
 * param regionId: 地域ID (Required)
 * param agId: 高可用组ID (Required)
 * param name: 伸缩规则名称，长度为1~32个字符，只允许中文、数字、大小写字母、英文下划线（_）、连字符（-） (Required)
 * param asRuleType: 伸缩规则类型，取值范围：[`Simple`,`Target`,`Step`]
- `Simple`：简单规则，直接设置调整方式和调整值
- `Target`：目标跟踪规则，根据监控项和目标值计算需要扩缩容的实例数量，尽量将监控项的指标值维持在目标值和目标值的90%之间
- `Step`：步进规则，根据阈值和指标值提供分步扩展方式
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateAsRuleRequest(
    regionId string,
    agId string,
    name string,
    asRuleType string,
) *CreateAsRuleRequest {

	return &CreateAsRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/asRules",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        AgId: agId,
        Name: name,
        AsRuleType: asRuleType,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param agId: 高可用组ID (Required)
 * param name: 伸缩规则名称，长度为1~32个字符，只允许中文、数字、大小写字母、英文下划线（_）、连字符（-） (Required)
 * param description: 伸缩规则描述，最大长度为256个字符 (Optional)
 * param asRuleType: 伸缩规则类型，取值范围：[`Simple`,`Target`,`Step`]
- `Simple`：简单规则，直接设置调整方式和调整值
- `Target`：目标跟踪规则，根据监控项和目标值计算需要扩缩容的实例数量，尽量将监控项的指标值维持在目标值和目标值的90%之间
- `Step`：步进规则，根据阈值和指标值提供分步扩展方式
 (Required)
 * param simpleAsRuleSpec: 简单规则相关参数，当`asRuleType`为`Simple`时必填 (Optional)
 * param targetAsRuleSpec: 目标跟踪规则相关参数，当`asRuleType`为`Target`时必填 (Optional)
 * param stepAsRuleSpec: 步进规则相关参数，当`asRuleType`为`Step`时必填 (Optional)
 */
func NewCreateAsRuleRequestWithAllParams(
    regionId string,
    agId string,
    name string,
    description *string,
    asRuleType string,
    simpleAsRuleSpec *ag.CreateSimpleAsRuleSpec,
    targetAsRuleSpec *ag.CreateTargetAsRuleSpec,
    stepAsRuleSpec *ag.CreateStepAsRuleSpec,
) *CreateAsRuleRequest {

    return &CreateAsRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/asRules",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        AgId: agId,
        Name: name,
        Description: description,
        AsRuleType: asRuleType,
        SimpleAsRuleSpec: simpleAsRuleSpec,
        TargetAsRuleSpec: targetAsRuleSpec,
        StepAsRuleSpec: stepAsRuleSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateAsRuleRequestWithoutParam() *CreateAsRuleRequest {

    return &CreateAsRuleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/asRules",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *CreateAsRuleRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param agId: 高可用组ID(Required) */
func (r *CreateAsRuleRequest) SetAgId(agId string) {
    r.AgId = agId
}
/* param name: 伸缩规则名称，长度为1~32个字符，只允许中文、数字、大小写字母、英文下划线（_）、连字符（-）(Required) */
func (r *CreateAsRuleRequest) SetName(name string) {
    r.Name = name
}
/* param description: 伸缩规则描述，最大长度为256个字符(Optional) */
func (r *CreateAsRuleRequest) SetDescription(description string) {
    r.Description = &description
}
/* param asRuleType: 伸缩规则类型，取值范围：[`Simple`,`Target`,`Step`]
- `Simple`：简单规则，直接设置调整方式和调整值
- `Target`：目标跟踪规则，根据监控项和目标值计算需要扩缩容的实例数量，尽量将监控项的指标值维持在目标值和目标值的90%之间
- `Step`：步进规则，根据阈值和指标值提供分步扩展方式
(Required) */
func (r *CreateAsRuleRequest) SetAsRuleType(asRuleType string) {
    r.AsRuleType = asRuleType
}
/* param simpleAsRuleSpec: 简单规则相关参数，当`asRuleType`为`Simple`时必填(Optional) */
func (r *CreateAsRuleRequest) SetSimpleAsRuleSpec(simpleAsRuleSpec *ag.CreateSimpleAsRuleSpec) {
    r.SimpleAsRuleSpec = simpleAsRuleSpec
}
/* param targetAsRuleSpec: 目标跟踪规则相关参数，当`asRuleType`为`Target`时必填(Optional) */
func (r *CreateAsRuleRequest) SetTargetAsRuleSpec(targetAsRuleSpec *ag.CreateTargetAsRuleSpec) {
    r.TargetAsRuleSpec = targetAsRuleSpec
}
/* param stepAsRuleSpec: 步进规则相关参数，当`asRuleType`为`Step`时必填(Optional) */
func (r *CreateAsRuleRequest) SetStepAsRuleSpec(stepAsRuleSpec *ag.CreateStepAsRuleSpec) {
    r.StepAsRuleSpec = stepAsRuleSpec
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateAsRuleRequest) GetRegionId() string {
    return r.RegionId
}

type CreateAsRuleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateAsRuleResult `json:"result"`
}

type CreateAsRuleResult struct {
    AsRuleId string `json:"asRuleId"`
}