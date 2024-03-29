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


type VpcPolicy struct {

    /* 策略的Id,全局唯一 (Optional) */
    VpcPolicyId string `json:"vpcPolicyId"`

    /* 访问策略名称 (Optional) */
    VpcPolicyName string `json:"vpcPolicyName"`

    /* 访问策略所属vpc id (Optional) */
    VpcId string `json:"vpcId"`

    /* 具体策略内容(格式参考样例) (Optional) */
    PolicyDocument string `json:"policyDocument"`
}
