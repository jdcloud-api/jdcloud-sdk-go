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


type InvoicePreviewReqVo struct {

    /* 用户pin (Optional) */
    Pin string `json:"pin"`

    /* 订单号 多个用逗号,分开 (Optional) */
    OrderIds string `json:"orderIds"`

    /* 欠票抵扣的订单信息 (Optional) */
    DeductInfoList []DeductInfo `json:"deductInfoList"`

    /* 指定开票金额 (Optional) */
    Amount int `json:"amount"`

    /* 开票月份 仅按月开票生效 (Optional) */
    MonthGroups string `json:"monthGroups"`

    /* 发票类型[电子-electronic, 数电-digital] (Optional) */
    MediumType string `json:"mediumType"`

    /* 全部开票标识 1-全部开票 全部开票, 不需要传orderIds 需要传 开始、结束时间 (Optional) */
    InvoiceAll int `json:"invoiceAll"`

    /* 开票订单的开始时间(yyyy-MM-dd HH:mm:ss) (Optional) */
    OrderStartTime string `json:"orderStartTime"`

    /* 开票订单的结束时间(yyyy-MM-dd HH:mm:ss) (Optional) */
    OrderEndTime string `json:"orderEndTime"`
}
