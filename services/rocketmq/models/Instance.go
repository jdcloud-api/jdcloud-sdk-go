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
import resourcetag "github.com/jdcloud-api/jdcloud-sdk-go/services/resourcetag/models"

type Instance struct {

    /* 实例ID (Optional) */
    InstanceId string `json:"instanceId"`

    /* 实例名称 (Optional) */
    InstanceName string `json:"instanceName"`

    /* rocketmq实例版本 (Optional) */
    InstanceVersion string `json:"instanceVersion"`

    /* 实例状态，running：运行，error：错误，creating：创建中，changing：变配，stop：停止 (Optional) */
    InstanceStatus string `json:"instanceStatus"`

    /* 创建时间，遵循ISO8601标准，使用UTC时间，格式为：YYYY-MM-DDTHH:mm:ssZ (Optional) */
    CreateTime string `json:"createTime"`

    /* 所属VPC的ID (Optional) */
    VpcId string `json:"vpcId"`

    /* 所属子网的ID (Optional) */
    SubnetId string `json:"subnetId"`

    /* azId (Optional) */
    AzId []string `json:"azId"`

    /* 集群规格信息 (Optional) */
    InstanceClass []InstanceClass `json:"instanceClass"`

    /* 计费信息 (Optional) */
    Charge charge.Charge `json:"charge"`

    /* Tag信息 (Optional) */
    Tags []resourcetag.Tag `json:"tags"`

    /* 扩展参数 (Optional) */
    Extension RespExtension `json:"extension"`
}
