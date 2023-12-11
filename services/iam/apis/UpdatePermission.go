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
    iam "github.com/jdcloud-api/jdcloud-sdk-go/services/iam/models"
)

type UpdatePermissionRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 权限id  */
    PermissionId int `json:"permissionId"`

    /* 权限信息  */
    UpdatePermissionInfo *iam.UpdatePermissionInfo `json:"updatePermissionInfo"`
}

/*
 * param regionId: Region ID (Required)
 * param permissionId: 权限id (Required)
 * param updatePermissionInfo: 权限信息 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdatePermissionRequest(
    regionId string,
    permissionId int,
    updatePermissionInfo *iam.UpdatePermissionInfo,
) *UpdatePermissionRequest {

	return &UpdatePermissionRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/permission/{permissionId}",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        PermissionId: permissionId,
        UpdatePermissionInfo: updatePermissionInfo,
	}
}

/*
 * param regionId: Region ID (Required)
 * param permissionId: 权限id (Required)
 * param updatePermissionInfo: 权限信息 (Required)
 */
func NewUpdatePermissionRequestWithAllParams(
    regionId string,
    permissionId int,
    updatePermissionInfo *iam.UpdatePermissionInfo,
) *UpdatePermissionRequest {

    return &UpdatePermissionRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/permission/{permissionId}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PermissionId: permissionId,
        UpdatePermissionInfo: updatePermissionInfo,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdatePermissionRequestWithoutParam() *UpdatePermissionRequest {

    return &UpdatePermissionRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/permission/{permissionId}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *UpdatePermissionRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param permissionId: 权限id(Required) */
func (r *UpdatePermissionRequest) SetPermissionId(permissionId int) {
    r.PermissionId = permissionId
}
/* param updatePermissionInfo: 权限信息(Required) */
func (r *UpdatePermissionRequest) SetUpdatePermissionInfo(updatePermissionInfo *iam.UpdatePermissionInfo) {
    r.UpdatePermissionInfo = updatePermissionInfo
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdatePermissionRequest) GetRegionId() string {
    return r.RegionId
}

type UpdatePermissionResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdatePermissionResult `json:"result"`
}

type UpdatePermissionResult struct {
}