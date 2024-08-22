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

type DescribeInvoiceRequest struct {

    core.JDCloudRequest

    /* 地域编码，参考OpenAPI公共说明  */
    RegionId string `json:"regionId"`

    /* 发票申请单号  */
    InvoiceNumber string `json:"invoiceNumber"`

    /* 页码(默认1)  */
    PageIndex int `json:"pageIndex"`

    /* 每页展示数据量(默认：20)  */
    PageSize int `json:"pageSize"`

    /* 发票类型 1详情 2按月 (Optional) */
    InvoiceType *int `json:"invoiceType"`

    /* 用户pin (Optional) */
    Pin *string `json:"pin"`
}

/*
 * param regionId: 地域编码，参考OpenAPI公共说明 (Required)
 * param invoiceNumber: 发票申请单号 (Required)
 * param pageIndex: 页码(默认1) (Required)
 * param pageSize: 每页展示数据量(默认：20) (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeInvoiceRequest(
    regionId string,
    invoiceNumber string,
    pageIndex int,
    pageSize int,
) *DescribeInvoiceRequest {

	return &DescribeInvoiceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/invoice/{invoiceNumber}",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        InvoiceNumber: invoiceNumber,
        PageIndex: pageIndex,
        PageSize: pageSize,
	}
}

/*
 * param regionId: 地域编码，参考OpenAPI公共说明 (Required)
 * param invoiceNumber: 发票申请单号 (Required)
 * param pageIndex: 页码(默认1) (Required)
 * param pageSize: 每页展示数据量(默认：20) (Required)
 * param invoiceType: 发票类型 1详情 2按月 (Optional)
 * param pin: 用户pin (Optional)
 */
func NewDescribeInvoiceRequestWithAllParams(
    regionId string,
    invoiceNumber string,
    pageIndex int,
    pageSize int,
    invoiceType *int,
    pin *string,
) *DescribeInvoiceRequest {

    return &DescribeInvoiceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/invoice/{invoiceNumber}",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        InvoiceNumber: invoiceNumber,
        PageIndex: pageIndex,
        PageSize: pageSize,
        InvoiceType: invoiceType,
        Pin: pin,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeInvoiceRequestWithoutParam() *DescribeInvoiceRequest {

    return &DescribeInvoiceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/invoice/{invoiceNumber}",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域编码，参考OpenAPI公共说明(Required) */
func (r *DescribeInvoiceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param invoiceNumber: 发票申请单号(Required) */
func (r *DescribeInvoiceRequest) SetInvoiceNumber(invoiceNumber string) {
    r.InvoiceNumber = invoiceNumber
}
/* param pageIndex: 页码(默认1)(Required) */
func (r *DescribeInvoiceRequest) SetPageIndex(pageIndex int) {
    r.PageIndex = pageIndex
}
/* param pageSize: 每页展示数据量(默认：20)(Required) */
func (r *DescribeInvoiceRequest) SetPageSize(pageSize int) {
    r.PageSize = pageSize
}
/* param invoiceType: 发票类型 1详情 2按月(Optional) */
func (r *DescribeInvoiceRequest) SetInvoiceType(invoiceType int) {
    r.InvoiceType = &invoiceType
}
/* param pin: 用户pin(Optional) */
func (r *DescribeInvoiceRequest) SetPin(pin string) {
    r.Pin = &pin
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeInvoiceRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeInvoiceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeInvoiceResult `json:"result"`
}

type DescribeInvoiceResult struct {
    Id int `json:"id"`
    InvoiceNumber string `json:"invoiceNumber"`
    Pin string `json:"pin"`
    InvoiceType string `json:"invoiceType"`
    InvoiceTitle string `json:"invoiceTitle"`
    TaxNo string `json:"taxNo"`
    TotalPrice int `json:"totalPrice"`
    Remark string `json:"remark"`
    InvoiceContent string `json:"invoiceContent"`
    UserType int `json:"userType"`
    AddressId int `json:"addressId"`
    TransportId int `json:"transportId"`
    CreateTime string `json:"createTime"`
    UpdateTime string `json:"updateTime"`
    Status string `json:"status"`
    BillingTime string `json:"billingTime"`
    Reason string `json:"reason"`
    InvoiceOrg string `json:"invoiceOrg"`
    RegisterAddress string `json:"registerAddress"`
    RegisterPhone string `json:"registerPhone"`
    AccountBank string `json:"accountBank"`
    Account string `json:"account"`
    MediumType string `json:"mediumType"`
    PdfUrl string `json:"pdfUrl"`
    XmlUrl string `json:"xmlUrl"`
    OfdUrl string `json:"ofdUrl"`
    Logistics invoice.Logistics `json:"logistics"`
    PostAddress invoice.PostAddress `json:"postAddress"`
    InvoiceOrders []invoice.InvoiceOrder `json:"invoiceOrders"`
    MonthGroupOrders []invoice.MonthGroupOrder `json:"monthGroupOrders"`
    IsHistoryData int `json:"isHistoryData"`
}