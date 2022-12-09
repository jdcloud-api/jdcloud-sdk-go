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
    dms "github.com/jdcloud-api/jdcloud-sdk-go/services/dms/models"
)

type AuthPrivilegeRequest struct {

    core.JDCloudRequest

    /* 用户pin列表信息 (Optional) */
    PinList []string `json:"pinList"`

    /* 权限名称,枚举值：PrivilegeLogin(实例登录权限) (Optional) */
    PrivilegeName *string `json:"privilegeName"`

    /* 授权过期时间(yyyy-MM-dd'T'HH:mm:ss.SSS'Z') (Optional) */
    ExpireDate *string `json:"expireDate"`

    /* 授权实例的信息，主要包括用户实例ID和实例名称 (Optional) */
    DmsPrivilegeInstanceParamList []dms.DmsPrivilegeInstanceParam `json:"dmsPrivilegeInstanceParamList"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAuthPrivilegeRequest(
) *AuthPrivilegeRequest {

	return &AuthPrivilegeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/management:authPrivilege",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param pinList: 用户pin列表信息 (Optional)
 * param privilegeName: 权限名称,枚举值：PrivilegeLogin(实例登录权限) (Optional)
 * param expireDate: 授权过期时间(yyyy-MM-dd'T'HH:mm:ss.SSS'Z') (Optional)
 * param dmsPrivilegeInstanceParamList: 授权实例的信息，主要包括用户实例ID和实例名称 (Optional)
 */
func NewAuthPrivilegeRequestWithAllParams(
    pinList []string,
    privilegeName *string,
    expireDate *string,
    dmsPrivilegeInstanceParamList []dms.DmsPrivilegeInstanceParam,
) *AuthPrivilegeRequest {

    return &AuthPrivilegeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/management:authPrivilege",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        PinList: pinList,
        PrivilegeName: privilegeName,
        ExpireDate: expireDate,
        DmsPrivilegeInstanceParamList: dmsPrivilegeInstanceParamList,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAuthPrivilegeRequestWithoutParam() *AuthPrivilegeRequest {

    return &AuthPrivilegeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/management:authPrivilege",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param pinList: 用户pin列表信息(Optional) */
func (r *AuthPrivilegeRequest) SetPinList(pinList []string) {
    r.PinList = pinList
}
/* param privilegeName: 权限名称,枚举值：PrivilegeLogin(实例登录权限)(Optional) */
func (r *AuthPrivilegeRequest) SetPrivilegeName(privilegeName string) {
    r.PrivilegeName = &privilegeName
}
/* param expireDate: 授权过期时间(yyyy-MM-dd'T'HH:mm:ss.SSS'Z')(Optional) */
func (r *AuthPrivilegeRequest) SetExpireDate(expireDate string) {
    r.ExpireDate = &expireDate
}
/* param dmsPrivilegeInstanceParamList: 授权实例的信息，主要包括用户实例ID和实例名称(Optional) */
func (r *AuthPrivilegeRequest) SetDmsPrivilegeInstanceParamList(dmsPrivilegeInstanceParamList []dms.DmsPrivilegeInstanceParam) {
    r.DmsPrivilegeInstanceParamList = dmsPrivilegeInstanceParamList
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AuthPrivilegeRequest) GetRegionId() string {
    return ""
}

type AuthPrivilegeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AuthPrivilegeResult `json:"result"`
}

type AuthPrivilegeResult struct {
}