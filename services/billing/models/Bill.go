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


type Bill struct {

    /* 数据来源：10、物联网 11、视频云 12、CDN 13、PCDN 14、IDC 15、通信云 (Optional) */
    DataSource int `json:"dataSource"`

    /* 账单唯一ID (Optional) */
    SourceId string `json:"sourceId"`

    /* 全局唯一ID (Optional) */
    GlobalId string `json:"globalId"`

    /* 用户pin (Optional) */
    Pin string `json:"pin"`

    /* 开始时间 (Optional) */
    StartTime string `json:"startTime"`

    /* 结束时间 (Optional) */
    EndTime string `json:"endTime"`

    /* 地域 (Optional) */
    Region string `json:"region"`

    /* 业务线 (Optional) */
    AppCode string `json:"appCode"`

    /* 产品编码 (Optional) */
    ServiceCode string `json:"serviceCode"`

    /* 资源id (Optional) */
    ResourceId string `json:"resourceId"`

    /* 计费类型： 1、按配置 2、按用量 3、包年包月 4、按次（一次性） (Optional) */
    BillingType int `json:"billingType"`

    /* 账单原价,6位精度 (Optional) */
    BillFee int `json:"billFee"`

    /* 应付金额，2位精度 (Optional) */
    ActualFee int `json:"actualFee"`

    /* 出账时间 (Optional) */
    BillTime string `json:"billTime"`

    /* 币种 CNY 人民币， USD 美元， HKD 港元， IDR 印尼卢比 (Optional) */
    Currency string `json:"currency"`

    /* 支付状态 0、未支付 1、己支付 (Optional) */
    PayState int `json:"payState"`

    /* 支付时间 (Optional) */
    PayTime string `json:"payTime"`

    /* 折扣金额，6位精度 (Optional) */
    DiscountFee int `json:"discountFee"`

    /* 抹零金额，6位精度 (Optional) */
    EraseFee int `json:"eraseFee"`

    /* 余额支付金额：2位精度 (Optional) */
    BalancePayFee int `json:"balancePayFee"`

    /* 现金支付金额：2位精度 (Optional) */
    CashPayFee int `json:"cashPayFee"`

    /* 代金券支付金额，2位精度 (Optional) */
    CashCouponFee int `json:"cashCouponFee"`

    /* 免费代金券金额，2位精度 (Optional) */
    FreeCashCouponFee int `json:"freeCashCouponFee"`

    /* 付费代金券金额，2位精度 (Optional) */
    PayCashCouponFee int `json:"payCashCouponFee"`

    /* 消费时间 (Optional) */
    ConsumeTime string `json:"consumeTime"`

    /* 交易单号 (Optional) */
    TransactionNo string `json:"transactionNo"`

    /* 退款单号 (Optional) */
    RefundNo string `json:"refundNo"`

    /* 站点，0:国内 (Optional) */
    Site int `json:"site"`

    /* 组织机构代码 (Optional) */
    Org string `json:"org"`

    /* 交易类型 1、使用 2、 新购 3、续费 4、配置变更 5、退款 (Optional) */
    TradeType int `json:"tradeType"`

    /* 账单类型 0-普通账单 1-退款账单 2-调账账单 3-保底账单 (Optional) */
    BillType int `json:"billType"`

    /* 配置描述 (Optional) */
    FormulaDesc string `json:"formulaDesc"`

    /* 是否删除 0:未删除 1:己删除 (Optional) */
    IsDeleted int `json:"isDeleted"`

    /* 优惠明细 (Optional) */
    FavorableInfo string `json:"favorableInfo"`

    /* 可用区 (Optional) */
    Az string `json:"az"`
}