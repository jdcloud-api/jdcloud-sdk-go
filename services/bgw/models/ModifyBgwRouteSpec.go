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


type ModifyBgwRouteSpec struct {

    /* Bgw上路由的Id  */
    BgwRouteId string `json:"bgwRouteId"`

    /* Bgw上路由的目的地址,格式如：172.10.2.15/16，参见RFC 4632  */
    BgwRouteDestination string `json:"bgwRouteDestination"`

    /* Bgw上路由的下一跳类型：目前支持托管通道(hosted_private_vif),专线通道(private_vif),vpc接口(vpc_attachment),vpn连接(vpn_connection)和vpc(vpc)  */
    BgwNexthopType string `json:"bgwNexthopType"`

    /* Bgw上路由的下一跳设备资源Id，目前支持vifId,vpcAttachmentId,vpnConnectionId  */
    BgwRouteNexthop string `json:"bgwRouteNexthop"`

    /* Bgw路由的描述，仅支持静态类型的route更改description (Optional) */
    Description string `json:"description"`
}
