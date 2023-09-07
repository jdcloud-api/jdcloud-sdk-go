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

type ModifyWhiteListRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* RDS 实例ID，唯一标识一个RDS实例  */
    InstanceId string `json:"instanceId"`

    /* 白名单名称，默认Default (Optional) */
    WhiteListName *string `json:"whiteListName"`

    /* IP或IP段，不同的IP/IP段之间用英文逗号分隔，例如0.0.0.0/0,192.168.0.10  */
    Ips string `json:"ips"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: RDS 实例ID，唯一标识一个RDS实例 (Required)
 * param ips: IP或IP段，不同的IP/IP段之间用英文逗号分隔，例如0.0.0.0/0,192.168.0.10 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyWhiteListRequest(
    regionId string,
    instanceId string,
    ips string,
) *ModifyWhiteListRequest {

	return &ModifyWhiteListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/whiteList",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        Ips: ips,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: RDS 实例ID，唯一标识一个RDS实例 (Required)
 * param whiteListName: 白名单名称，默认Default (Optional)
 * param ips: IP或IP段，不同的IP/IP段之间用英文逗号分隔，例如0.0.0.0/0,192.168.0.10 (Required)
 */
func NewModifyWhiteListRequestWithAllParams(
    regionId string,
    instanceId string,
    whiteListName *string,
    ips string,
) *ModifyWhiteListRequest {

    return &ModifyWhiteListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/whiteList",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        WhiteListName: whiteListName,
        Ips: ips,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyWhiteListRequestWithoutParam() *ModifyWhiteListRequest {

    return &ModifyWhiteListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/whiteList",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *ModifyWhiteListRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param instanceId: RDS 实例ID，唯一标识一个RDS实例(Required) */
func (r *ModifyWhiteListRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}
/* param whiteListName: 白名单名称，默认Default(Optional) */
func (r *ModifyWhiteListRequest) SetWhiteListName(whiteListName string) {
    r.WhiteListName = &whiteListName
}
/* param ips: IP或IP段，不同的IP/IP段之间用英文逗号分隔，例如0.0.0.0/0,192.168.0.10(Required) */
func (r *ModifyWhiteListRequest) SetIps(ips string) {
    r.Ips = ips
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyWhiteListRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyWhiteListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyWhiteListResult `json:"result"`
}

type ModifyWhiteListResult struct {
}