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


type DomainInfo struct {

    /* 域名 (Optional) */
    DomainName string `json:"domainName"`

    /* 域名状态 0失败 1正常 2预注册 3过期 4转出中 5已转出 6过户中 (Optional) */
    DomainState int `json:"domainState"`

    /* 模板ID (Optional) */
    TemplateId int64 `json:"templateId"`

    /* 域名起始时间 (Optional) */
    StartDate string `json:"startDate"`

    /* 域名结束时间 (Optional) */
    EndDate string `json:"endDate"`

    /* 京东云解析域名ID (Optional) */
    CloudDnsId string `json:"cloudDnsId"`

    /* 最后变更时间 (Optional) */
    Modified string `json:"modified"`
}
