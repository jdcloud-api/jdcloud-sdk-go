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


type NodeGroupSpec struct {

    /* 名称  */
    Name string `json:"name"`

    /*  (Optional) */
    Description *string `json:"description"`

    /* 工作节点组的信息  */
    NodeConfig *NodeConfigSpec `json:"nodeConfig"`

    /* 工作节点组的 az，必须为集群az的子集，默认为集群az (Optional) */
    Azs []string `json:"azs"`

    /* 工作节点组初始化大小，至少为1个  */
    InitialNodeCount int `json:"initialNodeCount"`

    /* 自动伸缩配置 (Optional) */
    CaConfig *CAConfigSpec `json:"caConfig"`
}
