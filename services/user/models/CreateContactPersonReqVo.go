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


type CreateContactPersonReqVo struct {

    /* 联系人姓名  */
    UserName string `json:"userName"`

    /* 邮箱 (Optional) */
    Email string `json:"email"`

    /* 手机号  */
    Mobile string `json:"mobile"`

    /* 接收消息类型：1账户消息 2产品消息 3故障消息 4活动消息（可多个，英文逗号隔开） (Optional) */
    SubscriptionType string `json:"subscriptionType"`
}
