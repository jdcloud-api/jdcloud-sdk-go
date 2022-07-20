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
    edcps "github.com/lshuining/jdcloud-sdk-go/services/edcps/models"
)

type DescribeDeviceRaidsRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* 实例类型，可调用（describeDeviceTypes）接口获取指定地域的实例类型，例如：edcps.c.normal1  */
    DeviceType string `json:"deviceType"`

    /* 磁盘类型，取值范围：system、data (Optional) */
    VolumeType *string `json:"volumeType"`
}

/*
 * param regionId: 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域 (Required)
 * param deviceType: 实例类型，可调用（describeDeviceTypes）接口获取指定地域的实例类型，例如：edcps.c.normal1 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeDeviceRaidsRequest(
    regionId string,
    deviceType string,
) *DescribeDeviceRaidsRequest {

	return &DescribeDeviceRaidsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/raids",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        DeviceType: deviceType,
	}
}

/*
 * param regionId: 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域 (Required)
 * param deviceType: 实例类型，可调用（describeDeviceTypes）接口获取指定地域的实例类型，例如：edcps.c.normal1 (Required)
 * param volumeType: 磁盘类型，取值范围：system、data (Optional)
 */
func NewDescribeDeviceRaidsRequestWithAllParams(
    regionId string,
    deviceType string,
    volumeType *string,
) *DescribeDeviceRaidsRequest {

    return &DescribeDeviceRaidsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/raids",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DeviceType: deviceType,
        VolumeType: volumeType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeDeviceRaidsRequestWithoutParam() *DescribeDeviceRaidsRequest {

    return &DescribeDeviceRaidsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/raids",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域(Required) */
func (r *DescribeDeviceRaidsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param deviceType: 实例类型，可调用（describeDeviceTypes）接口获取指定地域的实例类型，例如：edcps.c.normal1(Required) */
func (r *DescribeDeviceRaidsRequest) SetDeviceType(deviceType string) {
    r.DeviceType = deviceType
}

/* param volumeType: 磁盘类型，取值范围：system、data(Optional) */
func (r *DescribeDeviceRaidsRequest) SetVolumeType(volumeType string) {
    r.VolumeType = &volumeType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeDeviceRaidsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeDeviceRaidsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeDeviceRaidsResult `json:"result"`
}

type DescribeDeviceRaidsResult struct {
    Raids []edcps.Raid `json:"raids"`
}