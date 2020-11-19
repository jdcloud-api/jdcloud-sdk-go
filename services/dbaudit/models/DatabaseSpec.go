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


type DatabaseSpec struct {

    /* 数据库名称，库名,长度限制32字节,允许英文字母,数字,下划线,中划线和中文 (Optional) */
    DbName *string `json:"dbName"`

    /* 数据库地址, 可以是IP或域名，支持IPv6 (Optional) */
    DbAddr *string `json:"dbAddr"`

    /* 数据库端口 (Optional) */
    DbPort *int `json:"dbPort"`

    /* 数据库类型: 
0->Oracle
1->SQLServer
2->DB2
3->MySQL
4->Cache
5->Sybase
6->DM7
7->Informix
9->人大金仓
10->Teradata
11->Postgresql
12->Gbase
16->Hive
17->MongoDB
 (Optional) */
    DbType *int `json:"dbType"`

    /* 数据库版本 (Optional) */
    DbVersion *string `json:"dbVersion"`

    /* 用户名，SQLServer获取登录信息使用 (Optional) */
    Username *string `json:"username"`

    /* 密码，SQLServer获取登录信息使用 (Optional) */
    Password *string `json:"password"`

    /* 数据库的描述 (Optional) */
    DbDesc *string `json:"dbDesc"`

    /* 是否对数据进行掩码 (Optional) */
    DataMask *bool `json:"dataMask"`

    /* 是否对响应进行审计 (Optional) */
    AuditResponse *bool `json:"auditResponse"`
}