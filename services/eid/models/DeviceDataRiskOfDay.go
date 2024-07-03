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


type DeviceDataRiskOfDay struct {

    /* 设备总量 (Optional) */
    AllDevice int `json:"allDevice"`

    /* 正常 (Optional) */
    Normal int `json:"normal"`

    /* 脱机挂 (Optional) */
    TuoJiGua int `json:"tuoJiGua"`

    /* 伪造 (Optional) */
    Forgery int `json:"forgery"`

    /* 云手机 (Optional) */
    CloudPhone int `json:"cloudPhone"`

    /* root/越狱 (Optional) */
    Root int `json:"root"`

    /* hook (Optional) */
    Hook int `json:"hook"`

    /* 多开 (Optional) */
    DuoKai int `json:"duoKai"`

    /* 模拟器 (Optional) */
    MoNi int `json:"moNi"`

    /* 时间，yyyy-mm-dd hh:mm:ss格式 (Optional) */
    Time string `json:"time"`
}
