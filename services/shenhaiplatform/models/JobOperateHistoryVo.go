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


type JobOperateHistoryVo struct {

    /* id (Optional) */
    Id int `json:"id"`

    /* 作业id (Optional) */
    JobId int `json:"jobId"`

    /* 操作类型 (Optional) */
    OperateType string `json:"operateType"`

    /* 操作人 (Optional) */
    Operator string `json:"operator"`

    /* 操作时间 (Optional) */
    OperateTime string `json:"operateTime"`

    /* 表id (Optional) */
    TableId int `json:"tableId"`

    /* 备注 (Optional) */
    Memo string `json:"memo"`
}
