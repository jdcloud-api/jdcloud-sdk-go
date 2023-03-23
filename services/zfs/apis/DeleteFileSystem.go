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

type DeleteFileSystemRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 文件系统ID  */
    FileSystemId string `json:"fileSystemId"`
}

/*
 * param regionId: 地域ID (Required)
 * param fileSystemId: 文件系统ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteFileSystemRequest(
    regionId string,
    fileSystemId string,
) *DeleteFileSystemRequest {

	return &DeleteFileSystemRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/fileSystems/{fileSystemId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        FileSystemId: fileSystemId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param fileSystemId: 文件系统ID (Required)
 */
func NewDeleteFileSystemRequestWithAllParams(
    regionId string,
    fileSystemId string,
) *DeleteFileSystemRequest {

    return &DeleteFileSystemRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/fileSystems/{fileSystemId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        FileSystemId: fileSystemId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteFileSystemRequestWithoutParam() *DeleteFileSystemRequest {

    return &DeleteFileSystemRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/fileSystems/{fileSystemId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DeleteFileSystemRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param fileSystemId: 文件系统ID(Required) */
func (r *DeleteFileSystemRequest) SetFileSystemId(fileSystemId string) {
    r.FileSystemId = fileSystemId
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteFileSystemRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteFileSystemResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteFileSystemResult `json:"result"`
}

type DeleteFileSystemResult struct {
}