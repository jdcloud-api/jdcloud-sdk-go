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
    jdccs "github.com/lshuining/jdcloud-sdk-go/services/jdccs/models"
    common "github.com/lshuining/jdcloud-sdk-go/services/common/models"
)

type DescribeDevicesRequest struct {

    core.JDCloudRequest

    /* IDC机房ID  */
    Idc string `json:"idc"`

    /* 页码, 默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小，默认为20 (Optional) */
    PageSize *int `json:"pageSize"`

    /* 是否查询全部，默认分页 (Optional) */
    All *int `json:"all"`

    /* 机柜ID (Optional) */
    CabinetId *string `json:"cabinetId"`

    /* 设备类型 server:服务器 network:网络设备 storage:存储设备 other:其他设备 (Optional) */
    DeviceType *string `json:"deviceType"`

    /* 资产状态 launched:已上架 opened:已开通 canceling:退订中 operating:操作中 modifing:变更中 (Optional) */
    AssetStatus *string `json:"assetStatus"`

    /* 资产归属 own:自备 lease:租赁 (Optional) */
    AssetBelong *string `json:"assetBelong"`

    /* 设备编码 (Optional) */
    DeviceNo *string `json:"deviceNo"`

    /* 设备SN号 (Optional) */
    SnNo *string `json:"snNo"`

    /* 机柜编码 (Optional) */
    CabinetNo *string `json:"cabinetNo"`

    /* 工单模板CODE (Optional) */
    TicketTemplateCode *string `json:"ticketTemplateCode"`

    /* deviceId - 设备实例ID，精确匹配，支持多个
snNo - 设备SN号，精确匹配，支持多个
deviceNo - 设备编码，精确匹配，支持多个
cabinetNo - 机柜编码，精确匹配，支持多个
 (Optional) */
    Filters []common.Filter `json:"filters"`

    /* deviceNo - 设备编码 cabinetNo - 机柜编码 (Optional) */
    Sorts []common.Sort `json:"sorts"`
}

/*
 * param idc: IDC机房ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeDevicesRequest(
    idc string,
) *DescribeDevicesRequest {

	return &DescribeDevicesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/idcs/{idc}/devices",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        Idc: idc,
	}
}

/*
 * param idc: IDC机房ID (Required)
 * param pageNumber: 页码, 默认为1 (Optional)
 * param pageSize: 分页大小，默认为20 (Optional)
 * param all: 是否查询全部，默认分页 (Optional)
 * param cabinetId: 机柜ID (Optional)
 * param deviceType: 设备类型 server:服务器 network:网络设备 storage:存储设备 other:其他设备 (Optional)
 * param assetStatus: 资产状态 launched:已上架 opened:已开通 canceling:退订中 operating:操作中 modifing:变更中 (Optional)
 * param assetBelong: 资产归属 own:自备 lease:租赁 (Optional)
 * param deviceNo: 设备编码 (Optional)
 * param snNo: 设备SN号 (Optional)
 * param cabinetNo: 机柜编码 (Optional)
 * param ticketTemplateCode: 工单模板CODE (Optional)
 * param filters: deviceId - 设备实例ID，精确匹配，支持多个
snNo - 设备SN号，精确匹配，支持多个
deviceNo - 设备编码，精确匹配，支持多个
cabinetNo - 机柜编码，精确匹配，支持多个
 (Optional)
 * param sorts: deviceNo - 设备编码 cabinetNo - 机柜编码 (Optional)
 */
func NewDescribeDevicesRequestWithAllParams(
    idc string,
    pageNumber *int,
    pageSize *int,
    all *int,
    cabinetId *string,
    deviceType *string,
    assetStatus *string,
    assetBelong *string,
    deviceNo *string,
    snNo *string,
    cabinetNo *string,
    ticketTemplateCode *string,
    filters []common.Filter,
    sorts []common.Sort,
) *DescribeDevicesRequest {

    return &DescribeDevicesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/idcs/{idc}/devices",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        Idc: idc,
        PageNumber: pageNumber,
        PageSize: pageSize,
        All: all,
        CabinetId: cabinetId,
        DeviceType: deviceType,
        AssetStatus: assetStatus,
        AssetBelong: assetBelong,
        DeviceNo: deviceNo,
        SnNo: snNo,
        CabinetNo: cabinetNo,
        TicketTemplateCode: ticketTemplateCode,
        Filters: filters,
        Sorts: sorts,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeDevicesRequestWithoutParam() *DescribeDevicesRequest {

    return &DescribeDevicesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/idcs/{idc}/devices",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param idc: IDC机房ID(Required) */
func (r *DescribeDevicesRequest) SetIdc(idc string) {
    r.Idc = idc
}

/* param pageNumber: 页码, 默认为1(Optional) */
func (r *DescribeDevicesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小，默认为20(Optional) */
func (r *DescribeDevicesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param all: 是否查询全部，默认分页(Optional) */
func (r *DescribeDevicesRequest) SetAll(all int) {
    r.All = &all
}

/* param cabinetId: 机柜ID(Optional) */
func (r *DescribeDevicesRequest) SetCabinetId(cabinetId string) {
    r.CabinetId = &cabinetId
}

/* param deviceType: 设备类型 server:服务器 network:网络设备 storage:存储设备 other:其他设备(Optional) */
func (r *DescribeDevicesRequest) SetDeviceType(deviceType string) {
    r.DeviceType = &deviceType
}

/* param assetStatus: 资产状态 launched:已上架 opened:已开通 canceling:退订中 operating:操作中 modifing:变更中(Optional) */
func (r *DescribeDevicesRequest) SetAssetStatus(assetStatus string) {
    r.AssetStatus = &assetStatus
}

/* param assetBelong: 资产归属 own:自备 lease:租赁(Optional) */
func (r *DescribeDevicesRequest) SetAssetBelong(assetBelong string) {
    r.AssetBelong = &assetBelong
}

/* param deviceNo: 设备编码(Optional) */
func (r *DescribeDevicesRequest) SetDeviceNo(deviceNo string) {
    r.DeviceNo = &deviceNo
}

/* param snNo: 设备SN号(Optional) */
func (r *DescribeDevicesRequest) SetSnNo(snNo string) {
    r.SnNo = &snNo
}

/* param cabinetNo: 机柜编码(Optional) */
func (r *DescribeDevicesRequest) SetCabinetNo(cabinetNo string) {
    r.CabinetNo = &cabinetNo
}

/* param ticketTemplateCode: 工单模板CODE(Optional) */
func (r *DescribeDevicesRequest) SetTicketTemplateCode(ticketTemplateCode string) {
    r.TicketTemplateCode = &ticketTemplateCode
}

/* param filters: deviceId - 设备实例ID，精确匹配，支持多个
snNo - 设备SN号，精确匹配，支持多个
deviceNo - 设备编码，精确匹配，支持多个
cabinetNo - 机柜编码，精确匹配，支持多个
(Optional) */
func (r *DescribeDevicesRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

/* param sorts: deviceNo - 设备编码 cabinetNo - 机柜编码(Optional) */
func (r *DescribeDevicesRequest) SetSorts(sorts []common.Sort) {
    r.Sorts = sorts
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeDevicesRequest) GetRegionId() string {
    return ""
}

type DescribeDevicesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeDevicesResult `json:"result"`
}

type DescribeDevicesResult struct {
    Devices []jdccs.DescribeDevice `json:"devices"`
    PageNumber int `json:"pageNumber"`
    PageSize int `json:"pageSize"`
    TotalCount int `json:"totalCount"`
}