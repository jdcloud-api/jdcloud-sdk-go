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


type SecurityPolicy struct {

    /* 安全策略的Id (Optional) */
    SecurityPolicyId string `json:"securityPolicyId"`

    /* 安全策略的名称 (Optional) */
    SecurityPolicyName string `json:"securityPolicyName"`

    /* 安全策略类型，SYSTEM：系统安全策略，CUSTOM：自定义安全策略 (Optional) */
    SecurityPolicyType string `json:"securityPolicyType"`

    /* 关联的监听器ID列表，目前安全策略只支持应用型负载均衡的HTTPS和TLS监听器 (Optional) */
    ListenerIds []string `json:"listenerIds"`

    /* 安全策略的描述信息 (Optional) */
    Description string `json:"description"`

    /* TLS协议版本列表，目前支持TLSv1、TLSv1.1、TLSv1.2和TLSv1.3，传入的每个protocol至少需要传入一个支持的cipher (Optional) */
    Protocols []string `json:"protocols"`

    /* TLS加密套件列表，传入的每个cipher至少需要传入一个能够支持的protocol
TLSv1和TLSv1.1支持的加密套件：
AES128-SHA
AES256-SHA
CAMELLIA128-SHA
CAMELLIA256-SHA
DES-CBC3-SHA
ECDHE-RSA-AES128-SHA
ECDHE-RSA-AES256-SHA
ECDHE-RSA-DES-CBC3-SHA
IDEA-CBC-SHA
SEED-SHA
ECDHE-ECDSA-AES256-SHA
ECDHE-ECDSA-AES128-SHA
ECDHE-ECDSA-DES-CBC3-SHA
TLSv1.2支持的加密套件：
AES128-CCM
AES128-CCM8
AES128-GCM-SHA256
AES128-SHA
AES128-SHA256
AES256-CCM
AES256-CCM8
AES256-GCM-SHA384
AES256-SHA
AES256-SHA256
ARIA128-GCM-SHA256
ARIA256-GCM-SHA384
CAMELLIA128-SHA
CAMELLIA128-SHA256
CAMELLIA256-SHA
CAMELLIA256-SHA256
DES-CBC3-SHA
ECDHE-ARIA128-GCM-SHA256
ECDHE-ARIA256-GCM-SHA384
ECDHE-RSA-AES128-GCM-SHA256
ECDHE-RSA-AES128-SHA
ECDHE-RSA-AES128-SHA256
ECDHE-RSA-AES256-GCM-SHA384
ECDHE-RSA-AES256-SHA
ECDHE-RSA-AES256-SHA384
ECDHE-RSA-CAMELLIA128-SHA256
ECDHE-RSA-CAMELLIA256-SHA384
ECDHE-RSA-CHACHA20-POLY1305
ECDHE-RSA-DES-CBC3-SHA
SEED-SHA
ECDHE-ECDSA-AES256-GCM-SHA384
ECDHE-ECDSA-CHACHA20-POLY1305
ECDHE-ECDSA-AES256-CCM8
ECDHE-ECDSA-AES256-CCM
ECDHE-ECDSA-ARIA256-GCM-SHA384
ECDHE-ECDSA-AES128-GCM-SHA256
ECDHE-ECDSA-AES128-CCM8
ECDHE-ECDSA-AES128-CCM
ECDHE-ECDSA-ARIA128-GCM-SHA256
ECDHE-ECDSA-AES256-SHA384
ECDHE-ECDSA-CAMELLIA256-SHA384
ECDHE-ECDSA-AES128-SHA256
ECDHE-ECDSA-CAMELLIA128-SHA256
ECDHE-ECDSA-AES256-SHA
ECDHE-ECDSA-AES128-SHA
ECDHE-ECDSA-DES-CBC3-SHA
TLSv1.3支持的加密套件：
TLS_AES_256_GCM_SHA384
TLS_CHACHA20_POLY1305_SHA256
TLS_AES_128_GCM_SHA256
 (Optional) */
    Ciphers []string `json:"ciphers"`

    /* 安全策略的创建时间 (Optional) */
    CreatedTime string `json:"createdTime"`
}