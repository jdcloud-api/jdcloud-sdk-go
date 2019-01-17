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
    jdfusion "github.com/jdcloud-api/jdcloud-sdk-go/services/jdfusion/models"
)

type ReverseDeploymentRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 根据已有资源反向生成deployment  */
    Reverse *jdfusion.ReverseDeploymentInfo `json:"reverse"`
}

/*
 * param regionId: 地域ID (Required)
 * param reverse: 根据已有资源反向生成deployment (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewReverseDeploymentRequest(
    regionId string,
    reverse *jdfusion.ReverseDeploymentInfo,
) *ReverseDeploymentRequest {

	return &ReverseDeploymentRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/deployments:reverse",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Reverse: reverse,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param reverse: 根据已有资源反向生成deployment (Required)
 */
func NewReverseDeploymentRequestWithAllParams(
    regionId string,
    reverse *jdfusion.ReverseDeploymentInfo,
) *ReverseDeploymentRequest {

    return &ReverseDeploymentRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/deployments:reverse",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Reverse: reverse,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewReverseDeploymentRequestWithoutParam() *ReverseDeploymentRequest {

    return &ReverseDeploymentRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/deployments:reverse",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *ReverseDeploymentRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param reverse: 根据已有资源反向生成deployment(Required) */
func (r *ReverseDeploymentRequest) SetReverse(reverse *jdfusion.ReverseDeploymentInfo) {
    r.Reverse = reverse
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ReverseDeploymentRequest) GetRegionId() string {
    return r.RegionId
}

type ReverseDeploymentResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ReverseDeploymentResult `json:"result"`
}

type ReverseDeploymentResult struct {
    Deployment jdfusion.DeploymentInfo `json:"deployment"`
}