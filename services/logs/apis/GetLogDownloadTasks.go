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
    logs "github.com/jdcloud-api/jdcloud-sdk-go/services/logs/models"
)

type GetLogDownloadTasksRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 日志主题 UID  */
    LogtopicUID string `json:"logtopicUID"`
}

/*
 * param regionId: 地域 Id (Required)
 * param logtopicUID: 日志主题 UID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetLogDownloadTasksRequest(
    regionId string,
    logtopicUID string,
) *GetLogDownloadTasksRequest {

	return &GetLogDownloadTasksRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/logtopics/{logtopicUID}/downloadtasks",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        LogtopicUID: logtopicUID,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param logtopicUID: 日志主题 UID (Required)
 */
func NewGetLogDownloadTasksRequestWithAllParams(
    regionId string,
    logtopicUID string,
) *GetLogDownloadTasksRequest {

    return &GetLogDownloadTasksRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/logtopics/{logtopicUID}/downloadtasks",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        LogtopicUID: logtopicUID,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetLogDownloadTasksRequestWithoutParam() *GetLogDownloadTasksRequest {

    return &GetLogDownloadTasksRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/logtopics/{logtopicUID}/downloadtasks",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *GetLogDownloadTasksRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param logtopicUID: 日志主题 UID(Required) */
func (r *GetLogDownloadTasksRequest) SetLogtopicUID(logtopicUID string) {
    r.LogtopicUID = logtopicUID
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetLogDownloadTasksRequest) GetRegionId() string {
    return r.RegionId
}

type GetLogDownloadTasksResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetLogDownloadTasksResult `json:"result"`
}

type GetLogDownloadTasksResult struct {
    Data []logs.LogDownloadTask `json:"data"`
}