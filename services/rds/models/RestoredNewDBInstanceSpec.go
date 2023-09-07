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

import charge "github.com/jdcloud-api/jdcloud-sdk-go/services/charge/models"

type RestoredNewDBInstanceSpec struct {

    /* 数据库实例名，名称的限制可参考[帮助中心文档](../../../documentation/Database-and-Cache-Service/RDS/Introduction/Restrictions/SQLServer-Restrictions.md) (Optional) */
    InstanceName *string `json:"instanceName"`

    /* 实例规格代码，可以查看文档[MySQL 实例规格](../Instance-Specifications/Instance-Specifications-MySQL.md)、[SQL Server实例规格](../Instance-Specifications/Instance-Specifications-SQLServer.md)  */
    InstanceClass string `json:"instanceClass"`

    /* 磁盘大小，单位GB  */
    InstanceStorageGB int `json:"instanceStorageGB"`

    /* 可用区ID， 第一个ID必须为主实例所在的可用区。如两个可用区一样，也需输入两个azId  */
    AzId []string `json:"azId"`

    /* VPC的ID  */
    VpcId string `json:"vpcId"`

    /* 子网ID  */
    SubnetId string `json:"subnetId"`

    /* 参数组ID, 缺省系统会创建一个默认参数组<br>- 仅支持MySQL (Optional) */
    ParameterGroup *string `json:"parameterGroup"`

    /* 计费规格，包括计费类型，计费周期等  */
    ChargeSpec *charge.ChargeSpec `json:"chargeSpec"`

    /* 存储类型，参见[枚举参数定义](../Enum-Definitions/Enum-Definitions.md), 缺省值为：LOCAL_SSD<br>- 仅支持MySQL (Optional) */
    InstanceStorageType *string `json:"instanceStorageType"`

    /* 应用访问端口，支持的端口范围：1150～5999。MySQL、Percona、MariaDB的默认值为 3306；PostgreSQL的默认端口号为5432； (Optional) */
    InstancePort *string `json:"instancePort"`

    /* 实例数据加密(存储类型为云硬盘才支持数据加密)。false：不加密，true：加密，缺省为false<br>- 仅支持MySQL (Optional) */
    StorageEncrypted *bool `json:"storageEncrypted"`

    /* 实例的高可用架构。standalone：单机，cluster：主备双机架构，缺省为cluster，multi-replica：三副本<br>- 仅支持SQL Server (Optional) */
    InstanceType *string `json:"instanceType"`

    /* 标签信息 (Optional) */
    TagSpec []Tag `json:"tagSpec"`

    /* 资源组id (Optional) */
    ResourceGroupId *string `json:"resourceGroupId"`
}
