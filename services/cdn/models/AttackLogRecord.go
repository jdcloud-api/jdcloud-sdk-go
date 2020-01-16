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


type AttackLogRecord struct {

    /* 攻击峰值 (Optional) */
    AttackTraffic float32 `json:"attackTraffic"`

    /* 攻击类型 (Optional) */
    AttackType []string `json:"attackType"`

    /* 黑洞封禁 0->未封禁 1->封禁 (Optional) */
    BlackHole int `json:"blackHole"`

    /* 攻击结束时间 utc时间 (Optional) */
    EndTime string `json:"endTime"`

    /* 用户pin (Optional) */
    Pin string `json:"pin"`

    /* 攻击开始时间 utc时间 (Optional) */
    StartTime string `json:"startTime"`

    /* 攻击峰值单位 流量单位 (Optional) */
    Unit string `json:"unit"`
}