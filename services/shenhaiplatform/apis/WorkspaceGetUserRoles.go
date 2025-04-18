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

type WorkspaceGetUserRolesRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewWorkspaceGetUserRolesRequest(
    regionId string,
    appName string,
) *WorkspaceGetUserRolesRequest {

	return &WorkspaceGetUserRolesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/workspaceGetUserRoles",
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
 */
func NewWorkspaceGetUserRolesRequestWithAllParams(
    regionId string,
    appName string,
) *WorkspaceGetUserRolesRequest {

    return &WorkspaceGetUserRolesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/workspaceGetUserRoles",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewWorkspaceGetUserRolesRequestWithoutParam() *WorkspaceGetUserRolesRequest {

    return &WorkspaceGetUserRolesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/workspaceGetUserRoles",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *WorkspaceGetUserRolesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *WorkspaceGetUserRolesRequest) SetAppName(appName string) {
    r.AppName = appName
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r WorkspaceGetUserRolesRequest) GetRegionId() string {
    return r.RegionId
}

type WorkspaceGetUserRolesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result WorkspaceGetUserRolesResult `json:"result"`
}

type WorkspaceGetUserRolesResult struct {
    Code string `json:"code"`
    Message string `json:"message"`
    Result []shenhaiplatform.WorkspaceUserRoleResp `json:"result"`
}