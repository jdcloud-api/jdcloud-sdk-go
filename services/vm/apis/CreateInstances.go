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
    vm "github.com/lshuining/jdcloud-sdk-go/services/vm/models"
)

type CreateInstancesRequest struct {

    core.JDCloudRequest

    /* 地域ID。  */
    RegionId string `json:"regionId"`

    /* 实例配置。
  */
    InstanceSpec *vm.InstanceSpec `json:"instanceSpec"`

    /* 创建实例的数量，不能超过用户配额。
取值范围：[1,100]；默认值：1。
如果在弹性网卡中指定了内网IP地址，那么单次创建 `maxCount` 只能是 1。
 (Optional) */
    MaxCount *int `json:"maxCount"`

    /* 用于保证请求的幂等性。由客户端生成，并确保不同请求中该参数唯一，长度不能超过64个字符。
 (Optional) */
    ClientToken *string `json:"clientToken"`
}

/*
 * param regionId: 地域ID。 (Required)
 * param instanceSpec: 实例配置。
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateInstancesRequest(
    regionId string,
    instanceSpec *vm.InstanceSpec,
) *CreateInstancesRequest {

	return &CreateInstancesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceSpec: instanceSpec,
	}
}

/*
 * param regionId: 地域ID。 (Required)
 * param instanceSpec: 实例配置。
 (Required)
 * param maxCount: 创建实例的数量，不能超过用户配额。
取值范围：[1,100]；默认值：1。
如果在弹性网卡中指定了内网IP地址，那么单次创建 `maxCount` 只能是 1。
 (Optional)
 * param clientToken: 用于保证请求的幂等性。由客户端生成，并确保不同请求中该参数唯一，长度不能超过64个字符。
 (Optional)
 */
func NewCreateInstancesRequestWithAllParams(
    regionId string,
    instanceSpec *vm.InstanceSpec,
    maxCount *int,
    clientToken *string,
) *CreateInstancesRequest {

    return &CreateInstancesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceSpec: instanceSpec,
        MaxCount: maxCount,
        ClientToken: clientToken,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateInstancesRequestWithoutParam() *CreateInstancesRequest {

    return &CreateInstancesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID。(Required) */
func (r *CreateInstancesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceSpec: 实例配置。
(Required) */
func (r *CreateInstancesRequest) SetInstanceSpec(instanceSpec *vm.InstanceSpec) {
    r.InstanceSpec = instanceSpec
}

/* param maxCount: 创建实例的数量，不能超过用户配额。
取值范围：[1,100]；默认值：1。
如果在弹性网卡中指定了内网IP地址，那么单次创建 `maxCount` 只能是 1。
(Optional) */
func (r *CreateInstancesRequest) SetMaxCount(maxCount int) {
    r.MaxCount = &maxCount
}

/* param clientToken: 用于保证请求的幂等性。由客户端生成，并确保不同请求中该参数唯一，长度不能超过64个字符。
(Optional) */
func (r *CreateInstancesRequest) SetClientToken(clientToken string) {
    r.ClientToken = &clientToken
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateInstancesRequest) GetRegionId() string {
    return r.RegionId
}

type CreateInstancesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateInstancesResult `json:"result"`
}

type CreateInstancesResult struct {
    InstanceIds []string `json:"instanceIds"`
}