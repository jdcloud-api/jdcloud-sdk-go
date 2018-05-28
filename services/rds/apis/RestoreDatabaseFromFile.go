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

type RestoreDatabaseFromFileRequest struct {

    core.JDCloudRequest

    /* 区域代码  */
    RegionId string `json:"regionId"`

    /* 实例ID  */
    InstanceId string `json:"instanceId"`

    /* 库名称  */
    DbName string `json:"dbName"`

    /* 共享文件的全局ID，可从上传文件查询接口describeImportFiles获取；如果该文件不是共享文件，则全局ID为空 (Optional) */
    SharedFileGid *string `json:""`

    /* 用户在单库上云中上传的文件名称  */
    FileName string `json:""`
}

/*
 * param regionId: 区域代码 
 * param instanceId: 实例ID 
 * param dbName: 库名称 
 * param sharedFileGid: 共享文件的全局ID，可从上传文件查询接口describeImportFiles获取；如果该文件不是共享文件，则全局ID为空 (Optional)
 * param fileName: 用户在单库上云中上传的文件名称 
 */
func NewRestoreDatabaseFromFileRequest(
    regionId string,
    instanceId string,
    dbName string,
    fileName string,
) *RestoreDatabaseFromFileRequest {

	return &RestoreDatabaseFromFileRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/databases/{dbName}:restoreDatabaseFromFile",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        DbName: dbName,
        FileName: fileName,
	}
}

func (r *RestoreDatabaseFromFileRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *RestoreDatabaseFromFileRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

func (r *RestoreDatabaseFromFileRequest) SetDbName(dbName string) {
    r.DbName = dbName
}

func (r *RestoreDatabaseFromFileRequest) SetSharedFileGid(sharedFileGid string) {
    r.SharedFileGid = &sharedFileGid
}

func (r *RestoreDatabaseFromFileRequest) SetFileName(fileName string) {
    r.FileName = fileName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r RestoreDatabaseFromFileRequest) GetRegionId() string {
    return r.RegionId
}

type RestoreDatabaseFromFileResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result RestoreDatabaseFromFileResult `json:"result"`
}

type RestoreDatabaseFromFileResult struct {
}