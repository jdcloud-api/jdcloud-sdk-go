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


type GpjmListJobInstanceResultDTO struct {

    /* 作业名称 (Optional) */
    JobName string `json:"jobName"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 运行开始时间 (Optional) */
    StartTime string `json:"startTime"`

    /* 运行结束时间 (Optional) */
    EndTime string `json:"endTime"`

    /* 数据日期 (Optional) */
    TxDate string `json:"txDate"`

    /* 返回信息 (Optional) */
    CurrentStatusMsg string `json:"currentStatusMsg"`

    /* 返回编码 (Optional) */
    LastReturnCode string `json:"lastReturnCode"`

    /* 运行时间 (Optional) */
    RunTime string `json:"runTime"`

    /* Session ID (Optional) */
    Sessionid string `json:"sessionid"`

    /* 状态 (Optional) */
    Status string `json:"status"`

    /* 是否补数任务：Y 是 N 否 (Optional) */
    IsComplement string `json:"isComplement"`

    /* 计划执行时间 (Optional) */
    PlanExecTime string `json:"planExecTime"`

    /* 运行周期 ,F 分钟、H 小时、D 天、W 周、M 月、O 一次性运行、N 无周期 (Optional) */
    JobId string `json:"jobId"`

    /* 作业描述 (Optional) */
    Description string `json:"description"`

    /* 处理类型（calc：数据计算，extract：数据抽取，ods：ODS加工，load：数据推送，hdfs：数据同步,dqim:质量作业） (Optional) */
    ProcessType string `json:"processType"`

    /* 执行类型（fusing：熔断,warning:预警） (Optional) */
    ExeType string `json:"exeType"`

    /* 是否启用，0未上线、1已上线、2已下线 (Optional) */
    Enable string `json:"enable"`

    /* 负责人 (Optional) */
    Manager string `json:"manager"`

    /* 数据渠道来源：新模型（MODEL），老模型（OLD_MODEL）、集成开发（IDE）、数据管道（PIPE）、数据质量（DQ）、AI(KUAI) (Optional) */
    Channel string `json:"channel"`

    /* 任务类型 (Optional) */
    JobType string `json:"jobType"`

    /* 作业子类型 (Optional) */
    JobChildType string `json:"jobChildType"`

    /*  (Optional) */
    ManagerFlag bool `json:"managerFlag"`

    /* 客户作业名 (Optional) */
    CstJobName string `json:"cstJobName"`
}
