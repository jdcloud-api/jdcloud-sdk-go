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


type UranusCatalogInfoList struct {

    /* 前端排序唯一ID (Optional) */
    UnrealId string `json:"unrealId"`

    /* 目录code  */
    CatalogCode string `json:"catalogCode"`

    /* 目录名称 (Optional) */
    CatalogName string `json:"catalogName"`

    /* 0：非叶子目录， 1：叶子目录 (Optional) */
    CatalogType int `json:"catalogType"`

    /* 父目录code (Optional) */
    ParentCode string `json:"parentCode"`

    /* 子节点数量 (Optional) */
    ChildrenNum int `json:"childrenNum"`

    /* 流程列表 (Optional) */
    TaskFlowList []UranusTaskFlowListRes `json:"taskFlowList"`
}
