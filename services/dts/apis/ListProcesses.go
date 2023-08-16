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
    dts "github.com/jdcloud-api/jdcloud-sdk-go/services/dts/models"
)

type ListProcessesRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》]  */
    RegionId string `json:"regionId"`

    /* 获取Process列表的资源ID  */
    Id string `json:"id"`

    /* 获取Process列表的资源类型 (Optional) */
    ResourceType *string `json:"resourceType"`

    /* 页码；默认为1，取值范围：-1，[1,∞)；-1时返回全部页码。 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认为10；取值范围[1, 100] (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》] (Required)
 * param id: 获取Process列表的资源ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewListProcessesRequest(
    regionId string,
    id string,
) *ListProcessesRequest {

	return &ListProcessesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/process",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        Id: id,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》] (Required)
 * param id: 获取Process列表的资源ID (Required)
 * param resourceType: 获取Process列表的资源类型 (Optional)
 * param pageNumber: 页码；默认为1，取值范围：-1，[1,∞)；-1时返回全部页码。 (Optional)
 * param pageSize: 分页大小；默认为10；取值范围[1, 100] (Optional)
 */
func NewListProcessesRequestWithAllParams(
    regionId string,
    id string,
    resourceType *string,
    pageNumber *int,
    pageSize *int,
) *ListProcessesRequest {

    return &ListProcessesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/process",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        Id: id,
        ResourceType: resourceType,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewListProcessesRequestWithoutParam() *ListProcessesRequest {

    return &ListProcessesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/process",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](Required) */
func (r *ListProcessesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param id: 获取Process列表的资源ID(Required) */
func (r *ListProcessesRequest) SetId(id string) {
    r.Id = id
}
/* param resourceType: 获取Process列表的资源类型(Optional) */
func (r *ListProcessesRequest) SetResourceType(resourceType string) {
    r.ResourceType = &resourceType
}
/* param pageNumber: 页码；默认为1，取值范围：-1，[1,∞)；-1时返回全部页码。(Optional) */
func (r *ListProcessesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}
/* param pageSize: 分页大小；默认为10；取值范围[1, 100](Optional) */
func (r *ListProcessesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ListProcessesRequest) GetRegionId() string {
    return r.RegionId
}

type ListProcessesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ListProcessesResult `json:"result"`
}

type ListProcessesResult struct {
    Process []dts.Process `json:"process"`
    TotalCount int `json:"totalCount"`
}