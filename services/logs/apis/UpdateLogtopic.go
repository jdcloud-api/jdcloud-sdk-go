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

type UpdateLogtopicRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 日志主题 UID  */
    LogtopicUID string `json:"logtopicUID"`

    /* 日志主题名称 (Optional) */
    Name *string `json:"name"`

    /* 日志主题描述  */
    Description string `json:"description"`

    /* 保存周期，只能是 7， 15， 30 (Optional) */
    LifeCycle *int `json:"lifeCycle"`

    /* 保序 (Optional) */
    InOrder *bool `json:"inOrder"`
}

/*
 * param regionId: 地域 Id (Required)
 * param logtopicUID: 日志主题 UID (Required)
 * param description: 日志主题描述 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateLogtopicRequest(
    regionId string,
    logtopicUID string,
    description string,
) *UpdateLogtopicRequest {

	return &UpdateLogtopicRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/logtopics/{logtopicUID}",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        LogtopicUID: logtopicUID,
        Description: description,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param logtopicUID: 日志主题 UID (Required)
 * param name: 日志主题名称 (Optional)
 * param description: 日志主题描述 (Required)
 * param lifeCycle: 保存周期，只能是 7， 15， 30 (Optional)
 * param inOrder: 保序 (Optional)
 */
func NewUpdateLogtopicRequestWithAllParams(
    regionId string,
    logtopicUID string,
    name *string,
    description string,
    lifeCycle *int,
    inOrder *bool,
) *UpdateLogtopicRequest {

    return &UpdateLogtopicRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/logtopics/{logtopicUID}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        LogtopicUID: logtopicUID,
        Name: name,
        Description: description,
        LifeCycle: lifeCycle,
        InOrder: inOrder,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateLogtopicRequestWithoutParam() *UpdateLogtopicRequest {

    return &UpdateLogtopicRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/logtopics/{logtopicUID}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *UpdateLogtopicRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param logtopicUID: 日志主题 UID(Required) */
func (r *UpdateLogtopicRequest) SetLogtopicUID(logtopicUID string) {
    r.LogtopicUID = logtopicUID
}
/* param name: 日志主题名称(Optional) */
func (r *UpdateLogtopicRequest) SetName(name string) {
    r.Name = &name
}
/* param description: 日志主题描述(Required) */
func (r *UpdateLogtopicRequest) SetDescription(description string) {
    r.Description = description
}
/* param lifeCycle: 保存周期，只能是 7， 15， 30(Optional) */
func (r *UpdateLogtopicRequest) SetLifeCycle(lifeCycle int) {
    r.LifeCycle = &lifeCycle
}
/* param inOrder: 保序(Optional) */
func (r *UpdateLogtopicRequest) SetInOrder(inOrder bool) {
    r.InOrder = &inOrder
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateLogtopicRequest) GetRegionId() string {
    return r.RegionId
}

type UpdateLogtopicResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateLogtopicResult `json:"result"`
}

type UpdateLogtopicResult struct {
}