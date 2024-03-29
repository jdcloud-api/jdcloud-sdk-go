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


type OrganizationUser struct {

    /* 组织ID (Optional) */
    OrgId string `json:"orgId"`

    /* 组织用户标识，如：erp (Optional) */
    UserId string `json:"userId"`

    /* 组织部门路径名称 (Optional) */
    DepartmentPathName string `json:"departmentPathName"`

    /* 手机号码 (Optional) */
    Phone string `json:"phone"`

    /* 邮箱 (Optional) */
    Email string `json:"email"`

    /* 预校验结果：0-预校验通过; 1-子用户名格式不合法; 2-不允许与主账号同名; 3-BusinessType=1说明子用户名已存在, BusinessType=2说明关联关系已存在 (Optional) */
    ValidateState int `json:"validateState"`
}
