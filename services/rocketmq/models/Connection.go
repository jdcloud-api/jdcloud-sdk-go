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


type Connection struct {

    /* 客户端Id（客户端上报的） (Optional) */
    ClientId string `json:"clientId"`

    /* 客户端IP（客户端上报的） (Optional) */
    ClientIp string `json:"clientIp"`

    /* 公网IP (Optional) */
    ClientAddr string `json:"clientAddr"`

    /* 语言 (Optional) */
    Language string `json:"language"`

    /* 版本 (Optional) */
    VersionDesc string `json:"versionDesc"`
}
