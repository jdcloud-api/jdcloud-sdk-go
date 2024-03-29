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


type DescribeZonesRes struct {

    /* zone id (Optional) */
    ZoneId string `json:"zoneId"`

    /* zone name (Optional) */
    Name string `json:"name"`

    /* zone状态(active->已激活 deleted->已删除 pending->未激活) (Optional) */
    Status string `json:"status"`

    /* 是否暂停 (Optional) */
    Paused bool `json:"paused"`

    /* zone接入类型(full->全接入 partial->半接入) (Optional) */
    Ty_pe string `json:"ty_pe"`

    /* 开发者模式 (Optional) */
    Development_mode int `json:"development_mode"`

    /* 校验key (Optional) */
    Verification_key string `json:"verification_key"`

    /* cname后缀 (Optional) */
    Cname_suffix string `json:"cname_suffix"`

    /* 源ns (Optional) */
    Original_name_servers []string `json:"original_name_servers"`

    /* 源注册商 (Optional) */
    Original_registrar string `json:"original_registrar"`

    /* 源dns host (Optional) */
    Original_dnshost string `json:"original_dnshost"`

    /* 修改时间 (Optional) */
    Modified_on string `json:"modified_on"`

    /* 创建时间 (Optional) */
    Created_on string `json:"created_on"`

    /* 激活时间 (Optional) */
    Activated_on string `json:"activated_on"`

    /* name servers (Optional) */
    Name_servers []string `json:"name_servers"`

    /* 域名对应的ip信息 (Optional) */
    Ips []string `json:"ips"`

    /* ip类型 (Optional) */
    IpType string `json:"ipType"`
}
