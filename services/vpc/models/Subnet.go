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


type Subnet struct {

    /* Subnet的Id (Optional) */
    SubnetId string `json:"subnetId"`

    /* 子网名称 (Optional) */
    SubnetName string `json:"subnetName"`

    /* 子网所属VPC的Id (Optional) */
    VpcId string `json:"vpcId"`

    /* 子网网段，vpc内子网网段不能重叠，cidr的取值范围：10.0.0.0/8、172.16.0.0/12和192.168.0.0/16及它们包含的子网，且子网掩码长度为16-28之间，如果VPC含有Cidr，则必须为VPC所在Cidr的子网 (Optional) */
    AddressPrefix string `json:"addressPrefix"`

    /* 子网可用ip数量 (Optional) */
    AvailableIpCount int `json:"availableIpCount"`

    /* 子网内预留网段掩码长度，此网段IP地址按照单个申请，子网内其余部分IP地址以网段形式分配。此参数非必选，缺省值为0，代表子网内所有IP地址都按照单个申请(范围是[max{24, 子网掩码}, 28]) (Optional) */
    IpMaskLen int `json:"ipMaskLen"`

    /* 子网描述信息 (Optional) */
    Description string `json:"description"`

    /* 子网关联的路由表Id (Optional) */
    RouteTableId string `json:"routeTableId"`

    /* 子网关联的acl Id (Optional) */
    AclId string `json:"aclId"`

    /* 子网的起始地址，子网第1个地位为路由器网关保留，第2个地址为dhcp服务保留 (Optional) */
    StartIp string `json:"startIp"`

    /* 子网的结束地址，子网第1个地位为路由器网关保留，第2个地址为dhcp服务保留 (Optional) */
    EndIp string `json:"endIp"`

    /* 子网创建时间 (Optional) */
    CreatedTime string `json:"createdTime"`

    /* 子网类型，取值：standard(标准子网)，edge(边缘子网) (Optional) */
    SubnetType string `json:"subnetType"`

    /* 子网可用区 (Optional) */
    Az string `json:"az"`

    /* 子网是否为外部子网（即子网路由表中存在下一跳为internet的路由）。true表示外部子网，false表示内部子网 (Optional) */
    PublicSubnet bool `json:"publicSubnet"`

    /* 域名后缀，不限制个数。总长度最长254个字符，仅支持字母，数字，中划线，下划线和点。 (Optional) */
    DomainNames []string `json:"domainNames"`

    /* 域名服务器地址。最多支持5个IPv4地址，不同IPv4地址使用逗号分隔。如不输入，默认使用京东云默认DNS域名服务器地址。如不添加默认DNS域名服务器，可能会导致您无法访问京东云云上基础服务，请谨慎操作 (Optional) */
    DomainNameServers []string `json:"domainNameServers"`
}
