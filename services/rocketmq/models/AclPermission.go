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


type AclPermission struct {

    /* 被授权的用户名 (Optional) */
    UserName string `json:"userName"`

    /* 资源类型(Topic 或 Group) (Optional) */
    ResourceType string `json:"resourceType"`

    /* 资源名 (Optional) */
    ResourceName string `json:"resourceName"`

    /* 权限类型(DENY,PUB,SUB,PUB|SUB) (Optional) */
    Permission string `json:"permission"`
}
