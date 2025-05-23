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


type K8sWatchHeartbeat struct {

    /* 集群 (Optional) */
    Cluster string `json:"cluster"`

    /* 设备id (Optional) */
    DevId string `json:"devId"`

    /* 公有云集群标识 (Optional) */
    Cloud bool `json:"cloud"`

    /* node数量 (Optional) */
    NodeCount int `json:"nodeCount"`

    /* 容器数量 (Optional) */
    ContainerCount int `json:"containerCount"`

    /* 最新容器变化时间 (Optional) */
    LatestPodChangeTime int64 `json:"latestPodChangeTime"`
}
