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

type DeleteNetworkInterfaceRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* networkInterface ID  */
    NetworkInterfaceId string `json:"networkInterfaceId"`
}

/*
 * param regionId: Region ID (Required)
 * param networkInterfaceId: networkInterface ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteNetworkInterfaceRequest(
    regionId string,
    networkInterfaceId string,
) *DeleteNetworkInterfaceRequest {

	return &DeleteNetworkInterfaceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/ydNetworkInterfaces/{networkInterfaceId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        NetworkInterfaceId: networkInterfaceId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param networkInterfaceId: networkInterface ID (Required)
 */
func NewDeleteNetworkInterfaceRequestWithAllParams(
    regionId string,
    networkInterfaceId string,
) *DeleteNetworkInterfaceRequest {

    return &DeleteNetworkInterfaceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/ydNetworkInterfaces/{networkInterfaceId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        NetworkInterfaceId: networkInterfaceId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteNetworkInterfaceRequestWithoutParam() *DeleteNetworkInterfaceRequest {

    return &DeleteNetworkInterfaceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/ydNetworkInterfaces/{networkInterfaceId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DeleteNetworkInterfaceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param networkInterfaceId: networkInterface ID(Required) */
func (r *DeleteNetworkInterfaceRequest) SetNetworkInterfaceId(networkInterfaceId string) {
    r.NetworkInterfaceId = networkInterfaceId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteNetworkInterfaceRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteNetworkInterfaceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteNetworkInterfaceResult `json:"result"`
}

type DeleteNetworkInterfaceResult struct {
}