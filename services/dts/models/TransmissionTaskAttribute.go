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

import charge "github.com/jdcloud-api/jdcloud-sdk-go/services/charge/models"

type TransmissionTaskAttribute struct {

    /* 数据传输任务ID (Optional) */
    TaskId string `json:"taskId"`

    /* 数据传输任务名称 (Optional) */
    TaskName string `json:"taskName"`

    /* vpc (Optional) */
    Vpc string `json:"vpc"`

    /*  (Optional) */
    Subnet string `json:"subnet"`

    /* 实例类型，取值为：Sync：同步、Subscribe：订阅、Migration：迁移、DisasterRecovery 灾备 (Optional) */
    TransmissionMethod string `json:"transmissionMethod"`

    /* 数据传输拓扑，支持：oneway - 单向数据传输、bidirectional - 双向数据传输，默认取值为：oneway (Optional) */
    Topology string `json:"topology"`

    /* 数据传输实例规格，支持：dts.m1.medium、dts.m1.large、dts.m1.xlarge (Optional) */
    InstanceClass string `json:"instanceClass"`

    /* 数据传输任务所处阶段及描述 (Optional) */
    Phase Phase `json:"phase"`

    /* 计费配置 (Optional) */
    Charge charge.Charge `json:"charge"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* Creating - 创建中 CreateFailed - 创建失败 NotStarted - 未启动 Modifying - 修改同步对象中 ModifyFailed - 修改同步对象失败 PreChecking - 预检查中 PreCheckFailed - 预检查失败 PreCheckSucceed - 预检查成功 Initializing - 任务初始化中 InitializeFailed - 任务初始化中失败 Running - 运行中 Failed - 同步失败 Suspending - 暂停 Retrying - 重试中 Expired - 锁定中 Finished - 结束 Deleting - 删除中 Deleted - 删除 (Optional) */
    Status string `json:"status"`

    /* 源实例信息 (Optional) */
    Source TransmissionEndpoint `json:"source"`

    /* 目标实例信息 (Optional) */
    Destination TransmissionEndpoint `json:"destination"`

    /* 标签信息 (Optional) */
    Tags []Tag `json:"tags"`

    /* 数据传输初始化类型 (Optional) */
    TransmissionMode TransmissionMode `json:"transmissionMode"`

    /* 目标端执行策略：MySQL：mysql-replace-engine（覆盖写入）、conflict-detect-engine（冲突报错）；MongoDB：conflict-replace、conflict-detect (Optional) */
    DestExecutionStrategy string `json:"destExecutionStrategy"`

    /* 自定义设置 (Optional) */
    Options PipelineOptions `json:"options"`

    /* 任务自定义设置 (Optional) */
    CustomOptions CustomOptions `json:"customOptions"`

    /* 数据传输任务检查点 (Optional) */
    Checkpoint TransmissionCheckpoint `json:"checkpoint"`

    /* 数据传输任务检查点 (Optional) */
    StartCheckPoint TransmissionCheckpoint `json:"startCheckPoint"`
}
