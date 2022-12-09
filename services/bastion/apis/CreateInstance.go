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
    bastion "github.com/jdcloud-api/jdcloud-sdk-go/services/bastion/models"
    charge "github.com/jdcloud-api/jdcloud-sdk-go/services/charge/models"
)

type CreateInstanceRequest struct {

    core.JDCloudRequest

    /* regionId  */
    RegionId string `json:"regionId"`

    /* 实例的相关配置  */
    InstanceSpec *bastion.InstanceSpec `json:"instanceSpec"`

    /* 计费信息的相关配置 (Optional) */
    ChargeSpec *charge.ChargeSpec `json:"chargeSpec"`
}

/*
 * param regionId: regionId (Required)
 * param instanceSpec: 实例的相关配置 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateInstanceRequest(
    regionId string,
    instanceSpec *bastion.InstanceSpec,
) *CreateInstanceRequest {

	return &CreateInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceSpec: instanceSpec,
	}
}

/*
 * param regionId: regionId (Required)
 * param instanceSpec: 实例的相关配置 (Required)
 * param chargeSpec: 计费信息的相关配置 (Optional)
 */
func NewCreateInstanceRequestWithAllParams(
    regionId string,
    instanceSpec *bastion.InstanceSpec,
    chargeSpec *charge.ChargeSpec,
) *CreateInstanceRequest {

    return &CreateInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceSpec: instanceSpec,
        ChargeSpec: chargeSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateInstanceRequestWithoutParam() *CreateInstanceRequest {

    return &CreateInstanceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: regionId(Required) */
func (r *CreateInstanceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param instanceSpec: 实例的相关配置(Required) */
func (r *CreateInstanceRequest) SetInstanceSpec(instanceSpec *bastion.InstanceSpec) {
    r.InstanceSpec = instanceSpec
}
/* param chargeSpec: 计费信息的相关配置(Optional) */
func (r *CreateInstanceRequest) SetChargeSpec(chargeSpec *charge.ChargeSpec) {
    r.ChargeSpec = chargeSpec
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateInstanceRequest) GetRegionId() string {
    return r.RegionId
}

type CreateInstanceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateInstanceResult `json:"result"`
}

type CreateInstanceResult struct {
    BuyId string `json:"buyId"`
    SourceIds []bastion.Source `json:"sourceIds"`
}