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


type DescribedAlarmHistory struct {

    /* 报警规则ID (Optional) */
    AlarmId string `json:"alarmId"`

    /* 资源维度 (Optional) */
    Dimension string `json:"dimension"`

    /* 资源维度名称 (Optional) */
    DimensionName string `json:"dimensionName"`

    /* 告警持续次数 (Optional) */
    DurationTimes int64 `json:"durationTimes"`

    /* 是否是一键告警 (1-一键告警  0-非一键告警) (Optional) */
    IsOneClickAlarm int64 `json:"isOneClickAlarm"`

    /* 告警持续时间，单位分钟 (Optional) */
    NoticeDurationTime int64 `json:"noticeDurationTime"`

    /* 用于前端显示的‘触发告警级别’。从低到高分别为‘普通’, ‘紧急’, ‘严重’ (Optional) */
    NoticeLevel string `json:"noticeLevel"`

    /* 触发的告警级别。从低到高分别为‘common’, ‘critical’, ‘fatal’ (Optional) */
    NoticeLevelTriggered string `json:"noticeLevelTriggered"`

    /* 告警时间 (Optional) */
    NoticeTime string `json:"noticeTime"`

    /* 告警时间对应的时间戳 (Optional) */
    NoticeTimeUnix int64 `json:"noticeTimeUnix"`

    /* 资源类型 (Optional) */
    Product string `json:"product"`

    /* 资源类型名称 (Optional) */
    ProductName string `json:"productName"`

    /* 告警通知人 (Optional) */
    Receivers []NoticeReceiver `json:"receivers"`

    /* 资源Id对应的region (Optional) */
    Region string `json:"region"`

    /* 资源Id (Optional) */
    ResourceId string `json:"resourceId"`

    /*  (Optional) */
    Rule BasicRuleDetail `json:"rule"`

    /* 规则类型 (Optional) */
    RuleType string `json:"ruleType"`

    /* 告警类型  1-告警恢复  2-告警 4-数据不足 (Optional) */
    Status int64 `json:"status"`

    /* 资源tags (Optional) */
    Tags interface{} `json:"tags"`

    /* 告警值 (Optional) */
    Value float64 `json:"value"`
}
