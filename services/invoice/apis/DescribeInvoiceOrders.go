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
    invoice "github.com/jdcloud-api/jdcloud-sdk-go/services/invoice/models"
)

type DescribeInvoiceOrdersRequest struct {

    core.JDCloudRequest

    /* 地域编码，参考OpenAPI公共说明  */
    RegionId string `json:"regionId"`

    /* 开始时间 (Optional) */
    SearchStartTime *string `json:"searchStartTime"`

    /* 结束时间 (Optional) */
    SearchEndTime *string `json:"searchEndTime"`

    /* 页码 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 页大小 (Optional) */
    PageSize *int `json:"pageSize"`

    /* 单据类型[订单-order，账单-bill 月结算单-month] (Optional) */
    ReceiptType *string `json:"receiptType"`

    /* 交易(支付)类型 1-代付 2-自付 (Optional) */
    PayType *int `json:"payType"`
}

/*
 * param regionId: 地域编码，参考OpenAPI公共说明 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeInvoiceOrdersRequest(
    regionId string,
) *DescribeInvoiceOrdersRequest {

	return &DescribeInvoiceOrdersRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/invoiceorder:list",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域编码，参考OpenAPI公共说明 (Required)
 * param searchStartTime: 开始时间 (Optional)
 * param searchEndTime: 结束时间 (Optional)
 * param pageNumber: 页码 (Optional)
 * param pageSize: 页大小 (Optional)
 * param receiptType: 单据类型[订单-order，账单-bill 月结算单-month] (Optional)
 * param payType: 交易(支付)类型 1-代付 2-自付 (Optional)
 */
func NewDescribeInvoiceOrdersRequestWithAllParams(
    regionId string,
    searchStartTime *string,
    searchEndTime *string,
    pageNumber *int,
    pageSize *int,
    receiptType *string,
    payType *int,
) *DescribeInvoiceOrdersRequest {

    return &DescribeInvoiceOrdersRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/invoiceorder:list",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        SearchStartTime: searchStartTime,
        SearchEndTime: searchEndTime,
        PageNumber: pageNumber,
        PageSize: pageSize,
        ReceiptType: receiptType,
        PayType: payType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeInvoiceOrdersRequestWithoutParam() *DescribeInvoiceOrdersRequest {

    return &DescribeInvoiceOrdersRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/invoiceorder:list",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域编码，参考OpenAPI公共说明(Required) */
func (r *DescribeInvoiceOrdersRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param searchStartTime: 开始时间(Optional) */
func (r *DescribeInvoiceOrdersRequest) SetSearchStartTime(searchStartTime string) {
    r.SearchStartTime = &searchStartTime
}
/* param searchEndTime: 结束时间(Optional) */
func (r *DescribeInvoiceOrdersRequest) SetSearchEndTime(searchEndTime string) {
    r.SearchEndTime = &searchEndTime
}
/* param pageNumber: 页码(Optional) */
func (r *DescribeInvoiceOrdersRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}
/* param pageSize: 页大小(Optional) */
func (r *DescribeInvoiceOrdersRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}
/* param receiptType: 单据类型[订单-order，账单-bill 月结算单-month](Optional) */
func (r *DescribeInvoiceOrdersRequest) SetReceiptType(receiptType string) {
    r.ReceiptType = &receiptType
}
/* param payType: 交易(支付)类型 1-代付 2-自付(Optional) */
func (r *DescribeInvoiceOrdersRequest) SetPayType(payType int) {
    r.PayType = &payType
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeInvoiceOrdersRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeInvoiceOrdersResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeInvoiceOrdersResult `json:"result"`
}

type DescribeInvoiceOrdersResult struct {
    EnableInvoiceFee int `json:"enableInvoiceFee"`
    EnableAllInvoiceFee int `json:"enableAllInvoiceFee"`
    OweInvoiceFee int `json:"oweInvoiceFee"`
    IsSetInvoiceMSGTemplate bool `json:"isSetInvoiceMSGTemplate"`
    InvoiceOrderList invoice.InvoiceOrderItem `json:"invoiceOrderList"`
}