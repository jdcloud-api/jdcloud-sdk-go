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
    vm "github.com/jdcloud-api/jdcloud-sdk-go/services/vm/models"
)

type ModifyInstanceNetworkAttributeRequest struct {

    core.JDCloudRequest

    /* 地域ID。  */
    RegionId string `json:"regionId"`

    /* 云主机ID。  */
    InstanceId string `json:"instanceId"`

    /* 弹性网卡列表。  */
    Networks []vm.InstanceNetworkAttribute `json:"networks"`
}

/*
 * param regionId: 地域ID。 (Required)
 * param instanceId: 云主机ID。 (Required)
 * param networks: 弹性网卡列表。 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyInstanceNetworkAttributeRequest(
    regionId string,
    instanceId string,
    networks []vm.InstanceNetworkAttribute,
) *ModifyInstanceNetworkAttributeRequest {

	return &ModifyInstanceNetworkAttributeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:modifyInstanceNetworkAttribute",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        Networks: networks,
	}
}

/*
 * param regionId: 地域ID。 (Required)
 * param instanceId: 云主机ID。 (Required)
 * param networks: 弹性网卡列表。 (Required)
 */
func NewModifyInstanceNetworkAttributeRequestWithAllParams(
    regionId string,
    instanceId string,
    networks []vm.InstanceNetworkAttribute,
) *ModifyInstanceNetworkAttributeRequest {

    return &ModifyInstanceNetworkAttributeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:modifyInstanceNetworkAttribute",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        Networks: networks,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyInstanceNetworkAttributeRequestWithoutParam() *ModifyInstanceNetworkAttributeRequest {

    return &ModifyInstanceNetworkAttributeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:modifyInstanceNetworkAttribute",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID。(Required) */
func (r *ModifyInstanceNetworkAttributeRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param instanceId: 云主机ID。(Required) */
func (r *ModifyInstanceNetworkAttributeRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}
/* param networks: 弹性网卡列表。(Required) */
func (r *ModifyInstanceNetworkAttributeRequest) SetNetworks(networks []vm.InstanceNetworkAttribute) {
    r.Networks = networks
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyInstanceNetworkAttributeRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyInstanceNetworkAttributeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyInstanceNetworkAttributeResult `json:"result"`
}

type ModifyInstanceNetworkAttributeResult struct {
}