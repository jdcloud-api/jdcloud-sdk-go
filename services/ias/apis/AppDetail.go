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

type AppDetailRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /*   */
    ClientId string `json:"clientId"`
}

/*
 * param regionId:  (Required)
 * param clientId:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAppDetailRequest(
    regionId string,
    clientId string,
) *AppDetailRequest {

	return &AppDetailRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/operate_backend/app/{clientId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ClientId: clientId,
	}
}

/*
 * param regionId:  (Required)
 * param clientId:  (Required)
 */
func NewAppDetailRequestWithAllParams(
    regionId string,
    clientId string,
) *AppDetailRequest {

    return &AppDetailRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/operate_backend/app/{clientId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ClientId: clientId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAppDetailRequestWithoutParam() *AppDetailRequest {

    return &AppDetailRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/operate_backend/app/{clientId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *AppDetailRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param clientId: (Required) */
func (r *AppDetailRequest) SetClientId(clientId string) {
    r.ClientId = clientId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AppDetailRequest) GetRegionId() string {
    return r.RegionId
}

type AppDetailResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AppDetailResult `json:"result"`
}

type AppDetailResult struct {
    ClientId string `json:"clientId"`
    ClientName string `json:"clientName"`
    TokenEndpointAuthMethod string `json:"tokenEndpointAuthMethod"`
    GrantTypes string `json:"grantTypes"`
    ResponseTypes string `json:"responseTypes"`
    RedirectUris string `json:"redirectUris"`
    ClientUri string `json:"clientUri"`
    LogoUri string `json:"logoUri"`
    TosUri string `json:"tosUri"`
    PolicyUri string `json:"policyUri"`
    Scope string `json:"scope"`
    JwksUri string `json:"jwksUri"`
    Jwks string `json:"jwks"`
    Contacts string `json:"contacts"`
    Extension string `json:"extension"`
    AccessTokenValiditySeconds int `json:"accessTokenValiditySeconds"`
    RefreshTokenValiditySeconds int `json:"refreshTokenValiditySeconds"`
    MultiTenant bool `json:"multiTenant"`
    SecretUpdateTime int64 `json:"secretUpdateTime"`
    UpdateTime int64 `json:"updateTime"`
    CreateTime int64 `json:"createTime"`
    Account string `json:"account"`
    UserType string `json:"userType"`
    State string `json:"state"`
}