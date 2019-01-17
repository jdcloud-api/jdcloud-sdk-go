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
    jdfusion "github.com/jdcloud-api/jdcloud-sdk-go/services/jdfusion/models"
)

type GetTransferTaskFailedFilesRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 任务ID  */
    Id string `json:"id"`
}

/*
 * param regionId: 地域ID (Required)
 * param id: 任务ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetTransferTaskFailedFilesRequest(
    regionId string,
    id string,
) *GetTransferTaskFailedFilesRequest {

	return &GetTransferTaskFailedFilesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/oss_transferTasks/{id}/failed-files",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Id: id,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param id: 任务ID (Required)
 */
func NewGetTransferTaskFailedFilesRequestWithAllParams(
    regionId string,
    id string,
) *GetTransferTaskFailedFilesRequest {

    return &GetTransferTaskFailedFilesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/oss_transferTasks/{id}/failed-files",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Id: id,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetTransferTaskFailedFilesRequestWithoutParam() *GetTransferTaskFailedFilesRequest {

    return &GetTransferTaskFailedFilesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/oss_transferTasks/{id}/failed-files",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *GetTransferTaskFailedFilesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param id: 任务ID(Required) */
func (r *GetTransferTaskFailedFilesRequest) SetId(id string) {
    r.Id = id
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetTransferTaskFailedFilesRequest) GetRegionId() string {
    return r.RegionId
}

type GetTransferTaskFailedFilesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetTransferTaskFailedFilesResult `json:"result"`
}

type GetTransferTaskFailedFilesResult struct {
    Bucket jdfusion.TransferTaskFailedDetailInfo `json:"bucket"`
}