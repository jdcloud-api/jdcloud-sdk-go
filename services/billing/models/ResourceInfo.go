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


type ResourceInfo struct {

    /* 资源ID  */
    ResourceId string `json:"resourceId"`

    /* 用户PIN  */
    Pin string `json:"pin"`

    /* 产品编码  */
    ServiceCode string `json:"serviceCode"`

    /* 退款单号  */
    RefundNo string `json:"refundNo"`

    /* 订单编号列表,退续费订单时必传，一次退多个续费订单请按续费时间倒序传入 (Optional) */
    OrderNumberList []string `json:"orderNumberList"`
}