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


type IotDataStatsVO struct {

    /* 物类型总数 (Optional) */
    TotalThingTypes int `json:"totalThingTypes"`

    /* 直连终端物类型总数 (Optional) */
    DirectConnectThingTypes int `json:"directConnectThingTypes"`

    /* 直连终端物类型总数 (Optional) */
    ConnectAgentThingTypes int `json:"connectAgentThingTypes"`

    /* 非直连在线终端总数 (Optional) */
    IndirectConnectAgentThingTypes int `json:"indirectConnectAgentThingTypes"`

    /* 总设备数 (Optional) */
    TotalDevices int `json:"totalDevices"`

    /* 直连总设备数 (Optional) */
    DirectConnectDevices int `json:"directConnectDevices"`

    /* 连接终端设备数 (Optional) */
    ConnectAgentDevices int `json:"connectAgentDevices"`

    /* 总在线设备数 (Optional) */
    TotalOnlineDevices int `json:"totalOnlineDevices"`

    /* 在线直接设备数 (Optional) */
    OnlineDirectConnectDevices int `json:"onlineDirectConnectDevices"`

    /* 在线终端数 (Optional) */
    OnlineConnectAgentDevices int `json:"onlineConnectAgentDevices"`

    /* 在线非直连设备数 (Optional) */
    OnlineIndirectConnectDevices int `json:"onlineIndirectConnectDevices"`
}
