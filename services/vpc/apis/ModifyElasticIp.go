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
)

type ModifyElasticIpRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* ElasticIp ID  */
    ElasticIpId string `json:"elasticIpId"`

    /* 弹性公网IP的限速（单位：Mbps），取值范围为[1-200]  */
    BandwidthMbps int `json:"bandwidthMbps"`
}

/*
 * param regionId: Region ID (Required)
 * param elasticIpId: ElasticIp ID (Required)
 * param bandwidthMbps: 弹性公网IP的限速（单位：Mbps），取值范围为[1-200] (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyElasticIpRequest(
    regionId string,
    elasticIpId string,
    bandwidthMbps int,
) *ModifyElasticIpRequest {

	return &ModifyElasticIpRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/elasticIps/{elasticIpId}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ElasticIpId: elasticIpId,
        BandwidthMbps: bandwidthMbps,
	}
}

/*
 * param regionId: Region ID (Required)
 * param elasticIpId: ElasticIp ID (Required)
 * param bandwidthMbps: 弹性公网IP的限速（单位：Mbps），取值范围为[1-200] (Required)
 */
func NewModifyElasticIpRequestWithAllParams(
    regionId string,
    elasticIpId string,
    bandwidthMbps int,
) *ModifyElasticIpRequest {

    return &ModifyElasticIpRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/elasticIps/{elasticIpId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ElasticIpId: elasticIpId,
        BandwidthMbps: bandwidthMbps,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyElasticIpRequestWithoutParam() *ModifyElasticIpRequest {

    return &ModifyElasticIpRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/elasticIps/{elasticIpId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *ModifyElasticIpRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param elasticIpId: ElasticIp ID(Required) */
func (r *ModifyElasticIpRequest) SetElasticIpId(elasticIpId string) {
    r.ElasticIpId = elasticIpId
}

/* param bandwidthMbps: 弹性公网IP的限速（单位：Mbps），取值范围为[1-200](Required) */
func (r *ModifyElasticIpRequest) SetBandwidthMbps(bandwidthMbps int) {
    r.BandwidthMbps = bandwidthMbps
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyElasticIpRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyElasticIpResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyElasticIpResult `json:"result"`
}

type ModifyElasticIpResult struct {
}