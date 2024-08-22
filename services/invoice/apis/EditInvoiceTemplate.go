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
)

type EditInvoiceTemplateRequest struct {

    core.JDCloudRequest

    /* 地域编码，参考OpenAPI公共说明  */
    RegionId string `json:"regionId"`

    /* 发票类型:[个人增值税普通发票-Personal_VAT_Ordinary_Invoice,企业增值税专用发票-Enterprise_VAT_Special_Invoice,企业增值税普通发票-Enterprise_VAT_Ordinary_Invoice] (Optional) */
    InvoiceType *string `json:"invoiceType"`

    /* 发票抬头 (Optional) */
    InvoiceTitle *string `json:"invoiceTitle"`

    /* 开票内容（按类别开票/按明细开票） (Optional) */
    InvoiceContent *string `json:"invoiceContent"`

    /* 纳税人识别码（发票类型为企业增值税专用发票和企业增值税普通发票时必填） (Optional) */
    TaxId *string `json:"taxId"`

    /* 单位名称（发票类型为企业增值税专用发票时必填） (Optional) */
    Company *string `json:"company"`

    /* 注册电话（发票类型为个人增值税普通发票和企业增值税普通发票时作为收票人手机号） (Optional) */
    Phone *string `json:"phone"`

    /* 开户银行（发票类型为企业增值税专用发票时必填） (Optional) */
    Bank *string `json:"bank"`

    /* 银行账户（发票类型为企业增值税专用发票时必填） (Optional) */
    Account *string `json:"account"`

    /* 注册地址（发票类型为企业增值税专用发票时必填） (Optional) */
    Address *string `json:"address"`

    /* 邮箱 (Optional) */
    Email *string `json:"email"`

    /* [电子-electronic] (Optional) */
    MediumType *string `json:"mediumType"`
}

/*
 * param regionId: 地域编码，参考OpenAPI公共说明 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewEditInvoiceTemplateRequest(
    regionId string,
) *EditInvoiceTemplateRequest {

	return &EditInvoiceTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/editInvoiceTemplate",
			Method:  "PUT",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域编码，参考OpenAPI公共说明 (Required)
 * param invoiceType: 发票类型:[个人增值税普通发票-Personal_VAT_Ordinary_Invoice,企业增值税专用发票-Enterprise_VAT_Special_Invoice,企业增值税普通发票-Enterprise_VAT_Ordinary_Invoice] (Optional)
 * param invoiceTitle: 发票抬头 (Optional)
 * param invoiceContent: 开票内容（按类别开票/按明细开票） (Optional)
 * param taxId: 纳税人识别码（发票类型为企业增值税专用发票和企业增值税普通发票时必填） (Optional)
 * param company: 单位名称（发票类型为企业增值税专用发票时必填） (Optional)
 * param phone: 注册电话（发票类型为个人增值税普通发票和企业增值税普通发票时作为收票人手机号） (Optional)
 * param bank: 开户银行（发票类型为企业增值税专用发票时必填） (Optional)
 * param account: 银行账户（发票类型为企业增值税专用发票时必填） (Optional)
 * param address: 注册地址（发票类型为企业增值税专用发票时必填） (Optional)
 * param email: 邮箱 (Optional)
 * param mediumType: [电子-electronic] (Optional)
 */
func NewEditInvoiceTemplateRequestWithAllParams(
    regionId string,
    invoiceType *string,
    invoiceTitle *string,
    invoiceContent *string,
    taxId *string,
    company *string,
    phone *string,
    bank *string,
    account *string,
    address *string,
    email *string,
    mediumType *string,
) *EditInvoiceTemplateRequest {

    return &EditInvoiceTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/editInvoiceTemplate",
            Method:  "PUT",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        InvoiceType: invoiceType,
        InvoiceTitle: invoiceTitle,
        InvoiceContent: invoiceContent,
        TaxId: taxId,
        Company: company,
        Phone: phone,
        Bank: bank,
        Account: account,
        Address: address,
        Email: email,
        MediumType: mediumType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewEditInvoiceTemplateRequestWithoutParam() *EditInvoiceTemplateRequest {

    return &EditInvoiceTemplateRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/editInvoiceTemplate",
            Method:  "PUT",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域编码，参考OpenAPI公共说明(Required) */
func (r *EditInvoiceTemplateRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param invoiceType: 发票类型:[个人增值税普通发票-Personal_VAT_Ordinary_Invoice,企业增值税专用发票-Enterprise_VAT_Special_Invoice,企业增值税普通发票-Enterprise_VAT_Ordinary_Invoice](Optional) */
func (r *EditInvoiceTemplateRequest) SetInvoiceType(invoiceType string) {
    r.InvoiceType = &invoiceType
}
/* param invoiceTitle: 发票抬头(Optional) */
func (r *EditInvoiceTemplateRequest) SetInvoiceTitle(invoiceTitle string) {
    r.InvoiceTitle = &invoiceTitle
}
/* param invoiceContent: 开票内容（按类别开票/按明细开票）(Optional) */
func (r *EditInvoiceTemplateRequest) SetInvoiceContent(invoiceContent string) {
    r.InvoiceContent = &invoiceContent
}
/* param taxId: 纳税人识别码（发票类型为企业增值税专用发票和企业增值税普通发票时必填）(Optional) */
func (r *EditInvoiceTemplateRequest) SetTaxId(taxId string) {
    r.TaxId = &taxId
}
/* param company: 单位名称（发票类型为企业增值税专用发票时必填）(Optional) */
func (r *EditInvoiceTemplateRequest) SetCompany(company string) {
    r.Company = &company
}
/* param phone: 注册电话（发票类型为个人增值税普通发票和企业增值税普通发票时作为收票人手机号）(Optional) */
func (r *EditInvoiceTemplateRequest) SetPhone(phone string) {
    r.Phone = &phone
}
/* param bank: 开户银行（发票类型为企业增值税专用发票时必填）(Optional) */
func (r *EditInvoiceTemplateRequest) SetBank(bank string) {
    r.Bank = &bank
}
/* param account: 银行账户（发票类型为企业增值税专用发票时必填）(Optional) */
func (r *EditInvoiceTemplateRequest) SetAccount(account string) {
    r.Account = &account
}
/* param address: 注册地址（发票类型为企业增值税专用发票时必填）(Optional) */
func (r *EditInvoiceTemplateRequest) SetAddress(address string) {
    r.Address = &address
}
/* param email: 邮箱(Optional) */
func (r *EditInvoiceTemplateRequest) SetEmail(email string) {
    r.Email = &email
}
/* param mediumType: [电子-electronic](Optional) */
func (r *EditInvoiceTemplateRequest) SetMediumType(mediumType string) {
    r.MediumType = &mediumType
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r EditInvoiceTemplateRequest) GetRegionId() string {
    return r.RegionId
}

type EditInvoiceTemplateResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result EditInvoiceTemplateResult `json:"result"`
}

type EditInvoiceTemplateResult struct {
    Code int `json:"code"`
    Msg string `json:"msg"`
}