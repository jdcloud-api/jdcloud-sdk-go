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

type GetBillDetailRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /*   */
    BillId int `json:"billId"`

    /*  (Optional) */
    SystemType *int `json:"systemType"`
}

/*
 * param regionId:  (Required)
 * param billId:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetBillDetailRequest(
    regionId string,
    billId int,
) *GetBillDetailRequest {

	return &GetBillDetailRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/consumeBills/{billId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        BillId: billId,
	}
}

/*
 * param regionId:  (Required)
 * param billId:  (Required)
 * param systemType:  (Optional)
 */
func NewGetBillDetailRequestWithAllParams(
    regionId string,
    billId int,
    systemType *int,
) *GetBillDetailRequest {

    return &GetBillDetailRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/consumeBills/{billId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        BillId: billId,
        SystemType: systemType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetBillDetailRequestWithoutParam() *GetBillDetailRequest {

    return &GetBillDetailRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/consumeBills/{billId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *GetBillDetailRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param billId: (Required) */
func (r *GetBillDetailRequest) SetBillId(billId int) {
    r.BillId = billId
}

/* param systemType: (Optional) */
func (r *GetBillDetailRequest) SetSystemType(systemType int) {
    r.SystemType = &systemType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetBillDetailRequest) GetRegionId() string {
    return r.RegionId
}

type GetBillDetailResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetBillDetailResult `json:"result"`
}

type GetBillDetailResult struct {
    BillId int `json:"billId"`
    Pin string `json:"pin"`
    Site int `json:"site"`
    Region string `json:"region"`
    AppCode string `json:"appCode"`
    AppCodeName string `json:"appCodeName"`
    ServiceCode string `json:"serviceCode"`
    ServiceCodeName string `json:"serviceCodeName"`
    ResourceId string `json:"resourceId"`
    BillingType int `json:"billingType"`
    BillingTypeName string `json:"billingTypeName"`
    Formula string `json:"formula"`
    FormulaStr string `json:"formulaStr"`
    StartTime string `json:"startTime"`
    EndTime string `json:"endTime"`
    CreateTime string `json:"createTime"`
    BillFee int `json:"billFee"`
    BillFee2 int `json:"billFee2"`
    DiscountFee int `json:"discountFee"`
    CouponId string `json:"couponId"`
    CouponFee int `json:"couponFee"`
    ActualFee int `json:"actualFee"`
    CashCouponFee int `json:"cashCouponFee"`
    BalancePayFee int `json:"balancePayFee"`
    CashPayFee int `json:"cashPayFee"`
    ArrearFee int `json:"arrearFee"`
    PaySate int `json:"paySate"`
    SystemType int `json:"systemType"`
    ResourceName string `json:"resourceName"`
}