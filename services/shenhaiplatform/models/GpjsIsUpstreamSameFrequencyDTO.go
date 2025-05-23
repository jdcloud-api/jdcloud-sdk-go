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


type GpjsIsUpstreamSameFrequencyDTO struct {

    /* 作业名称 (Optional) */
    JobName string `json:"jobName"`

    /* 依赖作业名称清单 (Optional) */
    DependJobName []string `json:"dependJobName"`

    /* 作业开始时间 (Optional) */
    StartTime string `json:"startTime"`

    /* 作业结束时间 (Optional) */
    EndTime string `json:"endTime"`

    /* 作业时间间隔 (Optional) */
    Interval int `json:"interval"`

    /* 作业运行周期 (Optional) */
    Cycle string `json:"cycle"`
}
