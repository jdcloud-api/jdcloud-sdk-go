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

type GetResultRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 多级筛选任务ID  */
    ProfileSelectRecordId string `json:"profileSelectRecordId"`
}

/*
 * param regionId: 地域ID (Required)
 * param profileSelectRecordId: 多级筛选任务ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetResultRequest(
    regionId string,
    profileSelectRecordId string,
) *GetResultRequest {

	return &GetResultRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/profileMultiLevel:getResult",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ProfileSelectRecordId: profileSelectRecordId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param profileSelectRecordId: 多级筛选任务ID (Required)
 */
func NewGetResultRequestWithAllParams(
    regionId string,
    profileSelectRecordId string,
) *GetResultRequest {

    return &GetResultRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/profileMultiLevel:getResult",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ProfileSelectRecordId: profileSelectRecordId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetResultRequestWithoutParam() *GetResultRequest {

    return &GetResultRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/profileMultiLevel:getResult",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *GetResultRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param profileSelectRecordId: 多级筛选任务ID(Required) */
func (r *GetResultRequest) SetProfileSelectRecordId(profileSelectRecordId string) {
    r.ProfileSelectRecordId = profileSelectRecordId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetResultRequest) GetRegionId() string {
    return r.RegionId
}

type GetResultResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetResultResult `json:"result"`
}

type GetResultResult struct {
    Status bool `json:"status"`
    Message string `json:"message"`
    Data string `json:"data"`
}