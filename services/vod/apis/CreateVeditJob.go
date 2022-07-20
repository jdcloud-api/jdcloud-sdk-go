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
    vod "github.com/lshuining/jdcloud-sdk-go/services/vod/models"
)

type CreateVeditJobRequest struct {

    core.JDCloudRequest

    /* 工程名称  */
    ProjectName string `json:"projectName"`

    /* 工程描述 (Optional) */
    Description *string `json:"description"`

    /* 时间线信息  */
    Timeline *vod.Timeline `json:"timeline"`

    /* 剪辑合成媒资元数据 (Optional) */
    MediaMetadata *vod.MediaMetadata `json:"mediaMetadata"`

    /* 用户数据，JSON格式的字符串。
在Timeline中的所有MediaClip中，若有2个或以上的不同MediaId，即素材片段来源于2个或以上不同视频，则在提交剪辑作业时，必须在UserData中指明合并后的视频画面的宽高。
如 {\"extendData\": {\"width\": 720, \"height\": 500}}，其中width和height必须为[16, 4096]之间的偶数
videoMode 支持 normal 普通模式 screen_record 屏幕录制模式 两种模式，默认为 normal。
如 "{\"extendData\":{\"videoMode\":\"screen_record\"}}"
 (Optional) */
    UserData *string `json:"userData"`
}

/*
 * param projectName: 工程名称 (Required)
 * param timeline: 时间线信息 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateVeditJobRequest(
    projectName string,
    timeline *vod.Timeline,
) *CreateVeditJobRequest {

	return &CreateVeditJobRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/veditJobs",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        ProjectName: projectName,
        Timeline: timeline,
	}
}

/*
 * param projectName: 工程名称 (Required)
 * param description: 工程描述 (Optional)
 * param timeline: 时间线信息 (Required)
 * param mediaMetadata: 剪辑合成媒资元数据 (Optional)
 * param userData: 用户数据，JSON格式的字符串。
在Timeline中的所有MediaClip中，若有2个或以上的不同MediaId，即素材片段来源于2个或以上不同视频，则在提交剪辑作业时，必须在UserData中指明合并后的视频画面的宽高。
如 {\"extendData\": {\"width\": 720, \"height\": 500}}，其中width和height必须为[16, 4096]之间的偶数
videoMode 支持 normal 普通模式 screen_record 屏幕录制模式 两种模式，默认为 normal。
如 "{\"extendData\":{\"videoMode\":\"screen_record\"}}"
 (Optional)
 */
func NewCreateVeditJobRequestWithAllParams(
    projectName string,
    description *string,
    timeline *vod.Timeline,
    mediaMetadata *vod.MediaMetadata,
    userData *string,
) *CreateVeditJobRequest {

    return &CreateVeditJobRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/veditJobs",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        ProjectName: projectName,
        Description: description,
        Timeline: timeline,
        MediaMetadata: mediaMetadata,
        UserData: userData,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateVeditJobRequestWithoutParam() *CreateVeditJobRequest {

    return &CreateVeditJobRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/veditJobs",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param projectName: 工程名称(Required) */
func (r *CreateVeditJobRequest) SetProjectName(projectName string) {
    r.ProjectName = projectName
}

/* param description: 工程描述(Optional) */
func (r *CreateVeditJobRequest) SetDescription(description string) {
    r.Description = &description
}

/* param timeline: 时间线信息(Required) */
func (r *CreateVeditJobRequest) SetTimeline(timeline *vod.Timeline) {
    r.Timeline = timeline
}

/* param mediaMetadata: 剪辑合成媒资元数据(Optional) */
func (r *CreateVeditJobRequest) SetMediaMetadata(mediaMetadata *vod.MediaMetadata) {
    r.MediaMetadata = mediaMetadata
}

/* param userData: 用户数据，JSON格式的字符串。
在Timeline中的所有MediaClip中，若有2个或以上的不同MediaId，即素材片段来源于2个或以上不同视频，则在提交剪辑作业时，必须在UserData中指明合并后的视频画面的宽高。
如 {\"extendData\": {\"width\": 720, \"height\": 500}}，其中width和height必须为[16, 4096]之间的偶数
videoMode 支持 normal 普通模式 screen_record 屏幕录制模式 两种模式，默认为 normal。
如 "{\"extendData\":{\"videoMode\":\"screen_record\"}}"
(Optional) */
func (r *CreateVeditJobRequest) SetUserData(userData string) {
    r.UserData = &userData
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateVeditJobRequest) GetRegionId() string {
    return ""
}

type CreateVeditJobResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateVeditJobResult `json:"result"`
}

type CreateVeditJobResult struct {
    JobId int64 `json:"jobId"`
    ProjectId int64 `json:"projectId"`
}