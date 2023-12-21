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


type OCRInfo struct {

    /* 识别状态 (Optional) */
    Status string `json:"status"`

    /* 姓名 (Optional) */
    Name string `json:"name"`

    /* 民族 (Optional) */
    Nation string `json:"nation"`

    /* 住址 (Optional) */
    Address string `json:"address"`

    /* 身份证号 (Optional) */
    IdNumber string `json:"idNumber"`

    /* 出生日期 (Optional) */
    Birthday string `json:"birthday"`

    /* 性别 (Optional) */
    Gender string `json:"gender"`

    /* 签发机构 (Optional) */
    Authority string `json:"authority"`

    /* 签发时间 (Optional) */
    IssueTime string `json:"issueTime"`

    /* 到期时间 (Optional) */
    DueTime string `json:"dueTime"`

    /* 1收费，0不收费 (Optional) */
    ChargeFlag string `json:"chargeFlag"`
}
