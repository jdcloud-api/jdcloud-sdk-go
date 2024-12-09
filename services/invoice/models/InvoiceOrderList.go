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

package models


type InvoiceOrderList struct {

    /* 开始时间 (Optional) */
    SearchStartTime string `json:"searchStartTime"`

    /* 结束时间 (Optional) */
    SearchEndTime string `json:"searchEndTime"`

    /* 页码 (Optional) */
    PageNumber int `json:"pageNumber"`

    /* 页大小 (Optional) */
    PageSize int `json:"pageSize"`

    /* 单据类型[订单-order，账单-bill 月结算单-month] (Optional) */
    ReceiptType string `json:"receiptType"`

    /* 交易(支付)类型 1-代付 2-自付 (Optional) */
    PayType int `json:"payType"`
}