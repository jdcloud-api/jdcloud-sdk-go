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


type DebugReturnMessage struct {

    /* 请求地址 (Optional) */
    RequestUrl string `json:"requestUrl"`

    /* request中header信息 (Optional) */
    RequestHeader string `json:"requestHeader"`

    /* request中body信息 (Optional) */
    RequestBody string `json:"requestBody"`

    /* 响应状态码 (Optional) */
    ResponseCodeStatus string `json:"responseCodeStatus"`

    /* header返回值 (Optional) */
    ResponseHeaderValue string `json:"responseHeaderValue"`

    /* body返回值 (Optional) */
    ResponseBody string `json:"responseBody"`
}