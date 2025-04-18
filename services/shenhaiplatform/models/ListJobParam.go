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


type ListJobParam struct {

    /* 分页-页码 (Optional) */
    PageNum int `json:"pageNum"`

    /* 分页-每页数量 (Optional) */
    PageSize int `json:"pageSize"`

    /* 任务名称 (Optional) */
    CstJobName string `json:"cstJobName"`

    /* 任务编码 (Optional) */
    JobName string `json:"jobName"`

    /* 租户编码 (Optional) */
    CompanyCode string `json:"companyCode"`

    /* 负责人pin (Optional) */
    ManagerPin string `json:"managerPin"`

    /* 工作空间编码 (Optional) */
    WorkspaceCode string `json:"workspaceCode"`

    /* 任务创建-开始日期: yyyy-MM-dd (Optional) */
    CreateTimeAfter string `json:"createTimeAfter"`

    /* 任务创建-结束日期: yyyy-MM-dd (Optional) */
    CreateTimeBefore string `json:"createTimeBefore"`

    /* 任务更新-开始日期: yyyy-MM-dd (Optional) */
    UpdateTimeAfter string `json:"updateTimeAfter"`

    /* 任务更新-结束日期: yyyy-MM-dd (Optional) */
    UpdateTimeBefore string `json:"updateTimeBefore"`

    /* 任务运行-开始时间: yyyy-MM-dd HH:mm (Optional) */
    JobStartRunTime string `json:"jobStartRunTime"`

    /* 任务运行-结束时间: yyyy-MM-dd HH:mm (Optional) */
    JobEndRunTime string `json:"jobEndRunTime"`

    /* 任务上下线状态: 1-上线 2-下线 (Optional) */
    Enable int `json:"enable"`

    /* 任务周期 (Optional) */
    Cycle string `json:"cycle"`

    /* 任务父类型 (Optional) */
    JobType string `json:"jobType"`

    /* 任务子类型 (Optional) */
    JobChildType string `json:"jobChildType"`

    /* 任务最后状态: not,Done,Pending,Running,Failed,Timeout,Clean,Ready (Optional) */
    LastJobStatus []string `json:"lastJobStatus"`

    /* 任务最后数据日期 (Optional) */
    LastTxdate string `json:"lastTxdate"`

    /* 任务负责人 (Optional) */
    Manager string `json:"manager"`

    /* 过滤模式: exact-精确过滤，否则为模糊过滤 (Optional) */
    QueryMode string `json:"queryMode"`
}
