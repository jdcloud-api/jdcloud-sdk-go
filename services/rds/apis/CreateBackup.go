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
    rds "github.com/lshuining/jdcloud-sdk-go/services/rds/models"
)

type CreateBackupRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* RDS实例ID，唯一标识一个实例 (Optional) */
    InstanceId *string `json:"instanceId"`

    /* 备份规格 (Optional) */
    BackupSpec *rds.BackupSpec `json:"backupSpec"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateBackupRequest(
    regionId string,
) *CreateBackupRequest {

	return &CreateBackupRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/backups",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: RDS实例ID，唯一标识一个实例 (Optional)
 * param backupSpec: 备份规格 (Optional)
 */
func NewCreateBackupRequestWithAllParams(
    regionId string,
    instanceId *string,
    backupSpec *rds.BackupSpec,
) *CreateBackupRequest {

    return &CreateBackupRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/backups",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        BackupSpec: backupSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateBackupRequestWithoutParam() *CreateBackupRequest {

    return &CreateBackupRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/backups",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *CreateBackupRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: RDS实例ID，唯一标识一个实例(Optional) */
func (r *CreateBackupRequest) SetInstanceId(instanceId string) {
    r.InstanceId = &instanceId
}

/* param backupSpec: 备份规格(Optional) */
func (r *CreateBackupRequest) SetBackupSpec(backupSpec *rds.BackupSpec) {
    r.BackupSpec = backupSpec
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateBackupRequest) GetRegionId() string {
    return r.RegionId
}

type CreateBackupResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateBackupResult `json:"result"`
}

type CreateBackupResult struct {
    BackupId string `json:"backupId"`
}