// Copyright 2018-2025 JDCLOUD.COM
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
    . "github.com/jdcloud-api/jdcloud-sdk-go/core"
    "reflect"
    mongodb "github.com/jdcloud-api/jdcloud-sdk-go/services/mongodb/models"
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type DescribeInstancesRequest struct {

    JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 页码；默认为1，取值范围：[1,∞) (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认为10；取值范围[1, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* instanceId - 实例ID, 精确匹配
instanceName - 实例名称, 模糊匹配
instanceStatus - mongodb状态，精确匹配，支持多个.RUNNING：运行, ERROR：错误 ,BUILDING：创建中, DELETING：删除中, RESTORING：恢复中, RESIZING：变配中
 (Optional) */
    Filters []common.Filter `json:"filters"`

    /* createTime - 创建时间,asc（正序），desc（倒序）
 (Optional) */
    Sorts []common.Sort `json:"sorts"`
}

/*
 * param regionId: Region ID 
 * param pageNumber: 页码；默认为1，取值范围：[1,∞) (Optional)
 * param pageSize: 分页大小；默认为10；取值范围[1, 100] (Optional)
 * param filters: instanceId - 实例ID, 精确匹配
instanceName - 实例名称, 模糊匹配
instanceStatus - mongodb状态，精确匹配，支持多个.RUNNING：运行, ERROR：错误 ,BUILDING：创建中, DELETING：删除中, RESTORING：恢复中, RESIZING：变配中
 (Optional)
 * param sorts: createTime - 创建时间,asc（正序），desc（倒序）
 (Optional)
 */
func NewDescribeInstancesRequest(
    regionId string,
) *DescribeInstancesRequest {

	return &DescribeInstancesRequest{
        JDCloudRequest: JDCloudRequest{
			URL:     "/regions/{regionId}/instances",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

func (r *DescribeInstancesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *DescribeInstancesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

func (r *DescribeInstancesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

func (r *DescribeInstancesRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

func (r *DescribeInstancesRequest) SetSorts(sorts []common.Sort) {
    r.Sorts = sorts
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeInstancesRequest) GetRegionId() string {
    fieldName := "RegionId"
    reqType := reflect.TypeOf(r)
    value := reflect.ValueOf(r)
    _, ok := reqType.FieldByName(fieldName)
    if ok {
        return value.FieldByName(fieldName).String()
    }

    return ""
}

type DescribeInstancesResponse struct {
    RequestID string `json:"requestId"`
    Error ErrorResponse `json:"error"`
    Result DescribeInstancesResult `json:"result"`
}

type DescribeInstancesResult struct {
    DbInstances []mongodb.DBInstance `json:"dbInstances"`
    TotalCount int `json:"totalCount"`
    PageNumber int `json:"pageNumber"`
}