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
    streamcomputer "github.com/lshuining/jdcloud-sdk-go/services/streamcomputer/models"
)

type AddOrUpdateStorageRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 创建或者更新storage的详情  */
    StorageStr *streamcomputer.Storage `json:"storageStr"`
}

/*
 * param regionId: Region ID (Required)
 * param storageStr: 创建或者更新storage的详情 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAddOrUpdateStorageRequest(
    regionId string,
    storageStr *streamcomputer.Storage,
) *AddOrUpdateStorageRequest {

	return &AddOrUpdateStorageRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/storage",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        StorageStr: storageStr,
	}
}

/*
 * param regionId: Region ID (Required)
 * param storageStr: 创建或者更新storage的详情 (Required)
 */
func NewAddOrUpdateStorageRequestWithAllParams(
    regionId string,
    storageStr *streamcomputer.Storage,
) *AddOrUpdateStorageRequest {

    return &AddOrUpdateStorageRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/storage",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        StorageStr: storageStr,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAddOrUpdateStorageRequestWithoutParam() *AddOrUpdateStorageRequest {

    return &AddOrUpdateStorageRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/storage",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *AddOrUpdateStorageRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param storageStr: 创建或者更新storage的详情(Required) */
func (r *AddOrUpdateStorageRequest) SetStorageStr(storageStr *streamcomputer.Storage) {
    r.StorageStr = storageStr
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AddOrUpdateStorageRequest) GetRegionId() string {
    return r.RegionId
}

type AddOrUpdateStorageResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AddOrUpdateStorageResult `json:"result"`
}

type AddOrUpdateStorageResult struct {
    Message string `json:"message"`
    Status bool `json:"status"`
}