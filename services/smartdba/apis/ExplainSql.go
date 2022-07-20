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
    smartdba "github.com/lshuining/jdcloud-sdk-go/services/smartdba/models"
)

type ExplainSqlRequest struct {

    core.JDCloudRequest

    /* 地域代码  */
    RegionId string `json:"regionId"`

    /* 实例ID  */
    InstanceGid string `json:"instanceGid"`

    /* 数据库名 (Optional) */
    Database *string `json:"database"`

    /* SQL语句 (Optional) */
    Sql *string `json:"sql"`
}

/*
 * param regionId: 地域代码 (Required)
 * param instanceGid: 实例ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewExplainSqlRequest(
    regionId string,
    instanceGid string,
) *ExplainSqlRequest {

	return &ExplainSqlRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instance/{instanceGid}/explainSql",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        InstanceGid: instanceGid,
	}
}

/*
 * param regionId: 地域代码 (Required)
 * param instanceGid: 实例ID (Required)
 * param database: 数据库名 (Optional)
 * param sql: SQL语句 (Optional)
 */
func NewExplainSqlRequestWithAllParams(
    regionId string,
    instanceGid string,
    database *string,
    sql *string,
) *ExplainSqlRequest {

    return &ExplainSqlRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instance/{instanceGid}/explainSql",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        InstanceGid: instanceGid,
        Database: database,
        Sql: sql,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewExplainSqlRequestWithoutParam() *ExplainSqlRequest {

    return &ExplainSqlRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instance/{instanceGid}/explainSql",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域代码(Required) */
func (r *ExplainSqlRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceGid: 实例ID(Required) */
func (r *ExplainSqlRequest) SetInstanceGid(instanceGid string) {
    r.InstanceGid = instanceGid
}

/* param database: 数据库名(Optional) */
func (r *ExplainSqlRequest) SetDatabase(database string) {
    r.Database = &database
}

/* param sql: SQL语句(Optional) */
func (r *ExplainSqlRequest) SetSql(sql string) {
    r.Sql = &sql
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ExplainSqlRequest) GetRegionId() string {
    return r.RegionId
}

type ExplainSqlResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ExplainSqlResult `json:"result"`
}

type ExplainSqlResult struct {
    Data []smartdba.ExecutionPlan `json:"data"`
}