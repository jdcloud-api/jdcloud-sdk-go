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

type RestoreInstanceRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* RDS 实例ID，唯一标识一个RDS实例  */
    InstanceId string `json:"instanceId"`

    /* 用于恢复的备份Id，仅限于本实例生成的备份 (Optional) */
    BackupId *string `json:"backupId"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: RDS 实例ID，唯一标识一个RDS实例 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewRestoreInstanceRequest(
    regionId string,
    instanceId string,
) *RestoreInstanceRequest {

	return &RestoreInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:restoreInstance",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: RDS 实例ID，唯一标识一个RDS实例 (Required)
 * param backupId: 用于恢复的备份Id，仅限于本实例生成的备份 (Optional)
 */
func NewRestoreInstanceRequestWithAllParams(
    regionId string,
    instanceId string,
    backupId *string,
) *RestoreInstanceRequest {

    return &RestoreInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:restoreInstance",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        BackupId: backupId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewRestoreInstanceRequestWithoutParam() *RestoreInstanceRequest {

    return &RestoreInstanceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:restoreInstance",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *RestoreInstanceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: RDS 实例ID，唯一标识一个RDS实例(Required) */
func (r *RestoreInstanceRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param backupId: 用于恢复的备份Id，仅限于本实例生成的备份(Optional) */
func (r *RestoreInstanceRequest) SetBackupId(backupId string) {
    r.BackupId = &backupId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r RestoreInstanceRequest) GetRegionId() string {
    return r.RegionId
}

type RestoreInstanceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result RestoreInstanceResult `json:"result"`
}

type RestoreInstanceResult struct {
}