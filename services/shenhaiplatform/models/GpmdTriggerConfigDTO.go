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


type GpmdTriggerConfigDTO struct {

    /* 条件类型 (Optional) */
    ConditionType string `json:"conditionType"`

    /* 次要类型 (Optional) */
    SecondaryType string `json:"secondaryType"`

    /* 服务器域名 (Optional) */
    ServerDomain string `json:"serverDomain"`

    /* 服务器端口 (Optional) */
    ServerPort string `json:"serverPort"`

    /* 用户名 (Optional) */
    Username string `json:"username"`

    /* 用户密码 (Optional) */
    UsernamePwd string `json:"usernamePwd"`

    /* 协议 (Optional) */
    Protocol string `json:"protocol"`

    /* 数据库名称 (Optional) */
    DatabaseName string `json:"databaseName"`

    /* 数据源 (Optional) */
    DataSource string `json:"dataSource"`

    /* 对象信息 (Optional) */
    Object string `json:"object"`

    /* 条件配置 (Optional) */
    ConditionConfig string `json:"conditionConfig"`

    /* 触发时间 (Optional) */
    TriggerTime string `json:"triggerTime"`

    /* 触发间隔 (Optional) */
    TriggerInterval int `json:"triggerInterval"`

    /* 开始时间 (Optional) */
    BeginTime string `json:"beginTime"`

    /* 结束时间 (Optional) */
    EndTime string `json:"endTime"`
}
