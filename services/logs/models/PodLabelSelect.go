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


type PodLabelSelect struct {

    /* 操作类型 (Optional) */
    Operate string `json:"operate"`

    /* pod详情 (Optional) */
    PodInfo string `json:"podInfo"`

    /* 租户 (Optional) */
    Tenant string `json:"tenant"`

    /* pod详情数组 (Optional) */
    PodListStr string `json:"podListStr"`

    /* 产品名 (Optional) */
    ProductName string `json:"productName"`

    /* 是否所有租户可见 (Optional) */
    AllTenantDistributability bool `json:"allTenantDistributability"`
}
