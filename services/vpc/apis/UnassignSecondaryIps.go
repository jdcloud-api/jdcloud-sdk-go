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

type UnassignSecondaryIpsRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* networkInterface ID  */
    NetworkInterfaceId string `json:"networkInterfaceId"`

    /* 指定删除的secondaryIp地址 (Optional) */
    SecondaryIps []string `json:"secondaryIps"`

    /* 指定删除的secondaryIp网段 (Optional) */
    SecondaryCidrs []string `json:"secondaryCidrs"`
}

/*
 * param regionId: Region ID (Required)
 * param networkInterfaceId: networkInterface ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUnassignSecondaryIpsRequest(
    regionId string,
    networkInterfaceId string,
) *UnassignSecondaryIpsRequest {

	return &UnassignSecondaryIpsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/networkInterfaces/{networkInterfaceId}:unassignSecondaryIps",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        NetworkInterfaceId: networkInterfaceId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param networkInterfaceId: networkInterface ID (Required)
 * param secondaryIps: 指定删除的secondaryIp地址 (Optional)
 * param secondaryCidrs: 指定删除的secondaryIp网段 (Optional)
 */
func NewUnassignSecondaryIpsRequestWithAllParams(
    regionId string,
    networkInterfaceId string,
    secondaryIps []string,
    secondaryCidrs []string,
) *UnassignSecondaryIpsRequest {

    return &UnassignSecondaryIpsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/networkInterfaces/{networkInterfaceId}:unassignSecondaryIps",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        NetworkInterfaceId: networkInterfaceId,
        SecondaryIps: secondaryIps,
        SecondaryCidrs: secondaryCidrs,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUnassignSecondaryIpsRequestWithoutParam() *UnassignSecondaryIpsRequest {

    return &UnassignSecondaryIpsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/networkInterfaces/{networkInterfaceId}:unassignSecondaryIps",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *UnassignSecondaryIpsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param networkInterfaceId: networkInterface ID(Required) */
func (r *UnassignSecondaryIpsRequest) SetNetworkInterfaceId(networkInterfaceId string) {
    r.NetworkInterfaceId = networkInterfaceId
}
/* param secondaryIps: 指定删除的secondaryIp地址(Optional) */
func (r *UnassignSecondaryIpsRequest) SetSecondaryIps(secondaryIps []string) {
    r.SecondaryIps = secondaryIps
}
/* param secondaryCidrs: 指定删除的secondaryIp网段(Optional) */
func (r *UnassignSecondaryIpsRequest) SetSecondaryCidrs(secondaryCidrs []string) {
    r.SecondaryCidrs = secondaryCidrs
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UnassignSecondaryIpsRequest) GetRegionId() string {
    return r.RegionId
}

type UnassignSecondaryIpsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UnassignSecondaryIpsResult `json:"result"`
}

type UnassignSecondaryIpsResult struct {
}