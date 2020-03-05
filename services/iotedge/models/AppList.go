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


type AppList struct {

    /* App业务编号 (Optional) */
    AppId string `json:"appId"`

    /* 云翼编译编号 (Optional) */
    ArkId string `json:"arkId"`

    /* APP版本号 (Optional) */
    AppVersion string `json:"appVersion"`

    /* 发布时间 (Optional) */
    ReleaseTime string `json:"releaseTime"`

    /* APP状态，0-发布成功，1-发布失败，2-审核通过，3-审核不通过，4-上线，5-下线，99-发布中 (Optional) */
    AppStatus string `json:"appStatus"`

    /* 上线时间 (Optional) */
    OnlineTime string `json:"onlineTime"`

    /* 物模型编号 (Optional) */
    TmId string `json:"tmId"`

    /* 物模型名称 (Optional) */
    TmName string `json:"tmName"`
}
