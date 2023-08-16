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


type CreatePrivateVifSpec struct {

    /* PrivateVif的名称参考  */
    PrivateVifName string `json:"privateVifName"`

    /* 物理连接的Id  */
    ConnectionId string `json:"connectionId"`

    /* 通道的owner:用户pin (Optional) */
    PrivateVifOwner string `json:"privateVifOwner"`

    /* 边界网关的Id  */
    BgwId string `json:"bgwId"`

    /* 通道的业务vlan,取值范围 [1,4000]  */
    Vlan int `json:"vlan"`

    /* 通道涉及BgpPeer信息(即将弃用) (Optional) */
    BgpPeers []BgpPeerSpec `json:"bgpPeers"`

    /* PrivateVif的描述 (Optional) */
    Description string `json:"description"`

    /* 是否开启bgp，仅支持开启 (Optional) */
    EnableBgp bool `json:"enableBgp"`

    /* 通道Peer的ip信息 (Optional) */
    PeerIps []PeerIpsSpec `json:"peerIps"`

    /* 通道Bgp相关信息详情 (Optional) */
    BgpSpec BgpSpec `json:"bgpSpec"`

    /* 通道健康检查相关信息详情 (Optional) */
    HealthCheck HealthCheckSpec `json:"healthCheck"`
}
