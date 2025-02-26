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


type PodBrief struct {

    /* pod ID (Optional) */
    PodId string `json:"podId"`

    /* pod 名称 (Optional) */
    Name string `json:"name"`

    /* 描述信息，默认为空。 (Optional) */
    Description string `json:"description"`

    /* 可用区 (Optional) */
    Az string `json:"az"`

    /* 主机名 (Optional) */
    Hostname string `json:"hostname"`

    /* 高可用组 (Optional) */
    Ag AvailablityGroup `json:"ag"`

    /* pod 所需的计算资源规格 (Optional) */
    InstanceType string `json:"instanceType"`

    /* pod重启策略 (Optional) */
    RestartPolicy string `json:"restartPolicy"`

    /* 优雅关闭的时间 (Optional) */
    TerminationGracePeriodSeconds int `json:"terminationGracePeriodSeconds"`

    /* 主网卡所属vpcId (Optional) */
    VpcId string `json:"vpcId"`

    /* 主网卡所属子网的ID (Optional) */
    SubnetId string `json:"subnetId"`

    /* 主网卡主IP地址 (Optional) */
    PrivateIpAddress string `json:"privateIpAddress"`

    /* pod内容器的/etc/resolv.conf配置 (Optional) */
    DnsConfig DnsConfig `json:"dnsConfig"`

    /* 容器日志配置信息；默认会在本地分配10MB的存储空间 (Optional) */
    LogConfig LogConfig `json:"logConfig"`

    /* pod内容器的/etc/hosts配置 (Optional) */
    HostAliases []HostAlias `json:"hostAliases"`

    /* pod状态信息 (Optional) */
    PodStatus PodStatus `json:"podStatus"`

    /* 主网卡配置信息 (Optional) */
    PrimaryNetworkInterface NetworkInterfaceAttachment `json:"primaryNetworkInterface"`

    /* Pod创建时间 (Optional) */
    CreateTime string `json:"createTime"`
}
