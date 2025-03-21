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

type StopInstanceTimerRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* 由客户端生成，用于保证请求的幂等性，长度不能超过36个字符；<br/>
如果多个请求使用了相同的clientToken，只会执行第一个请求，之后的请求直接返回第一个请求的结果<br/>
 (Optional) */
    ClientToken *string `json:"clientToken"`

    /* 实例ID  */
    InstanceId string `json:"instanceId"`

    /* 关机时间，格式：2024-09-04 14:22:00  */
    StopTime string `json:"stopTime"`
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
 * param instanceId: 实例ID (Required)
 * param stopTime: 关机时间，格式：2024-09-04 14:22:00 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewStopInstanceTimerRequest(
    regionId string,
    instanceId string,
    stopTime string,
) *StopInstanceTimerRequest {

	return &StopInstanceTimerRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/stopInstanceTimer",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        StopTime: stopTime,
	}
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
 * param clientToken: 由客户端生成，用于保证请求的幂等性，长度不能超过36个字符；<br/>
如果多个请求使用了相同的clientToken，只会执行第一个请求，之后的请求直接返回第一个请求的结果<br/>
 (Optional)
 * param instanceId: 实例ID (Required)
 * param stopTime: 关机时间，格式：2024-09-04 14:22:00 (Required)
 */
func NewStopInstanceTimerRequestWithAllParams(
    regionId string,
    clientToken *string,
    instanceId string,
    stopTime string,
) *StopInstanceTimerRequest {

    return &StopInstanceTimerRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/stopInstanceTimer",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ClientToken: clientToken,
        InstanceId: instanceId,
        StopTime: stopTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewStopInstanceTimerRequestWithoutParam() *StopInstanceTimerRequest {

    return &StopInstanceTimerRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/stopInstanceTimer",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域(Required) */
func (r *StopInstanceTimerRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param clientToken: 由客户端生成，用于保证请求的幂等性，长度不能超过36个字符；<br/>
如果多个请求使用了相同的clientToken，只会执行第一个请求，之后的请求直接返回第一个请求的结果<br/>
(Optional) */
func (r *StopInstanceTimerRequest) SetClientToken(clientToken string) {
    r.ClientToken = &clientToken
}
/* param instanceId: 实例ID(Required) */
func (r *StopInstanceTimerRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}
/* param stopTime: 关机时间，格式：2024-09-04 14:22:00(Required) */
func (r *StopInstanceTimerRequest) SetStopTime(stopTime string) {
    r.StopTime = stopTime
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r StopInstanceTimerRequest) GetRegionId() string {
    return r.RegionId
}

type StopInstanceTimerResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result StopInstanceTimerResult `json:"result"`
}

type StopInstanceTimerResult struct {
    Success bool `json:"success"`
}