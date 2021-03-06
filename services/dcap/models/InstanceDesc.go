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


type InstanceDesc struct {

    /* 实例ID (Optional) */
    InstanceId string `json:"instanceId"`

    /* 实例名称 (Optional) */
    InstanceName string `json:"instanceName"`

    /* 实例类型-规格 1->网关版本, 2->插件版本 (Optional) */
    InstanceType int `json:"instanceType"`

    /* 区域 (Optional) */
    InsRegion string `json:"insRegion"`

    /* 可用区 (Optional) */
    InsZone string `json:"insZone"`

    /* 是否为透传模式 (Optional) */
    IsBypass bool `json:"isBypass"`

    /* 实例运行状态: 1->创建中, 2->运行中, 3->欠费停服 (Optional) */
    Status int `json:"status"`

    /* 保护的数据源ID (Optional) */
    DataSourceId string `json:"dataSourceId"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 更新时间 (Optional) */
    UpdateTime string `json:"updateTime"`

    /* 域名（网关版本） (Optional) */
    Domain string `json:"domain"`

    /* AccessKey (Optional) */
    AccessKey string `json:"accessKey"`

    /* SecretKey (Optional) */
    SecretKey string `json:"secretKey"`
}
