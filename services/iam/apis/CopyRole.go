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

type CopyRoleRequest struct {

    core.JDCloudRequest

    /* 角色信息  */
    CopyRoleInfo *iam.CopyRoleInfo `json:"copyRoleInfo"`
}

/*
 * param copyRoleInfo: 角色信息 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCopyRoleRequest(
    copyRoleInfo *iam.CopyRoleInfo,
) *CopyRoleRequest {

	return &CopyRoleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/copyRole",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        CopyRoleInfo: copyRoleInfo,
	}
}

/*
 * param copyRoleInfo: 角色信息 (Required)
 */
func NewCopyRoleRequestWithAllParams(
    copyRoleInfo *iam.CopyRoleInfo,
) *CopyRoleRequest {

    return &CopyRoleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/copyRole",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        CopyRoleInfo: copyRoleInfo,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCopyRoleRequestWithoutParam() *CopyRoleRequest {

    return &CopyRoleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/copyRole",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param copyRoleInfo: 角色信息(Required) */
func (r *CopyRoleRequest) SetCopyRoleInfo(copyRoleInfo *iam.CopyRoleInfo) {
    r.CopyRoleInfo = copyRoleInfo
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CopyRoleRequest) GetRegionId() string {
    return ""
}

type CopyRoleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CopyRoleResult `json:"result"`
}

type CopyRoleResult struct {
    RoleInfo iam.RoleInfo `json:"roleInfo"`
}