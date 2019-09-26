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


type SubnetSpec struct {

    /* 可用区, 如 cn-east-tz1a  */
    Az string `json:"az"`

    /* 私有网络ID  */
    VpcId string `json:"vpcId"`

    /* 子网的网络范围  */
    Cidr string `json:"cidr"`

    /* 名称  */
    Name string `json:"name"`

    /* 描述  */
    Description string `json:"description"`
}