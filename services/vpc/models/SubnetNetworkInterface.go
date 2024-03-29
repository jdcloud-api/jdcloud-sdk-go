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


type SubnetNetworkInterface struct {

    /* 弹性网卡ID (Optional) */
    NetworkInterfaceId string `json:"networkInterfaceId"`

    /* 子网ID (Optional) */
    SubnetId string `json:"subnetId"`

    /* 内网ip (Optional) */
    PrivateIpAddress string `json:"privateIpAddress"`

    /* 预分配的辅助ip段 (Optional) */
    SecondaryCidrs []string `json:"secondaryCidrs"`

    /* 资源ID (Optional) */
    ResourceId string `json:"resourceId"`

    /* 关联服务类型 (Optional) */
    ServiceType string `json:"serviceType"`

    /* 资源名称 (Optional) */
    ResourceName string `json:"resourceName"`
}
