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


type DescribeBandwidth struct {

    /* 机房英文标识 (Optional) */
    Idc string `json:"idc"`

    /* 机房名称 (Optional) */
    IdcName string `json:"idcName"`

    /* 带宽实例ID (Optional) */
    BandwidthId string `json:"bandwidthId"`

    /* 带宽（出口）名称 (Optional) */
    BandwidthName string `json:"bandwidthName"`

    /* 状态 normal:正常 abnormal:异常 (Optional) */
    Status string `json:"status"`

    /* 线路类型 dynamicBGP:动态BGP thirdLineBGP:三线BGP telecom:电信单线 unicom:联通单线 mobile:移动单线 (Optional) */
    LineType string `json:"lineType"`

    /* 计费方式
fixedBandwidth:固定带宽
95thPercentile:95峰值（IN，OUT统一计算95）
merge95thPercentile:95峰值（多出口合并计费）
95thPercentileSeparate:95峰值（IN，OUT分别计算95，取较大者）
merge95thPercentileAvg:日95峰值月平均（多出口合并计费）
 (Optional) */
    ChargeType string `json:"chargeType"`

    /* 合同带宽（Mbps） (Optional) */
    Bandwidth int `json:"bandwidth"`

    /* 关联的公网IP (Optional) */
    RelatedIp []RelatedIp `json:"relatedIp"`

    /* 交换机信息 (Optional) */
    Switchboard []Switchboard `json:"switchboard"`
}
