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


type InstanceVoucherSpec struct {

    /* 实例抵扣券名称，支持中文、数字、大小写字母及英文下划线“_”及中划线“-”，不超过32个字符  */
    Name string `json:"name"`

    /* 描述，不超过256个字符 (Optional) */
    Description *string `json:"description"`

    /* 产品类型 支持[vm, nativecontainer, pod]  */
    ResourceType string `json:"resourceType"`

    /* 资源分配方式 支持[nonReserved]  */
    ReservedType string `json:"reservedType"`

    /* 非资源预留型实例抵扣券参数，reservedType 为 nonReserved 时生效 (Optional) */
    NonReservedVoucherSpec *NonReservedVoucherSpec `json:"nonReservedVoucherSpec"`

    /* 计费模式 (Optional) */
    Charge *ChargeSpec `json:"charge"`
}
