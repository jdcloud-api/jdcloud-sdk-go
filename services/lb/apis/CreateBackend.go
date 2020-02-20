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
    lb "github.com/jdcloud-api/jdcloud-sdk-go/services/lb/models"
)

type CreateBackendRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 后端服务名字,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符  */
    BackendName string `json:"backendName"`

    /* 后端服务所属负载均衡的Id  */
    LoadBalancerId string `json:"loadBalancerId"`

    /* 后端服务的协议 <br>【alb】取值范围：Http、Tcp <br>【nlb】取值范围：Tcp <br>【dnlb】取值范围：Tcp  */
    Protocol string `json:"protocol"`

    /* 后端服务的端口，取值范围为[1, 65535]，如指定了TargetSpec中的port，实际按照target指定的port进行转发  */
    Port int `json:"port"`

    /* 健康检查信息  */
    HealthCheckSpec *lb.HealthCheckSpec `json:"healthCheckSpec"`

    /* 调度算法 <br>【alb,nlb】取值范围为[IpHash, RoundRobin, LeastConn]（取值范围的含义：加权源Ip哈希，加权轮询和加权最小连接），alb和nlb默认为加权轮询 <br>【dnlb】取值范围为[IpHash, QuintupleHash]（取值范围的含义分别为：加权源Ip哈希和加权五元组哈希），dnlb默认为加权源Ip哈希 (Optional) */
    Algorithm *string `json:"algorithm"`

    /* 虚拟服务器组的Id列表，目前只支持一个，且与agIds不能同时存在 (Optional) */
    TargetGroupIds []string `json:"targetGroupIds"`

    /* 高可用组的Id列表，目前只支持一个，且与targetGroupIds不能同时存在 (Optional) */
    AgIds []string `json:"agIds"`

    /* 【alb Tcp协议】获取真实ip, 取值为False(不获取)或者True(获取,支持Proxy Protocol v1版本)，默认为False (Optional) */
    ProxyProtocol *bool `json:"proxyProtocol"`

    /* 描述,允许输入UTF-8编码下的全部字符，不超过256字符 (Optional) */
    Description *string `json:"description"`

    /* 会话保持, 取值为false(不开启)或者true(开启)，默认为false <br>【alb Http协议，RoundRobin算法】支持基于cookie的会话保持 <br>【nlb】支持基于报文源目的IP的会话保持 (Optional) */
    SessionStickiness *bool `json:"sessionStickiness"`

    /* 【nlb】会话保持超时时间，sessionStickiness开启时生效，默认300s, 取值范围[1-3600] (Optional) */
    SessionStickyTimeout *int `json:"sessionStickyTimeout"`

    /* 【nlb】连接耗尽超时。移除target前，连接的最大保持时间，默认300s，取值范围[0-3600] (Optional) */
    ConnectionDrainingSeconds *int `json:"connectionDrainingSeconds"`

    /* 【alb Http协议】cookie的过期时间,sessionStickiness开启时生效，取值范围为[0-86400], 默认为0（表示cookie与浏览器同生命周期） (Optional) */
    HttpCookieExpireSeconds *int `json:"httpCookieExpireSeconds"`

    /* 【alb Http协议】获取负载均衡的协议, 取值为False(不获取)或True(获取), 默认为False (Optional) */
    HttpForwardedProtocol *bool `json:"httpForwardedProtocol"`

    /* 【alb Http协议】获取负载均衡的端口, 取值为False(不获取)或True(获取), 默认为False (Optional) */
    HttpForwardedPort *bool `json:"httpForwardedPort"`

    /* 【alb Http协议】获取负载均衡的host信息, 取值为False(不获取)或True(获取), 默认为False (Optional) */
    HttpForwardedHost *bool `json:"httpForwardedHost"`

    /* 【alb Http协议】获取负载均衡的vip, 取值为False(不获取)或True(获取), 默认为False (Optional) */
    HttpForwardedVip *bool `json:"httpForwardedVip"`
}

/*
 * param regionId: Region ID (Required)
 * param backendName: 后端服务名字,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符 (Required)
 * param loadBalancerId: 后端服务所属负载均衡的Id (Required)
 * param protocol: 后端服务的协议 <br>【alb】取值范围：Http、Tcp <br>【nlb】取值范围：Tcp <br>【dnlb】取值范围：Tcp (Required)
 * param port: 后端服务的端口，取值范围为[1, 65535]，如指定了TargetSpec中的port，实际按照target指定的port进行转发 (Required)
 * param healthCheckSpec: 健康检查信息 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateBackendRequest(
    regionId string,
    backendName string,
    loadBalancerId string,
    protocol string,
    port int,
    healthCheckSpec *lb.HealthCheckSpec,
) *CreateBackendRequest {

	return &CreateBackendRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/backends/",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        BackendName: backendName,
        LoadBalancerId: loadBalancerId,
        Protocol: protocol,
        Port: port,
        HealthCheckSpec: healthCheckSpec,
	}
}

/*
 * param regionId: Region ID (Required)
 * param backendName: 后端服务名字,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符 (Required)
 * param loadBalancerId: 后端服务所属负载均衡的Id (Required)
 * param protocol: 后端服务的协议 <br>【alb】取值范围：Http、Tcp <br>【nlb】取值范围：Tcp <br>【dnlb】取值范围：Tcp (Required)
 * param port: 后端服务的端口，取值范围为[1, 65535]，如指定了TargetSpec中的port，实际按照target指定的port进行转发 (Required)
 * param healthCheckSpec: 健康检查信息 (Required)
 * param algorithm: 调度算法 <br>【alb,nlb】取值范围为[IpHash, RoundRobin, LeastConn]（取值范围的含义：加权源Ip哈希，加权轮询和加权最小连接），alb和nlb默认为加权轮询 <br>【dnlb】取值范围为[IpHash, QuintupleHash]（取值范围的含义分别为：加权源Ip哈希和加权五元组哈希），dnlb默认为加权源Ip哈希 (Optional)
 * param targetGroupIds: 虚拟服务器组的Id列表，目前只支持一个，且与agIds不能同时存在 (Optional)
 * param agIds: 高可用组的Id列表，目前只支持一个，且与targetGroupIds不能同时存在 (Optional)
 * param proxyProtocol: 【alb Tcp协议】获取真实ip, 取值为False(不获取)或者True(获取,支持Proxy Protocol v1版本)，默认为False (Optional)
 * param description: 描述,允许输入UTF-8编码下的全部字符，不超过256字符 (Optional)
 * param sessionStickiness: 会话保持, 取值为false(不开启)或者true(开启)，默认为false <br>【alb Http协议，RoundRobin算法】支持基于cookie的会话保持 <br>【nlb】支持基于报文源目的IP的会话保持 (Optional)
 * param sessionStickyTimeout: 【nlb】会话保持超时时间，sessionStickiness开启时生效，默认300s, 取值范围[1-3600] (Optional)
 * param connectionDrainingSeconds: 【nlb】连接耗尽超时。移除target前，连接的最大保持时间，默认300s，取值范围[0-3600] (Optional)
 * param httpCookieExpireSeconds: 【alb Http协议】cookie的过期时间,sessionStickiness开启时生效，取值范围为[0-86400], 默认为0（表示cookie与浏览器同生命周期） (Optional)
 * param httpForwardedProtocol: 【alb Http协议】获取负载均衡的协议, 取值为False(不获取)或True(获取), 默认为False (Optional)
 * param httpForwardedPort: 【alb Http协议】获取负载均衡的端口, 取值为False(不获取)或True(获取), 默认为False (Optional)
 * param httpForwardedHost: 【alb Http协议】获取负载均衡的host信息, 取值为False(不获取)或True(获取), 默认为False (Optional)
 * param httpForwardedVip: 【alb Http协议】获取负载均衡的vip, 取值为False(不获取)或True(获取), 默认为False (Optional)
 */
func NewCreateBackendRequestWithAllParams(
    regionId string,
    backendName string,
    loadBalancerId string,
    protocol string,
    port int,
    healthCheckSpec *lb.HealthCheckSpec,
    algorithm *string,
    targetGroupIds []string,
    agIds []string,
    proxyProtocol *bool,
    description *string,
    sessionStickiness *bool,
    sessionStickyTimeout *int,
    connectionDrainingSeconds *int,
    httpCookieExpireSeconds *int,
    httpForwardedProtocol *bool,
    httpForwardedPort *bool,
    httpForwardedHost *bool,
    httpForwardedVip *bool,
) *CreateBackendRequest {

    return &CreateBackendRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/backends/",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        BackendName: backendName,
        LoadBalancerId: loadBalancerId,
        Protocol: protocol,
        Port: port,
        HealthCheckSpec: healthCheckSpec,
        Algorithm: algorithm,
        TargetGroupIds: targetGroupIds,
        AgIds: agIds,
        ProxyProtocol: proxyProtocol,
        Description: description,
        SessionStickiness: sessionStickiness,
        SessionStickyTimeout: sessionStickyTimeout,
        ConnectionDrainingSeconds: connectionDrainingSeconds,
        HttpCookieExpireSeconds: httpCookieExpireSeconds,
        HttpForwardedProtocol: httpForwardedProtocol,
        HttpForwardedPort: httpForwardedPort,
        HttpForwardedHost: httpForwardedHost,
        HttpForwardedVip: httpForwardedVip,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateBackendRequestWithoutParam() *CreateBackendRequest {

    return &CreateBackendRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/backends/",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CreateBackendRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param backendName: 后端服务名字,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符(Required) */
func (r *CreateBackendRequest) SetBackendName(backendName string) {
    r.BackendName = backendName
}

/* param loadBalancerId: 后端服务所属负载均衡的Id(Required) */
func (r *CreateBackendRequest) SetLoadBalancerId(loadBalancerId string) {
    r.LoadBalancerId = loadBalancerId
}

/* param protocol: 后端服务的协议 <br>【alb】取值范围：Http、Tcp <br>【nlb】取值范围：Tcp <br>【dnlb】取值范围：Tcp(Required) */
func (r *CreateBackendRequest) SetProtocol(protocol string) {
    r.Protocol = protocol
}

/* param port: 后端服务的端口，取值范围为[1, 65535]，如指定了TargetSpec中的port，实际按照target指定的port进行转发(Required) */
func (r *CreateBackendRequest) SetPort(port int) {
    r.Port = port
}

/* param healthCheckSpec: 健康检查信息(Required) */
func (r *CreateBackendRequest) SetHealthCheckSpec(healthCheckSpec *lb.HealthCheckSpec) {
    r.HealthCheckSpec = healthCheckSpec
}

/* param algorithm: 调度算法 <br>【alb,nlb】取值范围为[IpHash, RoundRobin, LeastConn]（取值范围的含义：加权源Ip哈希，加权轮询和加权最小连接），alb和nlb默认为加权轮询 <br>【dnlb】取值范围为[IpHash, QuintupleHash]（取值范围的含义分别为：加权源Ip哈希和加权五元组哈希），dnlb默认为加权源Ip哈希(Optional) */
func (r *CreateBackendRequest) SetAlgorithm(algorithm string) {
    r.Algorithm = &algorithm
}

/* param targetGroupIds: 虚拟服务器组的Id列表，目前只支持一个，且与agIds不能同时存在(Optional) */
func (r *CreateBackendRequest) SetTargetGroupIds(targetGroupIds []string) {
    r.TargetGroupIds = targetGroupIds
}

/* param agIds: 高可用组的Id列表，目前只支持一个，且与targetGroupIds不能同时存在(Optional) */
func (r *CreateBackendRequest) SetAgIds(agIds []string) {
    r.AgIds = agIds
}

/* param proxyProtocol: 【alb Tcp协议】获取真实ip, 取值为False(不获取)或者True(获取,支持Proxy Protocol v1版本)，默认为False(Optional) */
func (r *CreateBackendRequest) SetProxyProtocol(proxyProtocol bool) {
    r.ProxyProtocol = &proxyProtocol
}

/* param description: 描述,允许输入UTF-8编码下的全部字符，不超过256字符(Optional) */
func (r *CreateBackendRequest) SetDescription(description string) {
    r.Description = &description
}

/* param sessionStickiness: 会话保持, 取值为false(不开启)或者true(开启)，默认为false <br>【alb Http协议，RoundRobin算法】支持基于cookie的会话保持 <br>【nlb】支持基于报文源目的IP的会话保持(Optional) */
func (r *CreateBackendRequest) SetSessionStickiness(sessionStickiness bool) {
    r.SessionStickiness = &sessionStickiness
}

/* param sessionStickyTimeout: 【nlb】会话保持超时时间，sessionStickiness开启时生效，默认300s, 取值范围[1-3600](Optional) */
func (r *CreateBackendRequest) SetSessionStickyTimeout(sessionStickyTimeout int) {
    r.SessionStickyTimeout = &sessionStickyTimeout
}

/* param connectionDrainingSeconds: 【nlb】连接耗尽超时。移除target前，连接的最大保持时间，默认300s，取值范围[0-3600](Optional) */
func (r *CreateBackendRequest) SetConnectionDrainingSeconds(connectionDrainingSeconds int) {
    r.ConnectionDrainingSeconds = &connectionDrainingSeconds
}

/* param httpCookieExpireSeconds: 【alb Http协议】cookie的过期时间,sessionStickiness开启时生效，取值范围为[0-86400], 默认为0（表示cookie与浏览器同生命周期）(Optional) */
func (r *CreateBackendRequest) SetHttpCookieExpireSeconds(httpCookieExpireSeconds int) {
    r.HttpCookieExpireSeconds = &httpCookieExpireSeconds
}

/* param httpForwardedProtocol: 【alb Http协议】获取负载均衡的协议, 取值为False(不获取)或True(获取), 默认为False(Optional) */
func (r *CreateBackendRequest) SetHttpForwardedProtocol(httpForwardedProtocol bool) {
    r.HttpForwardedProtocol = &httpForwardedProtocol
}

/* param httpForwardedPort: 【alb Http协议】获取负载均衡的端口, 取值为False(不获取)或True(获取), 默认为False(Optional) */
func (r *CreateBackendRequest) SetHttpForwardedPort(httpForwardedPort bool) {
    r.HttpForwardedPort = &httpForwardedPort
}

/* param httpForwardedHost: 【alb Http协议】获取负载均衡的host信息, 取值为False(不获取)或True(获取), 默认为False(Optional) */
func (r *CreateBackendRequest) SetHttpForwardedHost(httpForwardedHost bool) {
    r.HttpForwardedHost = &httpForwardedHost
}

/* param httpForwardedVip: 【alb Http协议】获取负载均衡的vip, 取值为False(不获取)或True(获取), 默认为False(Optional) */
func (r *CreateBackendRequest) SetHttpForwardedVip(httpForwardedVip bool) {
    r.HttpForwardedVip = &httpForwardedVip
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateBackendRequest) GetRegionId() string {
    return r.RegionId
}

type CreateBackendResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateBackendResult `json:"result"`
}

type CreateBackendResult struct {
    BackendId string `json:"backendId"`
}