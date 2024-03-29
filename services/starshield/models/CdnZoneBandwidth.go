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


type CdnZoneBandwidth struct {

    /* 域名 (Optional) */
    ZoneName string `json:"zoneName"`

    /* 带宽峰值，单位bps（bit per second） (Optional) */
    MaxBPS float64 `json:"maxBPS"`

    /* 带宽峰值的发生时间。值为时间戳对应的long值。 (Optional) */
    MaxBpsTs int `json:"maxBpsTs"`

    /* 总流量，单位Byte (Optional) */
    TrafficSum float64 `json:"trafficSum"`

    /* 总请求量，单位次数 (Optional) */
    RequestSum float64 `json:"requestSum"`
}
