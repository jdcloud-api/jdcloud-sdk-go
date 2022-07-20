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
    live "github.com/lshuining/jdcloud-sdk-go/services/live/models"
)

type DescribeLiveSnapshotDataRequest struct {

    core.JDCloudRequest

    /* 推流域名 (Optional) */
    PublishDomain *string `json:"publishDomain"`

    /* 应用名称 (Optional) */
    AppName *string `json:"appName"`

    /* 流名称 (Optional) */
    StreamName *string `json:"streamName"`

    /* 截图模式：1表示采样截图；2表示关键帧截图(默认为2) (Optional) */
    ShotMode *int `json:"shotMode"`

    /* 起始时间:
- UTC时间
  格式: yyyy-MM-dd'T'HH:mm:ss'Z'
  示例: 2018-10-21T10:00:00Z
- 支持最大查询90天以内的数据
  */
    StartTime string `json:"startTime"`

    /* 结束时间:
- UTC时间
  格式: yyyy-MM-dd'T'HH:mm:ss'Z'
  示例: 2018-10-21T10:00:00Z
- 为空,默认当前时间
 (Optional) */
    EndTime *string `json:"endTime"`
}

/*
 * param startTime: 起始时间:
- UTC时间
  格式: yyyy-MM-dd'T'HH:mm:ss'Z'
  示例: 2018-10-21T10:00:00Z
- 支持最大查询90天以内的数据
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeLiveSnapshotDataRequest(
    startTime string,
) *DescribeLiveSnapshotDataRequest {

	return &DescribeLiveSnapshotDataRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/liveSnapshotData",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        StartTime: startTime,
	}
}

/*
 * param publishDomain: 推流域名 (Optional)
 * param appName: 应用名称 (Optional)
 * param streamName: 流名称 (Optional)
 * param shotMode: 截图模式：1表示采样截图；2表示关键帧截图(默认为2) (Optional)
 * param startTime: 起始时间:
- UTC时间
  格式: yyyy-MM-dd'T'HH:mm:ss'Z'
  示例: 2018-10-21T10:00:00Z
- 支持最大查询90天以内的数据
 (Required)
 * param endTime: 结束时间:
- UTC时间
  格式: yyyy-MM-dd'T'HH:mm:ss'Z'
  示例: 2018-10-21T10:00:00Z
- 为空,默认当前时间
 (Optional)
 */
func NewDescribeLiveSnapshotDataRequestWithAllParams(
    publishDomain *string,
    appName *string,
    streamName *string,
    shotMode *int,
    startTime string,
    endTime *string,
) *DescribeLiveSnapshotDataRequest {

    return &DescribeLiveSnapshotDataRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveSnapshotData",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PublishDomain: publishDomain,
        AppName: appName,
        StreamName: streamName,
        ShotMode: shotMode,
        StartTime: startTime,
        EndTime: endTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeLiveSnapshotDataRequestWithoutParam() *DescribeLiveSnapshotDataRequest {

    return &DescribeLiveSnapshotDataRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveSnapshotData",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param publishDomain: 推流域名(Optional) */
func (r *DescribeLiveSnapshotDataRequest) SetPublishDomain(publishDomain string) {
    r.PublishDomain = &publishDomain
}

/* param appName: 应用名称(Optional) */
func (r *DescribeLiveSnapshotDataRequest) SetAppName(appName string) {
    r.AppName = &appName
}

/* param streamName: 流名称(Optional) */
func (r *DescribeLiveSnapshotDataRequest) SetStreamName(streamName string) {
    r.StreamName = &streamName
}

/* param shotMode: 截图模式：1表示采样截图；2表示关键帧截图(默认为2)(Optional) */
func (r *DescribeLiveSnapshotDataRequest) SetShotMode(shotMode int) {
    r.ShotMode = &shotMode
}

/* param startTime: 起始时间:
- UTC时间
  格式: yyyy-MM-dd'T'HH:mm:ss'Z'
  示例: 2018-10-21T10:00:00Z
- 支持最大查询90天以内的数据
(Required) */
func (r *DescribeLiveSnapshotDataRequest) SetStartTime(startTime string) {
    r.StartTime = startTime
}

/* param endTime: 结束时间:
- UTC时间
  格式: yyyy-MM-dd'T'HH:mm:ss'Z'
  示例: 2018-10-21T10:00:00Z
- 为空,默认当前时间
(Optional) */
func (r *DescribeLiveSnapshotDataRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeLiveSnapshotDataRequest) GetRegionId() string {
    return ""
}

type DescribeLiveSnapshotDataResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeLiveSnapshotDataResult `json:"result"`
}

type DescribeLiveSnapshotDataResult struct {
    SnapshotData []live.SnapshotCountStatisticResult `json:"snapshotData"`
}