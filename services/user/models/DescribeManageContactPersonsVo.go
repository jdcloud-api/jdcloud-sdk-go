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


type DescribeManageContactPersonsVo struct {

    /* pin (Optional) */
    Pin string `json:"pin"`

    /* 个人姓名 (Optional) */
    UserName string `json:"userName"`

    /* 所在联系人组 (Optional) */
    ContactGroups string `json:"contactGroups"`

    /* 手机 (Optional) */
    Mobile string `json:"mobile"`

    /* 邮箱 (Optional) */
    Email string `json:"email"`
}
