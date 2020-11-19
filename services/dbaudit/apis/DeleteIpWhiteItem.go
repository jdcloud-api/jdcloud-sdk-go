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

type DeleteIpWhiteItemRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 审计实例ID  */
    InsId string `json:"insId"`

    /* IP白名单记录，支持掩码  */
    Cidr string `json:"cidr"`
}

/*
 * param regionId: 地域 Id (Required)
 * param insId: 审计实例ID (Required)
 * param cidr: IP白名单记录，支持掩码 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteIpWhiteItemRequest(
    regionId string,
    insId string,
    cidr string,
) *DeleteIpWhiteItemRequest {

	return &DeleteIpWhiteItemRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{insId}/ipwhitelist",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InsId: insId,
        Cidr: cidr,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param insId: 审计实例ID (Required)
 * param cidr: IP白名单记录，支持掩码 (Required)
 */
func NewDeleteIpWhiteItemRequestWithAllParams(
    regionId string,
    insId string,
    cidr string,
) *DeleteIpWhiteItemRequest {

    return &DeleteIpWhiteItemRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{insId}/ipwhitelist",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InsId: insId,
        Cidr: cidr,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteIpWhiteItemRequestWithoutParam() *DeleteIpWhiteItemRequest {

    return &DeleteIpWhiteItemRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{insId}/ipwhitelist",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *DeleteIpWhiteItemRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param insId: 审计实例ID(Required) */
func (r *DeleteIpWhiteItemRequest) SetInsId(insId string) {
    r.InsId = insId
}

/* param cidr: IP白名单记录，支持掩码(Required) */
func (r *DeleteIpWhiteItemRequest) SetCidr(cidr string) {
    r.Cidr = cidr
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteIpWhiteItemRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteIpWhiteItemResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteIpWhiteItemResult `json:"result"`
}

type DeleteIpWhiteItemResult struct {
}