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
    privatezone "github.com/jdcloud-api/jdcloud-sdk-go/services/privatezone/models"
)

type DescribeResourceRecordsRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* zone id  */
    ZoneId string `json:"zoneId"`

    /* 主机记录左右模糊查询 (Optional) */
    HostRecord *string `json:"hostRecord"`

    /* 页容量，默认10，取值范围(1 - 100) (Optional) */
    PageSize *int `json:"pageSize"`

    /* 页序号，默认1，不能小于1 (Optional) */
    PageNumber *int `json:"pageNumber"`
}

/*
 * param regionId: 地域ID (Required)
 * param zoneId: zone id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeResourceRecordsRequest(
    regionId string,
    zoneId string,
) *DescribeResourceRecordsRequest {

	return &DescribeResourceRecordsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/zone/{zoneId}/resourceRecords",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ZoneId: zoneId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param zoneId: zone id (Required)
 * param hostRecord: 主机记录左右模糊查询 (Optional)
 * param pageSize: 页容量，默认10，取值范围(1 - 100) (Optional)
 * param pageNumber: 页序号，默认1，不能小于1 (Optional)
 */
func NewDescribeResourceRecordsRequestWithAllParams(
    regionId string,
    zoneId string,
    hostRecord *string,
    pageSize *int,
    pageNumber *int,
) *DescribeResourceRecordsRequest {

    return &DescribeResourceRecordsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/zone/{zoneId}/resourceRecords",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ZoneId: zoneId,
        HostRecord: hostRecord,
        PageSize: pageSize,
        PageNumber: pageNumber,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeResourceRecordsRequestWithoutParam() *DescribeResourceRecordsRequest {

    return &DescribeResourceRecordsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/zone/{zoneId}/resourceRecords",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DescribeResourceRecordsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param zoneId: zone id(Required) */
func (r *DescribeResourceRecordsRequest) SetZoneId(zoneId string) {
    r.ZoneId = zoneId
}

/* param hostRecord: 主机记录左右模糊查询(Optional) */
func (r *DescribeResourceRecordsRequest) SetHostRecord(hostRecord string) {
    r.HostRecord = &hostRecord
}

/* param pageSize: 页容量，默认10，取值范围(1 - 100)(Optional) */
func (r *DescribeResourceRecordsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param pageNumber: 页序号，默认1，不能小于1(Optional) */
func (r *DescribeResourceRecordsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeResourceRecordsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeResourceRecordsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeResourceRecordsResult `json:"result"`
}

type DescribeResourceRecordsResult struct {
    DataList []privatezone.DescribeResourceRecordsRes `json:"dataList"`
    CurrentCount int `json:"currentCount"`
    TotalCount int `json:"totalCount"`
    TotalPage int `json:"totalPage"`
}