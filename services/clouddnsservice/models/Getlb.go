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


type Getlb struct {

    /* 负载均衡的解析记录的列表中解析记录是否是相同的权重<br>
true: 按权重分配负载<br>
false: 均等负载
 (Optional) */
    IsBalance bool `json:"isBalance"`

    /* 主机记录 (Optional) */
    Record string `json:"record"`

    /* 解析的类型 (Optional) */
    Type string `json:"type"`

    /* 解析线路的名称 (Optional) */
    ViewName string `json:"viewName"`

    /* 解析线路的ID (Optional) */
    ViewValue int `json:"viewValue"`

    /* 负载均衡的解析记录的列表 (Optional) */
    Items []HostRRlb `json:"items"`
}
