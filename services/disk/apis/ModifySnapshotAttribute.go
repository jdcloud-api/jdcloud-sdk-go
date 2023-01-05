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

type ModifySnapshotAttributeRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 快照ID  */
    SnapshotId string `json:"snapshotId"`

    /* 快照名称 (Optional) */
    Name *string `json:"name"`

    /* 快照描述 (Optional) */
    Description *string `json:"description"`

    /* 快照过期时间，三者至少指定一个 (Optional) */
    ExpireTime *string `json:"expireTime"`
}

/*
 * param regionId: 地域ID (Required)
 * param snapshotId: 快照ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifySnapshotAttributeRequest(
    regionId string,
    snapshotId string,
) *ModifySnapshotAttributeRequest {

	return &ModifySnapshotAttributeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/snapshots/{snapshotId}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        SnapshotId: snapshotId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param snapshotId: 快照ID (Required)
 * param name: 快照名称 (Optional)
 * param description: 快照描述 (Optional)
 * param expireTime: 快照过期时间，三者至少指定一个 (Optional)
 */
func NewModifySnapshotAttributeRequestWithAllParams(
    regionId string,
    snapshotId string,
    name *string,
    description *string,
    expireTime *string,
) *ModifySnapshotAttributeRequest {

    return &ModifySnapshotAttributeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/snapshots/{snapshotId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        SnapshotId: snapshotId,
        Name: name,
        Description: description,
        ExpireTime: expireTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifySnapshotAttributeRequestWithoutParam() *ModifySnapshotAttributeRequest {

    return &ModifySnapshotAttributeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/snapshots/{snapshotId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *ModifySnapshotAttributeRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param snapshotId: 快照ID(Required) */
func (r *ModifySnapshotAttributeRequest) SetSnapshotId(snapshotId string) {
    r.SnapshotId = snapshotId
}
/* param name: 快照名称(Optional) */
func (r *ModifySnapshotAttributeRequest) SetName(name string) {
    r.Name = &name
}
/* param description: 快照描述(Optional) */
func (r *ModifySnapshotAttributeRequest) SetDescription(description string) {
    r.Description = &description
}
/* param expireTime: 快照过期时间，三者至少指定一个(Optional) */
func (r *ModifySnapshotAttributeRequest) SetExpireTime(expireTime string) {
    r.ExpireTime = &expireTime
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifySnapshotAttributeRequest) GetRegionId() string {
    return r.RegionId
}

type ModifySnapshotAttributeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifySnapshotAttributeResult `json:"result"`
}

type ModifySnapshotAttributeResult struct {
}