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


type DescribeCabinet struct {

    /* 机房英文标识 (Optional) */
    Idc string `json:"idc"`

    /* 机房名称 (Optional) */
    IdcName string `json:"idcName"`

    /* 机柜Id (Optional) */
    CabinetId string `json:"cabinetId"`

    /* 机柜编码 (Optional) */
    CabinetNo string `json:"cabinetNo"`

    /* 房间号 (Optional) */
    RoomNo string `json:"roomNo"`

    /* 机柜空间(U) (Optional) */
    CabinetSpace int `json:"cabinetSpace"`

    /* 额定电流(A) (Optional) */
    CabinetPower int `json:"cabinetPower"`

    /* 机柜类型 formal:正式机柜 reserved:预留机柜 (Optional) */
    CabinetType string `json:"cabinetType"`

    /* 机柜开通状态 disabled:未开通 enabling:开通中 enabled:已开通 disabling:关闭中 (Optional) */
    CabinetOpenStatus string `json:"cabinetOpenStatus"`

    /* 开通时间，遵循ISO8601标准，使用UTC时间，格式为：YYYY-MM-DDTHH:mm:ssZ (Optional) */
    CabinetOpenTime string `json:"cabinetOpenTime"`

    /* 起租时间，遵循ISO8601标准，使用UTC时间，格式为：YYYY-MM-DDTHH:mm:ssZ (Optional) */
    StartTime string `json:"startTime"`

    /* 退租时间，遵循ISO8601标准，使用UTC时间，格式为：YYYY-MM-DDTHH:mm:ssZ (Optional) */
    EndTime string `json:"endTime"`
}
