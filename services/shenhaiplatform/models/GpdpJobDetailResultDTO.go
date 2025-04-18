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


type GpdpJobDetailResultDTO struct {

    /* 数据日期 (Optional) */
    TxDate string `json:"txDate"`

    /* 计划执行时间 (Optional) */
    PlanExecTime string `json:"planExecTime"`

    /* 负责人 (Optional) */
    Manager string `json:"manager"`

    /* 作业名称 (Optional) */
    JobName string `json:"jobName"`

    /* 最后一次运行结束时间 (Optional) */
    LastEndTime string `json:"lastEndTime"`

    /* 最后一次运行开始时间 (Optional) */
    LastStartTime string `json:"lastStartTime"`

    /* 作业运行耗时/秒 (Optional) */
    CostTime int `json:"costTime"`

    /* 表名 (Optional) */
    TargetTable int `json:"targetTable"`
}
