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


type AssumeRoleInfo struct {

    /* 角色资源标识(jrn)  */
    RoleJrn string `json:"roleJrn"`

    /* 角色会话名称  */
    RoleSessionName string `json:"roleSessionName"`

    /* 会话策略，策略描述需遵循 IAM 策略语法，但不可包含 Principal元素，长度限制 2048 字节。会话策略限制临时凭证的权限；如不指定，则临时凭证默认拥有附加到角色的所有权限。 (Optional) */
    Policy *string `json:"policy"`

    /* 临时凭证有效期，单位秒，取值范围：3600~您所扮演的角色设置的maxSessionDuration，默认3600 (Optional) */
    DurationSeconds *int `json:"durationSeconds"`
}
