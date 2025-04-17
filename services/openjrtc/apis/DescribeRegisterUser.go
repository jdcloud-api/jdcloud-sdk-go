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

type DescribeRegisterUserRequest struct {

    core.JDCloudRequest

    /* 应用ID  */
    AppId string `json:"appId"`

    /* 业务接入方用户体系定义的且在JRTC系统内注册过的userId  */
    UserId string `json:"userId"`
}

/*
 * param appId: 应用ID (Required)
 * param userId: 业务接入方用户体系定义的且在JRTC系统内注册过的userId (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeRegisterUserRequest(
    appId string,
    userId string,
) *DescribeRegisterUserRequest {

	return &DescribeRegisterUserRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/describeRegisterUser/{appId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        AppId: appId,
        UserId: userId,
	}
}

/*
 * param appId: 应用ID (Required)
 * param userId: 业务接入方用户体系定义的且在JRTC系统内注册过的userId (Required)
 */
func NewDescribeRegisterUserRequestWithAllParams(
    appId string,
    userId string,
) *DescribeRegisterUserRequest {

    return &DescribeRegisterUserRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/describeRegisterUser/{appId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        AppId: appId,
        UserId: userId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeRegisterUserRequestWithoutParam() *DescribeRegisterUserRequest {

    return &DescribeRegisterUserRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/describeRegisterUser/{appId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param appId: 应用ID(Required) */
func (r *DescribeRegisterUserRequest) SetAppId(appId string) {
    r.AppId = appId
}
/* param userId: 业务接入方用户体系定义的且在JRTC系统内注册过的userId(Required) */
func (r *DescribeRegisterUserRequest) SetUserId(userId string) {
    r.UserId = userId
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeRegisterUserRequest) GetRegionId() string {
    return ""
}

type DescribeRegisterUserResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeRegisterUserResult `json:"result"`
}

type DescribeRegisterUserResult struct {
    AppId string `json:"appId"`
    PeerId int64 `json:"peerId"`
    UserId string `json:"userId"`
    UserName string `json:"userName"`
    Temporary bool `json:"temporary"`
    CreateTime string `json:"createTime"`
}