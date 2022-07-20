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
    cps "github.com/lshuining/jdcloud-sdk-go/services/cps/models"
    common "github.com/lshuining/jdcloud-sdk-go/services/common/models"
)

type DescribeIpv6AddressesRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* 页码；默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认为20；取值范围[20, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* IPv6网关ID (Optional) */
    Ipv6GatewayId *string `json:"ipv6GatewayId"`

    /* IPv6地址 (Optional) */
    Ipv6Address *string `json:"ipv6Address"`

    /* 是否已开通公网 (Optional) */
    EnableInternet *bool `json:"enableInternet"`

    /* ipv6AddressId - IPv6地址ID，精确匹配，支持多个
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeIpv6AddressesRequest(
    regionId string,
) *DescribeIpv6AddressesRequest {

	return &DescribeIpv6AddressesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/ipv6Addresses",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
 * param pageNumber: 页码；默认为1 (Optional)
 * param pageSize: 分页大小；默认为20；取值范围[20, 100] (Optional)
 * param ipv6GatewayId: IPv6网关ID (Optional)
 * param ipv6Address: IPv6地址 (Optional)
 * param enableInternet: 是否已开通公网 (Optional)
 * param filters: ipv6AddressId - IPv6地址ID，精确匹配，支持多个
 (Optional)
 */
func NewDescribeIpv6AddressesRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    ipv6GatewayId *string,
    ipv6Address *string,
    enableInternet *bool,
    filters []common.Filter,
) *DescribeIpv6AddressesRequest {

    return &DescribeIpv6AddressesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/ipv6Addresses",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Ipv6GatewayId: ipv6GatewayId,
        Ipv6Address: ipv6Address,
        EnableInternet: enableInternet,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeIpv6AddressesRequestWithoutParam() *DescribeIpv6AddressesRequest {

    return &DescribeIpv6AddressesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/ipv6Addresses",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域(Required) */
func (r *DescribeIpv6AddressesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 页码；默认为1(Optional) */
func (r *DescribeIpv6AddressesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小；默认为20；取值范围[20, 100](Optional) */
func (r *DescribeIpv6AddressesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param ipv6GatewayId: IPv6网关ID(Optional) */
func (r *DescribeIpv6AddressesRequest) SetIpv6GatewayId(ipv6GatewayId string) {
    r.Ipv6GatewayId = &ipv6GatewayId
}

/* param ipv6Address: IPv6地址(Optional) */
func (r *DescribeIpv6AddressesRequest) SetIpv6Address(ipv6Address string) {
    r.Ipv6Address = &ipv6Address
}

/* param enableInternet: 是否已开通公网(Optional) */
func (r *DescribeIpv6AddressesRequest) SetEnableInternet(enableInternet bool) {
    r.EnableInternet = &enableInternet
}

/* param filters: ipv6AddressId - IPv6地址ID，精确匹配，支持多个
(Optional) */
func (r *DescribeIpv6AddressesRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeIpv6AddressesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeIpv6AddressesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeIpv6AddressesResult `json:"result"`
}

type DescribeIpv6AddressesResult struct {
    Ipv6Addresses []cps.Ipv6Address `json:"ipv6Addresses"`
    PageNumber int `json:"pageNumber"`
    PageSize int `json:"pageSize"`
    TotalCount int `json:"totalCount"`
}