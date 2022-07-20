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
    dms "github.com/lshuining/jdcloud-sdk-go/services/dms/models"
)

type ExportStructRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* 数据源id (Optional) */
    DataSourceId *int `json:"dataSourceId"`

    /* 数据库名称 (Optional) */
    DbName *string `json:"dbName"`

    /* 导出表名，为空时导出库中所有表 (Optional) */
    TableName *string `json:"tableName"`

    /* 导出文件字符编码 (Optional) */
    Charset *string `json:"charset"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewExportStructRequest(
    regionId string,
) *ExportStructRequest {

	return &ExportStructRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/console:exportStruct",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param dataSourceId: 数据源id (Optional)
 * param dbName: 数据库名称 (Optional)
 * param tableName: 导出表名，为空时导出库中所有表 (Optional)
 * param charset: 导出文件字符编码 (Optional)
 */
func NewExportStructRequestWithAllParams(
    regionId string,
    dataSourceId *int,
    dbName *string,
    tableName *string,
    charset *string,
) *ExportStructRequest {

    return &ExportStructRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/console:exportStruct",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DataSourceId: dataSourceId,
        DbName: dbName,
        TableName: tableName,
        Charset: charset,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewExportStructRequestWithoutParam() *ExportStructRequest {

    return &ExportStructRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/console:exportStruct",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *ExportStructRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param dataSourceId: 数据源id(Optional) */
func (r *ExportStructRequest) SetDataSourceId(dataSourceId int) {
    r.DataSourceId = &dataSourceId
}

/* param dbName: 数据库名称(Optional) */
func (r *ExportStructRequest) SetDbName(dbName string) {
    r.DbName = &dbName
}

/* param tableName: 导出表名，为空时导出库中所有表(Optional) */
func (r *ExportStructRequest) SetTableName(tableName string) {
    r.TableName = &tableName
}

/* param charset: 导出文件字符编码(Optional) */
func (r *ExportStructRequest) SetCharset(charset string) {
    r.Charset = &charset
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ExportStructRequest) GetRegionId() string {
    return r.RegionId
}

type ExportStructResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ExportStructResult `json:"result"`
}

type ExportStructResult struct {
    TableStructs []dms.TableStruct `json:"tableStructs"`
}