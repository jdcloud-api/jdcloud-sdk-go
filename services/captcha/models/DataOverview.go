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


type DataOverview struct {

    /* 应用数 (Optional) */
    AppCount int `json:"appCount"`

    /* 场景数 (Optional) */
    SceneCount int `json:"sceneCount"`

    /* 请求量 (Optional) */
    RequestCount int `json:"requestCount"`

    /* 之前请求量，用于计算环比 (Optional) */
    RequestCountPre int `json:"requestCountPre"`

    /* 通过量 (Optional) */
    PassCount int `json:"passCount"`

    /* 之前通过量，用于计算环比 (Optional) */
    PassCountPre int `json:"passCountPre"`

    /* 防御量 (Optional) */
    BlockCount int `json:"blockCount"`

    /* 之前的防御量，用于计算环比 (Optional) */
    BlockCountPre int `json:"blockCountPre"`
}
