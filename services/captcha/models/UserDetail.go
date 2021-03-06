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


type UserDetail struct {

    /* 用户名称 (Optional) */
    Pin string `json:"pin"`

    /* 更新时间, yyyy-mm-dd hh:mm:ss格式 (Optional) */
    UpdateTime string `json:"updateTime"`

    /* 创建时间, yyyy-mm-dd hh:mm:ss格式 (Optional) */
    CreateTime string `json:"createTime"`

    /* 用户类型，0-全部类型，1-付费用户，2-体验用户，3-无效用户 (Optional) */
    UserType string `json:"userType"`

    /* 剩余调用量 (Optional) */
    SurplusAmount int `json:"surplusAmount"`

    /* 累计调用量 (Optional) */
    UsedAmount int `json:"usedAmount"`

    /* 总购买量 (Optional) */
    PackagesAmount int `json:"packagesAmount"`

    /* 资源包数量 (Optional) */
    PackagesNum int `json:"packagesNum"`

    /* 跟踪描述 (Optional) */
    Tracking string `json:"tracking"`
}
