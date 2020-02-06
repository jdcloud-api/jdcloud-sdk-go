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

type ModifyRouterByLowerRequest struct {

    core.JDCloudRequest

    /* 申请接入id  */
    AccessId int `json:"accessId"`

    /* 商家编码  */
    CustomerId string `json:"customerId"`

    /* 商家服务后端地址  */
    BackendUrl string `json:"backendUrl"`
}

/*
 * param accessId: 申请接入id (Required)
 * param customerId: 商家编码 (Required)
 * param backendUrl: 商家服务后端地址 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyRouterByLowerRequest(
    accessId int,
    customerId string,
    backendUrl string,
) *ModifyRouterByLowerRequest {

	return &ModifyRouterByLowerRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/accesses/{accessId}:update",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        AccessId: accessId,
        CustomerId: customerId,
        BackendUrl: backendUrl,
	}
}

/*
 * param accessId: 申请接入id (Required)
 * param customerId: 商家编码 (Required)
 * param backendUrl: 商家服务后端地址 (Required)
 */
func NewModifyRouterByLowerRequestWithAllParams(
    accessId int,
    customerId string,
    backendUrl string,
) *ModifyRouterByLowerRequest {

    return &ModifyRouterByLowerRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/accesses/{accessId}:update",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        AccessId: accessId,
        CustomerId: customerId,
        BackendUrl: backendUrl,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyRouterByLowerRequestWithoutParam() *ModifyRouterByLowerRequest {

    return &ModifyRouterByLowerRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/accesses/{accessId}:update",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param accessId: 申请接入id(Required) */
func (r *ModifyRouterByLowerRequest) SetAccessId(accessId int) {
    r.AccessId = accessId
}

/* param customerId: 商家编码(Required) */
func (r *ModifyRouterByLowerRequest) SetCustomerId(customerId string) {
    r.CustomerId = customerId
}

/* param backendUrl: 商家服务后端地址(Required) */
func (r *ModifyRouterByLowerRequest) SetBackendUrl(backendUrl string) {
    r.BackendUrl = backendUrl
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyRouterByLowerRequest) GetRegionId() string {
    return ""
}

type ModifyRouterByLowerResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyRouterByLowerResult `json:"result"`
}

type ModifyRouterByLowerResult struct {
    Success bool `json:"success"`
}