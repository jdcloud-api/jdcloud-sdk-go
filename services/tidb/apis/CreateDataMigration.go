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
    tidb "github.com/jdcloud-api/jdcloud-sdk-go/services/tidb/models"
)

type CreateDataMigrationRequest struct {

    core.JDCloudRequest

    /* 地域代码  */
    RegionId string `json:"regionId"`

    /* 实例ID  */
    InstanceId string `json:"instanceId"`

    /* 迁移任务类型，支持以下类型（大小写不敏感）：-FULL_IMPORT:全量数据导入  */
    MigrationType string `json:"migrationType"`

    /* 使用 TiDB Lightning 进行的数据迁移任务  */
    ImportTask *tidb.FullImportTask `json:"importTask"`
}

/*
 * param regionId: 地域代码 (Required)
 * param instanceId: 实例ID (Required)
 * param migrationType: 迁移任务类型，支持以下类型（大小写不敏感）：-FULL_IMPORT:全量数据导入 (Required)
 * param importTask: 使用 TiDB Lightning 进行的数据迁移任务 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateDataMigrationRequest(
    regionId string,
    instanceId string,
    migrationType string,
    importTask *tidb.FullImportTask,
) *CreateDataMigrationRequest {

	return &CreateDataMigrationRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/migration",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        MigrationType: migrationType,
        ImportTask: importTask,
	}
}

/*
 * param regionId: 地域代码 (Required)
 * param instanceId: 实例ID (Required)
 * param migrationType: 迁移任务类型，支持以下类型（大小写不敏感）：-FULL_IMPORT:全量数据导入 (Required)
 * param importTask: 使用 TiDB Lightning 进行的数据迁移任务 (Required)
 */
func NewCreateDataMigrationRequestWithAllParams(
    regionId string,
    instanceId string,
    migrationType string,
    importTask *tidb.FullImportTask,
) *CreateDataMigrationRequest {

    return &CreateDataMigrationRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/migration",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        MigrationType: migrationType,
        ImportTask: importTask,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateDataMigrationRequestWithoutParam() *CreateDataMigrationRequest {

    return &CreateDataMigrationRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/migration",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码(Required) */
func (r *CreateDataMigrationRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 实例ID(Required) */
func (r *CreateDataMigrationRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param migrationType: 迁移任务类型，支持以下类型（大小写不敏感）：-FULL_IMPORT:全量数据导入(Required) */
func (r *CreateDataMigrationRequest) SetMigrationType(migrationType string) {
    r.MigrationType = migrationType
}

/* param importTask: 使用 TiDB Lightning 进行的数据迁移任务(Required) */
func (r *CreateDataMigrationRequest) SetImportTask(importTask *tidb.FullImportTask) {
    r.ImportTask = importTask
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateDataMigrationRequest) GetRegionId() string {
    return r.RegionId
}

type CreateDataMigrationResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateDataMigrationResult `json:"result"`
}

type CreateDataMigrationResult struct {
}