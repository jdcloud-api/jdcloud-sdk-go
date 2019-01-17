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


type ImportDNS struct {

    /* 解析记录对应的域名的ID  */
    DomainId int `json:"domainId"`

    /* 主机记录  */
    HostRecord string `json:"hostRecord"`

    /* 解析记录的值  */
    HostValue string `json:"hostValue"`

    /* 解析记录的ID (Optional) */
    Id int `json:"id"`

    /* 是否是京东云资源 (Optional) */
    JcloudRes bool `json:"jcloudRes"`

    /* 优先级，只存在于MX, SRV解析记录类型 (Optional) */
    MxPriority int `json:"mxPriority"`

    /* 端口，只存在于SRV解析记录类型 (Optional) */
    Port int `json:"port"`

    /* 解析记录的生存时间  */
    Ttl int `json:"ttl"`

    /* 解析的类型  */
    Type string `json:"type"`

    /* 解析记录的权重 (Optional) */
    Weight int `json:"weight"`

    /* 解析线路的ID，请调用getViewTree接口获取解析线路的ID。  */
    ViewValue int `json:"viewValue"`
}
