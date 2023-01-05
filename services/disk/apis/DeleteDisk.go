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

type DeleteDiskRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 云硬盘ID  */
    DiskId string `json:"diskId"`

    /* true 加入回收站 false 或者不传直接删除 (Optional) */
    PutInRecycleBin *bool `json:"putInRecycleBin"`
}

/*
 * param regionId: 地域ID (Required)
 * param diskId: 云硬盘ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteDiskRequest(
    regionId string,
    diskId string,
) *DeleteDiskRequest {

	return &DeleteDiskRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/disks/{diskId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        DiskId: diskId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param diskId: 云硬盘ID (Required)
 * param putInRecycleBin: true 加入回收站 false 或者不传直接删除 (Optional)
 */
func NewDeleteDiskRequestWithAllParams(
    regionId string,
    diskId string,
    putInRecycleBin *bool,
) *DeleteDiskRequest {

    return &DeleteDiskRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/disks/{diskId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DiskId: diskId,
        PutInRecycleBin: putInRecycleBin,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteDiskRequestWithoutParam() *DeleteDiskRequest {

    return &DeleteDiskRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/disks/{diskId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DeleteDiskRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param diskId: 云硬盘ID(Required) */
func (r *DeleteDiskRequest) SetDiskId(diskId string) {
    r.DiskId = diskId
}
/* param putInRecycleBin: true 加入回收站 false 或者不传直接删除(Optional) */
func (r *DeleteDiskRequest) SetPutInRecycleBin(putInRecycleBin bool) {
    r.PutInRecycleBin = &putInRecycleBin
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteDiskRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteDiskResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteDiskResult `json:"result"`
}

type DeleteDiskResult struct {
}