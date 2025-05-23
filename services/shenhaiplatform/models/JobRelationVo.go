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


type JobRelationVo struct {

    /* 作业名 (Optional) */
    JobName string `json:"jobName"`

    /* 客户作业名称 (Optional) */
    CstJobName string `json:"cstJobName"`

    /* 最后以此执行日期 (Optional) */
    LastTxDate string `json:"lastTxDate"`

    /* 最后一次运行开始时间 (Optional) */
    LastStartTime string `json:"lastStartTime"`

    /* 最后一次结束时间 (Optional) */
    LastEndTime string `json:"lastEndTime"`

    /* 最后一次运行状态，Pending、Ready、Running、Done、Failed、Clean (Optional) */
    LastStatus string `json:"lastStatus"`

    /* 是否可以自身并行，1启用、0关闭 (Optional) */
    FlagParallel string `json:"flagParallel"`

    /* 当前作业状态信息 (Optional) */
    CurrentStatusMsg string `json:"currentStatusMsg"`

    /* 周期开始时间(适用小时分钟) (Optional) */
    SequenceStartTime string `json:"sequenceStartTime"`

    /* 周期结束时间(适用小时分钟) (Optional) */
    SequenceEndTime string `json:"sequenceEndTime"`

    /* 周期间隔(适用小时分钟，当周期为小时，含义为间隔小时数，当周期为分钟，含义为间隔分钟数) (Optional) */
    SequenceInterval string `json:"sequenceInterval"`

    /* 是否启用，1启用、0关闭 (Optional) */
    Enable string `json:"enable"`

    /* 负责人，不超过10个 (Optional) */
    Manager string `json:"manager"`

    /* 作业所属系统，G gravity版本、A automation版本,用于迁移字段 (Optional) */
    JobBelong string `json:"jobBelong"`

    /* 是否强依赖，1启用、0关闭 (Optional) */
    Necessary string `json:"necessary"`

    /* 是否可以上下游并行，1启用、0关闭 (Optional) */
    DepFlagParallel string `json:"depFlagParallel"`

    /* 触发类型:dependency 依赖、time 时间、file 文件、manual 手工、once 一次性 (Optional) */
    TriggerType string `json:"triggerType"`

    /* 运行周期 ,F 分钟、H 小时、D 天、W 周、M 月、O 一次性运行、N 无周期 (Optional) */
    Cycle string `json:"cycle"`

    /* 周期具体日期 (Optional) */
    Sequence string `json:"sequence"`
}
