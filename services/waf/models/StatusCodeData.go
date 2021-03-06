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


type StatusCodeData struct {

    /* 异常状态码域名TOP10 (Optional) */
    DomainTop10 []TopCodeValue `json:"domainTop10"`

    /* 异常状态码URL的TOP10 (Optional) */
    UrlTop10 []TopCodeValue `json:"urlTop10"`

    /* 状态码响应统计，当请求字段isStaCode为true时返回 (Optional) */
    StatusCodeTotal ChartItemValue `json:"statusCodeTotal"`

    /* 状态码响应占比，当请求字段isStaCode为true时返回 (Optional) */
    StatusCodePie []TopValue `json:"statusCodePie"`
}
