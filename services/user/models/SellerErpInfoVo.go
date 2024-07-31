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


type SellerErpInfoVo struct {

    /* 部门 (Optional) */
    Dept string `json:"dept"`

    /* 邮箱 (Optional) */
    Email string `json:"email"`

    /* 创建时间 (Optional) */
    Created int `json:"created"`

    /* erp (Optional) */
    Erp string `json:"erp"`

    /* 手机 (Optional) */
    Mobile string `json:"mobile"`

    /* 修改时间 (Optional) */
    Modified int `json:"modified"`

    /* 验证方式 (Optional) */
    NoticeType string `json:"noticeType"`

    /* 个人邮箱 (Optional) */
    PersonalEmail string `json:"personalEmail"`

    /* 推荐码 (Optional) */
    RecommendedCode string `json:"recommendedCode"`

    /* 销售分组 (Optional) */
    SellerGroup int `json:"sellerGroup"`

    /* 销售名 (Optional) */
    SellerName string `json:"sellerName"`
}
