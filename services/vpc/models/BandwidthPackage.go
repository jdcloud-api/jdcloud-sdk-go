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

package models

import charge "github.com/jdcloud-api/jdcloud-sdk-go/services/charge/models"

type BandwidthPackage struct {

    /* 共享带宽包ID (Optional) */
    BandwidthPackageId string `json:"bandwidthPackageId"`

    /* 名称 (Optional) */
    BandwidthPackageName string `json:"bandwidthPackageName"`

    /* 描述 (Optional) */
    Description string `json:"description"`

    /* 共享带宽包带宽上限，取值范围200-5000，单位为Mbps，按用量计费模式的保底带宽 = 共享带宽包带宽上限 * 20% (Optional) */
    BandwidthMbps int `json:"bandwidthMbps"`

    /* 线路信息 (Optional) */
    Provider string `json:"provider"`

    /* 计费配置 (Optional) */
    Charge charge.Charge `json:"charge"`

    /* 创建时间，时间格式为UTC (Optional) */
    CreatedTime string `json:"createdTime"`

    /* 共享带宽包内公网IP信息 (Optional) */
    PublicIps []BwpIp `json:"publicIps"`

    /* 共享带宽包内加入公网IP个数 (Optional) */
    IpCount int `json:"ipCount"`

    /* 按用量计费模式的保底带宽百分比，-1代表无效值 ，目前保底带宽百分比为20% (Optional) */
    GuaranteedRatio int `json:"guaranteedRatio"`

    /* 按用量计费模式的保底带宽，-1代表无效值，保底带宽 = 共享带宽包带宽上限 * 20% (Optional) */
    GuaranteedBandwidth int `json:"guaranteedBandwidth"`

    /* 是否欠费停服，UP正常，DOWN停服 (Optional) */
    AdminStatus string `json:"adminStatus"`

    /* Tag信息 (Optional) */
    Tags []Tag `json:"tags"`

    /* 资源所属资源组ID (Optional) */
    ResourceGroupId string `json:"resourceGroupId"`
}