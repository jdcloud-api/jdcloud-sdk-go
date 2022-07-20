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
    jdfusion "github.com/lshuining/jdcloud-sdk-go/services/jdfusion/models"
)

type GetRdsDatabaseByInstIdAndDbNameRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* RDS实例ID  */
    InstId string `json:"instId"`

    /* 数据库名称  */
    DbName string `json:"dbName"`
}

/*
 * param regionId: 地域ID (Required)
 * param instId: RDS实例ID (Required)
 * param dbName: 数据库名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetRdsDatabaseByInstIdAndDbNameRequest(
    regionId string,
    instId string,
    dbName string,
) *GetRdsDatabaseByInstIdAndDbNameRequest {

	return &GetRdsDatabaseByInstIdAndDbNameRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/rds_instances/{instId}/databases/{dbName}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstId: instId,
        DbName: dbName,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param instId: RDS实例ID (Required)
 * param dbName: 数据库名称 (Required)
 */
func NewGetRdsDatabaseByInstIdAndDbNameRequestWithAllParams(
    regionId string,
    instId string,
    dbName string,
) *GetRdsDatabaseByInstIdAndDbNameRequest {

    return &GetRdsDatabaseByInstIdAndDbNameRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/rds_instances/{instId}/databases/{dbName}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstId: instId,
        DbName: dbName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetRdsDatabaseByInstIdAndDbNameRequestWithoutParam() *GetRdsDatabaseByInstIdAndDbNameRequest {

    return &GetRdsDatabaseByInstIdAndDbNameRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/rds_instances/{instId}/databases/{dbName}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *GetRdsDatabaseByInstIdAndDbNameRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instId: RDS实例ID(Required) */
func (r *GetRdsDatabaseByInstIdAndDbNameRequest) SetInstId(instId string) {
    r.InstId = instId
}

/* param dbName: 数据库名称(Required) */
func (r *GetRdsDatabaseByInstIdAndDbNameRequest) SetDbName(dbName string) {
    r.DbName = dbName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetRdsDatabaseByInstIdAndDbNameRequest) GetRegionId() string {
    return r.RegionId
}

type GetRdsDatabaseByInstIdAndDbNameResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetRdsDatabaseByInstIdAndDbNameResult `json:"result"`
}

type GetRdsDatabaseByInstIdAndDbNameResult struct {
    Db jdfusion.RdsDBInfo `json:"db"`
}