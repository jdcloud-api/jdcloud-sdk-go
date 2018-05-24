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

import . "github.com/jdcloud-api/jdcloud-sdk-go/services/charge/models"

type ElasticIp struct {

    /* 弹性IP的Id (Optional) */
    ElasticIpId string `json:"elasticIpId"`

    /* 弹性IP地址 (Optional) */
    ElasticIpAddress string `json:"elasticIpAddress"`

    /* 弹性ip的限速（单位：Mb) (Optional) */
    BandwidthMbps int `json:"bandwidthMbps"`

    /* IP服务商，取值为bgp或no_bgp (Optional) */
    Provider string `json:"provider"`

    /* 私有IP的IPV4地址 (Optional) */
    PrivateIpAddress string `json:"privateIpAddress"`

    /* 配置弹性网卡Id (Optional) */
    NetworkInterfaceId string `json:"networkInterfaceId"`

    /* 实例Id (Optional) */
    InstanceId string `json:"instanceId"`

    /* 实例类型 (Optional) */
    InstanceType string `json:"instanceType"`

    /* 计费配置 (Optional) */
    Charge Charge `json:"charge"`

    /* 弹性ip创建时间 (Optional) */
    CreatedTime string `json:"createdTime"`
}
