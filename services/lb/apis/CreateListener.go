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

type CreateListenerRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Listener的名字,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符  */
    ListenerName string `json:"listenerName"`

    /* 监听协议, 取值为Tcp, Tls, Http, Https <br>【alb】支持Http, Https，Tcp和Tls <br>【nlb】支持Tcp  <br>【dnlb】支持Tcp  */
    Protocol string `json:"protocol"`

    /* 监听端口，取值范围为[1, 65535]  */
    Port int `json:"port"`

    /* 默认的后端服务Id  */
    BackendId string `json:"backendId"`

    /* Listener所属loadBalancer的Id  */
    LoadBalancerId string `json:"loadBalancerId"`

    /* 【alb Https和Http协议】转发规则组Id (Optional) */
    UrlMapId *string `json:"urlMapId"`

    /* 默认后端服务的转发策略,取值为Forward或Redirect, 现只支持Forward, 默认为Forward (Optional) */
    Action *string `json:"action"`

    /* 【alb Https和Tls协议】ssl server证书列表，现只支持一个证书 (Optional) */
    CertificateSpecs []lb.CertificateSpec `json:"certificateSpecs"`

    /* 【alb、nlb】空闲连接超时时间, 范围为[1,86400]。 <br>（Tcp和Tls协议）默认为：1800s <br>（Http和Https协议）默认为：60s <br>【dnlb】不支持 (Optional) */
    ConnectionIdleTimeSeconds *int `json:"connectionIdleTimeSeconds"`

    /* 描述,允许输入UTF-8编码下的全部字符，不超过256字符 (Optional) */
    Description *string `json:"description"`
}

/*
 * param regionId: Region ID (Required)
 * param listenerName: Listener的名字,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符 (Required)
 * param protocol: 监听协议, 取值为Tcp, Tls, Http, Https <br>【alb】支持Http, Https，Tcp和Tls <br>【nlb】支持Tcp  <br>【dnlb】支持Tcp (Required)
 * param port: 监听端口，取值范围为[1, 65535] (Required)
 * param backendId: 默认的后端服务Id (Required)
 * param loadBalancerId: Listener所属loadBalancer的Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateListenerRequest(
    regionId string,
    listenerName string,
    protocol string,
    port int,
    backendId string,
    loadBalancerId string,
) *CreateListenerRequest {

	return &CreateListenerRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/listeners/",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ListenerName: listenerName,
        Protocol: protocol,
        Port: port,
        BackendId: backendId,
        LoadBalancerId: loadBalancerId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param listenerName: Listener的名字,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符 (Required)
 * param protocol: 监听协议, 取值为Tcp, Tls, Http, Https <br>【alb】支持Http, Https，Tcp和Tls <br>【nlb】支持Tcp  <br>【dnlb】支持Tcp (Required)
 * param port: 监听端口，取值范围为[1, 65535] (Required)
 * param backendId: 默认的后端服务Id (Required)
 * param loadBalancerId: Listener所属loadBalancer的Id (Required)
 * param urlMapId: 【alb Https和Http协议】转发规则组Id (Optional)
 * param action: 默认后端服务的转发策略,取值为Forward或Redirect, 现只支持Forward, 默认为Forward (Optional)
 * param certificateSpecs: 【alb Https和Tls协议】ssl server证书列表，现只支持一个证书 (Optional)
 * param connectionIdleTimeSeconds: 【alb、nlb】空闲连接超时时间, 范围为[1,86400]。 <br>（Tcp和Tls协议）默认为：1800s <br>（Http和Https协议）默认为：60s <br>【dnlb】不支持 (Optional)
 * param description: 描述,允许输入UTF-8编码下的全部字符，不超过256字符 (Optional)
 */
func NewCreateListenerRequestWithAllParams(
    regionId string,
    listenerName string,
    protocol string,
    port int,
    backendId string,
    loadBalancerId string,
    urlMapId *string,
    action *string,
    certificateSpecs []lb.CertificateSpec,
    connectionIdleTimeSeconds *int,
    description *string,
) *CreateListenerRequest {

    return &CreateListenerRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/listeners/",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ListenerName: listenerName,
        Protocol: protocol,
        Port: port,
        BackendId: backendId,
        LoadBalancerId: loadBalancerId,
        UrlMapId: urlMapId,
        Action: action,
        CertificateSpecs: certificateSpecs,
        ConnectionIdleTimeSeconds: connectionIdleTimeSeconds,
        Description: description,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateListenerRequestWithoutParam() *CreateListenerRequest {

    return &CreateListenerRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/listeners/",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CreateListenerRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param listenerName: Listener的名字,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符(Required) */
func (r *CreateListenerRequest) SetListenerName(listenerName string) {
    r.ListenerName = listenerName
}

/* param protocol: 监听协议, 取值为Tcp, Tls, Http, Https <br>【alb】支持Http, Https，Tcp和Tls <br>【nlb】支持Tcp  <br>【dnlb】支持Tcp(Required) */
func (r *CreateListenerRequest) SetProtocol(protocol string) {
    r.Protocol = protocol
}

/* param port: 监听端口，取值范围为[1, 65535](Required) */
func (r *CreateListenerRequest) SetPort(port int) {
    r.Port = port
}

/* param backendId: 默认的后端服务Id(Required) */
func (r *CreateListenerRequest) SetBackendId(backendId string) {
    r.BackendId = backendId
}

/* param loadBalancerId: Listener所属loadBalancer的Id(Required) */
func (r *CreateListenerRequest) SetLoadBalancerId(loadBalancerId string) {
    r.LoadBalancerId = loadBalancerId
}

/* param urlMapId: 【alb Https和Http协议】转发规则组Id(Optional) */
func (r *CreateListenerRequest) SetUrlMapId(urlMapId string) {
    r.UrlMapId = &urlMapId
}

/* param action: 默认后端服务的转发策略,取值为Forward或Redirect, 现只支持Forward, 默认为Forward(Optional) */
func (r *CreateListenerRequest) SetAction(action string) {
    r.Action = &action
}

/* param certificateSpecs: 【alb Https和Tls协议】ssl server证书列表，现只支持一个证书(Optional) */
func (r *CreateListenerRequest) SetCertificateSpecs(certificateSpecs []lb.CertificateSpec) {
    r.CertificateSpecs = certificateSpecs
}

/* param connectionIdleTimeSeconds: 【alb、nlb】空闲连接超时时间, 范围为[1,86400]。 <br>（Tcp和Tls协议）默认为：1800s <br>（Http和Https协议）默认为：60s <br>【dnlb】不支持(Optional) */
func (r *CreateListenerRequest) SetConnectionIdleTimeSeconds(connectionIdleTimeSeconds int) {
    r.ConnectionIdleTimeSeconds = &connectionIdleTimeSeconds
}

/* param description: 描述,允许输入UTF-8编码下的全部字符，不超过256字符(Optional) */
func (r *CreateListenerRequest) SetDescription(description string) {
    r.Description = &description
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateListenerRequest) GetRegionId() string {
    return r.RegionId
}

type CreateListenerResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateListenerResult `json:"result"`
}

type CreateListenerResult struct {
    ListenerId string `json:"listenerId"`
}