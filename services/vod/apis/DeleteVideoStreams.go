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
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
)

type DeleteVideoStreamsRequest struct {

    core.JDCloudRequest

    /* 视频ID  */
    VideoId string `json:"videoId"`

    /*  (Optional) */
    TaskIds []int64 `json:"taskIds"`
}

/*
 * param videoId: 视频ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteVideoStreamsRequest(
    videoId string,
) *DeleteVideoStreamsRequest {

	return &DeleteVideoStreamsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/videos/{videoId}:deleteStreams",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        VideoId: videoId,
	}
}

/*
 * param videoId: 视频ID (Required)
 * param taskIds:  (Optional)
 */
func NewDeleteVideoStreamsRequestWithAllParams(
    videoId string,
    taskIds []int64,
) *DeleteVideoStreamsRequest {

    return &DeleteVideoStreamsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/videos/{videoId}:deleteStreams",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        VideoId: videoId,
        TaskIds: taskIds,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteVideoStreamsRequestWithoutParam() *DeleteVideoStreamsRequest {

    return &DeleteVideoStreamsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/videos/{videoId}:deleteStreams",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param videoId: 视频ID(Required) */
func (r *DeleteVideoStreamsRequest) SetVideoId(videoId string) {
    r.VideoId = videoId
}

/* param taskIds: (Optional) */
func (r *DeleteVideoStreamsRequest) SetTaskIds(taskIds []int64) {
    r.TaskIds = taskIds
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteVideoStreamsRequest) GetRegionId() string {
    return ""
}

type DeleteVideoStreamsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteVideoStreamsResult `json:"result"`
}

type DeleteVideoStreamsResult struct {
}