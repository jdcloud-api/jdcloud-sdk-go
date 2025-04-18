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


type UranusTaskNodeDetailRes struct {

    /* 节点名称  */
    TaskName string `json:"taskName"`

    /* 子节点code List (Optional) */
    ChildrenCode []string `json:"childrenCode"`

    /* 父节点code List (Optional) */
    ParentCode []string `json:"parentCode"`

    /* 节点描述 (Optional) */
    TaskDesc *string `json:"taskDesc"`

    /* 任务类型  */
    TaskNodeId int `json:"taskNodeId"`

    /* 节点CODE  */
    TaskCode string `json:"taskCode"`

    /* 负责人  */
    Manager string `json:"manager"`

    /* 节点数据 以下前端需要用到的数据 (Optional) */
    TaskData *string `json:"taskData"`

    /* 节点名称 以下前端需要用到的数据 (Optional) */
    NodeName *string `json:"nodeName"`

    /* 节点类型 (Optional) */
    NodeTypeName *string `json:"nodeTypeName"`

    /* 节点图标 以下前端需要用到的数据 (Optional) */
    NodeIcon *string `json:"nodeIcon"`

    /* 工作流Code  */
    FlowCode string `json:"flowCode"`

    /* 脚本文件的业务Code (Optional) */
    FileCode *string `json:"fileCode"`

    /* 作业名称 (Optional) */
    JobName *string `json:"jobName"`

    /* 是否工作流任务 (Optional) */
    IsUranus *int `json:"isUranus"`

    /* 是否当前工作流任务 (Optional) */
    IsCurrentFlow *int `json:"isCurrentFlow"`

    /* 给前端用到的 (Optional) */
    Parent []UranusTaskNodeSaveReq `json:"parent"`

    /* 是否有子任务依赖前端用到 (Optional) */
    HaveChildren *int `json:"haveChildren"`

    /* 协作人 (Optional) */
    Worker *string `json:"worker"`

    /* 发布变更状态 (Optional) */
    TaskChangeStatus []int `json:"taskChangeStatus"`

    /* 发布变更状态描述 (Optional) */
    TaskChangeStatusDesc []string `json:"taskChangeStatusDesc"`

    /* 作业状态 (Optional) */
    GravityStatus *int `json:"gravityStatus"`

    /* 作业状态描述 (Optional) */
    GravityStatusDesc *string `json:"gravityStatusDesc"`

    /* 锁定状态 0 无锁 1 锁定 默认 为 1 (Optional) */
    LockStatus *int `json:"lockStatus"`

    /* 解锁人员 (Optional) */
    LockUser *string `json:"lockUser"`
}
