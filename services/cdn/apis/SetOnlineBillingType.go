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
    "github.com/lshuining/jdcloud-sdk-go/core"
)

type SetOnlineBillingTypeRequest struct {

    core.JDCloudRequest

    /* 计费方式,取值[0,1],0:日流量计费,1:日峰值带宽计费. (Optional) */
    AllType *int `json:"allType"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSetOnlineBillingTypeRequest(
) *SetOnlineBillingTypeRequest {

	return &SetOnlineBillingTypeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/onlineBillingType",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param allType: 计费方式,取值[0,1],0:日流量计费,1:日峰值带宽计费. (Optional)
 */
func NewSetOnlineBillingTypeRequestWithAllParams(
    allType *int,
) *SetOnlineBillingTypeRequest {

    return &SetOnlineBillingTypeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/onlineBillingType",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        AllType: allType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSetOnlineBillingTypeRequestWithoutParam() *SetOnlineBillingTypeRequest {

    return &SetOnlineBillingTypeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/onlineBillingType",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param allType: 计费方式,取值[0,1],0:日流量计费,1:日峰值带宽计费.(Optional) */
func (r *SetOnlineBillingTypeRequest) SetAllType(allType int) {
    r.AllType = &allType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SetOnlineBillingTypeRequest) GetRegionId() string {
    return ""
}

type SetOnlineBillingTypeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SetOnlineBillingTypeResult `json:"result"`
}

type SetOnlineBillingTypeResult struct {
}