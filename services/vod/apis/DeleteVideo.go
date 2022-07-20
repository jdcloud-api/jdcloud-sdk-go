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

type DeleteVideoRequest struct {

    core.JDCloudRequest

    /* 视频ID  */
    VideoId string `json:"videoId"`
}

/*
 * param videoId: 视频ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteVideoRequest(
    videoId string,
) *DeleteVideoRequest {

	return &DeleteVideoRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/videos/{videoId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        VideoId: videoId,
	}
}

/*
 * param videoId: 视频ID (Required)
 */
func NewDeleteVideoRequestWithAllParams(
    videoId string,
) *DeleteVideoRequest {

    return &DeleteVideoRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/videos/{videoId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        VideoId: videoId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteVideoRequestWithoutParam() *DeleteVideoRequest {

    return &DeleteVideoRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/videos/{videoId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param videoId: 视频ID(Required) */
func (r *DeleteVideoRequest) SetVideoId(videoId string) {
    r.VideoId = videoId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteVideoRequest) GetRegionId() string {
    return ""
}

type DeleteVideoResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteVideoResult `json:"result"`
}

type DeleteVideoResult struct {
}