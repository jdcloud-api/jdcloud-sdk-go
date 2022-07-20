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

import charge "github.com/lshuining/jdcloud-sdk-go/services/charge/models"

type Disk struct {

    /* 云硬盘ID (Optional) */
    DiskId string `json:"diskId"`

    /* 云硬盘所属AZ (Optional) */
    Az string `json:"az"`

    /* 云硬盘名称，只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符。 (Optional) */
    Name string `json:"name"`

    /* 云硬盘描述，允许输入UTF-8编码下的全部字符，不超过256字符。 (Optional) */
    Description string `json:"description"`

    /* 云硬盘类型，取值为 ssd,premium-hdd,ssd.gp1,ssd.io1,hdd.std1 (Optional) */
    DiskType string `json:"diskType"`

    /* 云硬盘大小，单位为 GiB (Optional) */
    DiskSizeGB int `json:"diskSizeGB"`

    /* 该云硬盘实际应用的iops值 (Optional) */
    Iops int `json:"iops"`

    /* 该云硬盘实际应用的吞吐量的数值 (Optional) */
    Throughput int `json:"throughput"`

    /* 云硬盘状态，取值为 creating、available、in-use、extending、restoring、deleting、deleted、error_create、error_delete、error_restore、error_extend 之一 (Optional) */
    Status string `json:"status"`

    /* 挂载信息 (Optional) */
    Attachments []DiskAttachment `json:"attachments"`

    /* 创建该云硬盘的快照ID (Optional) */
    SnapshotId string `json:"snapshotId"`

    /* 云盘是否支持多挂载 (Optional) */
    MultiAttachable bool `json:"multiAttachable"`

    /* 云盘是否为加密盘 (Optional) */
    Encrypted bool `json:"encrypted"`

    /* 云盘是否被暂停（IOPS限制为极低） (Optional) */
    Enabled bool `json:"enabled"`

    /* 创建云硬盘时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 云硬盘计费配置信息 (Optional) */
    Charge charge.Charge `json:"charge"`

    /* null (Optional) */
    Tags []Tag `json:"tags"`

    /*  (Optional) */
    SnapshotPolicies []SnapshotPolicy `json:"snapshotPolicies"`
}
