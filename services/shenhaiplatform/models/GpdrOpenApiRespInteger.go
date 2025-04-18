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


type GpdrOpenApiRespInteger struct {

    /* 成功标识，1成功，0失败 (Optional) */
    Success int `json:"success"`

    /* 返回结果对象 (Optional) */
    Result int `json:"result"`

    /* 返回状态码 (Optional) */
    Code string `json:"code"`

    /* 返回状态信息 (Optional) */
    Msg string `json:"msg"`

    /* 返回请求流水号 (Optional) */
    _REQ_SEQ_NO_ string `json:"_REQ_SEQ_NO_"`
}
