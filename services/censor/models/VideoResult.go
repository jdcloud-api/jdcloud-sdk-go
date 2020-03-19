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


type VideoResult struct {

    /* 错误码，和HTTP的status code一致 (Optional) */
    Code int `json:"code"`

    /* 错误描述信息 (Optional) */
    Msg string `json:"msg"`

    /* 对应请求的dataId (Optional) */
    DataId string `json:"dataId"`

    /* 该检测任务的ID (Optional) */
    TaskId string `json:"taskId"`

    /* 返回结果。调用成功时（code=200），返回结果中包含一个或多个元素。每个元素是个结构体，具体结构描述见VideoResultDetail (Optional) */
    Results []VideoResultDetail `json:"results"`

    /* 视频语音检测结果。具体结构描述见audioScanResult。 (Optional) */
    AudioResults []AudioResultDetail `json:"audioResults"`
}
