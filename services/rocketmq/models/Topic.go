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


type Topic struct {

    /* topic名称 (Optional) */
    Topic string `json:"topic"`

    /* 消息类型 (Optional) */
    TopicType string `json:"topicType"`

    /* 描述 (Optional) */
    Description string `json:"description"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 分区数 (Optional) */
    QueueNums int `json:"queueNums"`
}
