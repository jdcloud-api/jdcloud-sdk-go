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
    fc "github.com/jdcloud-api/jdcloud-sdk-go/services/fc/models"
)

type CreateTriggerRequest struct {

    core.JDCloudRequest

    /* 地域ID。  */
    RegionId string `json:"regionId"`

    /* 服务名称。  */
    ServiceName string `json:"serviceName"`

    /* 函数名称。  */
    FunctionName string `json:"functionName"`

    /* 触发器创建参数  */
    TriggerSpec *fc.TriggerSpec `json:"triggerSpec"`
}

/*
 * param regionId: 地域ID。 (Required)
 * param serviceName: 服务名称。 (Required)
 * param functionName: 函数名称。 (Required)
 * param triggerSpec: 触发器创建参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateTriggerRequest(
    regionId string,
    serviceName string,
    functionName string,
    triggerSpec *fc.TriggerSpec,
) *CreateTriggerRequest {

	return &CreateTriggerRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/services/{serviceName}/functions/{functionName}/triggers",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ServiceName: serviceName,
        FunctionName: functionName,
        TriggerSpec: triggerSpec,
	}
}

/*
 * param regionId: 地域ID。 (Required)
 * param serviceName: 服务名称。 (Required)
 * param functionName: 函数名称。 (Required)
 * param triggerSpec: 触发器创建参数 (Required)
 */
func NewCreateTriggerRequestWithAllParams(
    regionId string,
    serviceName string,
    functionName string,
    triggerSpec *fc.TriggerSpec,
) *CreateTriggerRequest {

    return &CreateTriggerRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/services/{serviceName}/functions/{functionName}/triggers",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ServiceName: serviceName,
        FunctionName: functionName,
        TriggerSpec: triggerSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateTriggerRequestWithoutParam() *CreateTriggerRequest {

    return &CreateTriggerRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/services/{serviceName}/functions/{functionName}/triggers",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID。(Required) */
func (r *CreateTriggerRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param serviceName: 服务名称。(Required) */
func (r *CreateTriggerRequest) SetServiceName(serviceName string) {
    r.ServiceName = serviceName
}
/* param functionName: 函数名称。(Required) */
func (r *CreateTriggerRequest) SetFunctionName(functionName string) {
    r.FunctionName = functionName
}
/* param triggerSpec: 触发器创建参数(Required) */
func (r *CreateTriggerRequest) SetTriggerSpec(triggerSpec *fc.TriggerSpec) {
    r.TriggerSpec = triggerSpec
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateTriggerRequest) GetRegionId() string {
    return r.RegionId
}

type CreateTriggerResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateTriggerResult `json:"result"`
}

type CreateTriggerResult struct {
    TriggerName string `json:"triggerName"`
}