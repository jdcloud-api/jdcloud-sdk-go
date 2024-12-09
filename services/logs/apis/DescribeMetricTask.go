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

type DescribeMetricTaskRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 日志集 UID  */
    LogsetUID string `json:"logsetUID"`

    /* 日志主题 UID  */
    LogtopicUID string `json:"logtopicUID"`

    /*   */
    LogmetrictaskUID string `json:"logmetrictaskUID"`
}

/*
 * param regionId: 地域 Id (Required)
 * param logsetUID: 日志集 UID (Required)
 * param logtopicUID: 日志主题 UID (Required)
 * param logmetrictaskUID:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeMetricTaskRequest(
    regionId string,
    logsetUID string,
    logtopicUID string,
    logmetrictaskUID string,
) *DescribeMetricTaskRequest {

	return &DescribeMetricTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/logsets/{logsetUID}/logtopics/{logtopicUID}/metrictasks/{logmetrictaskUID}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        LogsetUID: logsetUID,
        LogtopicUID: logtopicUID,
        LogmetrictaskUID: logmetrictaskUID,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param logsetUID: 日志集 UID (Required)
 * param logtopicUID: 日志主题 UID (Required)
 * param logmetrictaskUID:  (Required)
 */
func NewDescribeMetricTaskRequestWithAllParams(
    regionId string,
    logsetUID string,
    logtopicUID string,
    logmetrictaskUID string,
) *DescribeMetricTaskRequest {

    return &DescribeMetricTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/logsets/{logsetUID}/logtopics/{logtopicUID}/metrictasks/{logmetrictaskUID}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        LogsetUID: logsetUID,
        LogtopicUID: logtopicUID,
        LogmetrictaskUID: logmetrictaskUID,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeMetricTaskRequestWithoutParam() *DescribeMetricTaskRequest {

    return &DescribeMetricTaskRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/logsets/{logsetUID}/logtopics/{logtopicUID}/metrictasks/{logmetrictaskUID}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *DescribeMetricTaskRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param logsetUID: 日志集 UID(Required) */
func (r *DescribeMetricTaskRequest) SetLogsetUID(logsetUID string) {
    r.LogsetUID = logsetUID
}
/* param logtopicUID: 日志主题 UID(Required) */
func (r *DescribeMetricTaskRequest) SetLogtopicUID(logtopicUID string) {
    r.LogtopicUID = logtopicUID
}
/* param logmetrictaskUID: (Required) */
func (r *DescribeMetricTaskRequest) SetLogmetrictaskUID(logmetrictaskUID string) {
    r.LogmetrictaskUID = logmetrictaskUID
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeMetricTaskRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeMetricTaskResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeMetricTaskResult `json:"result"`
}

type DescribeMetricTaskResult struct {
    Data logs.MetrictaskDetailEnd `json:"data"`
}