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


type BfdSpec struct {

    /* BFD报文的最小发送间隔，取值200～1000ms，默认值1000ms (Optional) */
    MinTxInterval int `json:"minTxInterval"`

    /* BFD报文的最小接受间隔，取值200～1000ms，默认值1000ms (Optional) */
    MinRxInterval int `json:"minRxInterval"`

    /* 接收方允许发送方发送BFD控制报文的最大连续丢包数，取值范围为3-10，默认为3 (Optional) */
    DetectMultiplier int `json:"detectMultiplier"`
}
