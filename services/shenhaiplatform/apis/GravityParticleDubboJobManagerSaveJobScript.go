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

type GravityParticleDubboJobManagerSaveJobScriptRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 主键Id,为空时保存，不为空时更新 (Optional) */
    Id *int `json:"id"`

    /* 脚本名称 (Optional) */
    ScriptName *string `json:"scriptName"`

    /* 脚本类型 (Optional) */
    ScriptType *string `json:"scriptType"`

    /* 脚本描述 (Optional) */
    ScriptDesc *string `json:"scriptDesc"`

    /* 脚本来源， LOCAL(本地) GIT(git) SKYDRIVE(网盘) (Optional) */
    ScriptSourceType *string `json:"scriptSourceType"`

    /* 任务id (Optional) */
    JobId *int `json:"jobId"`

    /* 删除标识 (Optional) */
    DeleteFlag *int `json:"deleteFlag"`

    /* 创建人 (Optional) */
    Creator *string `json:"creator"`

    /* 创建时间 (Optional) */
    CreateTime *string `json:"createTime"`

    /* 修改人 (Optional) */
    UpdateUser *string `json:"updateUser"`

    /* 修改时间 (Optional) */
    UpdateTime *string `json:"updateTime"`

    /* 脚本内容 (Optional) */
    Content *string `json:"content"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGravityParticleDubboJobManagerSaveJobScriptRequest(
    regionId string,
    appName string,
) *GravityParticleDubboJobManagerSaveJobScriptRequest {

	return &GravityParticleDubboJobManagerSaveJobScriptRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/gravityParticleDubboJobManagerSaveJobScript",
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
 * param scriptName: 脚本名称 (Optional)
 * param scriptType: 脚本类型 (Optional)
 * param scriptDesc: 脚本描述 (Optional)
 * param scriptSourceType: 脚本来源， LOCAL(本地) GIT(git) SKYDRIVE(网盘) (Optional)
 * param jobId: 任务id (Optional)
 * param deleteFlag: 删除标识 (Optional)
 * param creator: 创建人 (Optional)
 * param createTime: 创建时间 (Optional)
 * param updateUser: 修改人 (Optional)
 * param updateTime: 修改时间 (Optional)
 * param content: 脚本内容 (Optional)
 */
func NewGravityParticleDubboJobManagerSaveJobScriptRequestWithAllParams(
    regionId string,
    appName string,
    id *int,
    scriptName *string,
    scriptType *string,
    scriptDesc *string,
    scriptSourceType *string,
    jobId *int,
    deleteFlag *int,
    creator *string,
    createTime *string,
    updateUser *string,
    updateTime *string,
    content *string,
) *GravityParticleDubboJobManagerSaveJobScriptRequest {

    return &GravityParticleDubboJobManagerSaveJobScriptRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/gravityParticleDubboJobManagerSaveJobScript",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
        Id: id,
        ScriptName: scriptName,
        ScriptType: scriptType,
        ScriptDesc: scriptDesc,
        ScriptSourceType: scriptSourceType,
        JobId: jobId,
        DeleteFlag: deleteFlag,
        Creator: creator,
        CreateTime: createTime,
        UpdateUser: updateUser,
        UpdateTime: updateTime,
        Content: content,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGravityParticleDubboJobManagerSaveJobScriptRequestWithoutParam() *GravityParticleDubboJobManagerSaveJobScriptRequest {

    return &GravityParticleDubboJobManagerSaveJobScriptRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/gravityParticleDubboJobManagerSaveJobScript",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *GravityParticleDubboJobManagerSaveJobScriptRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *GravityParticleDubboJobManagerSaveJobScriptRequest) SetAppName(appName string) {
    r.AppName = appName
}
/* param id: 主键Id,为空时保存，不为空时更新(Optional) */
func (r *GravityParticleDubboJobManagerSaveJobScriptRequest) SetId(id int) {
    r.Id = &id
}
/* param scriptName: 脚本名称(Optional) */
func (r *GravityParticleDubboJobManagerSaveJobScriptRequest) SetScriptName(scriptName string) {
    r.ScriptName = &scriptName
}
/* param scriptType: 脚本类型(Optional) */
func (r *GravityParticleDubboJobManagerSaveJobScriptRequest) SetScriptType(scriptType string) {
    r.ScriptType = &scriptType
}
/* param scriptDesc: 脚本描述(Optional) */
func (r *GravityParticleDubboJobManagerSaveJobScriptRequest) SetScriptDesc(scriptDesc string) {
    r.ScriptDesc = &scriptDesc
}
/* param scriptSourceType: 脚本来源， LOCAL(本地) GIT(git) SKYDRIVE(网盘)(Optional) */
func (r *GravityParticleDubboJobManagerSaveJobScriptRequest) SetScriptSourceType(scriptSourceType string) {
    r.ScriptSourceType = &scriptSourceType
}
/* param jobId: 任务id(Optional) */
func (r *GravityParticleDubboJobManagerSaveJobScriptRequest) SetJobId(jobId int) {
    r.JobId = &jobId
}
/* param deleteFlag: 删除标识(Optional) */
func (r *GravityParticleDubboJobManagerSaveJobScriptRequest) SetDeleteFlag(deleteFlag int) {
    r.DeleteFlag = &deleteFlag
}
/* param creator: 创建人(Optional) */
func (r *GravityParticleDubboJobManagerSaveJobScriptRequest) SetCreator(creator string) {
    r.Creator = &creator
}
/* param createTime: 创建时间(Optional) */
func (r *GravityParticleDubboJobManagerSaveJobScriptRequest) SetCreateTime(createTime string) {
    r.CreateTime = &createTime
}
/* param updateUser: 修改人(Optional) */
func (r *GravityParticleDubboJobManagerSaveJobScriptRequest) SetUpdateUser(updateUser string) {
    r.UpdateUser = &updateUser
}
/* param updateTime: 修改时间(Optional) */
func (r *GravityParticleDubboJobManagerSaveJobScriptRequest) SetUpdateTime(updateTime string) {
    r.UpdateTime = &updateTime
}
/* param content: 脚本内容(Optional) */
func (r *GravityParticleDubboJobManagerSaveJobScriptRequest) SetContent(content string) {
    r.Content = &content
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GravityParticleDubboJobManagerSaveJobScriptRequest) GetRegionId() string {
    return r.RegionId
}

type GravityParticleDubboJobManagerSaveJobScriptResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GravityParticleDubboJobManagerSaveJobScriptResult `json:"result"`
}

type GravityParticleDubboJobManagerSaveJobScriptResult struct {
    Success int `json:"success"`
    Code string `json:"code"`
    Msg string `json:"msg"`
    Result int `json:"result"`
}