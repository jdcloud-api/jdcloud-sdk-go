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
    vpc "github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/models"
)

type DescribeBandwidthPackageRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 共享带宽包ID  */
    BandwidthPackageId string `json:"bandwidthPackageId"`
}

/*
 * param regionId: Region ID (Required)
 * param bandwidthPackageId: 共享带宽包ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeBandwidthPackageRequest(
    regionId string,
    bandwidthPackageId string,
) *DescribeBandwidthPackageRequest {

	return &DescribeBandwidthPackageRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/bandwidthPackages/{bandwidthPackageId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        BandwidthPackageId: bandwidthPackageId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param bandwidthPackageId: 共享带宽包ID (Required)
 */
func NewDescribeBandwidthPackageRequestWithAllParams(
    regionId string,
    bandwidthPackageId string,
) *DescribeBandwidthPackageRequest {

    return &DescribeBandwidthPackageRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/bandwidthPackages/{bandwidthPackageId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        BandwidthPackageId: bandwidthPackageId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeBandwidthPackageRequestWithoutParam() *DescribeBandwidthPackageRequest {

    return &DescribeBandwidthPackageRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/bandwidthPackages/{bandwidthPackageId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeBandwidthPackageRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param bandwidthPackageId: 共享带宽包ID(Required) */
func (r *DescribeBandwidthPackageRequest) SetBandwidthPackageId(bandwidthPackageId string) {
    r.BandwidthPackageId = bandwidthPackageId
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeBandwidthPackageRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeBandwidthPackageResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeBandwidthPackageResult `json:"result"`
}

type DescribeBandwidthPackageResult struct {
    BandwidthPackage vpc.BandwidthPackage `json:"bandwidthPackage"`
}