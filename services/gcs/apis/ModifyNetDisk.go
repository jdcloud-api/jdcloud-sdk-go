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

type ModifyNetDiskRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* 由客户端生成，用于保证请求的幂等性，长度不能超过36个字符；<br/>
如果多个请求使用了相同的clientToken，只会执行第一个请求，之后的请求直接返回第一个请求的结果<br/>
 (Optional) */
    ClientToken *string `json:"clientToken"`

    /* 可用区名称 (Optional) */
    Az *string `json:"az"`
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyNetDiskRequest(
    regionId string,
) *ModifyNetDiskRequest {

	return &ModifyNetDiskRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/modifyNetDisk",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
 * param clientToken: 由客户端生成，用于保证请求的幂等性，长度不能超过36个字符；<br/>
如果多个请求使用了相同的clientToken，只会执行第一个请求，之后的请求直接返回第一个请求的结果<br/>
 (Optional)
 * param az: 可用区名称 (Optional)
 */
func NewModifyNetDiskRequestWithAllParams(
    regionId string,
    clientToken *string,
    az *string,
) *ModifyNetDiskRequest {

    return &ModifyNetDiskRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/modifyNetDisk",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ClientToken: clientToken,
        Az: az,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyNetDiskRequestWithoutParam() *ModifyNetDiskRequest {

    return &ModifyNetDiskRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/modifyNetDisk",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域(Required) */
func (r *ModifyNetDiskRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param clientToken: 由客户端生成，用于保证请求的幂等性，长度不能超过36个字符；<br/>
如果多个请求使用了相同的clientToken，只会执行第一个请求，之后的请求直接返回第一个请求的结果<br/>
(Optional) */
func (r *ModifyNetDiskRequest) SetClientToken(clientToken string) {
    r.ClientToken = &clientToken
}
/* param az: 可用区名称(Optional) */
func (r *ModifyNetDiskRequest) SetAz(az string) {
    r.Az = &az
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyNetDiskRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyNetDiskResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyNetDiskResult `json:"result"`
}

type ModifyNetDiskResult struct {
    Success bool `json:"success"`
}