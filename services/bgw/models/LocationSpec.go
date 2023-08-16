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


type LocationSpec struct {

    /* 专线创建的地域编码;只在创建自助连接时生效，通过调用[describeLocations](../Location/describeLocations.md)接口获取 (Optional) */
    LocationCode string `json:"locationCode"`

    /* 专线接入端口规格代码，在创建自助连接和托管专线时生效.通过调用[describeLocations](../Location/describeLocations.md)接口获取 (Optional) */
    LocationPortSpecCode string `json:"locationPortSpecCode"`

    /* 专线接入运营商代码，只在创建自助连接时生效.通过调用[describeLocations](../Location/describeLocations.md)接口获取 (Optional) */
    LocationISPCode string `json:"locationISPCode"`
}
