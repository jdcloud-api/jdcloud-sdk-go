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

type WorkspaceGetListRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /*  (Optional) */
    PageNum *int `json:"pageNum"`

    /*  (Optional) */
    PageSize *int `json:"pageSize"`

    /*  (Optional) */
    WorkspaceName *string `json:"workspaceName"`

    /*  (Optional) */
    WorkspaceId *string `json:"workspaceId"`

    /*  (Optional) */
    Manager *string `json:"manager"`

    /*  (Optional) */
    ImProjectManager *bool `json:"imProjectManager"`

    /*  (Optional) */
    ImInProject *bool `json:"imInProject"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewWorkspaceGetListRequest(
    regionId string,
    appName string,
) *WorkspaceGetListRequest {

	return &WorkspaceGetListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/workspaceGetList",
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
 * param pageNum:  (Optional)
 * param pageSize:  (Optional)
 * param workspaceName:  (Optional)
 * param workspaceId:  (Optional)
 * param manager:  (Optional)
 * param imProjectManager:  (Optional)
 * param imInProject:  (Optional)
 */
func NewWorkspaceGetListRequestWithAllParams(
    regionId string,
    appName string,
    pageNum *int,
    pageSize *int,
    workspaceName *string,
    workspaceId *string,
    manager *string,
    imProjectManager *bool,
    imInProject *bool,
) *WorkspaceGetListRequest {

    return &WorkspaceGetListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/workspaceGetList",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
        PageNum: pageNum,
        PageSize: pageSize,
        WorkspaceName: workspaceName,
        WorkspaceId: workspaceId,
        Manager: manager,
        ImProjectManager: imProjectManager,
        ImInProject: imInProject,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewWorkspaceGetListRequestWithoutParam() *WorkspaceGetListRequest {

    return &WorkspaceGetListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/workspaceGetList",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *WorkspaceGetListRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *WorkspaceGetListRequest) SetAppName(appName string) {
    r.AppName = appName
}
/* param pageNum: (Optional) */
func (r *WorkspaceGetListRequest) SetPageNum(pageNum int) {
    r.PageNum = &pageNum
}
/* param pageSize: (Optional) */
func (r *WorkspaceGetListRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}
/* param workspaceName: (Optional) */
func (r *WorkspaceGetListRequest) SetWorkspaceName(workspaceName string) {
    r.WorkspaceName = &workspaceName
}
/* param workspaceId: (Optional) */
func (r *WorkspaceGetListRequest) SetWorkspaceId(workspaceId string) {
    r.WorkspaceId = &workspaceId
}
/* param manager: (Optional) */
func (r *WorkspaceGetListRequest) SetManager(manager string) {
    r.Manager = &manager
}
/* param imProjectManager: (Optional) */
func (r *WorkspaceGetListRequest) SetImProjectManager(imProjectManager bool) {
    r.ImProjectManager = &imProjectManager
}
/* param imInProject: (Optional) */
func (r *WorkspaceGetListRequest) SetImInProject(imInProject bool) {
    r.ImInProject = &imInProject
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r WorkspaceGetListRequest) GetRegionId() string {
    return r.RegionId
}

type WorkspaceGetListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result WorkspaceGetListResult `json:"result"`
}

type WorkspaceGetListResult struct {
    Code string `json:"code"`
    Message string `json:"message"`
    Result shenhaiplatform.PageInfoWorkspaceListResp `json:"result"`
}