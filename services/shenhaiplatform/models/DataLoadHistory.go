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


type DataLoadHistory struct {

    /* 主键ID  */
    Id int `json:"id"`

    /* 文件名称  */
    FileName string `json:"fileName"`

    /* 项目编码  */
    ProjectCode string `json:"projectCode"`

    /* 目标表名  */
    TableName string `json:"tableName"`

    /* 任务id  */
    ApplicationId string `json:"applicationId"`

    /* 任务状态  */
    State string `json:"state"`

    /* 任务状态描述  */
    StateDesc string `json:"stateDesc"`

    /* 创建时间  */
    CreatedDate int `json:"createdDate"`
}
