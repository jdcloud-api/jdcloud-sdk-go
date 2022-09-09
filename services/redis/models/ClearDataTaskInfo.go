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


type ClearDataTaskInfo struct {

    /* 数据清理任务类型 (Optional) */
    ClearType string `json:"clearType"`

    /* 匹配模式, 如: test*、*test、ab*cc*, 当节点为AllData、ExpiredData时可以忽略此参数 (Optional) */
    KeyPattern string `json:"keyPattern"`

    /* key的过滤条件, 当节点为AllData、ExpiredData时可以忽略此参数 (Optional) */
    KeyFilter []KeyFilter `json:"keyFilter"`

    /* 数据遍历的速率 (Optional) */
    QpsLimit int `json:"qpsLimit"`

    /* 本次清理任务的完成进度 (Optional) */
    Progress int `json:"progress"`

    /* 本次清理任务已成功清理的key数量 (Optional) */
    EffectKeys int `json:"effectKeys"`

    /* 本次清理任务的状态:running、success、failed (Optional) */
    Status string `json:"status"`

    /* 本次清理任务失败的原因 (Optional) */
    Message string `json:"message"`

    /* 最近一次status转换的时间 (Optional) */
    LastTransitionTime string `json:"lastTransitionTime"`
}