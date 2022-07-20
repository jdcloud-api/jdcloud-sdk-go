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
    "github.com/lshuining/jdcloud-sdk-go/core"
    live "github.com/lshuining/jdcloud-sdk-go/services/live/models"
)

type DescribeLiveStatisticGroupByAreaIspRequest struct {

    core.JDCloudRequest

    /* 播放域名  */
    DomainName string `json:"domainName"`

    /* 应用名称 (Optional) */
    AppName *string `json:"appName"`

    /* 流名称 (Optional) */
    StreamName *string `json:"streamName"`

    /* 运营商
 (Optional) */
    IspName *string `json:"ispName"`

    /* 查询的区域，如beijing,shanghai。多个用逗号分隔
 (Optional) */
    LocationName *string `json:"locationName"`

    /* 查询的流协议类型，取值范围："rtmp,hdl,hls"，多个时以逗号分隔
 (Optional) */
    ProtocolType *string `json:"protocolType"`

    /* 查询周期，当前取值范围：“oneMin,fiveMin,halfHour,hour,twoHour,sixHour,day,followTime”，分别表示1min，5min，半小时，1小时，2小时，6小时，1天，跟随时间。默认为空，表示fiveMin。当传入followTime时，表示按Endtime-StartTime的周期，只返回一个点
 (Optional) */
    Period *string `json:"period"`

    /* 起始时间
- UTC时间
  格式:yyyy-MM-dd'T'HH:mm:ss'Z'
  示例:2018-10-21T10:00:00Z
  */
    StartTime string `json:"startTime"`

    /* 结束时间:
- UTC时间
  格式:yyyy-MM-dd'T'HH:mm:ss'Z'
  示例:2018-10-21T10:00:00Z
- 为空,默认为当前时间，查询时间跨度不超过1天
 (Optional) */
    EndTime *string `json:"endTime"`
}

/*
 * param domainName: 播放域名 (Required)
 * param startTime: 起始时间
- UTC时间
  格式:yyyy-MM-dd'T'HH:mm:ss'Z'
  示例:2018-10-21T10:00:00Z
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeLiveStatisticGroupByAreaIspRequest(
    domainName string,
    startTime string,
) *DescribeLiveStatisticGroupByAreaIspRequest {

	return &DescribeLiveStatisticGroupByAreaIspRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/describeLiveStatisticGroupByAreaIsp",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        DomainName: domainName,
        StartTime: startTime,
	}
}

/*
 * param domainName: 播放域名 (Required)
 * param appName: 应用名称 (Optional)
 * param streamName: 流名称 (Optional)
 * param ispName: 运营商
 (Optional)
 * param locationName: 查询的区域，如beijing,shanghai。多个用逗号分隔
 (Optional)
 * param protocolType: 查询的流协议类型，取值范围："rtmp,hdl,hls"，多个时以逗号分隔
 (Optional)
 * param period: 查询周期，当前取值范围：“oneMin,fiveMin,halfHour,hour,twoHour,sixHour,day,followTime”，分别表示1min，5min，半小时，1小时，2小时，6小时，1天，跟随时间。默认为空，表示fiveMin。当传入followTime时，表示按Endtime-StartTime的周期，只返回一个点
 (Optional)
 * param startTime: 起始时间
- UTC时间
  格式:yyyy-MM-dd'T'HH:mm:ss'Z'
  示例:2018-10-21T10:00:00Z
 (Required)
 * param endTime: 结束时间:
- UTC时间
  格式:yyyy-MM-dd'T'HH:mm:ss'Z'
  示例:2018-10-21T10:00:00Z
- 为空,默认为当前时间，查询时间跨度不超过1天
 (Optional)
 */
func NewDescribeLiveStatisticGroupByAreaIspRequestWithAllParams(
    domainName string,
    appName *string,
    streamName *string,
    ispName *string,
    locationName *string,
    protocolType *string,
    period *string,
    startTime string,
    endTime *string,
) *DescribeLiveStatisticGroupByAreaIspRequest {

    return &DescribeLiveStatisticGroupByAreaIspRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/describeLiveStatisticGroupByAreaIsp",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        DomainName: domainName,
        AppName: appName,
        StreamName: streamName,
        IspName: ispName,
        LocationName: locationName,
        ProtocolType: protocolType,
        Period: period,
        StartTime: startTime,
        EndTime: endTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeLiveStatisticGroupByAreaIspRequestWithoutParam() *DescribeLiveStatisticGroupByAreaIspRequest {

    return &DescribeLiveStatisticGroupByAreaIspRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/describeLiveStatisticGroupByAreaIsp",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domainName: 播放域名(Required) */
func (r *DescribeLiveStatisticGroupByAreaIspRequest) SetDomainName(domainName string) {
    r.DomainName = domainName
}

/* param appName: 应用名称(Optional) */
func (r *DescribeLiveStatisticGroupByAreaIspRequest) SetAppName(appName string) {
    r.AppName = &appName
}

/* param streamName: 流名称(Optional) */
func (r *DescribeLiveStatisticGroupByAreaIspRequest) SetStreamName(streamName string) {
    r.StreamName = &streamName
}

/* param ispName: 运营商
(Optional) */
func (r *DescribeLiveStatisticGroupByAreaIspRequest) SetIspName(ispName string) {
    r.IspName = &ispName
}

/* param locationName: 查询的区域，如beijing,shanghai。多个用逗号分隔
(Optional) */
func (r *DescribeLiveStatisticGroupByAreaIspRequest) SetLocationName(locationName string) {
    r.LocationName = &locationName
}

/* param protocolType: 查询的流协议类型，取值范围："rtmp,hdl,hls"，多个时以逗号分隔
(Optional) */
func (r *DescribeLiveStatisticGroupByAreaIspRequest) SetProtocolType(protocolType string) {
    r.ProtocolType = &protocolType
}

/* param period: 查询周期，当前取值范围：“oneMin,fiveMin,halfHour,hour,twoHour,sixHour,day,followTime”，分别表示1min，5min，半小时，1小时，2小时，6小时，1天，跟随时间。默认为空，表示fiveMin。当传入followTime时，表示按Endtime-StartTime的周期，只返回一个点
(Optional) */
func (r *DescribeLiveStatisticGroupByAreaIspRequest) SetPeriod(period string) {
    r.Period = &period
}

/* param startTime: 起始时间
- UTC时间
  格式:yyyy-MM-dd'T'HH:mm:ss'Z'
  示例:2018-10-21T10:00:00Z
(Required) */
func (r *DescribeLiveStatisticGroupByAreaIspRequest) SetStartTime(startTime string) {
    r.StartTime = startTime
}

/* param endTime: 结束时间:
- UTC时间
  格式:yyyy-MM-dd'T'HH:mm:ss'Z'
  示例:2018-10-21T10:00:00Z
- 为空,默认为当前时间，查询时间跨度不超过1天
(Optional) */
func (r *DescribeLiveStatisticGroupByAreaIspRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeLiveStatisticGroupByAreaIspRequest) GetRegionId() string {
    return ""
}

type DescribeLiveStatisticGroupByAreaIspResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeLiveStatisticGroupByAreaIspResult `json:"result"`
}

type DescribeLiveStatisticGroupByAreaIspResult struct {
    DataList []live.LiveStatisticGroupByAreaIspResult `json:"dataList"`
}