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


type CollectInfoDetailEnd struct {

    /* UID (Optional) */
    UID string `json:"uID"`

    /* 高可用组资源multi (Optional) */
    AgResource []AgResourceEnd `json:"agResource"`

    /* 日志来源 (Optional) */
    AppCode string `json:"appCode"`

    /* binlog规格 (Optional) */
    BinlogSpec interface{} `json:"binlogSpec"`

    /*  (Optional) */
    Detail CollectTempalteEnd `json:"detail"`

    /*  (Optional) */
    Enabled int64 `json:"enabled"`

    /* k8s规格 (Optional) */
    K8sSpec interface{} `json:"k8sSpec"`

    /* 采集配置名称 (Optional) */
    Name string `json:"name"`

    /* 采集资源时选择的模式，1.正常的选择实例模式（默认模式）；2.选择标签tag模式 3.选择高可用组ag模式 (Optional) */
    ResourceMode int64 `json:"resourceMode"`

    /* 采集实例类型, 只能是 all/part (Optional) */
    ResourceType string `json:"resourceType"`

    /* 采集实例数量 (Optional) */
    ResourcesCount int64 `json:"resourcesCount"`

    /* 产品线 (Optional) */
    ServiceCode string `json:"serviceCode"`

    /*  (Optional) */
    TagResource TagResourceEnd `json:"tagResource"`

    /* 日志类型名称 (Optional) */
    TemplateName string `json:"templateName"`

    /* 日志类型 (Optional) */
    TemplateUID string `json:"templateUID"`
}
