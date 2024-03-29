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


type PolicyDetailSpec struct {

    /* 策略类型，可选值如下：
1. "AutoImage" - 云主机自动创建镜像策略
  */
    PolicyType string `json:"policyType"`

    /* 触发时间，针对开启的策略生效，格式"yyyy-MM-dd HH24:mm:ss" (Optional) */
    FireTime *string `json:"fireTime"`

    /* 策略触发条件/周期，可选值如下：
1. "crond xxx" - 暂不支持，crond表达式，如每月1号0点[0 0 1 * *]
2. "interval xxx" - 间隔时间，支持"秒/分/时/天/周"，xxx格式如：[30s、20m、10h、1d、2w]
3. "condition xxx" - 暂不支持，满足某些条件
  */
    FireCondition string `json:"fireCondition"`

    /* 策略触发执行累计次数达到execNumLimit后，自动失效变为disable状态。该参数为-1时不生效，大于0时生效。 (Optional) */
    ExecNumLimit *int `json:"execNumLimit"`

    /* 该策略执行时的额外条件配置。 (Optional) */
    ExecConfig []Config `json:"execConfig"`

    /* 日志配置 (Optional) */
    Log *Log `json:"log"`
}
