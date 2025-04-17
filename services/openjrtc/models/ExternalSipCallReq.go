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


type ExternalSipCallReq struct {

    /* 目标sip域 (Optional) */
    DestDomain string `json:"destDomain"`

    /* 要呼叫的目标号码（会议号或者sip号） (Optional) */
    DestNumber string `json:"destNumber"`

    /* 京东会议所在的sip域 (Optional) */
    RoomDomain string `json:"roomDomain"`

    /* 业务接入方定义的且在JRTC系统内注册过的房间号 (Optional) */
    UserRoomId string `json:"userRoomId"`

    /* sip上显示的京东会议的名称 (Optional) */
    RoomDisplay string `json:"roomDisplay"`

    /* 呼叫的类型, 0：默认值，呼叫本地注册;1：通过sip网关呼叫;2：通过网关呼叫电话 (Optional) */
    DestType string `json:"destType"`

    /* 网关名称, 当destType为1或者2时，需要传该参数 (Optional) */
    Gateway string `json:"gateway"`

    /* 业务接入方用户体系的userId注册为jrtc系统内可识别和流转的用户id,当destType为2时,需要传该参数 (Optional) */
    UserId string `json:"userId"`

    /* 呼叫保持 (Optional) */
    Holdon bool `json:"holdon"`
}
