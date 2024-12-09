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


type LogInfo struct {

    /* 应用名 (Optional) */
    App string `json:"app"`

    /* 纳秒时间戳 (Optional) */
    Timestamp string `json:"timestamp"`

    /* ip (Optional) */
    Host string `json:"host"`

    /* 日志级别 (Optional) */
    Level string `json:"level"`

    /* sgm traceId (Optional) */
    TraceId string `json:"traceId"`

    /* filepath (Optional) */
    FilePath string `json:"filePath"`

    /* 前端默认展开 (Optional) */
    _expanded string `json:"_expanded"`

    /* t_h_r (Optional) */
    T_h_r string `json:"t_h_r"`

    /* c_l_s (Optional) */
    C_l_s string `json:"c_l_s"`

    /* m_s_g (Optional) */
    M_s_g string `json:"m_s_g"`
}