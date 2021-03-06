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


type QueryRemoteQuotaResVo struct {

    /* 用户pin (Optional) */
    Pin string `json:"pin"`

    /* 产品线 (Optional) */
    ServiceCode string `json:"serviceCode"`

    /* 产品线名称 (Optional) */
    ServiceName string `json:"serviceName"`

    /* 地域 (Optional) */
    Region string `json:"region"`

    /* 地域名称 (Optional) */
    RegionName string `json:"regionName"`

    /* 用户该地域该资源下的配额值 (Optional) */
    UserQuota int `json:"userQuota"`

    /* 可用配额 (Optional) */
    AvailableQuota int `json:"availableQuota"`

    /* 已用配额 (Optional) */
    UsedQuota int `json:"usedQuota"`

    /* 父资源id (Optional) */
    ParentResourceId string `json:"parentResourceId"`

    /* 配额超出标识,[已使用配额，大于等于配额总量，false]，[已使用配额，小于配额总量，返回true] (Optional) */
    QuotaBeyondFlag bool `json:"quotaBeyondFlag"`

    /* 失败原因 (Optional) */
    FailReason string `json:"failReason"`
}
