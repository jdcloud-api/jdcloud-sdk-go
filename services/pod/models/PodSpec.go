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

type PodSpec struct {

    /* Pod名称  */
    Name string `json:"name"`

    /* 描述信息，默认为空；允许输入UTF-8编码下的全部字符，不超过256字符。 (Optional) */
    Description *string `json:"description"`

    /* 主机名；范围：[1-63]个ASCII字符，默认值为 podId (Optional) */
    Hostname *string `json:"hostname"`

    /* pod中容器重启策略；Always, OnFailure, Never；默认：Always (Optional) */
    RestartPolicy *string `json:"restartPolicy"`

    /* 优雅关机宽限时长，如果超时，则触发强制关机。默认：30s，值不能是负数，范围：[0-300] (Optional) */
    TerminationGracePeriodSeconds *int `json:"terminationGracePeriodSeconds"`

    /* 实例类型；参考[文档](https://www.jdcloud.com/help/detail/1992/isCatalog/1)  */
    InstanceType string `json:"instanceType"`

    /* 容器所属可用区  */
    Az string `json:"az"`

    /* pod内容器的/etc/resolv.conf配置 (Optional) */
    DnsConfig *DnsConfigSpec `json:"dnsConfig"`

    /* 容器日志配置信息；默认会在本地分配10MB的存储空间 (Optional) */
    LogConfig *LogConfigSpec `json:"logConfig"`

    /* 域名和IP映射的信息；</br> 最大10个alias (Optional) */
    HostAliases []HostAliasSpec `json:"hostAliases"`

    /* Pod的volume列表，可以挂载到container上。长度范围：[0,7] (Optional) */
    Volumes []VolumeSpec `json:"volumes"`

    /* Pod的容器列表，至少一个容器。长度范围[1,8]  */
    Containers []ContainerSpec `json:"containers"`

    /* 计费模式：包年包月预付费（prepaid_by_duration）, 按配置后付费（postpaid_by_duration）。默认：按配置后付费 (Optional) */
    Charge *charge.ChargeSpec `json:"charge"`

    /* 主网卡主IP关联的弹性IP规格 (Optional) */
    ElasticIp *ElasticIpSpec `json:"elasticIp"`

    /* 主网卡配置信息  */
    PrimaryNetworkInterface *NetworkInterfaceAttachmentSpec `json:"primaryNetworkInterface"`
}
