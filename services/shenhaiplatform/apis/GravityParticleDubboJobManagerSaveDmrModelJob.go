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

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
)

type GravityParticleDubboJobManagerSaveDmrModelJobRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 主键Id,为空时保存，不为空时更新 (Optional) */
    Id *int `json:"id"`

    /* 作业Id (Optional) */
    JobId *int `json:"jobId"`

    /* 模型Id (Optional) */
    ModelTableId *int `json:"modelTableId"`

    /* 删除标识 (Optional) */
    DeleteFlag *int `json:"deleteFlag"`

    /* 更新时间 (Optional) */
    UpdateTime *string `json:"updateTime"`

    /* 表名 (Optional) */
    TableName *string `json:"tableName"`

    /* 表中文名称 (Optional) */
    TableNameCh *string `json:"tableNameCh"`

    /* 数据库名称 (Optional) */
    DatabaseName *string `json:"databaseName"`

    /* 工作空间code (Optional) */
    WorkspaceCode *string `json:"workspaceCode"`

    /* 数据来源渠道 (Optional) */
    Channel *string `json:"channel"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGravityParticleDubboJobManagerSaveDmrModelJobRequest(
    regionId string,
    appName string,
) *GravityParticleDubboJobManagerSaveDmrModelJobRequest {

	return &GravityParticleDubboJobManagerSaveDmrModelJobRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/gravityParticleDubboJobManagerSaveDmrModelJob",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        AppName: appName,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param id: 主键Id,为空时保存，不为空时更新 (Optional)
 * param jobId: 作业Id (Optional)
 * param modelTableId: 模型Id (Optional)
 * param deleteFlag: 删除标识 (Optional)
 * param updateTime: 更新时间 (Optional)
 * param tableName: 表名 (Optional)
 * param tableNameCh: 表中文名称 (Optional)
 * param databaseName: 数据库名称 (Optional)
 * param workspaceCode: 工作空间code (Optional)
 * param channel: 数据来源渠道 (Optional)
 */
func NewGravityParticleDubboJobManagerSaveDmrModelJobRequestWithAllParams(
    regionId string,
    appName string,
    id *int,
    jobId *int,
    modelTableId *int,
    deleteFlag *int,
    updateTime *string,
    tableName *string,
    tableNameCh *string,
    databaseName *string,
    workspaceCode *string,
    channel *string,
) *GravityParticleDubboJobManagerSaveDmrModelJobRequest {

    return &GravityParticleDubboJobManagerSaveDmrModelJobRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/gravityParticleDubboJobManagerSaveDmrModelJob",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
        Id: id,
        JobId: jobId,
        ModelTableId: modelTableId,
        DeleteFlag: deleteFlag,
        UpdateTime: updateTime,
        TableName: tableName,
        TableNameCh: tableNameCh,
        DatabaseName: databaseName,
        WorkspaceCode: workspaceCode,
        Channel: channel,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGravityParticleDubboJobManagerSaveDmrModelJobRequestWithoutParam() *GravityParticleDubboJobManagerSaveDmrModelJobRequest {

    return &GravityParticleDubboJobManagerSaveDmrModelJobRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/gravityParticleDubboJobManagerSaveDmrModelJob",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *GravityParticleDubboJobManagerSaveDmrModelJobRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *GravityParticleDubboJobManagerSaveDmrModelJobRequest) SetAppName(appName string) {
    r.AppName = appName
}
/* param id: 主键Id,为空时保存，不为空时更新(Optional) */
func (r *GravityParticleDubboJobManagerSaveDmrModelJobRequest) SetId(id int) {
    r.Id = &id
}
/* param jobId: 作业Id(Optional) */
func (r *GravityParticleDubboJobManagerSaveDmrModelJobRequest) SetJobId(jobId int) {
    r.JobId = &jobId
}
/* param modelTableId: 模型Id(Optional) */
func (r *GravityParticleDubboJobManagerSaveDmrModelJobRequest) SetModelTableId(modelTableId int) {
    r.ModelTableId = &modelTableId
}
/* param deleteFlag: 删除标识(Optional) */
func (r *GravityParticleDubboJobManagerSaveDmrModelJobRequest) SetDeleteFlag(deleteFlag int) {
    r.DeleteFlag = &deleteFlag
}
/* param updateTime: 更新时间(Optional) */
func (r *GravityParticleDubboJobManagerSaveDmrModelJobRequest) SetUpdateTime(updateTime string) {
    r.UpdateTime = &updateTime
}
/* param tableName: 表名(Optional) */
func (r *GravityParticleDubboJobManagerSaveDmrModelJobRequest) SetTableName(tableName string) {
    r.TableName = &tableName
}
/* param tableNameCh: 表中文名称(Optional) */
func (r *GravityParticleDubboJobManagerSaveDmrModelJobRequest) SetTableNameCh(tableNameCh string) {
    r.TableNameCh = &tableNameCh
}
/* param databaseName: 数据库名称(Optional) */
func (r *GravityParticleDubboJobManagerSaveDmrModelJobRequest) SetDatabaseName(databaseName string) {
    r.DatabaseName = &databaseName
}
/* param workspaceCode: 工作空间code(Optional) */
func (r *GravityParticleDubboJobManagerSaveDmrModelJobRequest) SetWorkspaceCode(workspaceCode string) {
    r.WorkspaceCode = &workspaceCode
}
/* param channel: 数据来源渠道(Optional) */
func (r *GravityParticleDubboJobManagerSaveDmrModelJobRequest) SetChannel(channel string) {
    r.Channel = &channel
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GravityParticleDubboJobManagerSaveDmrModelJobRequest) GetRegionId() string {
    return r.RegionId
}

type GravityParticleDubboJobManagerSaveDmrModelJobResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GravityParticleDubboJobManagerSaveDmrModelJobResult `json:"result"`
}

type GravityParticleDubboJobManagerSaveDmrModelJobResult struct {
    Success int `json:"success"`
    Code string `json:"code"`
    Msg string `json:"msg"`
    Result int `json:"result"`
}