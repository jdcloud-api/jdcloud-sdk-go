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


type NodeGroup struct {

    /* 集群 id (Optional) */
    ClusterId string `json:"clusterId"`

    /* 工作节点组 id (Optional) */
    NodeGroupId string `json:"nodeGroupId"`

    /* 工作节点组名称 (Optional) */
    Name string `json:"name"`

    /* 工作节点组描述 (Optional) */
    Description string `json:"description"`

    /* 工作节点组配置信息 (Optional) */
    NodeConfig NodeConfig `json:"nodeConfig"`

    /* 工作节点版本 (Optional) */
    Version string `json:"version"`

    /* 工作节点所属的网络信息 (Optional) */
    NodeNetwork NodeNetwork `json:"nodeNetwork"`

    /* 当前工作节点数量 (Optional) */
    CurrentCount int `json:"currentCount"`

    /* 期望的工作节点数量 (Optional) */
    ExpectCount int `json:"expectCount"`

    /* 工作节点组的ag id ，通过agid可以查询该工作节点组下的实例 (Optional) */
    AgId string `json:"agId"`

    /* 工作节点组所在的 az (Optional) */
    Azs []string `json:"azs"`

    /* 工作节点组的 ag 对应的实例模板 (Optional) */
    InstanceTemplateId string `json:"instanceTemplateId"`

    /* 状态  [pending,running,resizing,reconciling,deleting,deleted,error,running_with_error(部分节点有问题)] (Optional) */
    State string `json:"state"`

    /*  (Optional) */
    Tags []Tag `json:"tags"`

    /* 更新时间 (Optional) */
    UpdateTime string `json:"updateTime"`

    /* 状态变更原因 (Optional) */
    StateMessage string `json:"stateMessage"`

    /* 控制节点操作进度 (Optional) */
    Progress NodeGroupProgress `json:"progress"`

    /* 自动伸缩配置 (Optional) */
    CaConfig CAConfig `json:"caConfig"`

    /* 创建时间 (Optional) */
    CreatedTime string `json:"createdTime"`
}
