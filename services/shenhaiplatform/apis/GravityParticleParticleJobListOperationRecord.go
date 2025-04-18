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
    shenhaiplatform "github.com/jdcloud-api/jdcloud-sdk-go/services/shenhaiplatform/models"
)

type GravityParticleParticleJobListOperationRecordRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 页面大小  */
    PageSize int `json:"pageSize"`

    /* 页码  */
    PageNum int `json:"pageNum"`

    /* 作业id (Optional) */
    JobId *int `json:"jobId"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param pageSize: 页面大小 (Required)
 * param pageNum: 页码 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGravityParticleParticleJobListOperationRecordRequest(
    regionId string,
    appName string,
    pageSize int,
    pageNum int,
) *GravityParticleParticleJobListOperationRecordRequest {

	return &GravityParticleParticleJobListOperationRecordRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/gravityParticleParticleJobListOperationRecord",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        AppName: appName,
        PageSize: pageSize,
        PageNum: pageNum,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param pageSize: 页面大小 (Required)
 * param pageNum: 页码 (Required)
 * param jobId: 作业id (Optional)
 */
func NewGravityParticleParticleJobListOperationRecordRequestWithAllParams(
    regionId string,
    appName string,
    pageSize int,
    pageNum int,
    jobId *int,
) *GravityParticleParticleJobListOperationRecordRequest {

    return &GravityParticleParticleJobListOperationRecordRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/gravityParticleParticleJobListOperationRecord",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
        PageSize: pageSize,
        PageNum: pageNum,
        JobId: jobId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGravityParticleParticleJobListOperationRecordRequestWithoutParam() *GravityParticleParticleJobListOperationRecordRequest {

    return &GravityParticleParticleJobListOperationRecordRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/gravityParticleParticleJobListOperationRecord",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *GravityParticleParticleJobListOperationRecordRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *GravityParticleParticleJobListOperationRecordRequest) SetAppName(appName string) {
    r.AppName = appName
}
/* param pageSize: 页面大小(Required) */
func (r *GravityParticleParticleJobListOperationRecordRequest) SetPageSize(pageSize int) {
    r.PageSize = pageSize
}
/* param pageNum: 页码(Required) */
func (r *GravityParticleParticleJobListOperationRecordRequest) SetPageNum(pageNum int) {
    r.PageNum = pageNum
}
/* param jobId: 作业id(Optional) */
func (r *GravityParticleParticleJobListOperationRecordRequest) SetJobId(jobId int) {
    r.JobId = &jobId
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GravityParticleParticleJobListOperationRecordRequest) GetRegionId() string {
    return r.RegionId
}

type GravityParticleParticleJobListOperationRecordResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GravityParticleParticleJobListOperationRecordResult `json:"result"`
}

type GravityParticleParticleJobListOperationRecordResult struct {
    Success int `json:"success"`
    Result shenhaiplatform.GpjmPageInfoMgrOperationRecords `json:"result"`
    Code string `json:"code"`
    Msg string `json:"msg"`
    _REQ_SEQ_NO_ string `json:"_REQ_SEQ_NO_"`
}