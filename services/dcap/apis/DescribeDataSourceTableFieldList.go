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
    dcap "github.com/lshuining/jdcloud-sdk-go/services/dcap/models"
)

type DescribeDataSourceTableFieldListRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 数据源 ID  */
    DataSourceId string `json:"dataSourceId"`

    /* 表-名称  */
    TableName string `json:"tableName"`

    /* 页码；默认为1  */
    PageNumber int `json:"pageNumber"`

    /* 分页大小；默认为10；取值范围[10, 100]  */
    PageSize int `json:"pageSize"`
}

/*
 * param regionId: Region ID (Required)
 * param dataSourceId: 数据源 ID (Required)
 * param tableName: 表-名称 (Required)
 * param pageNumber: 页码；默认为1 (Required)
 * param pageSize: 分页大小；默认为10；取值范围[10, 100] (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeDataSourceTableFieldListRequest(
    regionId string,
    dataSourceId string,
    tableName string,
    pageNumber int,
    pageSize int,
) *DescribeDataSourceTableFieldListRequest {

	return &DescribeDataSourceTableFieldListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/dataSources/{dataSourceId}/tables/{tableName}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        DataSourceId: dataSourceId,
        TableName: tableName,
        PageNumber: pageNumber,
        PageSize: pageSize,
	}
}

/*
 * param regionId: Region ID (Required)
 * param dataSourceId: 数据源 ID (Required)
 * param tableName: 表-名称 (Required)
 * param pageNumber: 页码；默认为1 (Required)
 * param pageSize: 分页大小；默认为10；取值范围[10, 100] (Required)
 */
func NewDescribeDataSourceTableFieldListRequestWithAllParams(
    regionId string,
    dataSourceId string,
    tableName string,
    pageNumber int,
    pageSize int,
) *DescribeDataSourceTableFieldListRequest {

    return &DescribeDataSourceTableFieldListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dataSources/{dataSourceId}/tables/{tableName}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DataSourceId: dataSourceId,
        TableName: tableName,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeDataSourceTableFieldListRequestWithoutParam() *DescribeDataSourceTableFieldListRequest {

    return &DescribeDataSourceTableFieldListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/dataSources/{dataSourceId}/tables/{tableName}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeDataSourceTableFieldListRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param dataSourceId: 数据源 ID(Required) */
func (r *DescribeDataSourceTableFieldListRequest) SetDataSourceId(dataSourceId string) {
    r.DataSourceId = dataSourceId
}

/* param tableName: 表-名称(Required) */
func (r *DescribeDataSourceTableFieldListRequest) SetTableName(tableName string) {
    r.TableName = tableName
}

/* param pageNumber: 页码；默认为1(Required) */
func (r *DescribeDataSourceTableFieldListRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = pageNumber
}

/* param pageSize: 分页大小；默认为10；取值范围[10, 100](Required) */
func (r *DescribeDataSourceTableFieldListRequest) SetPageSize(pageSize int) {
    r.PageSize = pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeDataSourceTableFieldListRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeDataSourceTableFieldListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeDataSourceTableFieldListResult `json:"result"`
}

type DescribeDataSourceTableFieldListResult struct {
    List []dcap.FieldDesc `json:"list"`
    TotalCount int `json:"totalCount"`
}