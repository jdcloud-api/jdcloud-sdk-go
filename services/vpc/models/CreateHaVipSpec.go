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


type CreateHaVipSpec struct {

    /* 高可用虚拟ip名称。取值范围：1-32个中文、英文大小写的字母、数字和下划线分隔符中划线  */
    HaVipName string `json:"haVipName"`

    /* 高可用虚拟ip所属子网 id  */
    SubnetId string `json:"subnetId"`

    /* 该高可用虚拟ip描述 (Optional) */
    Description string `json:"description"`

    /* 高可用虚拟ip,选取统一子网下的内网ip (Optional) */
    IpAddress string `json:"ipAddress"`

    /* 预检标识，默认false，dryRun为true时只作检查，不做变更 (Optional) */
    DryRun bool `json:"dryRun"`
}
