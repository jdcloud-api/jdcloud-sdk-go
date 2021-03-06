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


type TestMetricTaskSpec struct {

    /* 聚合函数,支持 count sum max min avg  */
    Aggregate string `json:"aggregate"`

    /* 测试内容  */
    Content []string `json:"content"`

    /* 查询字段,支持 英文字母 数字 下划线 中划线 点（中文日志原文和各产品线的key）  */
    DataField string `json:"dataField"`

    /* 过滤语法，可以为空 (Optional) */
    FilterContent string `json:"filterContent"`

    /* 是否打开过滤  */
    FilterOpen string `json:"filterOpen"`

    /* 过滤类型，只能是fulltext和 advance  */
    FilterType string `json:"filterType"`
}
