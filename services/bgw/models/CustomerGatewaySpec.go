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


type CustomerGatewaySpec struct {

    /* 客户网关的名称，只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符。  */
    CustomerGatewayName string `json:"customerGatewayName"`

    /* 客户网关上公网可路由的IPv4地址列表，可以为单IP、双IP、四IP，IP不能相同  */
    CustomerIps []string `json:"customerIps"`

    /* 客户自治系统的ASN号，取值范围[65001,65499]，设置后不支持修改；  */
    BgpAsn int `json:"bgpAsn"`

    /* 客户网关的描述，允许输入UTF-8编码下的全部字符，不超过256字符。 (Optional) */
    Description string `json:"description"`
}
