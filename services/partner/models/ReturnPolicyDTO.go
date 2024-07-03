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


type ReturnPolicyDTO struct {

    /* ID (Optional) */
    Id int `json:"id"`

    /* 部门ID  */
    DeptId int `json:"deptId"`

    /* 部门名称 (Optional) */
    DeptName string `json:"deptName"`

    /* 渠道商类型  */
    DistributorType int `json:"distributorType"`

    /* 返还类型 (Optional) */
    ReturnType int `json:"returnType"`

    /* 返还政策主ID  */
    ReturnPolicyId int `json:"returnPolicyId"`

    /* 返还政策主名称 (Optional) */
    ReturnPolicyName string `json:"returnPolicyName"`

    /* 项目编码  */
    ItemId int `json:"itemId"`

    /* 项目名称 (Optional) */
    ItemName string `json:"itemName"`

    /* 返还依据类型  */
    ReturnRuleType int `json:"returnRuleType"`

    /* 审批邮件附件id，取上传邮件附件接口返回的approveFileId值  */
    ApproveFileId string `json:"approveFileId"`

    /*  (Optional) */
    ReturnPolicyProductDTOList []ReturnPolicyProductDTO `json:"returnPolicyProductDTOList"`

    /* 选择的产品列表
例："returnPolicyProductMap": {"1": ["500138","500139"],"2": ["719","720"]}
 (Optional) */
    ReturnPolicyProductMap interface{} `json:"returnPolicyProductMap"`

    /* 周期类型  */
    CircleType int `json:"circleType"`

    /* 周期名称 (Optional) */
    CircleName string `json:"circleName"`

    /* 指定周期标识 (Optional) */
    CircleFlag int `json:"circleFlag"`

    /* 周期值 (Optional) */
    CircleValue int `json:"circleValue"`

    /*  (Optional) */
    Conditions []ReturnConditionDTO `json:"conditions"`

    /* 状态 (Optional) */
    Status int `json:"status"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 创建人 (Optional) */
    CreateUser string `json:"createUser"`

    /* 修改时间 (Optional) */
    UpdateTime string `json:"updateTime"`

    /* 修改人 (Optional) */
    UpdateUser string `json:"updateUser"`

    /* 是否删除0未删除,1已删除 (Optional) */
    Yn int `json:"yn"`
}
