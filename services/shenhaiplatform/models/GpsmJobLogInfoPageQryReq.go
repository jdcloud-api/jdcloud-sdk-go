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


type GpsmJobLogInfoPageQryReq struct {

    /* 页面大小  */
    PageSize int `json:"pageSize"`

    /* 页码  */
    PageNum int `json:"pageNum"`

    /* 作业名称 (Optional) */
    JobName string `json:"jobName"`

    /* job 运行session id (Optional) */
    JobSessionId string `json:"jobSessionId"`

    /* job运行日期 (Optional) */
    TxDate string `json:"txDate"`
}
