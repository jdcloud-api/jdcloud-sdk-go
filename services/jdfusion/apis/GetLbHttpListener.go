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
    jdfusion "github.com/lshuining/jdcloud-sdk-go/services/jdfusion/models"
)

type GetLbHttpListenerRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 负载均衡id (Optional) */
    Slbid *string `json:"slbid"`
}

/*
 * param regionId: 地域ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetLbHttpListenerRequest(
    regionId string,
) *GetLbHttpListenerRequest {

	return &GetLbHttpListenerRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/vpc_lbHttpListener",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param slbid: 负载均衡id (Optional)
 */
func NewGetLbHttpListenerRequestWithAllParams(
    regionId string,
    slbid *string,
) *GetLbHttpListenerRequest {

    return &GetLbHttpListenerRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vpc_lbHttpListener",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Slbid: slbid,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetLbHttpListenerRequestWithoutParam() *GetLbHttpListenerRequest {

    return &GetLbHttpListenerRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vpc_lbHttpListener",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *GetLbHttpListenerRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param slbid: 负载均衡id(Optional) */
func (r *GetLbHttpListenerRequest) SetSlbid(slbid string) {
    r.Slbid = &slbid
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetLbHttpListenerRequest) GetRegionId() string {
    return r.RegionId
}

type GetLbHttpListenerResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetLbHttpListenerResult `json:"result"`
}

type GetLbHttpListenerResult struct {
    HttpListeners []jdfusion.LbHttpListener `json:"httpListeners"`
}