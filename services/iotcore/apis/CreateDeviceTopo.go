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
    iotcore "github.com/jdcloud-api/jdcloud-sdk-go/services/iotcore/models"
)

type CreateDeviceTopoRequest struct {

    core.JDCloudRequest

    /* 区域id  */
    RegionId string `json:"regionId"`

    /* 实例Id  */
    InstanceId string `json:"instanceId"`

    /* 方法查询请求  */
    DeviceTopoInfoVO *iotcore.DeviceTopoInfoVO `json:"deviceTopoInfoVO"`
}

/*
 * param regionId: 区域id (Required)
 * param instanceId: 实例Id (Required)
 * param deviceTopoInfoVO: 方法查询请求 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateDeviceTopoRequest(
    regionId string,
    instanceId string,
    deviceTopoInfoVO *iotcore.DeviceTopoInfoVO,
) *CreateDeviceTopoRequest {

	return &CreateDeviceTopoRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/coreinstances/{instanceId}/deviceTopo:create",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        DeviceTopoInfoVO: deviceTopoInfoVO,
	}
}

/*
 * param regionId: 区域id (Required)
 * param instanceId: 实例Id (Required)
 * param deviceTopoInfoVO: 方法查询请求 (Required)
 */
func NewCreateDeviceTopoRequestWithAllParams(
    regionId string,
    instanceId string,
    deviceTopoInfoVO *iotcore.DeviceTopoInfoVO,
) *CreateDeviceTopoRequest {

    return &CreateDeviceTopoRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/coreinstances/{instanceId}/deviceTopo:create",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        DeviceTopoInfoVO: deviceTopoInfoVO,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateDeviceTopoRequestWithoutParam() *CreateDeviceTopoRequest {

    return &CreateDeviceTopoRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/coreinstances/{instanceId}/deviceTopo:create",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 区域id(Required) */
func (r *CreateDeviceTopoRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 实例Id(Required) */
func (r *CreateDeviceTopoRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param deviceTopoInfoVO: 方法查询请求(Required) */
func (r *CreateDeviceTopoRequest) SetDeviceTopoInfoVO(deviceTopoInfoVO *iotcore.DeviceTopoInfoVO) {
    r.DeviceTopoInfoVO = deviceTopoInfoVO
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateDeviceTopoRequest) GetRegionId() string {
    return r.RegionId
}

type CreateDeviceTopoResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateDeviceTopoResult `json:"result"`
}

type CreateDeviceTopoResult struct {
    DeviceTopoResult iotcore.DeviceTopoResult `json:"deviceTopoResult"`
}