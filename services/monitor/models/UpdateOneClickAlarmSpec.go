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


type UpdateOneClickAlarmSpec struct {

    /* 一键告警规则id  */
    AlarmId string `json:"alarmId"`

    /* 一键告警规则配置
in: body (Optional) */
    AlarmOptions []UpdateOneClickAlarmOption `json:"alarmOptions"`

    /* 中文描述 (Optional) */
    DescriptionCN string `json:"descriptionCN"`

    /* 英文描述 (Optional) */
    DescriptionEN string `json:"descriptionEN"`
}
