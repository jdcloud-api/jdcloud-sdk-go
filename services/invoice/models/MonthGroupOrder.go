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


type MonthGroupOrder struct {

    /* 如果是月结订单，返回具体月份 例202301 (Optional) */
    MonthGroup string `json:"monthGroup"`

    /* 月份描述 例2023年01月 (Optional) */
    MonthGroupDesc string `json:"monthGroupDesc"`

    /* 可开票金额 (Optional) */
    EnableAmount int `json:"enableAmount"`

    /* 已开票金额 (Optional) */
    UseAmount int `json:"useAmount"`

    /* 总金额 (Optional) */
    Amount int `json:"amount"`

    /* 本次开票金额 (Optional) */
    NowAmount int `json:"nowAmount"`

    /* 按产品线的详情 (Optional) */
    InvoiceMonthDetailVos []InvoiceDetails `json:"invoiceMonthDetailVos"`
}
