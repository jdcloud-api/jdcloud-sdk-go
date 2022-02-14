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


type MetrictaskDetailEnd struct {

    /* 聚合函数 (Optional) */
    Aggregate string `json:"aggregate"`

    /*  (Optional) */
    CreatedTime string `json:"createdTime"`

    /* 开启自定义单位 (Optional) */
    CustomUnit string `json:"customUnit"`

    /* 查询字段 (Optional) */
    DataField string `json:"dataField"`

    /* 过滤语法 (Optional) */
    FilterContent string `json:"filterContent"`

    /*  (Optional) */
    FilterOpen string `json:"filterOpen"`

    /* 过滤类型 (Optional) */
    FilterType string `json:"filterType"`

    /* id (Optional) */
    Id string `json:"id"`

    /* 周期 (Optional) */
    Interval int64 `json:"interval"`

    /* 监控项名称 (Optional) */
    Metric string `json:"metric"`

    /* 监控任务名称 (Optional) */
    Name string `json:"name"`

    /* 配置方式:枚举值 visual，sql；分别代表可视化配置及sql配置方式 (Optional) */
    SettingType string `json:"settingType"`

    /*  (Optional) */
    SqlSpec MetricTaskSqlSpec `json:"sqlSpec"`

    /* 单位 (Optional) */
    Unit string `json:"unit"`

    /*  (Optional) */
    UpdateTime string `json:"updateTime"`
}
