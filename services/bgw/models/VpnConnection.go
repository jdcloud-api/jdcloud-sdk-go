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

package models

import charge "github.com/jdcloud-api/jdcloud-sdk-go/services/charge/models"

type VpnConnection struct {

    /* VPN connection的Id (Optional) */
    VpnConnectionId string `json:"vpnConnectionId"`

    /* VPN connection的名称 (Optional) */
    VpnConnectionName string `json:"vpnConnectionName"`

    /* 边界网关的Id (Optional) */
    BgwId string `json:"bgwId"`

    /* 客户网关的Id (Optional) */
    CgwId string `json:"cgwId"`

    /* 是否使能BGP路由 (Optional) */
    BgpEnabled bool `json:"bgpEnabled"`

    /* 本地的BGP ASN号 (Optional) */
    LocalAsn int `json:"localAsn"`

    /* VPN connection上分配的本端公网可路由的两个IPv4地址 (Optional) */
    CloudPublicIp []string `json:"cloudPublicIp"`

    /* VPN连接的2个公网IP线路信息。当VPN为标准VPN时，2个线路都为bgp。当VPN为边缘VPN时，显示使用的2个公网IP线路 (Optional) */
    Providers []string `json:"providers"`

    /* 当前隧道状态是否满足高可用, 取值范围为: redundancy, no_redundancy (Optional) */
    HaStatus string `json:"haStatus"`

    /* VPN connection的描述 (Optional) */
    Description string `json:"description"`

    /* 客户网关的创建时间 (Optional) */
    CreatedTime string `json:"createdTime"`

    /* 客户网关的更新时间 (Optional) */
    UpdatedTime string `json:"updatedTime"`

    /* 计费信息 (Optional) */
    Charge charge.Charge `json:"charge"`

    /* vpn connection policy，IPSec VPN的感兴趣流，用于第二阶段协商 (Optional) */
    TrafficSelectors []VpnTrafficSelectorSpec `json:"trafficSelectors"`

    /* VPN az类型，取值：standard(标准VPN)，edge(边缘VPN) (Optional) */
    AzType string `json:"azType"`

    /* VPN可用区 (Optional) */
    Az string `json:"az"`
}
