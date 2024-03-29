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


type DataDiskSpec struct {

    /* 云硬盘类型，取值为ssd、premium-hdd、ssd.gp1、ssd.io1、hdd.std1之一  */
    DiskType string `json:"diskType"`

    /* 云硬盘大小，单位为 GiB，ssd 类型取值范围[20,1000]GB，步长为10G，premium-hdd 类型取值范围[20,3000]GB，步长为10G, ssd.gp1, ssd.io1, hdd.std1 类型取值均是范围[20,16000]GB，步长为10G  */
    DiskSizeGB int `json:"diskSizeGB"`

    /* 云硬盘IOPS的大小，当且仅当云盘类型是ssd.io1型的云盘有效，步长是10. (Optional) */
    Iops *int `json:"iops"`

    /* 用于创建云硬盘的快照ID (Optional) */
    SnapshotId *string `json:"snapshotId"`

    /* 随虚机删除，默认true (Optional) */
    AutoDelete *bool `json:"autoDelete"`

    /* 自动format，默认true (Optional) */
    AutoFormat *bool `json:"autoFormat"`

    /* 支持ext4,xfs  */
    FsType string `json:"fsType"`

    /* 挂载点，默认 /var/lib/docker (Optional) */
    MountPoint *string `json:"mountPoint"`
}
