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

type UpdateListenerRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 监听器ID  */
    ListenerId string `json:"listenerId"`

    /* 监听器名称,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符 (Optional) */
    ListenerName *string `json:"listenerName"`

    /* Listener状态, 取值为On或者为Off (Optional) */
    Status *string `json:"status"`

    /* 【alb使用https时支持】是否开启HSTS，True(开启)， False(关闭)，缺省为不改变原值 (Optional) */
    HstsEnable *bool `json:"hstsEnable"`

    /* 【alb使用https时支持】HSTS过期时间(秒)，取值范围为[1, 94608000(3年)]，缺省为不改变原值 (Optional) */
    HstsMaxAge *int `json:"hstsMaxAge"`

    /* 【alb Https和Tls协议】Listener绑定的默认证书，最多支持两个，两个证书的加密算法需要不同 (Optional) */
    CertificateSpecs []lb.CertificateSpec `json:"certificateSpecs"`

    /* 【仅ALB支持】限速配置 (Optional) */
    Limitation *lb.LimitationSpec `json:"limitation"`

    /* 【alb、nlb】空闲连接超时时间, 范围为[1,86400]。 <br>（Tcp和Tls协议）默认为：1800s <br>（Http和Https协议）默认为：60s <br>【dnlb】不支持该功能 (Optional) */
    ConnectionIdleTimeSeconds *int `json:"connectionIdleTimeSeconds"`

    /* 默认后端服务Id (Optional) */
    BackendId *string `json:"backendId"`

    /* 【alb Https和Http协议】转发规则组Id (Optional) */
    UrlMapId *string `json:"urlMapId"`

    /* 监听器描述,允许输入UTF-8编码下的全部字符，不超过256字符 (Optional) */
    Description *string `json:"description"`

    /* 绑定的安全策略id，仅支持应用负载均衡的HTTPS、TLS监听配置 (Optional) */
    SecurityPolicyId *string `json:"securityPolicyId"`
}

/*
 * param regionId: Region ID (Required)
 * param listenerId: 监听器ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateListenerRequest(
    regionId string,
    listenerId string,
) *UpdateListenerRequest {

	return &UpdateListenerRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/listeners/{listenerId}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ListenerId: listenerId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param listenerId: 监听器ID (Required)
 * param listenerName: 监听器名称,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符 (Optional)
 * param status: Listener状态, 取值为On或者为Off (Optional)
 * param hstsEnable: 【alb使用https时支持】是否开启HSTS，True(开启)， False(关闭)，缺省为不改变原值 (Optional)
 * param hstsMaxAge: 【alb使用https时支持】HSTS过期时间(秒)，取值范围为[1, 94608000(3年)]，缺省为不改变原值 (Optional)
 * param certificateSpecs: 【alb Https和Tls协议】Listener绑定的默认证书，最多支持两个，两个证书的加密算法需要不同 (Optional)
 * param limitation: 【仅ALB支持】限速配置 (Optional)
 * param connectionIdleTimeSeconds: 【alb、nlb】空闲连接超时时间, 范围为[1,86400]。 <br>（Tcp和Tls协议）默认为：1800s <br>（Http和Https协议）默认为：60s <br>【dnlb】不支持该功能 (Optional)
 * param backendId: 默认后端服务Id (Optional)
 * param urlMapId: 【alb Https和Http协议】转发规则组Id (Optional)
 * param description: 监听器描述,允许输入UTF-8编码下的全部字符，不超过256字符 (Optional)
 * param securityPolicyId: 绑定的安全策略id，仅支持应用负载均衡的HTTPS、TLS监听配置 (Optional)
 */
func NewUpdateListenerRequestWithAllParams(
    regionId string,
    listenerId string,
    listenerName *string,
    status *string,
    hstsEnable *bool,
    hstsMaxAge *int,
    certificateSpecs []lb.CertificateSpec,
    limitation *lb.LimitationSpec,
    connectionIdleTimeSeconds *int,
    backendId *string,
    urlMapId *string,
    description *string,
    securityPolicyId *string,
) *UpdateListenerRequest {

    return &UpdateListenerRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/listeners/{listenerId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ListenerId: listenerId,
        ListenerName: listenerName,
        Status: status,
        HstsEnable: hstsEnable,
        HstsMaxAge: hstsMaxAge,
        CertificateSpecs: certificateSpecs,
        Limitation: limitation,
        ConnectionIdleTimeSeconds: connectionIdleTimeSeconds,
        BackendId: backendId,
        UrlMapId: urlMapId,
        Description: description,
        SecurityPolicyId: securityPolicyId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateListenerRequestWithoutParam() *UpdateListenerRequest {

    return &UpdateListenerRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/listeners/{listenerId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *UpdateListenerRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param listenerId: 监听器ID(Required) */
func (r *UpdateListenerRequest) SetListenerId(listenerId string) {
    r.ListenerId = listenerId
}
/* param listenerName: 监听器名称,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符(Optional) */
func (r *UpdateListenerRequest) SetListenerName(listenerName string) {
    r.ListenerName = &listenerName
}
/* param status: Listener状态, 取值为On或者为Off(Optional) */
func (r *UpdateListenerRequest) SetStatus(status string) {
    r.Status = &status
}
/* param hstsEnable: 【alb使用https时支持】是否开启HSTS，True(开启)， False(关闭)，缺省为不改变原值(Optional) */
func (r *UpdateListenerRequest) SetHstsEnable(hstsEnable bool) {
    r.HstsEnable = &hstsEnable
}
/* param hstsMaxAge: 【alb使用https时支持】HSTS过期时间(秒)，取值范围为[1, 94608000(3年)]，缺省为不改变原值(Optional) */
func (r *UpdateListenerRequest) SetHstsMaxAge(hstsMaxAge int) {
    r.HstsMaxAge = &hstsMaxAge
}
/* param certificateSpecs: 【alb Https和Tls协议】Listener绑定的默认证书，最多支持两个，两个证书的加密算法需要不同(Optional) */
func (r *UpdateListenerRequest) SetCertificateSpecs(certificateSpecs []lb.CertificateSpec) {
    r.CertificateSpecs = certificateSpecs
}
/* param limitation: 【仅ALB支持】限速配置(Optional) */
func (r *UpdateListenerRequest) SetLimitation(limitation *lb.LimitationSpec) {
    r.Limitation = limitation
}
/* param connectionIdleTimeSeconds: 【alb、nlb】空闲连接超时时间, 范围为[1,86400]。 <br>（Tcp和Tls协议）默认为：1800s <br>（Http和Https协议）默认为：60s <br>【dnlb】不支持该功能(Optional) */
func (r *UpdateListenerRequest) SetConnectionIdleTimeSeconds(connectionIdleTimeSeconds int) {
    r.ConnectionIdleTimeSeconds = &connectionIdleTimeSeconds
}
/* param backendId: 默认后端服务Id(Optional) */
func (r *UpdateListenerRequest) SetBackendId(backendId string) {
    r.BackendId = &backendId
}
/* param urlMapId: 【alb Https和Http协议】转发规则组Id(Optional) */
func (r *UpdateListenerRequest) SetUrlMapId(urlMapId string) {
    r.UrlMapId = &urlMapId
}
/* param description: 监听器描述,允许输入UTF-8编码下的全部字符，不超过256字符(Optional) */
func (r *UpdateListenerRequest) SetDescription(description string) {
    r.Description = &description
}
/* param securityPolicyId: 绑定的安全策略id，仅支持应用负载均衡的HTTPS、TLS监听配置(Optional) */
func (r *UpdateListenerRequest) SetSecurityPolicyId(securityPolicyId string) {
    r.SecurityPolicyId = &securityPolicyId
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateListenerRequest) GetRegionId() string {
    return r.RegionId
}

type UpdateListenerResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateListenerResult `json:"result"`
}

type UpdateListenerResult struct {
}