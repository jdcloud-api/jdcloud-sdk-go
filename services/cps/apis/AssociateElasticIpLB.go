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

type AssociateElasticIpLBRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* 负载均衡实例ID  */
    LoadBalancerId string `json:"loadBalancerId"`

    /* 由客户端生成，用于保证请求的幂等性，长度不能超过36个字符；<br/>
如果多个请求使用了相同的clientToken，只会执行第一个请求，之后的请求直接返回第一个请求的结果<br/>
 (Optional) */
    ClientToken *string `json:"clientToken"`

    /* 弹性公网IPId  */
    ElasticIpId string `json:"elasticIpId"`
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
 * param loadBalancerId: 负载均衡实例ID (Required)
 * param elasticIpId: 弹性公网IPId (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAssociateElasticIpLBRequest(
    regionId string,
    loadBalancerId string,
    elasticIpId string,
) *AssociateElasticIpLBRequest {

	return &AssociateElasticIpLBRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/slbs/{loadBalancerId}:associateElasticIp",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        LoadBalancerId: loadBalancerId,
        ElasticIpId: elasticIpId,
	}
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
 * param loadBalancerId: 负载均衡实例ID (Required)
 * param clientToken: 由客户端生成，用于保证请求的幂等性，长度不能超过36个字符；<br/>
如果多个请求使用了相同的clientToken，只会执行第一个请求，之后的请求直接返回第一个请求的结果<br/>
 (Optional)
 * param elasticIpId: 弹性公网IPId (Required)
 */
func NewAssociateElasticIpLBRequestWithAllParams(
    regionId string,
    loadBalancerId string,
    clientToken *string,
    elasticIpId string,
) *AssociateElasticIpLBRequest {

    return &AssociateElasticIpLBRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/slbs/{loadBalancerId}:associateElasticIp",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        LoadBalancerId: loadBalancerId,
        ClientToken: clientToken,
        ElasticIpId: elasticIpId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAssociateElasticIpLBRequestWithoutParam() *AssociateElasticIpLBRequest {

    return &AssociateElasticIpLBRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/slbs/{loadBalancerId}:associateElasticIp",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域(Required) */
func (r *AssociateElasticIpLBRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param loadBalancerId: 负载均衡实例ID(Required) */
func (r *AssociateElasticIpLBRequest) SetLoadBalancerId(loadBalancerId string) {
    r.LoadBalancerId = loadBalancerId
}

/* param clientToken: 由客户端生成，用于保证请求的幂等性，长度不能超过36个字符；<br/>
如果多个请求使用了相同的clientToken，只会执行第一个请求，之后的请求直接返回第一个请求的结果<br/>
(Optional) */
func (r *AssociateElasticIpLBRequest) SetClientToken(clientToken string) {
    r.ClientToken = &clientToken
}

/* param elasticIpId: 弹性公网IPId(Required) */
func (r *AssociateElasticIpLBRequest) SetElasticIpId(elasticIpId string) {
    r.ElasticIpId = elasticIpId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AssociateElasticIpLBRequest) GetRegionId() string {
    return r.RegionId
}

type AssociateElasticIpLBResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AssociateElasticIpLBResult `json:"result"`
}

type AssociateElasticIpLBResult struct {
    Success bool `json:"success"`
}