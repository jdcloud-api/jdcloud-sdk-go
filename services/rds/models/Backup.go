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


type Backup struct {

    /* 备份ID (Optional) */
    BackupId string `json:"backupId"`

    /* 备份名称 (Optional) */
    BackupName string `json:"backupName"`

    /* 备份所属实例ID (Optional) */
    InstanceId string `json:"instanceId"`

    /* 备份状态，COMPLETED：备份完成，FAILED：备份失败，BUILDING：备份中，DELETING：删除中 (Optional) */
    BackupStatus string `json:"backupStatus"`

    /* 备份开始时间，遵循ISO8601标准，使用UTC时间，格式为：YYYY-MM-DDTHH:mm:ssZ (Optional) */
    BackupStartTime string `json:"backupStartTime"`

    /* 备份结束时间，遵循ISO8601标准，使用UTC时间，格式为：YYYY-MM-DDTHH:mm:ssZ (Optional) */
    BackupEndTime string `json:"backupEndTime"`

    /* 备份类型，全量备份或增量备份，full：全量，diff：增量 (Optional) */
    BackupType string `json:"backupType"`

    /* 备份模式，系统自动备份或手动备份，Automated：自动备份  Manual：手工备份 (Optional) */
    BackupMode string `json:"backupMode"`

    /* 备份粒度，实例备份或者多库备份，instance：实例备份 ，dbs：数据库备份 (Optional) */
    BackupUnit string `json:"backupUnit"`

    /* 备份文件列表，仅SQL Server支持该参数，文件名的命名规则为:全备:数据库名+.bak; 增量:数据库名+.diff (Optional) */
    BackupFiles []string `json:"backupFiles"`

    /* 整个备份集大小，单位：Byte (Optional) */
    BackupSizeByte int `json:"backupSizeByte"`
}
