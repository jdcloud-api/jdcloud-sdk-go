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


type CredentialInfo struct {

    /* 临时身份凭证AK (Optional) */
    AccessKey string `json:"accessKey"`

    /* 临时身份凭证SK (Optional) */
    SecretKey string `json:"secretKey"`

    /* 临时身份凭证令牌 (Optional) */
    SessionToken string `json:"sessionToken"`

    /* token失效时点（sdk内对token有效期做了校验，接入方不需要关注该字段） (Optional) */
    Expiration string `json:"expiration"`

    /* 角色所属主账号，产品线判断灰度使用（非内测公测产品线不用关注该字段） (Optional) */
    RolePin string `json:"rolePin"`
}
