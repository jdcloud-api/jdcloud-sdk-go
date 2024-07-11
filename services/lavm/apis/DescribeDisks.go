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
    lavm "github.com/jdcloud-api/jdcloud-sdk-go/services/lavm/models"
)

type DescribeDisksRequest struct {

    core.JDCloudRequest

    /* regionId
  */
    RegionId string `json:"regionId"`

    /* 指定的轻量应用云主机的实例ID
 (Optional) */
    InstanceId *string `json:"instanceId"`

    /* 磁盘ID, `[\"volume-xxx\", \"volume-yyy\"]`, json array 字串
 (Optional) */
    DiskIds *string `json:"diskIds"`

    /* 页码；默认为1。 (Optional) */
    PageSize *int `json:"pageSize"`

    /* 分页大小；<br>默认为20；取值范围[10, 100]。 (Optional) */
    PageNumber *int `json:"pageNumber"`
}

/*
 * param regionId: regionId
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeDisksRequest(
    regionId string,
) *DescribeDisksRequest {

	return &DescribeDisksRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/disks",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: regionId
 (Required)
 * param instanceId: 指定的轻量应用云主机的实例ID
 (Optional)
 * param diskIds: 磁盘ID, `[\"volume-xxx\", \"volume-yyy\"]`, json array 字串
 (Optional)
 * param pageSize: 页码；默认为1。 (Optional)
 * param pageNumber: 分页大小；<br>默认为20；取值范围[10, 100]。 (Optional)
 */
func NewDescribeDisksRequestWithAllParams(
    regionId string,
    instanceId *string,
    diskIds *string,
    pageSize *int,
    pageNumber *int,
) *DescribeDisksRequest {

    return &DescribeDisksRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/disks",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        DiskIds: diskIds,
        PageSize: pageSize,
        PageNumber: pageNumber,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeDisksRequestWithoutParam() *DescribeDisksRequest {

    return &DescribeDisksRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/disks",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: regionId
(Required) */
func (r *DescribeDisksRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param instanceId: 指定的轻量应用云主机的实例ID
(Optional) */
func (r *DescribeDisksRequest) SetInstanceId(instanceId string) {
    r.InstanceId = &instanceId
}
/* param diskIds: 磁盘ID, `[\"volume-xxx\", \"volume-yyy\"]`, json array 字串
(Optional) */
func (r *DescribeDisksRequest) SetDiskIds(diskIds string) {
    r.DiskIds = &diskIds
}
/* param pageSize: 页码；默认为1。(Optional) */
func (r *DescribeDisksRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}
/* param pageNumber: 分页大小；<br>默认为20；取值范围[10, 100]。(Optional) */
func (r *DescribeDisksRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeDisksRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeDisksResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeDisksResult `json:"result"`
}

type DescribeDisksResult struct {
    Disks []lavm.Disk `json:"disks"`
    Total int `json:"total"`
}