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


type DescribeGlobalAlarmRes struct {

    /* 告警规则ID (Optional) */
    Id int64 `json:"id"`

    /* 告警类型 WEB_ATTACK->网站攻击 CC_ATTACK->CC攻击 DDOS_ATTACK->DDOS攻击 STATUS_CODE_ERROR->状态码错误 (Optional) */
    AlarmType string `json:"alarmType"`

    /* 统计周期 (Optional) */
    StatCycle int `json:"statCycle"`

    /* 统计周期单位 (Optional) */
    StatCycleUnit string `json:"statCycleUnit"`

    /* 统计阈值 (Optional) */
    StatThreshold int64 `json:"statThreshold"`

    /* 告警次数限制 (Optional) */
    AlarmTimesLimit int `json:"alarmTimesLimit"`

    /* 发送短信开关 (Optional) */
    SmsEnable bool `json:"smsEnable"`

    /* 发送邮件开关 (Optional) */
    EmailEnable bool `json:"emailEnable"`

    /* 发送站内信开关 (Optional) */
    WebMsEnable bool `json:"webMsEnable"`

    /* 规则开关 (Optional) */
    Enable bool `json:"enable"`

    /* 告警联系人 (Optional) */
    ContactUsers []int64 `json:"contactUsers"`

    /* 告警联系组 (Optional) */
    ContactGroups []int64 `json:"contactGroups"`
}