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

type ModifyElasticIpBandwidthRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* 弹性公网IPID  */
    ElasticIpId string `json:"elasticIpId"`

    /* 由客户端生成，用于保证请求的幂等性，长度不能超过36个字符；<br/>
如果多个请求使用了相同的clientToken，只会执行第一个请求，之后的请求直接返回第一个请求的结果<br/>
 (Optional) */
    ClientToken *string `json:"clientToken"`

    /* 带宽，单位Mbps，取值范围[1,10240]  */
    Bandwidth int `json:"bandwidth"`

    /* 额外上行带宽，单位Mbps，取值范围[0,10240] (Optional) */
    ExtraUplinkBandwidth *int `json:"extraUplinkBandwidth"`
}

/*
 * param regionId: 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域 (Required)
 * param elasticIpId: 弹性公网IPID (Required)
 * param bandwidth: 带宽，单位Mbps，取值范围[1,10240] (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyElasticIpBandwidthRequest(
    regionId string,
    elasticIpId string,
    bandwidth int,
) *ModifyElasticIpBandwidthRequest {

	return &ModifyElasticIpBandwidthRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/elasticIps/{elasticIpId}:modifyElasticIpBandwidth",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ElasticIpId: elasticIpId,
        Bandwidth: bandwidth,
	}
}

/*
 * param regionId: 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域 (Required)
 * param elasticIpId: 弹性公网IPID (Required)
 * param clientToken: 由客户端生成，用于保证请求的幂等性，长度不能超过36个字符；<br/>
如果多个请求使用了相同的clientToken，只会执行第一个请求，之后的请求直接返回第一个请求的结果<br/>
 (Optional)
 * param bandwidth: 带宽，单位Mbps，取值范围[1,10240] (Required)
 * param extraUplinkBandwidth: 额外上行带宽，单位Mbps，取值范围[0,10240] (Optional)
 */
func NewModifyElasticIpBandwidthRequestWithAllParams(
    regionId string,
    elasticIpId string,
    clientToken *string,
    bandwidth int,
    extraUplinkBandwidth *int,
) *ModifyElasticIpBandwidthRequest {

    return &ModifyElasticIpBandwidthRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/elasticIps/{elasticIpId}:modifyElasticIpBandwidth",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ElasticIpId: elasticIpId,
        ClientToken: clientToken,
        Bandwidth: bandwidth,
        ExtraUplinkBandwidth: extraUplinkBandwidth,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyElasticIpBandwidthRequestWithoutParam() *ModifyElasticIpBandwidthRequest {

    return &ModifyElasticIpBandwidthRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/elasticIps/{elasticIpId}:modifyElasticIpBandwidth",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域(Required) */
func (r *ModifyElasticIpBandwidthRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param elasticIpId: 弹性公网IPID(Required) */
func (r *ModifyElasticIpBandwidthRequest) SetElasticIpId(elasticIpId string) {
    r.ElasticIpId = elasticIpId
}

/* param clientToken: 由客户端生成，用于保证请求的幂等性，长度不能超过36个字符；<br/>
如果多个请求使用了相同的clientToken，只会执行第一个请求，之后的请求直接返回第一个请求的结果<br/>
(Optional) */
func (r *ModifyElasticIpBandwidthRequest) SetClientToken(clientToken string) {
    r.ClientToken = &clientToken
}

/* param bandwidth: 带宽，单位Mbps，取值范围[1,10240](Required) */
func (r *ModifyElasticIpBandwidthRequest) SetBandwidth(bandwidth int) {
    r.Bandwidth = bandwidth
}

/* param extraUplinkBandwidth: 额外上行带宽，单位Mbps，取值范围[0,10240](Optional) */
func (r *ModifyElasticIpBandwidthRequest) SetExtraUplinkBandwidth(extraUplinkBandwidth int) {
    r.ExtraUplinkBandwidth = &extraUplinkBandwidth
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyElasticIpBandwidthRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyElasticIpBandwidthResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyElasticIpBandwidthResult `json:"result"`
}

type ModifyElasticIpBandwidthResult struct {
    Success bool `json:"success"`
}