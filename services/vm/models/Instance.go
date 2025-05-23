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
import disk "github.com/jdcloud-api/jdcloud-sdk-go/services/disk/models"

type Instance struct {

    /* 云主机ID。 (Optional) */
    InstanceId string `json:"instanceId"`

    /* 云主机名称。 (Optional) */
    InstanceName string `json:"instanceName"`

    /* 云主机hostname。 (Optional) */
    Hostname string `json:"hostname"`

    /* 实例规格。 (Optional) */
    InstanceType string `json:"instanceType"`

    /* 主网卡所属VPC的ID。 (Optional) */
    VpcId string `json:"vpcId"`

    /* 主网卡所属子网的ID。 (Optional) */
    SubnetId string `json:"subnetId"`

    /* 主网卡主内网IP地址。 (Optional) */
    PrivateIpAddress string `json:"privateIpAddress"`

    /* 主网卡主IP绑定弹性IP的ID。 (Optional) */
    ElasticIpId string `json:"elasticIpId"`

    /* 主网卡主IP绑定弹性IP的地址。 (Optional) */
    ElasticIpAddress string `json:"elasticIpAddress"`

    /* 云主机状态，参考 [云主机状态](https://docs.jdcloud.com/virtual-machines/api/vm_status)。 (Optional) */
    Status string `json:"status"`

    /* 云主机描述。 (Optional) */
    Description string `json:"description"`

    /* 云主机使用的镜像ID。 (Optional) */
    ImageId string `json:"imageId"`

    /* 系统盘配置。 (Optional) */
    SystemDisk InstanceDiskAttachment `json:"systemDisk"`

    /* 数据盘配置列表。 (Optional) */
    DataDisks []InstanceDiskAttachment `json:"dataDisks"`

    /* 主网卡主IP关联的弹性公网IP配置。 (Optional) */
    PrimaryNetworkInterface InstanceNetworkInterfaceAttachment `json:"primaryNetworkInterface"`

    /* 辅助网卡配置列表。 (Optional) */
    SecondaryNetworkInterfaces []InstanceNetworkInterfaceAttachment `json:"secondaryNetworkInterfaces"`

    /* RDMA网卡配置列表。 (Optional) */
    RdmaNetworkInterfaces []RdmaNetworkInterface `json:"rdmaNetworkInterfaces"`

    /* 云主机实例的创建时间。 (Optional) */
    LaunchTime string `json:"launchTime"`

    /* 云主机所在可用区。 (Optional) */
    Az string `json:"az"`

    /* 云主机使用的密钥对名称。 (Optional) */
    KeyNames []string `json:"keyNames"`

    /* 云主机的计费信息。 (Optional) */
    Charge charge.Charge `json:"charge"`

    /* 云主机使用镜像的计费配置与信息。 (Optional) */
    ImageCharge ImageCharge `json:"imageCharge"`

    /* 抢占实例状态机 (Optional) */
    SpotStatus string `json:"spotStatus"`

    /* 云主机关联的高可用组，如果创建云主机使用了高可用组，此处可展示高可用组名称。 (Optional) */
    Ag Ag `json:"ag"`

    /* 高可用组中的错误域。 (Optional) */
    FaultDomain string `json:"faultDomain"`

    /* Tag信息。 (Optional) */
    Tags []disk.Tag `json:"tags"`

    /* 停机不计费模式。该参数仅对按配置计费且系统盘为云硬盘的实例生效，并且不是专有宿主机中的实例。
`keepCharging`：关机后继续计费。
`stopCharging`：关机后停止计费。
 (Optional) */
    ChargeOnStopped string `json:"chargeOnStopped"`

    /* 自动任务策略，关联了自动任务策略时可获取相应信息。 (Optional) */
    Policies []Policy `json:"policies"`

    /* 云主机所属的专有宿主机池。 (Optional) */
    DedicatedPoolId string `json:"dedicatedPoolId"`

    /* 云主机所属的专有宿主机ID。 (Optional) */
    DedicatedHostId string `json:"dedicatedHostId"`

    /* 突发型实例参数信息 (Optional) */
    BurstInfo BurstInfo `json:"burstInfo"`

    /* 资源组ID (Optional) */
    ResourceGroupId string `json:"resourceGroupId"`

    /* 云主机操作系统类型，如linux或者windows (Optional) */
    OsType string `json:"osType"`

    /* 虚机CPU拓扑 (Optional) */
    CpuTopology CpuTopology `json:"cpuTopology"`

    /* 云主机操作系统版本，如7.6 (Optional) */
    OsVersion string `json:"osVersion"`

    /* 具体操作系统CentOS (Optional) */
    Platform string `json:"platform"`

    /* 架构信息，如x86_64 (Optional) */
    Architecture string `json:"architecture"`

    /* 主机组详情 (Optional) */
    HostGroup HostGroup `json:"hostGroup"`

    /* 定时删除时间，例如:"2025-01-01 00:00:00" (Optional) */
    AutoReleaseTime string `json:"autoReleaseTime"`

    /* 启动模式，支持 bios uefi (Optional) */
    BootMode string `json:"bootMode"`
}
