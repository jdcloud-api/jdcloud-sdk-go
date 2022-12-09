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


type DmsOnlineProxySubTask struct {

    /* 子任务id。 (Optional) */
    SubTaskId int `json:"subTaskId"`

    /* 执行sql。 (Optional) */
    Sql int `json:"sql"`

    /* 执行节点ip。 (Optional) */
    NodeIp string `json:"nodeIp"`

    /* 执行节点端口。 (Optional) */
    NodePort string `json:"nodePort"`

    /* 用户名。 (Optional) */
    UserName string `json:"userName"`

    /* 数据库名。 (Optional) */
    DbName string `json:"dbName"`

    /* 状态。 (Optional) */
    Status string `json:"status"`

    /* sql执行消息。 (Optional) */
    Ext bool `json:"ext"`

    /* 创建时间。 (Optional) */
    CreateTime string `json:"createTime"`

    /* 更新时间。 (Optional) */
    UpdateTime string `json:"updateTime"`
}