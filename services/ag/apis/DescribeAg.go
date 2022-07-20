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
    ag "github.com/lshuining/jdcloud-sdk-go/services/ag/models"
)

type DescribeAgRequest struct {

    core.JDCloudRequest

    /* 地域  */
    RegionId string `json:"regionId"`

    /* 高可用组 ID  */
    AgId string `json:"agId"`
}

/*
 * param regionId: 地域 (Required)
 * param agId: 高可用组 ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeAgRequest(
    regionId string,
    agId string,
) *DescribeAgRequest {

	return &DescribeAgRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/availabilityGroups/{agId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        AgId: agId,
	}
}

/*
 * param regionId: 地域 (Required)
 * param agId: 高可用组 ID (Required)
 */
func NewDescribeAgRequestWithAllParams(
    regionId string,
    agId string,
) *DescribeAgRequest {

    return &DescribeAgRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/availabilityGroups/{agId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        AgId: agId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeAgRequestWithoutParam() *DescribeAgRequest {

    return &DescribeAgRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/availabilityGroups/{agId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域(Required) */
func (r *DescribeAgRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param agId: 高可用组 ID(Required) */
func (r *DescribeAgRequest) SetAgId(agId string) {
    r.AgId = agId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeAgRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeAgResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeAgResult `json:"result"`
}

type DescribeAgResult struct {
    Ag ag.AvailabilityGroup `json:"ag"`
}