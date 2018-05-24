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
    "reflect"
    vpc "github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/models"
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type DescribeSubnetsRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 页码, 默认为1, 取值范围：[1,∞), 页码超过总页数时, 显示最后一页 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小，默认为20，取值范围：[10,500] (Optional) */
    PageSize *int `json:"pageSize"`

    /* subnetIds - subnet ID列表，支持多个
subnetNames - subnet名称列表，支持多个
routeTableId	- 子网关联路由表Id
aclId - 子网关联acl Id
vpcId - 子网所属VPC Id
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: Region ID 
 * param pageNumber: 页码, 默认为1, 取值范围：[1,∞), 页码超过总页数时, 显示最后一页 (Optional)
 * param pageSize: 分页大小，默认为20，取值范围：[10,500] (Optional)
 * param filters: subnetIds - subnet ID列表，支持多个
subnetNames - subnet名称列表，支持多个
routeTableId	- 子网关联路由表Id
aclId - 子网关联acl Id
vpcId - 子网所属VPC Id
 (Optional)
 */
func NewDescribeSubnetsRequest(
    regionId string,
) *DescribeSubnetsRequest {

	return &DescribeSubnetsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/subnets/",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

func (r *DescribeSubnetsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *DescribeSubnetsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

func (r *DescribeSubnetsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

func (r *DescribeSubnetsRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeSubnetsRequest) GetRegionId() string {
    fieldName := "RegionId"
    reqType := reflect.TypeOf(r)
    value := reflect.ValueOf(r)
    _, ok := reqType.FieldByName(fieldName)
    if ok {
        return value.FieldByName(fieldName).String()
    }

    return ""
}

type DescribeSubnetsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeSubnetsResult `json:"result"`
}

type DescribeSubnetsResult struct {
    Subnets []vpc.Subnet `json:"subnets"`
    TotalCount int `json:"totalCount"`
}