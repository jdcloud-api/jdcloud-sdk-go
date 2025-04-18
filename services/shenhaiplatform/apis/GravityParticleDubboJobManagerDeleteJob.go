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

type GravityParticleDubboJobManagerDeleteJobRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 作业名 (Optional) */
    JobName *string `json:"jobName"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGravityParticleDubboJobManagerDeleteJobRequest(
    regionId string,
    appName string,
) *GravityParticleDubboJobManagerDeleteJobRequest {

	return &GravityParticleDubboJobManagerDeleteJobRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/gravityParticleDubboJobManagerDeleteJob",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        AppName: appName,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param jobName: 作业名 (Optional)
 */
func NewGravityParticleDubboJobManagerDeleteJobRequestWithAllParams(
    regionId string,
    appName string,
    jobName *string,
) *GravityParticleDubboJobManagerDeleteJobRequest {

    return &GravityParticleDubboJobManagerDeleteJobRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/gravityParticleDubboJobManagerDeleteJob",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
        JobName: jobName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGravityParticleDubboJobManagerDeleteJobRequestWithoutParam() *GravityParticleDubboJobManagerDeleteJobRequest {

    return &GravityParticleDubboJobManagerDeleteJobRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/gravityParticleDubboJobManagerDeleteJob",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *GravityParticleDubboJobManagerDeleteJobRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *GravityParticleDubboJobManagerDeleteJobRequest) SetAppName(appName string) {
    r.AppName = appName
}
/* param jobName: 作业名(Optional) */
func (r *GravityParticleDubboJobManagerDeleteJobRequest) SetJobName(jobName string) {
    r.JobName = &jobName
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GravityParticleDubboJobManagerDeleteJobRequest) GetRegionId() string {
    return r.RegionId
}

type GravityParticleDubboJobManagerDeleteJobResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GravityParticleDubboJobManagerDeleteJobResult `json:"result"`
}

type GravityParticleDubboJobManagerDeleteJobResult struct {
    Success int `json:"success"`
    Code string `json:"code"`
    Msg string `json:"msg"`
    Result bool `json:"result"`
}