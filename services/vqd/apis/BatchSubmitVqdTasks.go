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
    vqd "github.com/lshuining/jdcloud-sdk-go/services/vqd/models"
)

type BatchSubmitVqdTasksRequest struct {

    core.JDCloudRequest

    /* 媒体列表  */
    MediaList []vqd.VqdMediaObject `json:"mediaList"`

    /* 检测模板ID  */
    TemplateId string `json:"templateId"`
}

/*
 * param mediaList: 媒体列表 (Required)
 * param templateId: 检测模板ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewBatchSubmitVqdTasksRequest(
    mediaList []vqd.VqdMediaObject,
    templateId string,
) *BatchSubmitVqdTasksRequest {

	return &BatchSubmitVqdTasksRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/vqdTasks:batchSubmit",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        MediaList: mediaList,
        TemplateId: templateId,
	}
}

/*
 * param mediaList: 媒体列表 (Required)
 * param templateId: 检测模板ID (Required)
 */
func NewBatchSubmitVqdTasksRequestWithAllParams(
    mediaList []vqd.VqdMediaObject,
    templateId string,
) *BatchSubmitVqdTasksRequest {

    return &BatchSubmitVqdTasksRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/vqdTasks:batchSubmit",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        MediaList: mediaList,
        TemplateId: templateId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewBatchSubmitVqdTasksRequestWithoutParam() *BatchSubmitVqdTasksRequest {

    return &BatchSubmitVqdTasksRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/vqdTasks:batchSubmit",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param mediaList: 媒体列表(Required) */
func (r *BatchSubmitVqdTasksRequest) SetMediaList(mediaList []vqd.VqdMediaObject) {
    r.MediaList = mediaList
}

/* param templateId: 检测模板ID(Required) */
func (r *BatchSubmitVqdTasksRequest) SetTemplateId(templateId string) {
    r.TemplateId = templateId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r BatchSubmitVqdTasksRequest) GetRegionId() string {
    return ""
}

type BatchSubmitVqdTasksResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result BatchSubmitVqdTasksResult `json:"result"`
}

type BatchSubmitVqdTasksResult struct {
    TaskIds []string `json:"taskIds"`
}