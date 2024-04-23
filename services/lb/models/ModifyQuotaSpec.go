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


type ModifyQuotaSpec struct {

    /* lb类型，取值范围：alb、nlb、dnlb，默认为alb (Optional) */
    LbType string `json:"lbType"`

    /* 资源类型，取值范围：loadbalancer、listener、target_group、target、backend、urlMap(仅alb支持)、rules(仅alb支持)、ruleExtendAction(仅alb支持)、extensionCertificate(仅alb支持)、customizedConfiguration  */
    Type string `json:"type"`

    /* type为loadbalancer、customizedConfiguration不设置, type为listener、backend、target_group、urlMap设置为loadbalancerId, type为target设置为targetGroupId, type为rules设置为urlMapId,type为extensionCertificate设置为listenerId (Optional) */
    ParentResourceId string `json:"parentResourceId"`

    /* 配额大小  */
    MaxLimit int `json:"maxLimit"`
}
