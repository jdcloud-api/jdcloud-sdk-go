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

type DescribePublishObjsRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 页面容量 (Optional) */
    PageSize *int `json:"pageSize"`

    /* 页号 (Optional) */
    PageNum *int `json:"pageNum"`

    /* 对象ID (Optional) */
    ObjUk *string `json:"objUk"`

    /* 对象名称 (Optional) */
    ObjName *string `json:"objName"`

    /* 对象类型编码 (Optional) */
    ObjTypeCode *string `json:"objTypeCode"`

    /* 变更类型编码 (Optional) */
    ObjChangeTypeCode *string `json:"objChangeTypeCode"`

    /* 提交人 (Optional) */
    SubmitterPin *string `json:"submitterPin"`

    /* 提交开始时间 (Optional) */
    SubmitTimeStart *int64 `json:"submitTimeStart"`

    /* 提交结束时间 (Optional) */
    SubmitTimeEnd *int64 `json:"submitTimeEnd"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribePublishObjsRequest(
    regionId string,
    appName string,
) *DescribePublishObjsRequest {

	return &DescribePublishObjsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/describePublishObjs",
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
 * param pageSize: 页面容量 (Optional)
 * param pageNum: 页号 (Optional)
 * param objUk: 对象ID (Optional)
 * param objName: 对象名称 (Optional)
 * param objTypeCode: 对象类型编码 (Optional)
 * param objChangeTypeCode: 变更类型编码 (Optional)
 * param submitterPin: 提交人 (Optional)
 * param submitTimeStart: 提交开始时间 (Optional)
 * param submitTimeEnd: 提交结束时间 (Optional)
 */
func NewDescribePublishObjsRequestWithAllParams(
    regionId string,
    appName string,
    pageSize *int,
    pageNum *int,
    objUk *string,
    objName *string,
    objTypeCode *string,
    objChangeTypeCode *string,
    submitterPin *string,
    submitTimeStart *int64,
    submitTimeEnd *int64,
) *DescribePublishObjsRequest {

    return &DescribePublishObjsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/describePublishObjs",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
        PageSize: pageSize,
        PageNum: pageNum,
        ObjUk: objUk,
        ObjName: objName,
        ObjTypeCode: objTypeCode,
        ObjChangeTypeCode: objChangeTypeCode,
        SubmitterPin: submitterPin,
        SubmitTimeStart: submitTimeStart,
        SubmitTimeEnd: submitTimeEnd,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribePublishObjsRequestWithoutParam() *DescribePublishObjsRequest {

    return &DescribePublishObjsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/describePublishObjs",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DescribePublishObjsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *DescribePublishObjsRequest) SetAppName(appName string) {
    r.AppName = appName
}
/* param pageSize: 页面容量(Optional) */
func (r *DescribePublishObjsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}
/* param pageNum: 页号(Optional) */
func (r *DescribePublishObjsRequest) SetPageNum(pageNum int) {
    r.PageNum = &pageNum
}
/* param objUk: 对象ID(Optional) */
func (r *DescribePublishObjsRequest) SetObjUk(objUk string) {
    r.ObjUk = &objUk
}
/* param objName: 对象名称(Optional) */
func (r *DescribePublishObjsRequest) SetObjName(objName string) {
    r.ObjName = &objName
}
/* param objTypeCode: 对象类型编码(Optional) */
func (r *DescribePublishObjsRequest) SetObjTypeCode(objTypeCode string) {
    r.ObjTypeCode = &objTypeCode
}
/* param objChangeTypeCode: 变更类型编码(Optional) */
func (r *DescribePublishObjsRequest) SetObjChangeTypeCode(objChangeTypeCode string) {
    r.ObjChangeTypeCode = &objChangeTypeCode
}
/* param submitterPin: 提交人(Optional) */
func (r *DescribePublishObjsRequest) SetSubmitterPin(submitterPin string) {
    r.SubmitterPin = &submitterPin
}
/* param submitTimeStart: 提交开始时间(Optional) */
func (r *DescribePublishObjsRequest) SetSubmitTimeStart(submitTimeStart int64) {
    r.SubmitTimeStart = &submitTimeStart
}
/* param submitTimeEnd: 提交结束时间(Optional) */
func (r *DescribePublishObjsRequest) SetSubmitTimeEnd(submitTimeEnd int64) {
    r.SubmitTimeEnd = &submitTimeEnd
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribePublishObjsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribePublishObjsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribePublishObjsResult `json:"result"`
}

type DescribePublishObjsResult struct {
    Code string `json:"code"`
    Message string `json:"message"`
    Result shenhaiplatform.PublishObjPageVo `json:"result"`
}