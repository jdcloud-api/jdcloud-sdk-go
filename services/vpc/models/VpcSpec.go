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


type VpcSpec struct {

    /* 私有网络名称,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符。  */
    VpcName string `json:"vpcName"`

    /* 如果为空，则不限制网段，如果不为空，10.0.0.0/8、172.16.0.0/12和192.168.0.0/16及它们包含的子网，且子网掩码长度为16-28之间 (Optional) */
    AddressPrefix string `json:"addressPrefix"`

    /* vpc描述，允许输入UTF-8编码下的全部字符，不超过256字符。 (Optional) */
    Description string `json:"description"`

    /* 取值包括true、false，默认为false，不开启组播。 (Optional) */
    EnableMulticast bool `json:"enableMulticast"`

    /* VPC az类型，取值：standard(标准VPC)，edge(边缘VPC) (Optional) */
    AzType string `json:"azType"`

    /* VPC可用区，边缘VPC必须指定可用区 (Optional) */
    Az string `json:"az"`
}
