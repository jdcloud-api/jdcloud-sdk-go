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

type ModifyVpcRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* 私有网络ID  */
    VpcId string `json:"vpcId"`

    /* 名称 (Optional) */
    Name *string `json:"name"`

    /* 描述 (Optional) */
    Description *string `json:"description"`
}

/*
 * param regionId: 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域 (Required)
 * param vpcId: 私有网络ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyVpcRequest(
    regionId string,
    vpcId string,
) *ModifyVpcRequest {

	return &ModifyVpcRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/vpcs/{vpcId}",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        VpcId: vpcId,
	}
}

/*
 * param regionId: 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域 (Required)
 * param vpcId: 私有网络ID (Required)
 * param name: 名称 (Optional)
 * param description: 描述 (Optional)
 */
func NewModifyVpcRequestWithAllParams(
    regionId string,
    vpcId string,
    name *string,
    description *string,
) *ModifyVpcRequest {

    return &ModifyVpcRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vpcs/{vpcId}",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        VpcId: vpcId,
        Name: name,
        Description: description,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyVpcRequestWithoutParam() *ModifyVpcRequest {

    return &ModifyVpcRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vpcs/{vpcId}",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域(Required) */
func (r *ModifyVpcRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param vpcId: 私有网络ID(Required) */
func (r *ModifyVpcRequest) SetVpcId(vpcId string) {
    r.VpcId = vpcId
}

/* param name: 名称(Optional) */
func (r *ModifyVpcRequest) SetName(name string) {
    r.Name = &name
}

/* param description: 描述(Optional) */
func (r *ModifyVpcRequest) SetDescription(description string) {
    r.Description = &description
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyVpcRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyVpcResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyVpcResult `json:"result"`
}

type ModifyVpcResult struct {
    Vpc edcps.Vpc `json:"vpc"`
}