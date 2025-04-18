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

type WorkspaceVerifyAuthRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /*  (Optional) */
    PageNum *int `json:"pageNum"`

    /*  (Optional) */
    PageSize *int `json:"pageSize"`

    /*  (Optional) */
    Id *int64 `json:"id"`

    /*  (Optional) */
    UserName *string `json:"userName"`

    /*  (Optional) */
    NickName *string `json:"nickName"`

    /*  (Optional) */
    UserPin *string `json:"userPin"`

    /*  (Optional) */
    WorkspaceCode *string `json:"workspaceCode"`

    /*  (Optional) */
    CompanyCode *string `json:"companyCode"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewWorkspaceVerifyAuthRequest(
    regionId string,
    appName string,
) *WorkspaceVerifyAuthRequest {

	return &WorkspaceVerifyAuthRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/workspaceVerifyAuth",
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
 * param pageNum:  (Optional)
 * param pageSize:  (Optional)
 * param id:  (Optional)
 * param userName:  (Optional)
 * param nickName:  (Optional)
 * param userPin:  (Optional)
 * param workspaceCode:  (Optional)
 * param companyCode:  (Optional)
 */
func NewWorkspaceVerifyAuthRequestWithAllParams(
    regionId string,
    appName string,
    pageNum *int,
    pageSize *int,
    id *int64,
    userName *string,
    nickName *string,
    userPin *string,
    workspaceCode *string,
    companyCode *string,
) *WorkspaceVerifyAuthRequest {

    return &WorkspaceVerifyAuthRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/workspaceVerifyAuth",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
        PageNum: pageNum,
        PageSize: pageSize,
        Id: id,
        UserName: userName,
        NickName: nickName,
        UserPin: userPin,
        WorkspaceCode: workspaceCode,
        CompanyCode: companyCode,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewWorkspaceVerifyAuthRequestWithoutParam() *WorkspaceVerifyAuthRequest {

    return &WorkspaceVerifyAuthRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/workspaceVerifyAuth",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *WorkspaceVerifyAuthRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *WorkspaceVerifyAuthRequest) SetAppName(appName string) {
    r.AppName = appName
}
/* param pageNum: (Optional) */
func (r *WorkspaceVerifyAuthRequest) SetPageNum(pageNum int) {
    r.PageNum = &pageNum
}
/* param pageSize: (Optional) */
func (r *WorkspaceVerifyAuthRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}
/* param id: (Optional) */
func (r *WorkspaceVerifyAuthRequest) SetId(id int64) {
    r.Id = &id
}
/* param userName: (Optional) */
func (r *WorkspaceVerifyAuthRequest) SetUserName(userName string) {
    r.UserName = &userName
}
/* param nickName: (Optional) */
func (r *WorkspaceVerifyAuthRequest) SetNickName(nickName string) {
    r.NickName = &nickName
}
/* param userPin: (Optional) */
func (r *WorkspaceVerifyAuthRequest) SetUserPin(userPin string) {
    r.UserPin = &userPin
}
/* param workspaceCode: (Optional) */
func (r *WorkspaceVerifyAuthRequest) SetWorkspaceCode(workspaceCode string) {
    r.WorkspaceCode = &workspaceCode
}
/* param companyCode: (Optional) */
func (r *WorkspaceVerifyAuthRequest) SetCompanyCode(companyCode string) {
    r.CompanyCode = &companyCode
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r WorkspaceVerifyAuthRequest) GetRegionId() string {
    return r.RegionId
}

type WorkspaceVerifyAuthResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result WorkspaceVerifyAuthResult `json:"result"`
}

type WorkspaceVerifyAuthResult struct {
    Code string `json:"code"`
    Message string `json:"message"`
    Result bool `json:"result"`
}