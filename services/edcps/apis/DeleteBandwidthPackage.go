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

type DeleteBandwidthPackageRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* 共享带宽ID  */
    BandwidthPackageId string `json:"bandwidthPackageId"`

    /* 由客户端生成，用于保证请求的幂等性，长度不能超过36个字符；<br/>
如果多个请求使用了相同的clientToken，只会执行第一个请求，之后的请求直接返回第一个请求的结果<br/>
 (Optional) */
    ClientToken *string `json:"clientToken"`
}

/*
 * param regionId: 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域 (Required)
 * param bandwidthPackageId: 共享带宽ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteBandwidthPackageRequest(
    regionId string,
    bandwidthPackageId string,
) *DeleteBandwidthPackageRequest {

	return &DeleteBandwidthPackageRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/bandwidthPackages/{bandwidthPackageId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        BandwidthPackageId: bandwidthPackageId,
	}
}

/*
 * param regionId: 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域 (Required)
 * param bandwidthPackageId: 共享带宽ID (Required)
 * param clientToken: 由客户端生成，用于保证请求的幂等性，长度不能超过36个字符；<br/>
如果多个请求使用了相同的clientToken，只会执行第一个请求，之后的请求直接返回第一个请求的结果<br/>
 (Optional)
 */
func NewDeleteBandwidthPackageRequestWithAllParams(
    regionId string,
    bandwidthPackageId string,
    clientToken *string,
) *DeleteBandwidthPackageRequest {

    return &DeleteBandwidthPackageRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/bandwidthPackages/{bandwidthPackageId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        BandwidthPackageId: bandwidthPackageId,
        ClientToken: clientToken,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteBandwidthPackageRequestWithoutParam() *DeleteBandwidthPackageRequest {

    return &DeleteBandwidthPackageRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/bandwidthPackages/{bandwidthPackageId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域(Required) */
func (r *DeleteBandwidthPackageRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param bandwidthPackageId: 共享带宽ID(Required) */
func (r *DeleteBandwidthPackageRequest) SetBandwidthPackageId(bandwidthPackageId string) {
    r.BandwidthPackageId = bandwidthPackageId
}

/* param clientToken: 由客户端生成，用于保证请求的幂等性，长度不能超过36个字符；<br/>
如果多个请求使用了相同的clientToken，只会执行第一个请求，之后的请求直接返回第一个请求的结果<br/>
(Optional) */
func (r *DeleteBandwidthPackageRequest) SetClientToken(clientToken string) {
    r.ClientToken = &clientToken
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteBandwidthPackageRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteBandwidthPackageResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteBandwidthPackageResult `json:"result"`
}

type DeleteBandwidthPackageResult struct {
    Success bool `json:"success"`
}