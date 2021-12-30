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
    vm "github.com/jdcloud-api/jdcloud-sdk-go/services/vm/models"
)

type ImageTasksRequest struct {

    core.JDCloudRequest

    /* 地域ID。  */
    RegionId string `json:"regionId"`

    /* 任务操作类型。支持范围：`ImportImage、ExportImage`。 (Optional) */
    TaskAction *string `json:"taskAction"`

    /* 任务id列表。 (Optional) */
    TaskIds []string `json:"taskIds"`

    /* 任务状态。支持范围：`pending、running、failed、finished`。 (Optional) */
    TaskStatus *string `json:"taskStatus"`

    /* 任务开始时间 (Optional) */
    StartTime *string `json:"startTime"`

    /* 任务结束时间 (Optional) */
    EndTime *string `json:"endTime"`

    /* 页码；默认为1。 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；取值范围[10, 100]。 (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: 地域ID。 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewImageTasksRequest(
    regionId string,
) *ImageTasksRequest {

	return &ImageTasksRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/imageTasks",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID。 (Required)
 * param taskAction: 任务操作类型。支持范围：`ImportImage、ExportImage`。 (Optional)
 * param taskIds: 任务id列表。 (Optional)
 * param taskStatus: 任务状态。支持范围：`pending、running、failed、finished`。 (Optional)
 * param startTime: 任务开始时间 (Optional)
 * param endTime: 任务结束时间 (Optional)
 * param pageNumber: 页码；默认为1。 (Optional)
 * param pageSize: 分页大小；取值范围[10, 100]。 (Optional)
 */
func NewImageTasksRequestWithAllParams(
    regionId string,
    taskAction *string,
    taskIds []string,
    taskStatus *string,
    startTime *string,
    endTime *string,
    pageNumber *int,
    pageSize *int,
) *ImageTasksRequest {

    return &ImageTasksRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/imageTasks",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        TaskAction: taskAction,
        TaskIds: taskIds,
        TaskStatus: taskStatus,
        StartTime: startTime,
        EndTime: endTime,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewImageTasksRequestWithoutParam() *ImageTasksRequest {

    return &ImageTasksRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/imageTasks",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID。(Required) */
func (r *ImageTasksRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param taskAction: 任务操作类型。支持范围：`ImportImage、ExportImage`。(Optional) */
func (r *ImageTasksRequest) SetTaskAction(taskAction string) {
    r.TaskAction = &taskAction
}

/* param taskIds: 任务id列表。(Optional) */
func (r *ImageTasksRequest) SetTaskIds(taskIds []string) {
    r.TaskIds = taskIds
}

/* param taskStatus: 任务状态。支持范围：`pending、running、failed、finished`。(Optional) */
func (r *ImageTasksRequest) SetTaskStatus(taskStatus string) {
    r.TaskStatus = &taskStatus
}

/* param startTime: 任务开始时间(Optional) */
func (r *ImageTasksRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param endTime: 任务结束时间(Optional) */
func (r *ImageTasksRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

/* param pageNumber: 页码；默认为1。(Optional) */
func (r *ImageTasksRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小；取值范围[10, 100]。(Optional) */
func (r *ImageTasksRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ImageTasksRequest) GetRegionId() string {
    return r.RegionId
}

type ImageTasksResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ImageTasksResult `json:"result"`
}

type ImageTasksResult struct {
    TaskSet []vm.TaskInfo `json:"taskSet"`
    TotalCount int `json:"totalCount"`
}