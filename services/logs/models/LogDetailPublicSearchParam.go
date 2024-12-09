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


type LogDetailPublicSearchParam struct {

    /* UUID，同一批次保持一致，必填 (Optional) */
    SearchId string `json:"searchId"`

    /* 日志类型 (Optional) */
    LogType string `json:"logType"`

    /* ip (Optional) */
    Ip []string `json:"ip"`

    /* pod name (Optional) */
    PodName []string `json:"podName"`

    /* 命名空间 (Optional) */
    Namespace string `json:"namespace"`

    /* 集群 (Optional) */
    Cluster string `json:"cluster"`

    /* 容器名 (Optional) */
    ContainerName string `json:"containerName"`

    /* 文件路径 (Optional) */
    FilePath []string `json:"filePath"`

    /* 开始时间纳秒数字，默认一小时前 (Optional) */
    StartTime int64 `json:"startTime"`

    /* 结束时间纳秒数字，默认现在 (Optional) */
    EndTime int64 `json:"endTime"`

    /* 正序：FORWARD、倒序：BACKWARD，默认BACKWARD (Optional) */
    Direction string `json:"direction"`

    /* 查询数量数字，默认100 (Optional) */
    Limit int `json:"limit"`

    /* 步长时间（单位：秒），10、0.5 (Optional) */
    Step float32 `json:"step"`

    /* label过滤 (Optional) */
    StreamFilter []StreamFilter `json:"streamFilter"`

    /* 行过滤 (Optional) */
    LineFilter LineFilter `json:"lineFilter"`

    /* 格式化过滤条件 (Optional) */
    FmtFilter FormatFilter `json:"fmtFilter"`
}
