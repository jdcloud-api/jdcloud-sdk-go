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
    iotlink "github.com/jdcloud-api/jdcloud-sdk-go/services/iotlink/models"
)

type GprsStatusByIMSIRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 物联网卡IMSI  */
    Imsi string `json:"imsi"`
}

/*
 * param regionId: Region ID (Required)
 * param imsi: 物联网卡IMSI (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGprsStatusByIMSIRequest(
    regionId string,
    imsi string,
) *GprsStatusByIMSIRequest {

	return &GprsStatusByIMSIRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/gprsStatusByIMSI",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Imsi: imsi,
	}
}

/*
 * param regionId: Region ID (Required)
 * param imsi: 物联网卡IMSI (Required)
 */
func NewGprsStatusByIMSIRequestWithAllParams(
    regionId string,
    imsi string,
) *GprsStatusByIMSIRequest {

    return &GprsStatusByIMSIRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/gprsStatusByIMSI",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Imsi: imsi,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGprsStatusByIMSIRequestWithoutParam() *GprsStatusByIMSIRequest {

    return &GprsStatusByIMSIRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/gprsStatusByIMSI",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *GprsStatusByIMSIRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param imsi: 物联网卡IMSI(Required) */
func (r *GprsStatusByIMSIRequest) SetImsi(imsi string) {
    r.Imsi = imsi
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GprsStatusByIMSIRequest) GetRegionId() string {
    return r.RegionId
}

type GprsStatusByIMSIResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GprsStatusByIMSIResult `json:"result"`
}

type GprsStatusByIMSIResult struct {
    Status string `json:"status"`
    Message string `json:"message"`
    Result iotlink.GprsStatusResp `json:"result"`
}