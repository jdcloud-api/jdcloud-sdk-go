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
    apigateway "github.com/lshuining/jdcloud-sdk-go/services/apigateway/models"
)

type QuerySubscriptionKeyRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* subscription key id  */
    SubscriptionKeyId string `json:"subscriptionKeyId"`
}

/*
 * param regionId: 地域ID (Required)
 * param subscriptionKeyId: subscription key id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQuerySubscriptionKeyRequest(
    regionId string,
    subscriptionKeyId string,
) *QuerySubscriptionKeyRequest {

	return &QuerySubscriptionKeyRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/subscriptionKeys/{subscriptionKeyId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        SubscriptionKeyId: subscriptionKeyId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param subscriptionKeyId: subscription key id (Required)
 */
func NewQuerySubscriptionKeyRequestWithAllParams(
    regionId string,
    subscriptionKeyId string,
) *QuerySubscriptionKeyRequest {

    return &QuerySubscriptionKeyRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/subscriptionKeys/{subscriptionKeyId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        SubscriptionKeyId: subscriptionKeyId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQuerySubscriptionKeyRequestWithoutParam() *QuerySubscriptionKeyRequest {

    return &QuerySubscriptionKeyRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/subscriptionKeys/{subscriptionKeyId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *QuerySubscriptionKeyRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param subscriptionKeyId: subscription key id(Required) */
func (r *QuerySubscriptionKeyRequest) SetSubscriptionKeyId(subscriptionKeyId string) {
    r.SubscriptionKeyId = subscriptionKeyId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QuerySubscriptionKeyRequest) GetRegionId() string {
    return r.RegionId
}

type QuerySubscriptionKeyResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QuerySubscriptionKeyResult `json:"result"`
}

type QuerySubscriptionKeyResult struct {
    SubscriptionKey apigateway.SubscriptionKey `json:"subscriptionKey"`
}