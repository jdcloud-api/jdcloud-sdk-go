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


type PrivateVif struct {

    /* 通道的Id (Optional) */
    PrivateVifId string `json:"privateVifId"`

    /* 通道的名称 (Optional) */
    PrivateVifName string `json:"privateVifName"`

    /* 连接Id (Optional) */
    ConnectionId string `json:"connectionId"`

    /* 通道类型,hosted:托管通道; directed:专线通道 (Optional) */
    Type string `json:"type"`

    /* 通道的owner:用户pin (Optional) */
    PrivateVifOwner string `json:"privateVifOwner"`

    /* 边界网关的Id (Optional) */
    BgwId string `json:"bgwId"`

    /* 通道涉及BgpPeer信息 (Optional) */
    BgpPeers []BgpPeer `json:"bgpPeers"`

    /* 通道的业务vlan,取值范围 [1,4000] (Optional) */
    Vlan int `json:"vlan"`

    /* 通道的状态，取值为：Confirming(待确认),Rejected（已拒绝）,Pending（配置中）,Active（可用）,Down（不可用）,Deleting（删除中）,Deleted（已删除）,Unknown（未知状态）,InActive(不可用) (Optional) */
    PrivateVifStatus string `json:"privateVifStatus"`

    /* 通道的描述 (Optional) */
    Description string `json:"description"`

    /* 通道创建的时间 (Optional) */
    CreatedTime string `json:"createdTime"`

    /* 是否开启bgp，仅支持开启 (Optional) */
    EnableBgp bool `json:"enableBgp"`

    /* 通道Peer相关信息详情 (Optional) */
    Peers []Peer `json:"peers"`

    /* 通道Bgp相关信息详情 (Optional) */
    Bgp Bgp `json:"bgp"`

    /* 通道健康检查相关信息详情 (Optional) */
    HealthCheck HealthCheck `json:"healthCheck"`
}
