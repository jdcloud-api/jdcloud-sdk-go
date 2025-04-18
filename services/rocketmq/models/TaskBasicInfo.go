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


type TaskBasicInfo struct {

    /* 任务开始执行时间 (Optional) */
    TaskStartTime string `json:"taskStartTime"`

    /* 任务结束执行时间 (Optional) */
    TaskEndTime string `json:"taskEndTime"`

    /* 任务执行人(主账户) (Optional) */
    User string `json:"user"`

    /* 任务执行人(子用户) (Optional) */
    SubUser string `json:"subUser"`

    /* 任务的执行状态
init - 初始化。代表任务刚初始化, 待调度运行。中间态, 允许中断
executing - 进行中。 代表任务被调度执行。 中间态, 允许中断
succ - 已完成。 终止态, 任务运行完成
fail - 变更中断。 终止态
error - 失败。 中间态, 允许重试
 (Optional) */
    TaskStatus string `json:"taskStatus"`
}
