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


type ModifyIpSetUpVo struct {

    /* 标题 (Optional) */
    Title string `json:"title"`

    /* 标题编码 用户ip白名单为空 (Optional) */
    TitleCode string `json:"titleCode"`

    /* 配置方式-0ip白名单，1固定ip白名单 (Optional) */
    Type int `json:"type"`

    /* pin集合 固定ip白名单入参 (Optional) */
    Pins []string `json:"pins"`

    /* pin集合 ip白名单入参 (Optional) */
    Pin string `json:"pin"`

    /* ip集合 (Optional) */
    Ip []string `json:"ip"`

    /* 是否控制网关sdk：0否，1是 (Optional) */
    IsApi int `json:"isApi"`
}
