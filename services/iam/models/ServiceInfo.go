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


type ServiceInfo struct {

    /* 产品线编码 (Optional) */
    ServiceCode string `json:"serviceCode"`

    /* 中文名称 (Optional) */
    CnName string `json:"cnName"`

    /* 英文名称 (Optional) */
    EnName string `json:"enName"`

    /* 可用区信息 (Optional) */
    Region string `json:"region"`

    /* 接入时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 操作人erp (Optional) */
    CreateUser string `json:"createUser"`

    /* 是否可删除 (Optional) */
    CanDelete bool `json:"canDelete"`
}
