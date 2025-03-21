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
    gcs "github.com/jdcloud-api/jdcloud-sdk-go/services/gcs/models"
)

type DescribeOrderInfosRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* GCS系统资源实例ID (Optional) */
    InstanceId *string `json:"instanceId"`

    /* 主订单编号 (Optional) */
    OrderNumber *string `json:"orderNumber"`

    /* 子订单编号 (Optional) */
    SubOrderNumber *string `json:"subOrderNumber"`

    /* 订单类型：1-新购|2-续费|3-配置变更 (Optional) */
    OrderType *string `json:"orderType"`

    /* 主单打包id号 (Optional) */
    TaskId *string `json:"taskId"`
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeOrderInfosRequest(
    regionId string,
) *DescribeOrderInfosRequest {

	return &DescribeOrderInfosRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/describeOrderInfos",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
 * param instanceId: GCS系统资源实例ID (Optional)
 * param orderNumber: 主订单编号 (Optional)
 * param subOrderNumber: 子订单编号 (Optional)
 * param orderType: 订单类型：1-新购|2-续费|3-配置变更 (Optional)
 * param taskId: 主单打包id号 (Optional)
 */
func NewDescribeOrderInfosRequestWithAllParams(
    regionId string,
    instanceId *string,
    orderNumber *string,
    subOrderNumber *string,
    orderType *string,
    taskId *string,
) *DescribeOrderInfosRequest {

    return &DescribeOrderInfosRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/describeOrderInfos",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        OrderNumber: orderNumber,
        SubOrderNumber: subOrderNumber,
        OrderType: orderType,
        TaskId: taskId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeOrderInfosRequestWithoutParam() *DescribeOrderInfosRequest {

    return &DescribeOrderInfosRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/describeOrderInfos",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域(Required) */
func (r *DescribeOrderInfosRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param instanceId: GCS系统资源实例ID(Optional) */
func (r *DescribeOrderInfosRequest) SetInstanceId(instanceId string) {
    r.InstanceId = &instanceId
}
/* param orderNumber: 主订单编号(Optional) */
func (r *DescribeOrderInfosRequest) SetOrderNumber(orderNumber string) {
    r.OrderNumber = &orderNumber
}
/* param subOrderNumber: 子订单编号(Optional) */
func (r *DescribeOrderInfosRequest) SetSubOrderNumber(subOrderNumber string) {
    r.SubOrderNumber = &subOrderNumber
}
/* param orderType: 订单类型：1-新购|2-续费|3-配置变更(Optional) */
func (r *DescribeOrderInfosRequest) SetOrderType(orderType string) {
    r.OrderType = &orderType
}
/* param taskId: 主单打包id号(Optional) */
func (r *DescribeOrderInfosRequest) SetTaskId(taskId string) {
    r.TaskId = &taskId
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeOrderInfosRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeOrderInfosResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeOrderInfosResult `json:"result"`
}

type DescribeOrderInfosResult struct {
    List []gcs.OrderInfo `json:"list"`
}