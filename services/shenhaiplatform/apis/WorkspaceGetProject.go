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
    shenhaiplatform "github.com/jdcloud-api/jdcloud-sdk-go/services/shenhaiplatform/models"
)

type WorkspaceGetProjectRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /*  (Optional) */
    Id *int `json:"id"`

    /*  (Optional) */
    WorkspaceCode *string `json:"workspaceCode"`

    /*  (Optional) */
    ProjectCode *string `json:"projectCode"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewWorkspaceGetProjectRequest(
    regionId string,
    appName string,
) *WorkspaceGetProjectRequest {

	return &WorkspaceGetProjectRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/workspaceGetProject",
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
 * param id:  (Optional)
 * param workspaceCode:  (Optional)
 * param projectCode:  (Optional)
 */
func NewWorkspaceGetProjectRequestWithAllParams(
    regionId string,
    appName string,
    id *int,
    workspaceCode *string,
    projectCode *string,
) *WorkspaceGetProjectRequest {

    return &WorkspaceGetProjectRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/workspaceGetProject",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
        Id: id,
        WorkspaceCode: workspaceCode,
        ProjectCode: projectCode,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewWorkspaceGetProjectRequestWithoutParam() *WorkspaceGetProjectRequest {

    return &WorkspaceGetProjectRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/workspaceGetProject",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *WorkspaceGetProjectRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *WorkspaceGetProjectRequest) SetAppName(appName string) {
    r.AppName = appName
}
/* param id: (Optional) */
func (r *WorkspaceGetProjectRequest) SetId(id int) {
    r.Id = &id
}
/* param workspaceCode: (Optional) */
func (r *WorkspaceGetProjectRequest) SetWorkspaceCode(workspaceCode string) {
    r.WorkspaceCode = &workspaceCode
}
/* param projectCode: (Optional) */
func (r *WorkspaceGetProjectRequest) SetProjectCode(projectCode string) {
    r.ProjectCode = &projectCode
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r WorkspaceGetProjectRequest) GetRegionId() string {
    return r.RegionId
}

type WorkspaceGetProjectResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result WorkspaceGetProjectResult `json:"result"`
}

type WorkspaceGetProjectResult struct {
    Code string `json:"code"`
    Message string `json:"message"`
    Result []shenhaiplatform.ProjectResp `json:"result"`
}