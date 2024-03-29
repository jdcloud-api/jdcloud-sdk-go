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

import common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"

type Invocation struct {

    /* 命令执行状态,取值为：[`waiting`,`pending`,`running`,`stopping`,`failed`,`partial_failed`,`stopped`,`finish`] (Optional) */
    Status string `json:"status"`

    /* 命令Id (Optional) */
    CommandId string `json:"commandId"`

    /* 命令名称 (Optional) */
    CommandName string `json:"commandName"`

    /* 命令来源，[`jdcloud`:公共命令, `self`:私有命令] (Optional) */
    SourceType string `json:"sourceType"`

    /* 命令内容调用Id (Optional) */
    InvokeId string `json:"invokeId"`

    /* 请求执行命令的云主机Id (Optional) */
    Instances []string `json:"instances"`

    /* 请求执行命令的云主机标签 (Optional) */
    Tags []common.TagFilter `json:"tags"`

    /* 执行命令的云主机 (Optional) */
    InvokeInstances []string `json:"invokeInstances"`

    /* 命令类型 (Optional) */
    CommandType string `json:"commandType"`

    /* 命令内容，base64编码 (Optional) */
    CommandContent string `json:"commandContent"`

    /* 命令描述 (Optional) */
    CommandDescription string `json:"commandDescription"`

    /* 每个云主机的执行详情 (Optional) */
    InvocationInstances []InvocationInstance `json:"invocationInstances"`

    /* 命令参数 (Optional) */
    Parameters string `json:"parameters"`

    /* 命令超时时间 (Optional) */
    Timeout string `json:"timeout"`

    /* 用户名 (Optional) */
    Username string `json:"username"`

    /* 命令执行路径 (Optional) */
    Workdir string `json:"workdir"`

    /* 命令调用的错误信息 (Optional) */
    ErrorInfo string `json:"errorInfo"`

    /* 该次命令调用的开始时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 该次命令配置的定时执行时间 (Optional) */
    ExecTime string `json:"execTime"`
}
