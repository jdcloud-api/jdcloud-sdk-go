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

import charge "github.com/lshuining/jdcloud-sdk-go/services/charge/models"

type Instance struct {

    /* 云物理服务器实例ID (Optional) */
    InstanceId string `json:"instanceId"`

    /* 区域代码, 如 cn-north-1 (Optional) */
    Region string `json:"region"`

    /* 可用区, 如 cn-north-1a (Optional) */
    Az string `json:"az"`

    /* 实例类型, 如 cps.c.normal (Optional) */
    DeviceType string `json:"deviceType"`

    /* 云物理服务器名称 (Optional) */
    Name string `json:"name"`

    /* 云物理服务器描述 (Optional) */
    Description string `json:"description"`

    /* 云物理服务器生命周期状态 (Optional) */
    Status string `json:"status"`

    /* 是否启用外网, 如 yes/no (Optional) */
    EnableInternet string `json:"enableInternet"`

    /* 是否启用IPv6, 如 yes/no (Optional) */
    EnableIpv6 string `json:"enableIpv6"`

    /* 带宽, 单位Mbps (Optional) */
    Bandwidth int `json:"bandwidth"`

    /* 镜像类型, 如 standard (Optional) */
    ImageType string `json:"imageType"`

    /* 机柜信息 (Optional) */
    Cabinet string `json:"cabinet"`

    /* 带外管理IP (Optional) */
    IloIp string `json:"iloIp"`

    /* 操作系统类型ID (Optional) */
    OsTypeId string `json:"osTypeId"`

    /* 操作系统名称 (Optional) */
    OsName string `json:"osName"`

    /* 操作系统类型, 如 ubuntu/centos (Optional) */
    OsType string `json:"osType"`

    /* 操作系统版本, 如 16.04 (Optional) */
    OsVersion string `json:"osVersion"`

    /* 系统盘RAID类型ID (Optional) */
    SysRaidTypeId string `json:"sysRaidTypeId"`

    /* 系统盘RAID类型, 如 NORAID, RAID0, RAID1 (Optional) */
    SysRaidType string `json:"sysRaidType"`

    /* 数据盘RAID类型ID (Optional) */
    DataRaidTypeId string `json:"dataRaidTypeId"`

    /* 数据盘RAID类型, 如 NORAID, RAID0, RAID1，RAID10 (Optional) */
    DataRaidType string `json:"dataRaidType"`

    /* 网络类型：basic（基础网络）、vpc（私有网络）、retail（零售中台网络） (Optional) */
    NetworkType string `json:"networkType"`

    /* 私有网络ID (Optional) */
    VpcId string `json:"vpcId"`

    /* 私有网络名称 (Optional) */
    VpcName string `json:"vpcName"`

    /* 私有网络IPv4 CIDR (Optional) */
    VpcIpv4Cidr string `json:"vpcIpv4Cidr"`

    /* 私有网络IPv6 CIDR (Optional) */
    VpcIpv6Cidr string `json:"vpcIpv6Cidr"`

    /* IPv6网关ID (Optional) */
    Ipv6GatewayId string `json:"ipv6GatewayId"`

    /* POD网络名称 (Optional) */
    PodName string `json:"podName"`

    /* POD机房地址描述 (Optional) */
    PodRoom string `json:"podRoom"`

    /* 子网编号 (Optional) */
    SubnetId string `json:"subnetId"`

    /* 子网名称 (Optional) */
    SubnetName string `json:"subnetName"`

    /* 子网IPv4 CIDR (Optional) */
    SubnetIpv4Cidr string `json:"subnetIpv4Cidr"`

    /* 子网IPv6 CIDR (Optional) */
    SubnetIpv6Cidr string `json:"subnetIpv6Cidr"`

    /* 内网IP (Optional) */
    PrivateIp string `json:"privateIp"`

    /* 外网链路类型, 如 bgp (Optional) */
    LineType string `json:"lineType"`

    /* 弹性公网IPID (Optional) */
    ElasticIpId string `json:"elasticIpId"`

    /* 公网IP (Optional) */
    PublicIp string `json:"publicIp"`

    /* IPv6地址 (Optional) */
    Ipv6Address string `json:"ipv6Address"`

    /* 公网IPv6地址ID (Optional) */
    Ipv6AddressId string `json:"ipv6AddressId"`

    /* 公网IPv6带宽, 单位Mbps (Optional) */
    Ipv6AddressBandwidth int `json:"ipv6AddressBandwidth"`

    /* 网络接口模式：bond（网口bond）、dual（双网口） (Optional) */
    InterfaceMode string `json:"interfaceMode"`

    /* 辅网口私有网络ID (Optional) */
    ExtensionVpcId string `json:"extensionVpcId"`

    /* 辅网口私有网络名称 (Optional) */
    ExtensionVpcName string `json:"extensionVpcName"`

    /* 辅网口私有网络IPv4 CIDR (Optional) */
    ExtensionVpcIpv4Cidr string `json:"extensionVpcIpv4Cidr"`

    /* 辅网口私有网络IPv6 CIDR (Optional) */
    ExtensionVpcIpv6Cidr string `json:"extensionVpcIpv6Cidr"`

    /* 辅网口子网ID (Optional) */
    ExtensionSubnetId string `json:"extensionSubnetId"`

    /* 辅网口子网名称 (Optional) */
    ExtensionSubnetName string `json:"extensionSubnetName"`

    /* 辅网口子网IPv4 CIDR (Optional) */
    ExtensionSubnetIpv4Cidr string `json:"extensionSubnetIpv4Cidr"`

    /* 辅网口子网IPv6 CIDR (Optional) */
    ExtensionSubnetIpv6Cidr string `json:"extensionSubnetIpv6Cidr"`

    /* 辅网口手动分配的内网ip (Optional) */
    ExtensionPrivateIp string `json:"extensionPrivateIp"`

    /* 辅网口是否启用外网 (Optional) */
    ExtensionEnableInternet string `json:"extensionEnableInternet"`

    /* 辅网口弹性公网ip id (Optional) */
    ExtensionElasticIpId string `json:"extensionElasticIpId"`

    /* 辅网口公网ip (Optional) */
    ExtensionPublicIp string `json:"extensionPublicIp"`

    /* 辅网口外网带宽，单位Mbps (Optional) */
    ExtensionBandwidth int `json:"extensionBandwidth"`

    /* 辅网口是否启用IPv6, 如 yes/no (Optional) */
    ExtensionEnableIpv6 string `json:"extensionEnableIpv6"`

    /* 辅网口IPv6地址 (Optional) */
    ExtensionIpv6Address string `json:"extensionIpv6Address"`

    /* 辅网口公网IPv6地址ID (Optional) */
    ExtensionIpv6AddressId string `json:"extensionIpv6AddressId"`

    /* 辅网口IPv6公网带宽, 单位Mbps (Optional) */
    ExtensionIpv6AddressBandwidth int `json:"extensionIpv6AddressBandwidth"`

    /* IPv6网关ID (Optional) */
    ExtensionIpv6GatewayId string `json:"extensionIpv6GatewayId"`

    /* 密钥对id (Optional) */
    KeypairId string `json:"keypairId"`

    /* agent状态 (Optional) */
    AgentStatus string `json:"agentStatus"`

    /* 计费信息 (Optional) */
    Charge charge.Charge `json:"charge"`
}
