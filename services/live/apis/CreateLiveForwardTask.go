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

package apis

import (
    "github.com/lshuining/jdcloud-sdk-go/core"
)

type CreateLiveForwardTaskRequest struct {

    core.JDCloudRequest

    /* 拉流地址
- 支持rtmp
  */
    SourceUrl string `json:"sourceUrl"`

    /* 转推地址
- 支持rtmp
  */
    PushUrl string `json:"pushUrl"`

    /* 执行方式
- StartNow: 立即执行
- StartAsScheduled: 定时执行，根据参数设定的时间
  */
    StartMode string `json:"startMode"`

    /* 开始时间
- UTC时间， ISO8601示例：2021-07-26T08:08:08Z
- 不填表示立即开始
 (Optional) */
    StartTime *string `json:"startTime"`

    /* 结束时间
- UTC时间， ISO8601示例：2021-07-26T08:08:08Z
- 最大支持365天，与开始时间间隔不超过7天。
  - 假设当前时间2021年03月30日11:50:01，则：结束时间不可超过2022年03月30日11:50:01。
- 不填拉不到流10分钟自动结束
 (Optional) */
    EndTime *string `json:"endTime"`

    /* 回调类型
- 不填发送全部回调
- TaskStart 任务开始
- TaskExit 任务结束
- callbackUrl非空的情况下，callbackEvents有效
 (Optional) */
    CallbackEvents []string `json:"callbackEvents"`

    /* 事件回调地址
 (Optional) */
    CallbackUrl *string `json:"callbackUrl"`

    /* 任务名称
- 最大255字符
 (Optional) */
    Name *string `json:"name"`
}

/*
 * param sourceUrl: 拉流地址
- 支持rtmp
 (Required)
 * param pushUrl: 转推地址
- 支持rtmp
 (Required)
 * param startMode: 执行方式
- StartNow: 立即执行
- StartAsScheduled: 定时执行，根据参数设定的时间
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateLiveForwardTaskRequest(
    sourceUrl string,
    pushUrl string,
    startMode string,
) *CreateLiveForwardTaskRequest {

	return &CreateLiveForwardTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/LiveForwardTask:create",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        SourceUrl: sourceUrl,
        PushUrl: pushUrl,
        StartMode: startMode,
	}
}

/*
 * param sourceUrl: 拉流地址
- 支持rtmp
 (Required)
 * param pushUrl: 转推地址
- 支持rtmp
 (Required)
 * param startMode: 执行方式
- StartNow: 立即执行
- StartAsScheduled: 定时执行，根据参数设定的时间
 (Required)
 * param startTime: 开始时间
- UTC时间， ISO8601示例：2021-07-26T08:08:08Z
- 不填表示立即开始
 (Optional)
 * param endTime: 结束时间
- UTC时间， ISO8601示例：2021-07-26T08:08:08Z
- 最大支持365天，与开始时间间隔不超过7天。
  - 假设当前时间2021年03月30日11:50:01，则：结束时间不可超过2022年03月30日11:50:01。
- 不填拉不到流10分钟自动结束
 (Optional)
 * param callbackEvents: 回调类型
- 不填发送全部回调
- TaskStart 任务开始
- TaskExit 任务结束
- callbackUrl非空的情况下，callbackEvents有效
 (Optional)
 * param callbackUrl: 事件回调地址
 (Optional)
 * param name: 任务名称
- 最大255字符
 (Optional)
 */
func NewCreateLiveForwardTaskRequestWithAllParams(
    sourceUrl string,
    pushUrl string,
    startMode string,
    startTime *string,
    endTime *string,
    callbackEvents []string,
    callbackUrl *string,
    name *string,
) *CreateLiveForwardTaskRequest {

    return &CreateLiveForwardTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/LiveForwardTask:create",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        SourceUrl: sourceUrl,
        PushUrl: pushUrl,
        StartMode: startMode,
        StartTime: startTime,
        EndTime: endTime,
        CallbackEvents: callbackEvents,
        CallbackUrl: callbackUrl,
        Name: name,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateLiveForwardTaskRequestWithoutParam() *CreateLiveForwardTaskRequest {

    return &CreateLiveForwardTaskRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/LiveForwardTask:create",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param sourceUrl: 拉流地址
- 支持rtmp
(Required) */
func (r *CreateLiveForwardTaskRequest) SetSourceUrl(sourceUrl string) {
    r.SourceUrl = sourceUrl
}

/* param pushUrl: 转推地址
- 支持rtmp
(Required) */
func (r *CreateLiveForwardTaskRequest) SetPushUrl(pushUrl string) {
    r.PushUrl = pushUrl
}

/* param startMode: 执行方式
- StartNow: 立即执行
- StartAsScheduled: 定时执行，根据参数设定的时间
(Required) */
func (r *CreateLiveForwardTaskRequest) SetStartMode(startMode string) {
    r.StartMode = startMode
}

/* param startTime: 开始时间
- UTC时间， ISO8601示例：2021-07-26T08:08:08Z
- 不填表示立即开始
(Optional) */
func (r *CreateLiveForwardTaskRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param endTime: 结束时间
- UTC时间， ISO8601示例：2021-07-26T08:08:08Z
- 最大支持365天，与开始时间间隔不超过7天。
  - 假设当前时间2021年03月30日11:50:01，则：结束时间不可超过2022年03月30日11:50:01。
- 不填拉不到流10分钟自动结束
(Optional) */
func (r *CreateLiveForwardTaskRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

/* param callbackEvents: 回调类型
- 不填发送全部回调
- TaskStart 任务开始
- TaskExit 任务结束
- callbackUrl非空的情况下，callbackEvents有效
(Optional) */
func (r *CreateLiveForwardTaskRequest) SetCallbackEvents(callbackEvents []string) {
    r.CallbackEvents = callbackEvents
}

/* param callbackUrl: 事件回调地址
(Optional) */
func (r *CreateLiveForwardTaskRequest) SetCallbackUrl(callbackUrl string) {
    r.CallbackUrl = &callbackUrl
}

/* param name: 任务名称
- 最大255字符
(Optional) */
func (r *CreateLiveForwardTaskRequest) SetName(name string) {
    r.Name = &name
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateLiveForwardTaskRequest) GetRegionId() string {
    return ""
}

type CreateLiveForwardTaskResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateLiveForwardTaskResult `json:"result"`
}

type CreateLiveForwardTaskResult struct {
    TaskId string `json:"taskId"`
}