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

type DeleteBackupPlanRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》]  */
    RegionId string `json:"regionId"`

    /* 备份计划ID  */
    BackupPlanId string `json:"backupPlanId"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》] (Required)
 * param backupPlanId: 备份计划ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteBackupPlanRequest(
    regionId string,
    backupPlanId string,
) *DeleteBackupPlanRequest {

	return &DeleteBackupPlanRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/backupPlans/{backupPlanId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        BackupPlanId: backupPlanId,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》] (Required)
 * param backupPlanId: 备份计划ID (Required)
 */
func NewDeleteBackupPlanRequestWithAllParams(
    regionId string,
    backupPlanId string,
) *DeleteBackupPlanRequest {

    return &DeleteBackupPlanRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/backupPlans/{backupPlanId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        BackupPlanId: backupPlanId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteBackupPlanRequestWithoutParam() *DeleteBackupPlanRequest {

    return &DeleteBackupPlanRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/backupPlans/{backupPlanId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](Required) */
func (r *DeleteBackupPlanRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param backupPlanId: 备份计划ID(Required) */
func (r *DeleteBackupPlanRequest) SetBackupPlanId(backupPlanId string) {
    r.BackupPlanId = backupPlanId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteBackupPlanRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteBackupPlanResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteBackupPlanResult `json:"result"`
}

type DeleteBackupPlanResult struct {
}