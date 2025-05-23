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


type PodTemplateCloudDisk struct {

    /* 云盘快照id (Optional) */
    SnapshotId string `json:"snapshotId"`

    /* 云盘类型：hdd.std1,ssd.gp1,ssd.io1 (Optional) */
    DiskType string `json:"diskType"`

    /* 云盘size,单位 GB (Optional) */
    SizeGB int `json:"sizeGB"`

    /* 指定volume文件系统类型，目前支持[xfs, ext4]；如果新创建的盘，不指定文件系统类型默认格式化成xfs。 (Optional) */
    FsType string `json:"fsType"`

    /* 云盘的 iops 值，目前只有 ssd.io1 类型有效。 (Optional) */
    Iops int `json:"iops"`

    /* 是否随pod删除。默认：true (Optional) */
    AutoDelete bool `json:"autoDelete"`
}
