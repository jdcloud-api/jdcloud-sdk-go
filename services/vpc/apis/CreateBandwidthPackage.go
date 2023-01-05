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
    charge "github.com/jdcloud-api/jdcloud-sdk-go/services/charge/models"
    vpc "github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/models"
)

type CreateBandwidthPackageRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 名称，只支持中文、数字、大小写字母、英文下划线“_”及中划线“-”，且长度不超过32个字符  */
    Name string `json:"name"`

    /* 描述，长度不超过256个字符 (Optional) */
    Description *string `json:"description"`

    /* 共享带宽包带宽上限，取值范围200-5000，单位为Mbps，保底带宽 = 共享带宽包带宽上限 * 20%  */
    BandwidthMbps int `json:"bandwidthMbps"`

    /* 线路信息，默认bgp，目前只支持中心节点的BGP线路 (Optional) */
    Provider *string `json:"provider"`

    /* 计费配置。支持包年包月、按配置、按用量计费模式 (Optional) */
    ChargeSpec *charge.ChargeSpec `json:"chargeSpec"`

    /* 用户标签 (Optional) */
    UserTags []vpc.Tag `json:"userTags"`

    /* 资源所属资源组ID (Optional) */
    ResourceGroupId *string `json:"resourceGroupId"`
}

/*
 * param regionId: Region ID (Required)
 * param name: 名称，只支持中文、数字、大小写字母、英文下划线“_”及中划线“-”，且长度不超过32个字符 (Required)
 * param bandwidthMbps: 共享带宽包带宽上限，取值范围200-5000，单位为Mbps，保底带宽 = 共享带宽包带宽上限 * 20% (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateBandwidthPackageRequest(
    regionId string,
    name string,
    bandwidthMbps int,
) *CreateBandwidthPackageRequest {

	return &CreateBandwidthPackageRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/bandwidthPackages/",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Name: name,
        BandwidthMbps: bandwidthMbps,
	}
}

/*
 * param regionId: Region ID (Required)
 * param name: 名称，只支持中文、数字、大小写字母、英文下划线“_”及中划线“-”，且长度不超过32个字符 (Required)
 * param description: 描述，长度不超过256个字符 (Optional)
 * param bandwidthMbps: 共享带宽包带宽上限，取值范围200-5000，单位为Mbps，保底带宽 = 共享带宽包带宽上限 * 20% (Required)
 * param provider: 线路信息，默认bgp，目前只支持中心节点的BGP线路 (Optional)
 * param chargeSpec: 计费配置。支持包年包月、按配置、按用量计费模式 (Optional)
 * param userTags: 用户标签 (Optional)
 * param resourceGroupId: 资源所属资源组ID (Optional)
 */
func NewCreateBandwidthPackageRequestWithAllParams(
    regionId string,
    name string,
    description *string,
    bandwidthMbps int,
    provider *string,
    chargeSpec *charge.ChargeSpec,
    userTags []vpc.Tag,
    resourceGroupId *string,
) *CreateBandwidthPackageRequest {

    return &CreateBandwidthPackageRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/bandwidthPackages/",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Name: name,
        Description: description,
        BandwidthMbps: bandwidthMbps,
        Provider: provider,
        ChargeSpec: chargeSpec,
        UserTags: userTags,
        ResourceGroupId: resourceGroupId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateBandwidthPackageRequestWithoutParam() *CreateBandwidthPackageRequest {

    return &CreateBandwidthPackageRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/bandwidthPackages/",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CreateBandwidthPackageRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param name: 名称，只支持中文、数字、大小写字母、英文下划线“_”及中划线“-”，且长度不超过32个字符(Required) */
func (r *CreateBandwidthPackageRequest) SetName(name string) {
    r.Name = name
}
/* param description: 描述，长度不超过256个字符(Optional) */
func (r *CreateBandwidthPackageRequest) SetDescription(description string) {
    r.Description = &description
}
/* param bandwidthMbps: 共享带宽包带宽上限，取值范围200-5000，单位为Mbps，保底带宽 = 共享带宽包带宽上限 * 20%(Required) */
func (r *CreateBandwidthPackageRequest) SetBandwidthMbps(bandwidthMbps int) {
    r.BandwidthMbps = bandwidthMbps
}
/* param provider: 线路信息，默认bgp，目前只支持中心节点的BGP线路(Optional) */
func (r *CreateBandwidthPackageRequest) SetProvider(provider string) {
    r.Provider = &provider
}
/* param chargeSpec: 计费配置。支持包年包月、按配置、按用量计费模式(Optional) */
func (r *CreateBandwidthPackageRequest) SetChargeSpec(chargeSpec *charge.ChargeSpec) {
    r.ChargeSpec = chargeSpec
}
/* param userTags: 用户标签(Optional) */
func (r *CreateBandwidthPackageRequest) SetUserTags(userTags []vpc.Tag) {
    r.UserTags = userTags
}
/* param resourceGroupId: 资源所属资源组ID(Optional) */
func (r *CreateBandwidthPackageRequest) SetResourceGroupId(resourceGroupId string) {
    r.ResourceGroupId = &resourceGroupId
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateBandwidthPackageRequest) GetRegionId() string {
    return r.RegionId
}

type CreateBandwidthPackageResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateBandwidthPackageResult `json:"result"`
}

type CreateBandwidthPackageResult struct {
    BandwidthPackageId string `json:"bandwidthPackageId"`
    RequestId string `json:"requestId"`
}