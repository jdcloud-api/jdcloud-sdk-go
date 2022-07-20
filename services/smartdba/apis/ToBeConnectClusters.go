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

type ToBeConnectClustersRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* 实例类型，支持MySQL/Percona，默认MySQL (Optional) */
    DbType *string `json:"dbType"`

    /* 显示数据的页码，默认为1，取值范围：[-1,∞)。pageIndex 为-1时，返回所有数据页码； (Optional) */
    PageIndex *int `json:"pageIndex"`

    /* 每页显示的数据条数，默认为10，取值范围：[1,100]，用于查询列表的接口 (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewToBeConnectClustersRequest(
    regionId string,
) *ToBeConnectClustersRequest {

	return &ToBeConnectClustersRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/waitcutin",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param dbType: 实例类型，支持MySQL/Percona，默认MySQL (Optional)
 * param pageIndex: 显示数据的页码，默认为1，取值范围：[-1,∞)。pageIndex 为-1时，返回所有数据页码； (Optional)
 * param pageSize: 每页显示的数据条数，默认为10，取值范围：[1,100]，用于查询列表的接口 (Optional)
 */
func NewToBeConnectClustersRequestWithAllParams(
    regionId string,
    dbType *string,
    pageIndex *int,
    pageSize *int,
) *ToBeConnectClustersRequest {

    return &ToBeConnectClustersRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/waitcutin",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        DbType: dbType,
        PageIndex: pageIndex,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewToBeConnectClustersRequestWithoutParam() *ToBeConnectClustersRequest {

    return &ToBeConnectClustersRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/waitcutin",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *ToBeConnectClustersRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param dbType: 实例类型，支持MySQL/Percona，默认MySQL(Optional) */
func (r *ToBeConnectClustersRequest) SetDbType(dbType string) {
    r.DbType = &dbType
}

/* param pageIndex: 显示数据的页码，默认为1，取值范围：[-1,∞)。pageIndex 为-1时，返回所有数据页码；(Optional) */
func (r *ToBeConnectClustersRequest) SetPageIndex(pageIndex int) {
    r.PageIndex = &pageIndex
}

/* param pageSize: 每页显示的数据条数，默认为10，取值范围：[1,100]，用于查询列表的接口(Optional) */
func (r *ToBeConnectClustersRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ToBeConnectClustersRequest) GetRegionId() string {
    return r.RegionId
}

type ToBeConnectClustersResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ToBeConnectClustersResult `json:"result"`
}

type ToBeConnectClustersResult struct {
    TotalCount int `json:"totalCount"`
    Data []smartdba.InstancesInfo `json:"data"`
}