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


type Invoices struct {

    /* 发票状态[已申请-applied ，处理中-processing ，已开票-invoiced ，已邮寄-mailed ，已驳回-dismissed ，已作废-obsolete ，已取消-cancelled，退票中-refund，已退票-refunded，退票驳回-refund_rejected] (Optional) */
    Status string `json:"status"`

    /* 开始时间 (Optional) */
    SearchStartDate string `json:"searchStartDate"`

    /* 结束时间 (Optional) */
    SearchEndDate string `json:"searchEndDate"`

    /* 页码 (Optional) */
    PageNumber int `json:"pageNumber"`

    /* 页大小 (Optional) */
    PageSize int `json:"pageSize"`

    /* 按明细开票 1 按月结算单开票 2 按指定金额开票 3 (Optional) */
    InvoiceType int `json:"invoiceType"`
}