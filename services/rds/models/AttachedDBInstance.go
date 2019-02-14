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


type AttachedDBInstance struct {

    /* 实例ID (Optional) */
    InstanceId string `json:"instanceId"`

    /* 实例名称，具体规则可参见帮助中心文档:[名称及密码限制](../../../documentation/Cloud-Database-and-Cache/RDS/Introduction/Restrictions/SQLServer-Restrictions.md) (Optional) */
    InstanceName string `json:"instanceName"`

    /* 实例类型，例如主实例，只读实例等，参见[枚举参数定义](../Enum-Definitions/Enum-Definitions.md) (Optional) */
    InstanceType string `json:"instanceType"`

    /* 实例引擎类型，如MySQL或SQL Server等，参见[枚举参数定义](../Enum-Definitions/Enum-Definitions.md) (Optional) */
    Engine string `json:"engine"`

    /* 实例引擎版本，参见[枚举参数定义](../Enum-Definitions/Enum-Definitions.md) (Optional) */
    EngineVersion string `json:"engineVersion"`

    /* 参数组ID (Optional) */
    ParameterGroupId string `json:"parameterGroupId"`

    /* 参数的状态，参见[枚举参数定义](../Enum-Definitions/Enum-Definitions.md) (Optional) */
    ParameterStatus string `json:"parameterStatus"`

    /* 实例状态，参见[枚举参数定义](../Enum-Definitions/Enum-Definitions.md) (Optional) */
    InstanceStatus string `json:"instanceStatus"`

    /* 实例创建时间 (Optional) */
    CreateTime string `json:"createTime"`
}
