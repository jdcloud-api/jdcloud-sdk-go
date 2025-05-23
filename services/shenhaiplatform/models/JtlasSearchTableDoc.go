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


type JtlasSearchTableDoc struct {

    /* 表id (Optional) */
    Id string `json:"id"`

    /* 项目id (Optional) */
    ProjectCode string `json:"projectCode"`

    /* 数据库名称 (Optional) */
    Database string `json:"database"`

    /* 表名称 (Optional) */
    TableName string `json:"tableName"`

    /* 表别名 (Optional) */
    TableAlias string `json:"tableAlias"`

    /* 表类型：MANAGED_TABLE，EXTERNAL_TABLE (Optional) */
    TableType string `json:"tableType"`

    /* 创建人 (Optional) */
    Creator string `json:"creator"`

    /* 元数据类型：HIVE，MYSQL (Optional) */
    MetaType string `json:"metaType"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 最后修改时间 (Optional) */
    LastModifyTime string `json:"lastModifyTime"`

    /* 表存储路径 (Optional) */
    Location string `json:"location"`

    /* 输入格式 (Optional) */
    InputFormat string `json:"inputFormat"`

    /* 输出格式 (Optional) */
    OutputFormat string `json:"outputFormat"`

    /* 用户打标 (Optional) */
    UserDefineTags []string `json:"userDefineTags"`

    /* 负责人 (Optional) */
    CollectPersons []string `json:"collectPersons"`

    /* 是否收藏 (Optional) */
    HasCollected bool `json:"hasCollected"`

    /* 点击数 (Optional) */
    ClickCount int `json:"clickCount"`

    /* 是否热门文章 (Optional) */
    FireFlag bool `json:"fireFlag"`

    /* 环境信息dev或prod (Optional) */
    Env string `json:"env"`

    /* 分桶字段名称 (Optional) */
    BucketCols []string `json:"bucketCols"`

    /* 分桶个数 (Optional) */
    BucketNum string `json:"bucketNum"`
}
