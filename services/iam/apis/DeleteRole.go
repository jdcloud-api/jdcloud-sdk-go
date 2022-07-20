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
    "github.com/lshuining/jdcloud-sdk-go/core"
)

type DeleteRoleRequest struct {

    core.JDCloudRequest

    /* 角色名称  */
    RoleName string `json:"roleName"`
}

/*
 * param roleName: 角色名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteRoleRequest(
    roleName string,
) *DeleteRoleRequest {

	return &DeleteRoleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/role/{roleName}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RoleName: roleName,
	}
}

/*
 * param roleName: 角色名称 (Required)
 */
func NewDeleteRoleRequestWithAllParams(
    roleName string,
) *DeleteRoleRequest {

    return &DeleteRoleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/role/{roleName}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RoleName: roleName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteRoleRequestWithoutParam() *DeleteRoleRequest {

    return &DeleteRoleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/role/{roleName}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param roleName: 角色名称(Required) */
func (r *DeleteRoleRequest) SetRoleName(roleName string) {
    r.RoleName = roleName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteRoleRequest) GetRegionId() string {
    return ""
}

type DeleteRoleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteRoleResult `json:"result"`
}

type DeleteRoleResult struct {
}