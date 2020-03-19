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


type VideoDetailList struct {

    /* 音频总时长,单位为分钟 (Optional) */
    Audio_time []int `json:"audio_time"`

    /* 涉政截帧量 (Optional) */
    Politics_frame_count []int `json:"politics_frame_count"`

    /* 涉黄截帧量 (Optional) */
    Porn_frame_count []int `json:"porn_frame_count"`

    /* 暴恐截帧量 (Optional) */
    Terrorism_frame_count []int `json:"terrorism_frame_count"`

    /* 涉政暴恐截帧量 (Optional) */
    Politics_terrorism_frame_count []int `json:"politics_terrorism_frame_count"`

    /* 总截帧量 (Optional) */
    Total_frame_count []int `json:"total_frame_count"`
}
