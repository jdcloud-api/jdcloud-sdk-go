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


type Table struct {

    /* 模式名称 (Optional) */
    Schema string `json:"schema"`

    /* 表名称 (Optional) */
    Name string `json:"name"`

    /* 表类型 (Optional) */
    Tp string `json:"tp"`

    /* 表引擎 (Optional) */
    Engine string `json:"engine"`

    /* 表行数 (Optional) */
    TableRows string `json:"tableRows"`

    /* 表数据大小 (Optional) */
    DataLength string `json:"dataLength"`

    /* 表索引大小 (Optional) */
    IndexLength string `json:"indexLength"`

    /* 表字符序 (Optional) */
    Collation string `json:"collation"`

    /* 主键、唯一键 (Optional) */
    UniqueKey []string `json:"uniqueKey"`

    /* 字段 (Optional) */
    Columns []Column `json:"columns"`

    /* 分布式表名称（Clickhouse） (Optional) */
    DistributedTable string `json:"distributedTable"`
}
