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
    edcps "github.com/lshuining/jdcloud-sdk-go/services/edcps/models"
)

type DescribeElasticIpRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* 弹性公网IPID  */
    ElasticIpId string `json:"elasticIpId"`
}

/*
 * param regionId: 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域 (Required)
 * param elasticIpId: 弹性公网IPID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeElasticIpRequest(
    regionId string,
    elasticIpId string,
) *DescribeElasticIpRequest {

	return &DescribeElasticIpRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/elasticIps/{elasticIpId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ElasticIpId: elasticIpId,
	}
}

/*
 * param regionId: 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域 (Required)
 * param elasticIpId: 弹性公网IPID (Required)
 */
func NewDescribeElasticIpRequestWithAllParams(
    regionId string,
    elasticIpId string,
) *DescribeElasticIpRequest {

    return &DescribeElasticIpRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/elasticIps/{elasticIpId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ElasticIpId: elasticIpId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeElasticIpRequestWithoutParam() *DescribeElasticIpRequest {

    return &DescribeElasticIpRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/elasticIps/{elasticIpId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域(Required) */
func (r *DescribeElasticIpRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param elasticIpId: 弹性公网IPID(Required) */
func (r *DescribeElasticIpRequest) SetElasticIpId(elasticIpId string) {
    r.ElasticIpId = elasticIpId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeElasticIpRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeElasticIpResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeElasticIpResult `json:"result"`
}

type DescribeElasticIpResult struct {
    ElasticIp edcps.ElasticIp `json:"elasticIp"`
}