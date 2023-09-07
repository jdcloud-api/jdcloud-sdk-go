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
    rds "github.com/jdcloud-api/jdcloud-sdk-go/services/rds/models"
)

type ModifyReadWriteProxyRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* 读写分离代理服务ID  */
    ReadWriteProxyId string `json:"readWriteProxyId"`

    /* 延迟阈值，范围是1~1000，单位：秒，默认为100，仅MySQL (Optional) */
    DelayThreshold *int `json:"delayThreshold"`

    /* wal日志延迟阈值，范围是1~1024，单位：MB，默认为200，仅PostgreSQL (Optional) */
    WalDelayThreshold *int `json:"walDelayThreshold"`

    /* 读写分离代理后端实例负载均衡策略，默认值为LEAST_CURRENT_OPERATIONS；当前支持的负载均衡策略请查看[枚举参数定义](../Enum-Definitions/Enum-Definitions.md) (Optional) */
    LoadBalancerPolicy *string `json:"loadBalancerPolicy"`

    /* 后端实例健康检查配置 (Optional) */
    HealthCheckSpec *rds.HealthCheckSpec `json:"healthCheckSpec"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param readWriteProxyId: 读写分离代理服务ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyReadWriteProxyRequest(
    regionId string,
    readWriteProxyId string,
) *ModifyReadWriteProxyRequest {

	return &ModifyReadWriteProxyRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/readWriteProxy/{readWriteProxyId}:modifyReadWriteProxy",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ReadWriteProxyId: readWriteProxyId,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param readWriteProxyId: 读写分离代理服务ID (Required)
 * param delayThreshold: 延迟阈值，范围是1~1000，单位：秒，默认为100，仅MySQL (Optional)
 * param walDelayThreshold: wal日志延迟阈值，范围是1~1024，单位：MB，默认为200，仅PostgreSQL (Optional)
 * param loadBalancerPolicy: 读写分离代理后端实例负载均衡策略，默认值为LEAST_CURRENT_OPERATIONS；当前支持的负载均衡策略请查看[枚举参数定义](../Enum-Definitions/Enum-Definitions.md) (Optional)
 * param healthCheckSpec: 后端实例健康检查配置 (Optional)
 */
func NewModifyReadWriteProxyRequestWithAllParams(
    regionId string,
    readWriteProxyId string,
    delayThreshold *int,
    walDelayThreshold *int,
    loadBalancerPolicy *string,
    healthCheckSpec *rds.HealthCheckSpec,
) *ModifyReadWriteProxyRequest {

    return &ModifyReadWriteProxyRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/readWriteProxy/{readWriteProxyId}:modifyReadWriteProxy",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ReadWriteProxyId: readWriteProxyId,
        DelayThreshold: delayThreshold,
        WalDelayThreshold: walDelayThreshold,
        LoadBalancerPolicy: loadBalancerPolicy,
        HealthCheckSpec: healthCheckSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyReadWriteProxyRequestWithoutParam() *ModifyReadWriteProxyRequest {

    return &ModifyReadWriteProxyRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/readWriteProxy/{readWriteProxyId}:modifyReadWriteProxy",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *ModifyReadWriteProxyRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param readWriteProxyId: 读写分离代理服务ID(Required) */
func (r *ModifyReadWriteProxyRequest) SetReadWriteProxyId(readWriteProxyId string) {
    r.ReadWriteProxyId = readWriteProxyId
}
/* param delayThreshold: 延迟阈值，范围是1~1000，单位：秒，默认为100，仅MySQL(Optional) */
func (r *ModifyReadWriteProxyRequest) SetDelayThreshold(delayThreshold int) {
    r.DelayThreshold = &delayThreshold
}
/* param walDelayThreshold: wal日志延迟阈值，范围是1~1024，单位：MB，默认为200，仅PostgreSQL(Optional) */
func (r *ModifyReadWriteProxyRequest) SetWalDelayThreshold(walDelayThreshold int) {
    r.WalDelayThreshold = &walDelayThreshold
}
/* param loadBalancerPolicy: 读写分离代理后端实例负载均衡策略，默认值为LEAST_CURRENT_OPERATIONS；当前支持的负载均衡策略请查看[枚举参数定义](../Enum-Definitions/Enum-Definitions.md)(Optional) */
func (r *ModifyReadWriteProxyRequest) SetLoadBalancerPolicy(loadBalancerPolicy string) {
    r.LoadBalancerPolicy = &loadBalancerPolicy
}
/* param healthCheckSpec: 后端实例健康检查配置(Optional) */
func (r *ModifyReadWriteProxyRequest) SetHealthCheckSpec(healthCheckSpec *rds.HealthCheckSpec) {
    r.HealthCheckSpec = healthCheckSpec
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyReadWriteProxyRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyReadWriteProxyResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyReadWriteProxyResult `json:"result"`
}

type ModifyReadWriteProxyResult struct {
}