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


type ServerlessCluster struct {

    /* 集群id (Optional) */
    ClusterId string `json:"clusterId"`

    /* 名称 (Optional) */
    Name string `json:"name"`

    /* 描述 (Optional) */
    Description string `json:"description"`

    /* kubernetes的版本 (Optional) */
    Version string `json:"version"`

    /* 集群所在的az (Optional) */
    Azs []string `json:"azs"`

    /* k8s的cluster的cidr (Optional) */
    ClusterCidr string `json:"clusterCidr"`

    /* 认证信息 (Optional) */
    MasterAuth MasterAuth `json:"masterAuth"`

    /* 状态  [pending,running,reconciling（升级时的状态）, deleting, deleted, error] (Optional) */
    ClusterState string `json:"clusterState"`

    /* 状态变更原因 (Optional) */
    StateMessage string `json:"stateMessage"`

    /* 用户的AccessKey，插件调用open-api时的认证凭证 (Optional) */
    AccessKey string `json:"accessKey"`

    /* 基本验证方式 (Optional) */
    BasicAuth bool `json:"basicAuth"`

    /* 证书验证方式 (Optional) */
    ClientCertificate bool `json:"clientCertificate"`

    /* 用户访问的内网ip (Optional) */
    PrivateEndpoint string `json:"privateEndpoint"`

    /* 用户访问的ip (Optional) */
    Endpoint string `json:"endpoint"`

    /* IPv6地址 (Optional) */
    EndpointIPV6 string `json:"endpointIPV6"`

    /* endpoint的port (Optional) */
    EndpointPort string `json:"endpointPort"`

    /* endpoint的dashboard port (Optional) */
    DashboardPort string `json:"dashboardPort"`

    /* deprecated 优先以addonsConfig中的配置为准 <br>用户是否启用集群自定义监控，true 表示开启用，false 表示未开启用 (Optional) */
    UserMetrics bool `json:"userMetrics"`

    /* 集群组件配置信息 (Optional) */
    AddonsConfig []AddonConfig `json:"addonsConfig"`

    /* 控制节点操作进度 (Optional) */
    MasterProgress MaintenanceWindow `json:"masterProgress"`

    /* 网络配置信息 (Optional) */
    ClusterNetwork ServerlessClusterNetworkConfig `json:"clusterNetwork"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 集群网络类型,可取值为auto和customized (Optional) */
    NetworkMode string `json:"networkMode"`

    /* 用户自定义的集群的环境信息，会影响到创建集群时的组件模版的渲染 (Optional) */
    ClusterEnvironments []StringKeyValuePair `json:"clusterEnvironments"`

    /* 是否是边缘计算集群 (Optional) */
    IsEdge bool `json:"isEdge"`
}
