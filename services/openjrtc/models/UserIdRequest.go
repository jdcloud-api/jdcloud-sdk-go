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


type UserIdRequest struct {

    /* UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2023-09-05T07:00:00Z (Optional) */
    StartTime string `json:"startTime"`

    /* UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2023-09-07T08:00:00Z (Optional) */
    EndTime string `json:"endTime"`

    /* 业务接入方定义的且在JRTC系统内注册过的房间号 (Optional) */
    UserRoomId string `json:"userRoomId"`

    /* 业务接入方定义的且在JRTC系统内注册过userId列表,最多支持20个userId (Optional) */
    UserIds []string `json:"userIds"`
}
