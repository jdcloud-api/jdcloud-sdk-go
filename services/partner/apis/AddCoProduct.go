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
    partner "github.com/jdcloud-api/jdcloud-sdk-go/services/partner/models"
)

type AddCoProductRequest struct {

    core.JDCloudRequest

    /* 区域(如:cn-north-1)  */
    RegionId string `json:"regionId"`

    /* 合作id (Optional) */
    CooperationId *string `json:"cooperationId"`

    /* 合作产品名称 (Optional) */
    ProductName *string `json:"productName"`

    /* 合作名称 (Optional) */
    Name *string `json:"name"`

    /* 产品类型 (Optional) */
    ProductType *int `json:"productType"`

    /* 产品模式 (Optional) */
    ProductMode *int `json:"productMode"`

    /* 产品简介 (Optional) */
    ProductDesc *string `json:"productDesc"`

    /* 目标客户 (Optional) */
    TargetCustomer *string `json:"targetCustomer"`

    /* 市场规模 (Optional) */
    MarketSize *string `json:"marketSize"`

    /* 主要竞品 (Optional) */
    ComparableProduct *string `json:"comparableProduct"`

    /* 售卖形态  1套/n年、2套/n月、3套、4次 (Optional) */
    SellingForm *int `json:"sellingForm"`

    /* 售卖方式  1直销，2渠道，3代理 (Optional) */
    SellingMode *string `json:"sellingMode"`

    /* 定价 (Optional) */
    Pricing *string `json:"pricing"`

    /* 产品状态 (Optional) */
    ProductStatus *int `json:"productStatus"`

    /* 收入预测 (Optional) */
    IncomeForecast *string `json:"incomeForecast"`

    /* 成本结构 (Optional) */
    CostStructure *string `json:"costStructure"`

    /* 毛利率预测 (Optional) */
    GrossMarginForecast *string `json:"grossMarginForecast"`

    /* 定价策略 1市场对标， 2总成本加成，3变动成本加成 (Optional) */
    PricingStrategy *int `json:"pricingStrategy"`

    /* 结算方式 1固定金额结算，2实际售价固定比例结算，3实际售价浮动比例结算 (Optional) */
    SettlementMode *int `json:"settlementMode"`

    /* 结算周期 1周结后付款、2月结后付款、3季结后付款、4年结后付款，5 PO预付款 (Optional) */
    SettlementCycle *int `json:"settlementCycle"`

    /* 风险及建议 (Optional) */
    RiskSuggestion *string `json:"riskSuggestion"`

    /* erp (Optional) */
    Erp *string `json:"erp"`

    /* 产品唯一标识id (Optional) */
    Uuid *string `json:"uuid"`
}

/*
 * param regionId: 区域(如:cn-north-1) (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAddCoProductRequest(
    regionId string,
) *AddCoProductRequest {

	return &AddCoProductRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/CooperationInfo:addCoProduct",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 区域(如:cn-north-1) (Required)
 * param cooperationId: 合作id (Optional)
 * param productName: 合作产品名称 (Optional)
 * param name: 合作名称 (Optional)
 * param productType: 产品类型 (Optional)
 * param productMode: 产品模式 (Optional)
 * param productDesc: 产品简介 (Optional)
 * param targetCustomer: 目标客户 (Optional)
 * param marketSize: 市场规模 (Optional)
 * param comparableProduct: 主要竞品 (Optional)
 * param sellingForm: 售卖形态  1套/n年、2套/n月、3套、4次 (Optional)
 * param sellingMode: 售卖方式  1直销，2渠道，3代理 (Optional)
 * param pricing: 定价 (Optional)
 * param productStatus: 产品状态 (Optional)
 * param incomeForecast: 收入预测 (Optional)
 * param costStructure: 成本结构 (Optional)
 * param grossMarginForecast: 毛利率预测 (Optional)
 * param pricingStrategy: 定价策略 1市场对标， 2总成本加成，3变动成本加成 (Optional)
 * param settlementMode: 结算方式 1固定金额结算，2实际售价固定比例结算，3实际售价浮动比例结算 (Optional)
 * param settlementCycle: 结算周期 1周结后付款、2月结后付款、3季结后付款、4年结后付款，5 PO预付款 (Optional)
 * param riskSuggestion: 风险及建议 (Optional)
 * param erp: erp (Optional)
 * param uuid: 产品唯一标识id (Optional)
 */
func NewAddCoProductRequestWithAllParams(
    regionId string,
    cooperationId *string,
    productName *string,
    name *string,
    productType *int,
    productMode *int,
    productDesc *string,
    targetCustomer *string,
    marketSize *string,
    comparableProduct *string,
    sellingForm *int,
    sellingMode *string,
    pricing *string,
    productStatus *int,
    incomeForecast *string,
    costStructure *string,
    grossMarginForecast *string,
    pricingStrategy *int,
    settlementMode *int,
    settlementCycle *int,
    riskSuggestion *string,
    erp *string,
    uuid *string,
) *AddCoProductRequest {

    return &AddCoProductRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/CooperationInfo:addCoProduct",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        CooperationId: cooperationId,
        ProductName: productName,
        Name: name,
        ProductType: productType,
        ProductMode: productMode,
        ProductDesc: productDesc,
        TargetCustomer: targetCustomer,
        MarketSize: marketSize,
        ComparableProduct: comparableProduct,
        SellingForm: sellingForm,
        SellingMode: sellingMode,
        Pricing: pricing,
        ProductStatus: productStatus,
        IncomeForecast: incomeForecast,
        CostStructure: costStructure,
        GrossMarginForecast: grossMarginForecast,
        PricingStrategy: pricingStrategy,
        SettlementMode: settlementMode,
        SettlementCycle: settlementCycle,
        RiskSuggestion: riskSuggestion,
        Erp: erp,
        Uuid: uuid,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAddCoProductRequestWithoutParam() *AddCoProductRequest {

    return &AddCoProductRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/CooperationInfo:addCoProduct",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域(如:cn-north-1)(Required) */
func (r *AddCoProductRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param cooperationId: 合作id(Optional) */
func (r *AddCoProductRequest) SetCooperationId(cooperationId string) {
    r.CooperationId = &cooperationId
}
/* param productName: 合作产品名称(Optional) */
func (r *AddCoProductRequest) SetProductName(productName string) {
    r.ProductName = &productName
}
/* param name: 合作名称(Optional) */
func (r *AddCoProductRequest) SetName(name string) {
    r.Name = &name
}
/* param productType: 产品类型(Optional) */
func (r *AddCoProductRequest) SetProductType(productType int) {
    r.ProductType = &productType
}
/* param productMode: 产品模式(Optional) */
func (r *AddCoProductRequest) SetProductMode(productMode int) {
    r.ProductMode = &productMode
}
/* param productDesc: 产品简介(Optional) */
func (r *AddCoProductRequest) SetProductDesc(productDesc string) {
    r.ProductDesc = &productDesc
}
/* param targetCustomer: 目标客户(Optional) */
func (r *AddCoProductRequest) SetTargetCustomer(targetCustomer string) {
    r.TargetCustomer = &targetCustomer
}
/* param marketSize: 市场规模(Optional) */
func (r *AddCoProductRequest) SetMarketSize(marketSize string) {
    r.MarketSize = &marketSize
}
/* param comparableProduct: 主要竞品(Optional) */
func (r *AddCoProductRequest) SetComparableProduct(comparableProduct string) {
    r.ComparableProduct = &comparableProduct
}
/* param sellingForm: 售卖形态  1套/n年、2套/n月、3套、4次(Optional) */
func (r *AddCoProductRequest) SetSellingForm(sellingForm int) {
    r.SellingForm = &sellingForm
}
/* param sellingMode: 售卖方式  1直销，2渠道，3代理(Optional) */
func (r *AddCoProductRequest) SetSellingMode(sellingMode string) {
    r.SellingMode = &sellingMode
}
/* param pricing: 定价(Optional) */
func (r *AddCoProductRequest) SetPricing(pricing string) {
    r.Pricing = &pricing
}
/* param productStatus: 产品状态(Optional) */
func (r *AddCoProductRequest) SetProductStatus(productStatus int) {
    r.ProductStatus = &productStatus
}
/* param incomeForecast: 收入预测(Optional) */
func (r *AddCoProductRequest) SetIncomeForecast(incomeForecast string) {
    r.IncomeForecast = &incomeForecast
}
/* param costStructure: 成本结构(Optional) */
func (r *AddCoProductRequest) SetCostStructure(costStructure string) {
    r.CostStructure = &costStructure
}
/* param grossMarginForecast: 毛利率预测(Optional) */
func (r *AddCoProductRequest) SetGrossMarginForecast(grossMarginForecast string) {
    r.GrossMarginForecast = &grossMarginForecast
}
/* param pricingStrategy: 定价策略 1市场对标， 2总成本加成，3变动成本加成(Optional) */
func (r *AddCoProductRequest) SetPricingStrategy(pricingStrategy int) {
    r.PricingStrategy = &pricingStrategy
}
/* param settlementMode: 结算方式 1固定金额结算，2实际售价固定比例结算，3实际售价浮动比例结算(Optional) */
func (r *AddCoProductRequest) SetSettlementMode(settlementMode int) {
    r.SettlementMode = &settlementMode
}
/* param settlementCycle: 结算周期 1周结后付款、2月结后付款、3季结后付款、4年结后付款，5 PO预付款(Optional) */
func (r *AddCoProductRequest) SetSettlementCycle(settlementCycle int) {
    r.SettlementCycle = &settlementCycle
}
/* param riskSuggestion: 风险及建议(Optional) */
func (r *AddCoProductRequest) SetRiskSuggestion(riskSuggestion string) {
    r.RiskSuggestion = &riskSuggestion
}
/* param erp: erp(Optional) */
func (r *AddCoProductRequest) SetErp(erp string) {
    r.Erp = &erp
}
/* param uuid: 产品唯一标识id(Optional) */
func (r *AddCoProductRequest) SetUuid(uuid string) {
    r.Uuid = &uuid
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AddCoProductRequest) GetRegionId() string {
    return r.RegionId
}

type AddCoProductResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AddCoProductResult `json:"result"`
}

type AddCoProductResult struct {
    Result bool `json:"result"`
    RequestId string `json:"requestId"`
}