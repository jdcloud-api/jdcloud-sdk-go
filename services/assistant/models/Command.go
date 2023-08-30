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


type Command struct {

    /* 命令名称，长度为1\~128个字符，只允许中文、数字、大小写字母、英文下划线（\_）、连字符（-）及点（.）。
 (Optional) */
    CommandName string `json:"commandName"`

    /* 命令类型，可选值：shell和powershell，默认shell
 (Optional) */
    CommandType string `json:"commandType"`

    /* 以base64编码的命令内容，编码后长度小于16KB
 (Optional) */
    CommandContent string `json:"commandContent"`

    /* 超时时间，取值范围：[10, 86400], 超过该时间后，尚未执行完的命令会置为失败。默认60s
 (Optional) */
    Timeout int `json:"timeout"`

    /* 用户名，执行该命令时的用户身份。在linux上默认是root，windows上默认是administrator。长度小于256。
 (Optional) */
    Username string `json:"username"`

    /* 命令执行路径。在linux上默认是/root，windows上默认是C:\Windows\System32。长度小于256。
 (Optional) */
    Workdir string `json:"workdir"`

    /* 命令描述，描述该命令详细信息，如功能、使用注意事项等。长度小于256。
 (Optional) */
    CommandDescription string `json:"commandDescription"`

    /* 是否使用参数, 默认false，不使用参数 (Optional) */
    EnableParameter bool `json:"enableParameter"`

    /* 脚本中用户自定义的参数，最多20个。 (Optional) */
    Parameters []Parameter `json:"parameters"`

    /* 命令创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 命令更新时间 (Optional) */
    UpdateTime string `json:"updateTime"`
}
