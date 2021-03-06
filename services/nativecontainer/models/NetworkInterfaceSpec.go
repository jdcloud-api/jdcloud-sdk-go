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


type NetworkInterfaceSpec struct {

    /* 子网ID  */
    SubnetId string `json:"subnetId"`

    /* 可用区，用户的默认可用区，暂不支持  */
    Az string `json:"az"`

    /* 网卡主IP (Optional) */
    PrimaryIpAddress *string `json:"primaryIpAddress"`

    /* 网卡辅助IP，暂不支持 (Optional) */
    SecondaryIpAddresses []string `json:"secondaryIpAddresses"`

    /* 自动分配的辅助Ip数量，暂不支持 (Optional) */
    SecondaryIpCount *int `json:"secondaryIpCount"`

    /* 要绑定的安全组ID列表，最多指定5个安全组 (Optional) */
    SecurityGroups []string `json:"securityGroups"`

    /* 源和目标IP地址校验，取值为0或者1，默认为1，暂不支持此功能 (Optional) */
    SanityCheck *bool `json:"sanityCheck"`

    /* 描述，最大长度256字符 (Optional) */
    Description *string `json:"description"`
}
