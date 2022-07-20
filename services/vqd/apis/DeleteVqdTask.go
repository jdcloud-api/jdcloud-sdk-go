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

type DeleteVqdTaskRequest struct {

    core.JDCloudRequest

    /* 任务ID，路径参数  */
    TaskId string `json:"taskId"`
}

/*
 * param taskId: 任务ID，路径参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteVqdTaskRequest(
    taskId string,
) *DeleteVqdTaskRequest {

	return &DeleteVqdTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/vqdTasks/{taskId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        TaskId: taskId,
	}
}

/*
 * param taskId: 任务ID，路径参数 (Required)
 */
func NewDeleteVqdTaskRequestWithAllParams(
    taskId string,
) *DeleteVqdTaskRequest {

    return &DeleteVqdTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/vqdTasks/{taskId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        TaskId: taskId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteVqdTaskRequestWithoutParam() *DeleteVqdTaskRequest {

    return &DeleteVqdTaskRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/vqdTasks/{taskId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param taskId: 任务ID，路径参数(Required) */
func (r *DeleteVqdTaskRequest) SetTaskId(taskId string) {
    r.TaskId = taskId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteVqdTaskRequest) GetRegionId() string {
    return ""
}

type DeleteVqdTaskResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteVqdTaskResult `json:"result"`
}

type DeleteVqdTaskResult struct {
}