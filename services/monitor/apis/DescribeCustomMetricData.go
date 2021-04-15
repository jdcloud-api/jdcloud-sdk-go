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
    monitor "github.com/jdcloud-api/jdcloud-sdk-go/services/monitor/models"
)

type DescribeCustomMetricDataRequest struct {

    core.JDCloudRequest

    /* region  */
    RegionId string `json:"regionId"`

    /* namespace  */
    NamespaceUID string `json:"namespaceUID"`

    /* 查询时间范围的结束时间， UTC时间，格式：2016-12-11T00:00:00+0800（为空时，将由startTime与timeInterval计算得出） (Optional) */
    EndTime *string `json:"endTime"`

    /*   */
    Query *monitor.QueryOption `json:"query"`

    /* 查询时间范围的开始时间， UTC时间，格式：2016-12-11T00:00:00+0800 (Optional) */
    StartTime *string `json:"startTime"`

    /* 时间间隔：1h，6h，12h，1d，3d，7d，14d，固定时间间隔，timeInterval默认为1h，当前时间往 前1h (Optional) */
    TimeInterval *string `json:"timeInterval"`
}

/*
 * param regionId: region (Required)
 * param namespaceUID: namespace (Required)
 * param query:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeCustomMetricDataRequest(
    regionId string,
    namespaceUID string,
    query *monitor.QueryOption,
) *DescribeCustomMetricDataRequest {

	return &DescribeCustomMetricDataRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/namespaces/{namespaceUID}/metricData",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        NamespaceUID: namespaceUID,
        Query: query,
	}
}

/*
 * param regionId: region (Required)
 * param namespaceUID: namespace (Required)
 * param endTime: 查询时间范围的结束时间， UTC时间，格式：2016-12-11T00:00:00+0800（为空时，将由startTime与timeInterval计算得出） (Optional)
 * param query:  (Required)
 * param startTime: 查询时间范围的开始时间， UTC时间，格式：2016-12-11T00:00:00+0800 (Optional)
 * param timeInterval: 时间间隔：1h，6h，12h，1d，3d，7d，14d，固定时间间隔，timeInterval默认为1h，当前时间往 前1h (Optional)
 */
func NewDescribeCustomMetricDataRequestWithAllParams(
    regionId string,
    namespaceUID string,
    endTime *string,
    query *monitor.QueryOption,
    startTime *string,
    timeInterval *string,
) *DescribeCustomMetricDataRequest {

    return &DescribeCustomMetricDataRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/namespaces/{namespaceUID}/metricData",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        NamespaceUID: namespaceUID,
        EndTime: endTime,
        Query: query,
        StartTime: startTime,
        TimeInterval: timeInterval,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeCustomMetricDataRequestWithoutParam() *DescribeCustomMetricDataRequest {

    return &DescribeCustomMetricDataRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/namespaces/{namespaceUID}/metricData",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: region(Required) */
func (r *DescribeCustomMetricDataRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param namespaceUID: namespace(Required) */
func (r *DescribeCustomMetricDataRequest) SetNamespaceUID(namespaceUID string) {
    r.NamespaceUID = namespaceUID
}

/* param endTime: 查询时间范围的结束时间， UTC时间，格式：2016-12-11T00:00:00+0800（为空时，将由startTime与timeInterval计算得出）(Optional) */
func (r *DescribeCustomMetricDataRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

/* param query: (Required) */
func (r *DescribeCustomMetricDataRequest) SetQuery(query *monitor.QueryOption) {
    r.Query = query
}

/* param startTime: 查询时间范围的开始时间， UTC时间，格式：2016-12-11T00:00:00+0800(Optional) */
func (r *DescribeCustomMetricDataRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param timeInterval: 时间间隔：1h，6h，12h，1d，3d，7d，14d，固定时间间隔，timeInterval默认为1h，当前时间往 前1h(Optional) */
func (r *DescribeCustomMetricDataRequest) SetTimeInterval(timeInterval string) {
    r.TimeInterval = &timeInterval
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeCustomMetricDataRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeCustomMetricDataResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeCustomMetricDataResult `json:"result"`
}

type DescribeCustomMetricDataResult struct {
    MetricDatas []monitor.MetricDataItemCm `json:"metricDatas"`
}