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


type BitRate struct {

    /* 业务接入方用户体系定义的且在JRTC系统内注册过的userId (Optional) */
    UserId string `json:"userId"`

    /* 时间戳-毫秒 (Optional) */
    Date int64 `json:"date"`

    /* 比特率 bps (Optional) */
    Bitrate float64 `json:"bitrate"`
}
