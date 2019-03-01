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


type CertInfo struct {

    /* 通用名称 (Optional) */
    Domain string `json:"domain"`

    /* 证书生效时间 (Optional) */
    From string `json:"from"`

    /* 证书到期时间 (Optional) */
    To string `json:"to"`

    /* 证书组织 (Optional) */
    User string `json:"user"`

    /* 加密算法 (Optional) */
    SigAlgName string `json:"sigAlgName"`

    /* 颁发者 (Optional) */
    Issuer string `json:"issuer"`
}