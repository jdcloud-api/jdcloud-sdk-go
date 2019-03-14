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


type PublishDomain struct {

    /* 推流域名 (Optional) */
    PublishDomain string `json:"publishDomain"`

    /* 推流域名(Cname) (Optional) */
    PublishDomainCname string `json:"publishDomainCname"`

    /* 直播域名状态：
  - online表示启用
  - offline表示停用
  - configuring表示配置中
  - configure_failed表示配置失败
  - checking表示正在审核
  - check_failed表示审核失败
 (Optional) */
    DomainStatus string `json:"domainStatus"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 更新时间 (Optional) */
    UpdateTime string `json:"updateTime"`
}
