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

type ExportCoProductsRequest struct {

    core.JDCloudRequest

    /* 区域(如:cn-north-1)  */
    RegionId string `json:"regionId"`

    /* 合作id (Optional) */
    CooperationId *string `json:"cooperationId"`

    /* 公司名称 (Optional) */
    CompanyName *string `json:"companyName"`

    /* 合作名称 (Optional) */
    Name *string `json:"name"`

    /* 产品名称 (Optional) */
    ProductName *string `json:"productName"`

    /* 合同编号 (Optional) */
    ContractNo *string `json:"contractNo"`

    /* 产品类型 (Optional) */
    ProductType *int `json:"productType"`

    /* 产品模式 (Optional) */
    ProductMode *int `json:"productMode"`

    /* 产品状态 (Optional) */
    ProductStatus *int `json:"productStatus"`

    /* 结算方式 (Optional) */
    SettlementMode *int `json:"settlementMode"`

    /* 结算周期 (Optional) */
    SettlementCycle *int `json:"settlementCycle"`

    /* 页码 (Optional) */
    PageIndex *int `json:"pageIndex"`

    /* 每页记录数 (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: 区域(如:cn-north-1) (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewExportCoProductsRequest(
    regionId string,
) *ExportCoProductsRequest {

	return &ExportCoProductsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/CooperationInfo:exportCoProducts",
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
 * param companyName: 公司名称 (Optional)
 * param name: 合作名称 (Optional)
 * param productName: 产品名称 (Optional)
 * param contractNo: 合同编号 (Optional)
 * param productType: 产品类型 (Optional)
 * param productMode: 产品模式 (Optional)
 * param productStatus: 产品状态 (Optional)
 * param settlementMode: 结算方式 (Optional)
 * param settlementCycle: 结算周期 (Optional)
 * param pageIndex: 页码 (Optional)
 * param pageSize: 每页记录数 (Optional)
 */
func NewExportCoProductsRequestWithAllParams(
    regionId string,
    cooperationId *string,
    companyName *string,
    name *string,
    productName *string,
    contractNo *string,
    productType *int,
    productMode *int,
    productStatus *int,
    settlementMode *int,
    settlementCycle *int,
    pageIndex *int,
    pageSize *int,
) *ExportCoProductsRequest {

    return &ExportCoProductsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/CooperationInfo:exportCoProducts",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        CooperationId: cooperationId,
        CompanyName: companyName,
        Name: name,
        ProductName: productName,
        ContractNo: contractNo,
        ProductType: productType,
        ProductMode: productMode,
        ProductStatus: productStatus,
        SettlementMode: settlementMode,
        SettlementCycle: settlementCycle,
        PageIndex: pageIndex,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewExportCoProductsRequestWithoutParam() *ExportCoProductsRequest {

    return &ExportCoProductsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/CooperationInfo:exportCoProducts",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域(如:cn-north-1)(Required) */
func (r *ExportCoProductsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param cooperationId: 合作id(Optional) */
func (r *ExportCoProductsRequest) SetCooperationId(cooperationId string) {
    r.CooperationId = &cooperationId
}
/* param companyName: 公司名称(Optional) */
func (r *ExportCoProductsRequest) SetCompanyName(companyName string) {
    r.CompanyName = &companyName
}
/* param name: 合作名称(Optional) */
func (r *ExportCoProductsRequest) SetName(name string) {
    r.Name = &name
}
/* param productName: 产品名称(Optional) */
func (r *ExportCoProductsRequest) SetProductName(productName string) {
    r.ProductName = &productName
}
/* param contractNo: 合同编号(Optional) */
func (r *ExportCoProductsRequest) SetContractNo(contractNo string) {
    r.ContractNo = &contractNo
}
/* param productType: 产品类型(Optional) */
func (r *ExportCoProductsRequest) SetProductType(productType int) {
    r.ProductType = &productType
}
/* param productMode: 产品模式(Optional) */
func (r *ExportCoProductsRequest) SetProductMode(productMode int) {
    r.ProductMode = &productMode
}
/* param productStatus: 产品状态(Optional) */
func (r *ExportCoProductsRequest) SetProductStatus(productStatus int) {
    r.ProductStatus = &productStatus
}
/* param settlementMode: 结算方式(Optional) */
func (r *ExportCoProductsRequest) SetSettlementMode(settlementMode int) {
    r.SettlementMode = &settlementMode
}
/* param settlementCycle: 结算周期(Optional) */
func (r *ExportCoProductsRequest) SetSettlementCycle(settlementCycle int) {
    r.SettlementCycle = &settlementCycle
}
/* param pageIndex: 页码(Optional) */
func (r *ExportCoProductsRequest) SetPageIndex(pageIndex int) {
    r.PageIndex = &pageIndex
}
/* param pageSize: 每页记录数(Optional) */
func (r *ExportCoProductsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ExportCoProductsRequest) GetRegionId() string {
    return r.RegionId
}

type ExportCoProductsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ExportCoProductsResult `json:"result"`
}

type ExportCoProductsResult struct {
    Result string `json:"result"`
    RequestId string `json:"requestId"`
}