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

type CreateFirewallRuleRequest struct {

    core.JDCloudRequest

    /* regionId
  */
    RegionId string `json:"regionId"`

    /* 指定的轻量应用云主机的实例ID
  */
    InstanceId string `json:"instanceId"`

    /* 源Ip的CRDI格式的地址
  */
    SourceAddress string `json:"sourceAddress"`

    /* 规则限定协议。取值范围：
TCP：TCP协议。
UDP：UDP协议。
ICMP：ICMP协议。
  */
    RuleProtocol string `json:"ruleProtocol"`

    /* 端口范围。若规则限定协议为传输层协议（TCP、UDP)，取值范围为1`~`65535，若规则限定协议为非传输层协议（ICMP协议），恒为0。使用正斜线（/）隔开起始端口和终止端口，例如：1024/1055表示端口范围为1024`~`1055。
  */
    Port string `json:"port"`

    /* 防火墙规则的备注, 不超过100个字符
 (Optional) */
    Remark *string `json:"remark"`

    /* 用于保证请求的幂等性。由客户端生成，并确保不同请求中该参数唯一，长度不能超过64个字符。
 (Optional) */
    ClientToken *string `json:"clientToken"`
}

/*
 * param regionId: regionId
 (Required)
 * param instanceId: 指定的轻量应用云主机的实例ID
 (Required)
 * param sourceAddress: 源Ip的CRDI格式的地址
 (Required)
 * param ruleProtocol: 规则限定协议。取值范围：
TCP：TCP协议。
UDP：UDP协议。
ICMP：ICMP协议。
 (Required)
 * param port: 端口范围。若规则限定协议为传输层协议（TCP、UDP)，取值范围为1`~`65535，若规则限定协议为非传输层协议（ICMP协议），恒为0。使用正斜线（/）隔开起始端口和终止端口，例如：1024/1055表示端口范围为1024`~`1055。
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateFirewallRuleRequest(
    regionId string,
    instanceId string,
    sourceAddress string,
    ruleProtocol string,
    port string,
) *CreateFirewallRuleRequest {

	return &CreateFirewallRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/firewallRule",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        SourceAddress: sourceAddress,
        RuleProtocol: ruleProtocol,
        Port: port,
	}
}

/*
 * param regionId: regionId
 (Required)
 * param instanceId: 指定的轻量应用云主机的实例ID
 (Required)
 * param sourceAddress: 源Ip的CRDI格式的地址
 (Required)
 * param ruleProtocol: 规则限定协议。取值范围：
TCP：TCP协议。
UDP：UDP协议。
ICMP：ICMP协议。
 (Required)
 * param port: 端口范围。若规则限定协议为传输层协议（TCP、UDP)，取值范围为1`~`65535，若规则限定协议为非传输层协议（ICMP协议），恒为0。使用正斜线（/）隔开起始端口和终止端口，例如：1024/1055表示端口范围为1024`~`1055。
 (Required)
 * param remark: 防火墙规则的备注, 不超过100个字符
 (Optional)
 * param clientToken: 用于保证请求的幂等性。由客户端生成，并确保不同请求中该参数唯一，长度不能超过64个字符。
 (Optional)
 */
func NewCreateFirewallRuleRequestWithAllParams(
    regionId string,
    instanceId string,
    sourceAddress string,
    ruleProtocol string,
    port string,
    remark *string,
    clientToken *string,
) *CreateFirewallRuleRequest {

    return &CreateFirewallRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/firewallRule",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        SourceAddress: sourceAddress,
        RuleProtocol: ruleProtocol,
        Port: port,
        Remark: remark,
        ClientToken: clientToken,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateFirewallRuleRequestWithoutParam() *CreateFirewallRuleRequest {

    return &CreateFirewallRuleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/firewallRule",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: regionId
(Required) */
func (r *CreateFirewallRuleRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param instanceId: 指定的轻量应用云主机的实例ID
(Required) */
func (r *CreateFirewallRuleRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}
/* param sourceAddress: 源Ip的CRDI格式的地址
(Required) */
func (r *CreateFirewallRuleRequest) SetSourceAddress(sourceAddress string) {
    r.SourceAddress = sourceAddress
}
/* param ruleProtocol: 规则限定协议。取值范围：
TCP：TCP协议。
UDP：UDP协议。
ICMP：ICMP协议。
(Required) */
func (r *CreateFirewallRuleRequest) SetRuleProtocol(ruleProtocol string) {
    r.RuleProtocol = ruleProtocol
}
/* param port: 端口范围。若规则限定协议为传输层协议（TCP、UDP)，取值范围为1`~`65535，若规则限定协议为非传输层协议（ICMP协议），恒为0。使用正斜线（/）隔开起始端口和终止端口，例如：1024/1055表示端口范围为1024`~`1055。
(Required) */
func (r *CreateFirewallRuleRequest) SetPort(port string) {
    r.Port = port
}
/* param remark: 防火墙规则的备注, 不超过100个字符
(Optional) */
func (r *CreateFirewallRuleRequest) SetRemark(remark string) {
    r.Remark = &remark
}
/* param clientToken: 用于保证请求的幂等性。由客户端生成，并确保不同请求中该参数唯一，长度不能超过64个字符。
(Optional) */
func (r *CreateFirewallRuleRequest) SetClientToken(clientToken string) {
    r.ClientToken = &clientToken
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateFirewallRuleRequest) GetRegionId() string {
    return r.RegionId
}

type CreateFirewallRuleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateFirewallRuleResult `json:"result"`
}

type CreateFirewallRuleResult struct {
    FirewallId string `json:"firewallId"`
}