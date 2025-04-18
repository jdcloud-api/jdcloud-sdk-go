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

type GravityParticleDubboJobManagerSearchJobIdRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 数据库名 (Optional) */
    DatabaseName *string `json:"databaseName"`

    /* 表名 (Optional) */
    TableName *string `json:"tableName"`

    /* 工作空间code (Optional) */
    WorkspaceCode *string `json:"workspaceCode"`

    /* 作业名 (Optional) */
    JobName *string `json:"jobName"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGravityParticleDubboJobManagerSearchJobIdRequest(
    regionId string,
    appName string,
) *GravityParticleDubboJobManagerSearchJobIdRequest {

	return &GravityParticleDubboJobManagerSearchJobIdRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/gravityParticleDubboJobManagerSearchJobId",
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
 * param databaseName: 数据库名 (Optional)
 * param tableName: 表名 (Optional)
 * param workspaceCode: 工作空间code (Optional)
 * param jobName: 作业名 (Optional)
 */
func NewGravityParticleDubboJobManagerSearchJobIdRequestWithAllParams(
    regionId string,
    appName string,
    databaseName *string,
    tableName *string,
    workspaceCode *string,
    jobName *string,
) *GravityParticleDubboJobManagerSearchJobIdRequest {

    return &GravityParticleDubboJobManagerSearchJobIdRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/gravityParticleDubboJobManagerSearchJobId",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
        DatabaseName: databaseName,
        TableName: tableName,
        WorkspaceCode: workspaceCode,
        JobName: jobName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGravityParticleDubboJobManagerSearchJobIdRequestWithoutParam() *GravityParticleDubboJobManagerSearchJobIdRequest {

    return &GravityParticleDubboJobManagerSearchJobIdRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/gravityParticleDubboJobManagerSearchJobId",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *GravityParticleDubboJobManagerSearchJobIdRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *GravityParticleDubboJobManagerSearchJobIdRequest) SetAppName(appName string) {
    r.AppName = appName
}
/* param databaseName: 数据库名(Optional) */
func (r *GravityParticleDubboJobManagerSearchJobIdRequest) SetDatabaseName(databaseName string) {
    r.DatabaseName = &databaseName
}
/* param tableName: 表名(Optional) */
func (r *GravityParticleDubboJobManagerSearchJobIdRequest) SetTableName(tableName string) {
    r.TableName = &tableName
}
/* param workspaceCode: 工作空间code(Optional) */
func (r *GravityParticleDubboJobManagerSearchJobIdRequest) SetWorkspaceCode(workspaceCode string) {
    r.WorkspaceCode = &workspaceCode
}
/* param jobName: 作业名(Optional) */
func (r *GravityParticleDubboJobManagerSearchJobIdRequest) SetJobName(jobName string) {
    r.JobName = &jobName
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GravityParticleDubboJobManagerSearchJobIdRequest) GetRegionId() string {
    return r.RegionId
}

type GravityParticleDubboJobManagerSearchJobIdResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GravityParticleDubboJobManagerSearchJobIdResult `json:"result"`
}

type GravityParticleDubboJobManagerSearchJobIdResult struct {
    Success int `json:"success"`
    Code string `json:"code"`
    Msg string `json:"msg"`
    Result []int `json:"result"`
}