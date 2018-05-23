// Copyright 2018-2025 JDCLOUD.COM
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
    . "github.com/jdcloud-api/jdcloud-sdk-go/core"
    "reflect"
)

type DeleteDatabaseRequest struct {

    JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 数据库名  */
    DatabaseName string `json:"databaseName"`

    /* 实例名称  */
    InstanceName string `json:"instanceName"`
}

/*
 * param regionId: 地域ID 
 * param databaseName: 数据库名 
 * param instanceName: 实例名称 
 */
func NewDeleteDatabaseRequest(
    regionId string,
    databaseName string,
    instanceName string,
) *DeleteDatabaseRequest {

	return &DeleteDatabaseRequest{
        JDCloudRequest: JDCloudRequest{
			URL:     "/regions/{regionId}/dwDatabase/{databaseName}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        DatabaseName: databaseName,
        InstanceName: instanceName,
	}
}

func (r *DeleteDatabaseRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *DeleteDatabaseRequest) SetDatabaseName(databaseName string) {
    r.DatabaseName = databaseName
}

func (r *DeleteDatabaseRequest) SetInstanceName(instanceName string) {
    r.InstanceName = instanceName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteDatabaseRequest) GetRegionId() string {
    fieldName := "RegionId"
    reqType := reflect.TypeOf(r)
    value := reflect.ValueOf(r)
    _, ok := reqType.FieldByName(fieldName)
    if ok {
        return value.FieldByName(fieldName).String()
    }

    return ""
}

type DeleteDatabaseResponse struct {
    RequestID string `json:"requestId"`
    Error ErrorResponse `json:"error"`
    Result DeleteDatabaseResult `json:"result"`
}

type DeleteDatabaseResult struct {
    Status bool `json:"status"`
    Message string `json:"message"`
}