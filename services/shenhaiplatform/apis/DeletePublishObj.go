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

type DeletePublishObjRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* ID列表 (Optional) */
    ObjIds []int64 `json:"objIds"`

    /* 当前工作空间编码 (Optional) */
    WorkspaceCode *string `json:"workspaceCode"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeletePublishObjRequest(
    regionId string,
    appName string,
) *DeletePublishObjRequest {

	return &DeletePublishObjRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/deletePublishObj",
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
 * param objIds: ID列表 (Optional)
 * param workspaceCode: 当前工作空间编码 (Optional)
 */
func NewDeletePublishObjRequestWithAllParams(
    regionId string,
    appName string,
    objIds []int64,
    workspaceCode *string,
) *DeletePublishObjRequest {

    return &DeletePublishObjRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/deletePublishObj",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
        ObjIds: objIds,
        WorkspaceCode: workspaceCode,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeletePublishObjRequestWithoutParam() *DeletePublishObjRequest {

    return &DeletePublishObjRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/deletePublishObj",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DeletePublishObjRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *DeletePublishObjRequest) SetAppName(appName string) {
    r.AppName = appName
}
/* param objIds: ID列表(Optional) */
func (r *DeletePublishObjRequest) SetObjIds(objIds []int64) {
    r.ObjIds = objIds
}
/* param workspaceCode: 当前工作空间编码(Optional) */
func (r *DeletePublishObjRequest) SetWorkspaceCode(workspaceCode string) {
    r.WorkspaceCode = &workspaceCode
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeletePublishObjRequest) GetRegionId() string {
    return r.RegionId
}

type DeletePublishObjResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeletePublishObjResult `json:"result"`
}

type DeletePublishObjResult struct {
    Code string `json:"code"`
    Message string `json:"message"`
}