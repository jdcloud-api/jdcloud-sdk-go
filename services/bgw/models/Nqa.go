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


type Nqa struct {

    /* 连续两次发送探测报文的时间间隔。 (Optional) */
    TxInterval int `json:"txInterval"`

    /* 指定发送连续探测报文的个数。 (Optional) */
    ThresholdCount int `json:"thresholdCount"`

    /* 一次NQA探测的超时时间。 (Optional) */
    ResponseTimeout int `json:"responseTimeout"`
}
