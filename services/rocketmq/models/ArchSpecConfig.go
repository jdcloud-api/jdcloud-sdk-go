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


type ArchSpecConfig struct {

    /* cpu架构类型(x86/arm64) (Optional) */
    ArchType string `json:"archType"`

    /* 该架构是否对外售卖 (Optional) */
    OnSale bool `json:"onSale"`

    /* 指定架构下包含的节点的可选配置信息 (Optional) */
    NodeSpecConfig []NodeSpecConfig `json:"nodeSpecConfig"`
}
