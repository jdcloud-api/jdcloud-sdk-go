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


type GpdjmcDmrModelJobRDTO struct {

    /* 主键Id,为空时保存，不为空时更新 (Optional) */
    Id int `json:"id"`

    /* 作业Id (Optional) */
    JobId int `json:"jobId"`

    /* 模型Id (Optional) */
    ModelTableId int `json:"modelTableId"`

    /* 删除标识 (Optional) */
    DeleteFlag int `json:"deleteFlag"`

    /* 更新时间 (Optional) */
    UpdateTime string `json:"updateTime"`

    /* 表名 (Optional) */
    TableName string `json:"tableName"`

    /* 表中文名称 (Optional) */
    TableNameCh string `json:"tableNameCh"`

    /* 数据库名称 (Optional) */
    DatabaseName string `json:"databaseName"`

    /* 工作空间code (Optional) */
    WorkspaceCode string `json:"workspaceCode"`

    /* 数据来源渠道 (Optional) */
    Channel string `json:"channel"`
}
