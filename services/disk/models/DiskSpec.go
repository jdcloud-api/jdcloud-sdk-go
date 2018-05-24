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

import . "github.com/jdcloud-api/jdcloud-sdk-go/services/charge/models"

type DiskSpec struct {

    /* 云硬盘所属的可用区  */
    Az string `json:"az"`

    /* 云硬盘名称  */
    Name string `json:"name"`

    /* 云硬盘描述 (Optional) */
    Description *string `json:"description"`

    /* 云硬盘类型，取值为ssd、premium-hdd之一  */
    DiskType string `json:"diskType"`

    /* 云硬盘大小，单位为 GiB，ssd 类型取值范围[20,1000]GB，步长为10G，premium-hdd 类型取值范围[20,3000]GB，步长为10G  */
    DiskSizeGB int `json:"diskSizeGB"`

    /* 用于创建云硬盘的快照ID (Optional) */
    SnapshotId *string `json:"snapshotId"`

    /* 计费配置；如不指定，默认计费类型是后付费-按使用时常付费 (Optional) */
    Charge *ChargeSpec `json:"charge"`
}
