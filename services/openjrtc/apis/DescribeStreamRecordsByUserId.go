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
    openjrtc "github.com/jdcloud-api/jdcloud-sdk-go/services/openjrtc/models"
)

type DescribeStreamRecordsByUserIdRequest struct {

    core.JDCloudRequest

    /* 应用ID  */
    AppId string `json:"appId"`

    /* 页码；默认值为 1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认值为 10；取值范围 [10, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 业务接入方定义的且在JRTC系统内注册过的房间号  */
    UserRoomId string `json:"userRoomId"`

    /* 业务接入方用户体系定义的且在JRTC系统内注册过的userId  */
    UserId string `json:"userId"`

    /* 传参字段描述:
- kind[eq] 在线状态 1-音频流 2-视频流 100-数据流
- startTime[eq] 用户推流开始时间-UTC时间  startTime,endTime同时指定时生效
- endTime[eq]   用户推流结束时间-UTC时间  startTime,endTime同时指定时生效
 (Optional) */
    Filters []openjrtc.Filter `json:"filters"`
}

/*
 * param appId: 应用ID (Required)
 * param userRoomId: 业务接入方定义的且在JRTC系统内注册过的房间号 (Required)
 * param userId: 业务接入方用户体系定义的且在JRTC系统内注册过的userId (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeStreamRecordsByUserIdRequest(
    appId string,
    userRoomId string,
    userId string,
) *DescribeStreamRecordsByUserIdRequest {

	return &DescribeStreamRecordsByUserIdRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/describeStreamRecordsByUserId/{appId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        AppId: appId,
        UserRoomId: userRoomId,
        UserId: userId,
	}
}

/*
 * param appId: 应用ID (Required)
 * param pageNumber: 页码；默认值为 1 (Optional)
 * param pageSize: 分页大小；默认值为 10；取值范围 [10, 100] (Optional)
 * param userRoomId: 业务接入方定义的且在JRTC系统内注册过的房间号 (Required)
 * param userId: 业务接入方用户体系定义的且在JRTC系统内注册过的userId (Required)
 * param filters: 传参字段描述:
- kind[eq] 在线状态 1-音频流 2-视频流 100-数据流
- startTime[eq] 用户推流开始时间-UTC时间  startTime,endTime同时指定时生效
- endTime[eq]   用户推流结束时间-UTC时间  startTime,endTime同时指定时生效
 (Optional)
 */
func NewDescribeStreamRecordsByUserIdRequestWithAllParams(
    appId string,
    pageNumber *int,
    pageSize *int,
    userRoomId string,
    userId string,
    filters []openjrtc.Filter,
) *DescribeStreamRecordsByUserIdRequest {

    return &DescribeStreamRecordsByUserIdRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/describeStreamRecordsByUserId/{appId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        AppId: appId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        UserRoomId: userRoomId,
        UserId: userId,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeStreamRecordsByUserIdRequestWithoutParam() *DescribeStreamRecordsByUserIdRequest {

    return &DescribeStreamRecordsByUserIdRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/describeStreamRecordsByUserId/{appId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param appId: 应用ID(Required) */
func (r *DescribeStreamRecordsByUserIdRequest) SetAppId(appId string) {
    r.AppId = appId
}
/* param pageNumber: 页码；默认值为 1(Optional) */
func (r *DescribeStreamRecordsByUserIdRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}
/* param pageSize: 分页大小；默认值为 10；取值范围 [10, 100](Optional) */
func (r *DescribeStreamRecordsByUserIdRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}
/* param userRoomId: 业务接入方定义的且在JRTC系统内注册过的房间号(Required) */
func (r *DescribeStreamRecordsByUserIdRequest) SetUserRoomId(userRoomId string) {
    r.UserRoomId = userRoomId
}
/* param userId: 业务接入方用户体系定义的且在JRTC系统内注册过的userId(Required) */
func (r *DescribeStreamRecordsByUserIdRequest) SetUserId(userId string) {
    r.UserId = userId
}
/* param filters: 传参字段描述:
- kind[eq] 在线状态 1-音频流 2-视频流 100-数据流
- startTime[eq] 用户推流开始时间-UTC时间  startTime,endTime同时指定时生效
- endTime[eq]   用户推流结束时间-UTC时间  startTime,endTime同时指定时生效
(Optional) */
func (r *DescribeStreamRecordsByUserIdRequest) SetFilters(filters []openjrtc.Filter) {
    r.Filters = filters
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeStreamRecordsByUserIdRequest) GetRegionId() string {
    return ""
}

type DescribeStreamRecordsByUserIdResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeStreamRecordsByUserIdResult `json:"result"`
}

type DescribeStreamRecordsByUserIdResult struct {
    PageNumber int `json:"pageNumber"`
    PageSize int `json:"pageSize"`
    TotalElements int `json:"totalElements"`
    TotalPages int `json:"totalPages"`
    Content []openjrtc.StreamRecordInfo `json:"content"`
}