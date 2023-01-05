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

type ModifySubnetRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Subnet ID  */
    SubnetId string `json:"subnetId"`

    /* 子网名称,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符。 (Optional) */
    SubnetName *string `json:"subnetName"`

    /* 子网描述信息，允许输入UTF-8编码下的全部字符，不超过256字符。 (Optional) */
    Description *string `json:"description"`

    /* 子网内预留网段掩码长度，此网段IP地址按照单个申请，子网内其余部分IP地址以网段形式分配。此参数非必选，缺省值为0，代表子网内所有IP地址都按照单个申请 (Optional) */
    IpMaskLen *int `json:"ipMaskLen"`
}

/*
 * param regionId: Region ID (Required)
 * param subnetId: Subnet ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifySubnetRequest(
    regionId string,
    subnetId string,
) *ModifySubnetRequest {

	return &ModifySubnetRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/subnets/{subnetId}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        SubnetId: subnetId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param subnetId: Subnet ID (Required)
 * param subnetName: 子网名称,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符。 (Optional)
 * param description: 子网描述信息，允许输入UTF-8编码下的全部字符，不超过256字符。 (Optional)
 * param ipMaskLen: 子网内预留网段掩码长度，此网段IP地址按照单个申请，子网内其余部分IP地址以网段形式分配。此参数非必选，缺省值为0，代表子网内所有IP地址都按照单个申请 (Optional)
 */
func NewModifySubnetRequestWithAllParams(
    regionId string,
    subnetId string,
    subnetName *string,
    description *string,
    ipMaskLen *int,
) *ModifySubnetRequest {

    return &ModifySubnetRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/subnets/{subnetId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        SubnetId: subnetId,
        SubnetName: subnetName,
        Description: description,
        IpMaskLen: ipMaskLen,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifySubnetRequestWithoutParam() *ModifySubnetRequest {

    return &ModifySubnetRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/subnets/{subnetId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *ModifySubnetRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param subnetId: Subnet ID(Required) */
func (r *ModifySubnetRequest) SetSubnetId(subnetId string) {
    r.SubnetId = subnetId
}
/* param subnetName: 子网名称,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符。(Optional) */
func (r *ModifySubnetRequest) SetSubnetName(subnetName string) {
    r.SubnetName = &subnetName
}
/* param description: 子网描述信息，允许输入UTF-8编码下的全部字符，不超过256字符。(Optional) */
func (r *ModifySubnetRequest) SetDescription(description string) {
    r.Description = &description
}
/* param ipMaskLen: 子网内预留网段掩码长度，此网段IP地址按照单个申请，子网内其余部分IP地址以网段形式分配。此参数非必选，缺省值为0，代表子网内所有IP地址都按照单个申请(Optional) */
func (r *ModifySubnetRequest) SetIpMaskLen(ipMaskLen int) {
    r.IpMaskLen = &ipMaskLen
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifySubnetRequest) GetRegionId() string {
    return r.RegionId
}

type ModifySubnetResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifySubnetResult `json:"result"`
}

type ModifySubnetResult struct {
}