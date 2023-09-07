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
    rds "github.com/jdcloud-api/jdcloud-sdk-go/services/rds/models"
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type DescribeAuditResultRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* RDS 实例ID，唯一标识一个RDS实例  */
    InstanceId string `json:"instanceId"`

    /* 查询开始时间，格式为：YYYY-MM-DD HH:mm:ss，开始时间不能早于当前时间3天  */
    StartTime string `json:"startTime"`

    /* 查询截止时间，格式为：YYYY-MM-DD HH:mm:ss，开始时间到结束时间不能超过3天  */
    EndTime string `json:"endTime"`

    /* 废弃，使用filter，数据库名 (Optional) */
    DbName *string `json:"dbName"`

    /* 废弃，使用filter，账号名 (Optional) */
    AccountName *string `json:"accountName"`

    /* 显示数据的页码，默认为1，取值范围：[-1,∞) (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 每页显示的数据条数，默认为10，取值范围：[10,100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 过滤参数，多个过滤参数之间的关系为“与”(and支持以下属性的过滤(默认等值)：)
- operation：仅第一个value生效，语句类型【create/alter/drop/truncate/select/insert/update/delete/replace/ddl/dml/disconnect/connect/failed_connect/query】,operator仅支持eq或者in
- account：实例账号名，operator仅支持eq或者in
- keyword：SQL 关键词，模糊查询，operator仅支持eq或者in
- database：实例库名，operator仅支持eq或者in
- threadId：会话id，operator仅支持eq或者in
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: RDS 实例ID，唯一标识一个RDS实例 (Required)
 * param startTime: 查询开始时间，格式为：YYYY-MM-DD HH:mm:ss，开始时间不能早于当前时间3天 (Required)
 * param endTime: 查询截止时间，格式为：YYYY-MM-DD HH:mm:ss，开始时间到结束时间不能超过3天 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeAuditResultRequest(
    regionId string,
    instanceId string,
    startTime string,
    endTime string,
) *DescribeAuditResultRequest {

	return &DescribeAuditResultRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/audit:describeAuditResult",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        StartTime: startTime,
        EndTime: endTime,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: RDS 实例ID，唯一标识一个RDS实例 (Required)
 * param startTime: 查询开始时间，格式为：YYYY-MM-DD HH:mm:ss，开始时间不能早于当前时间3天 (Required)
 * param endTime: 查询截止时间，格式为：YYYY-MM-DD HH:mm:ss，开始时间到结束时间不能超过3天 (Required)
 * param dbName: 废弃，使用filter，数据库名 (Optional)
 * param accountName: 废弃，使用filter，账号名 (Optional)
 * param pageNumber: 显示数据的页码，默认为1，取值范围：[-1,∞) (Optional)
 * param pageSize: 每页显示的数据条数，默认为10，取值范围：[10,100] (Optional)
 * param filters: 过滤参数，多个过滤参数之间的关系为“与”(and支持以下属性的过滤(默认等值)：)
- operation：仅第一个value生效，语句类型【create/alter/drop/truncate/select/insert/update/delete/replace/ddl/dml/disconnect/connect/failed_connect/query】,operator仅支持eq或者in
- account：实例账号名，operator仅支持eq或者in
- keyword：SQL 关键词，模糊查询，operator仅支持eq或者in
- database：实例库名，operator仅支持eq或者in
- threadId：会话id，operator仅支持eq或者in
 (Optional)
 */
func NewDescribeAuditResultRequestWithAllParams(
    regionId string,
    instanceId string,
    startTime string,
    endTime string,
    dbName *string,
    accountName *string,
    pageNumber *int,
    pageSize *int,
    filters []common.Filter,
) *DescribeAuditResultRequest {

    return &DescribeAuditResultRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/audit:describeAuditResult",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        StartTime: startTime,
        EndTime: endTime,
        DbName: dbName,
        AccountName: accountName,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeAuditResultRequestWithoutParam() *DescribeAuditResultRequest {

    return &DescribeAuditResultRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/audit:describeAuditResult",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *DescribeAuditResultRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param instanceId: RDS 实例ID，唯一标识一个RDS实例(Required) */
func (r *DescribeAuditResultRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}
/* param startTime: 查询开始时间，格式为：YYYY-MM-DD HH:mm:ss，开始时间不能早于当前时间3天(Required) */
func (r *DescribeAuditResultRequest) SetStartTime(startTime string) {
    r.StartTime = startTime
}
/* param endTime: 查询截止时间，格式为：YYYY-MM-DD HH:mm:ss，开始时间到结束时间不能超过3天(Required) */
func (r *DescribeAuditResultRequest) SetEndTime(endTime string) {
    r.EndTime = endTime
}
/* param dbName: 废弃，使用filter，数据库名(Optional) */
func (r *DescribeAuditResultRequest) SetDbName(dbName string) {
    r.DbName = &dbName
}
/* param accountName: 废弃，使用filter，账号名(Optional) */
func (r *DescribeAuditResultRequest) SetAccountName(accountName string) {
    r.AccountName = &accountName
}
/* param pageNumber: 显示数据的页码，默认为1，取值范围：[-1,∞)(Optional) */
func (r *DescribeAuditResultRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}
/* param pageSize: 每页显示的数据条数，默认为10，取值范围：[10,100](Optional) */
func (r *DescribeAuditResultRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}
/* param filters: 过滤参数，多个过滤参数之间的关系为“与”(and支持以下属性的过滤(默认等值)：)
- operation：仅第一个value生效，语句类型【create/alter/drop/truncate/select/insert/update/delete/replace/ddl/dml/disconnect/connect/failed_connect/query】,operator仅支持eq或者in
- account：实例账号名，operator仅支持eq或者in
- keyword：SQL 关键词，模糊查询，operator仅支持eq或者in
- database：实例库名，operator仅支持eq或者in
- threadId：会话id，operator仅支持eq或者in
(Optional) */
func (r *DescribeAuditResultRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeAuditResultRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeAuditResultResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeAuditResultResult `json:"result"`
}

type DescribeAuditResultResult struct {
    AuditResult []rds.AuditResult `json:"auditResult"`
    TotalCount int `json:"totalCount"`
}