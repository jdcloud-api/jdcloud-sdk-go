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
    cdn "github.com/jdcloud-api/jdcloud-sdk-go/services/cdn/models"
)

type QueryMixTrafficGroupSumRequest struct {

    core.JDCloudRequest

    /* 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional) */
    StartTime *string `json:"startTime"`

    /* 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional) */
    EndTime *string `json:"endTime"`

    /* 需要查询的域名, 必须为用户pin下有权限的域名 (Optional) */
    Domain *string `json:"domain"`

    /* 需要查询的字段 (Optional) */
    Fields *string `json:"fields"`

    /*  (Optional) */
    Area *string `json:"area"`

    /*  (Optional) */
    Isp *string `json:"isp"`

    /* 时间粒度，可选值:[oneMin,fiveMin,followTime],followTime只会返回一个汇总后的数据 (Optional) */
    Period *string `json:"period"`

    /* 分组依据,只能按域名分组 (Optional) */
    GroupBy *string `json:"groupBy"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryMixTrafficGroupSumRequest(
) *QueryMixTrafficGroupSumRequest {

	return &QueryMixTrafficGroupSumRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/statistics:groupSum",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
 * param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
 * param domain: 需要查询的域名, 必须为用户pin下有权限的域名 (Optional)
 * param fields: 需要查询的字段 (Optional)
 * param area:  (Optional)
 * param isp:  (Optional)
 * param period: 时间粒度，可选值:[oneMin,fiveMin,followTime],followTime只会返回一个汇总后的数据 (Optional)
 * param groupBy: 分组依据,只能按域名分组 (Optional)
 */
func NewQueryMixTrafficGroupSumRequestWithAllParams(
    startTime *string,
    endTime *string,
    domain *string,
    fields *string,
    area *string,
    isp *string,
    period *string,
    groupBy *string,
) *QueryMixTrafficGroupSumRequest {

    return &QueryMixTrafficGroupSumRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/statistics:groupSum",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        StartTime: startTime,
        EndTime: endTime,
        Domain: domain,
        Fields: fields,
        Area: area,
        Isp: isp,
        Period: period,
        GroupBy: groupBy,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryMixTrafficGroupSumRequestWithoutParam() *QueryMixTrafficGroupSumRequest {

    return &QueryMixTrafficGroupSumRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/statistics:groupSum",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z(Optional) */
func (r *QueryMixTrafficGroupSumRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z(Optional) */
func (r *QueryMixTrafficGroupSumRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

/* param domain: 需要查询的域名, 必须为用户pin下有权限的域名(Optional) */
func (r *QueryMixTrafficGroupSumRequest) SetDomain(domain string) {
    r.Domain = &domain
}

/* param fields: 需要查询的字段(Optional) */
func (r *QueryMixTrafficGroupSumRequest) SetFields(fields string) {
    r.Fields = &fields
}

/* param area: (Optional) */
func (r *QueryMixTrafficGroupSumRequest) SetArea(area string) {
    r.Area = &area
}

/* param isp: (Optional) */
func (r *QueryMixTrafficGroupSumRequest) SetIsp(isp string) {
    r.Isp = &isp
}

/* param period: 时间粒度，可选值:[oneMin,fiveMin,followTime],followTime只会返回一个汇总后的数据(Optional) */
func (r *QueryMixTrafficGroupSumRequest) SetPeriod(period string) {
    r.Period = &period
}

/* param groupBy: 分组依据,只能按域名分组(Optional) */
func (r *QueryMixTrafficGroupSumRequest) SetGroupBy(groupBy string) {
    r.GroupBy = &groupBy
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryMixTrafficGroupSumRequest) GetRegionId() string {
    return ""
}

type QueryMixTrafficGroupSumResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryMixTrafficGroupSumResult `json:"result"`
}

type QueryMixTrafficGroupSumResult struct {
    StartTime string `json:"startTime"`
    EndTime string `json:"endTime"`
    Domain string `json:"domain"`
    Statistics []cdn.StatisticsGroupSumDataItem `json:"statistics"`
}