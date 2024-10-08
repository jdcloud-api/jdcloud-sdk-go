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

type BillDetailRequest struct {

    core.JDCloudRequest

    /* 地域编码，参考OpenAPI公共说明  */
    RegionId string `json:"regionId"`

    /* 计费账单编号  */
    BillId string `json:"billId"`
}

/*
 * param regionId: 地域编码，参考OpenAPI公共说明 (Required)
 * param billId: 计费账单编号 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewBillDetailRequest(
    regionId string,
    billId string,
) *BillDetailRequest {

	return &BillDetailRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/billDetail/{billId}",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        BillId: billId,
	}
}

/*
 * param regionId: 地域编码，参考OpenAPI公共说明 (Required)
 * param billId: 计费账单编号 (Required)
 */
func NewBillDetailRequestWithAllParams(
    regionId string,
    billId string,
) *BillDetailRequest {

    return &BillDetailRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/billDetail/{billId}",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        BillId: billId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewBillDetailRequestWithoutParam() *BillDetailRequest {

    return &BillDetailRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/billDetail/{billId}",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域编码，参考OpenAPI公共说明(Required) */
func (r *BillDetailRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param billId: 计费账单编号(Required) */
func (r *BillDetailRequest) SetBillId(billId string) {
    r.BillId = billId
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r BillDetailRequest) GetRegionId() string {
    return r.RegionId
}

type BillDetailResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result BillDetailResult `json:"result"`
}

type BillDetailResult struct {
    BillDataList []invoice.BillDataList `json:"billDataList"`
}