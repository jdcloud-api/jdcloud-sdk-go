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

type DescribeMetricDataRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 监控项英文标识(id)  */
    Metric string `json:"metric"`

    /* 聚合方式，默认等于downSampleType或avg，可选值参考:sum、avg、last、min、max (Optional) */
    AggrType *string `json:"aggrType"`

    /* 采样方式，默认等于aggrType或avg，可选值参考：sum、avg、last、min、max (Optional) */
    DownSampleType *string `json:"downSampleType"`

    /* 查询时间范围的开始时间， UTC时间，格式：yyyy-MM-dd'T'HH:mm:ssZ (Optional) */
    StartTime *string `json:"startTime"`

    /* 查询时间范围的结束时间， UTC时间，格式：2016-12- yyyy-MM-dd'T'HH:mm:ssZ（为空时，将由startTime与timeInterval计算得出） (Optional) */
    EndTime *string `json:"endTime"`

    /* 时间间隔：1h，6h，12h，1d，3d，7d，14d，固定时间间隔，timeInterval默认为1h，当前时间往 前1h (Optional) */
    TimeInterval *string `json:"timeInterval"`

    /* 自定义标签/tag；至少要传一个tag，且tag.Values不为空 (Optional) */
    Tags []monitor.TagFilter `json:"tags"`

    /* 是否对查询的tags分组 (Optional) */
    GroupBy *bool `json:"groupBy"`

    /* 是否求速率 (Optional) */
    Rate *bool `json:"rate"`

    /* 资源的类型，取值vm, lb, ip, database 等  */
    ServiceCode string `json:"serviceCode"`

    /* 资源的uuid  */
    ResourceId string `json:"resourceId"`
}

/*
 * param regionId: 地域 Id (Required)
 * param metric: 监控项英文标识(id) (Required)
 * param serviceCode: 资源的类型，取值vm, lb, ip, database 等 (Required)
 * param resourceId: 资源的uuid (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeMetricDataRequest(
    regionId string,
    metric string,
    serviceCode string,
    resourceId string,
) *DescribeMetricDataRequest {

	return &DescribeMetricDataRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/metrics/{metric}/metricData",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Metric: metric,
        ServiceCode: serviceCode,
        ResourceId: resourceId,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param metric: 监控项英文标识(id) (Required)
 * param aggrType: 聚合方式，默认等于downSampleType或avg，可选值参考:sum、avg、last、min、max (Optional)
 * param downSampleType: 采样方式，默认等于aggrType或avg，可选值参考：sum、avg、last、min、max (Optional)
 * param startTime: 查询时间范围的开始时间， UTC时间，格式：yyyy-MM-dd'T'HH:mm:ssZ (Optional)
 * param endTime: 查询时间范围的结束时间， UTC时间，格式：2016-12- yyyy-MM-dd'T'HH:mm:ssZ（为空时，将由startTime与timeInterval计算得出） (Optional)
 * param timeInterval: 时间间隔：1h，6h，12h，1d，3d，7d，14d，固定时间间隔，timeInterval默认为1h，当前时间往 前1h (Optional)
 * param tags: 自定义标签/tag；至少要传一个tag，且tag.Values不为空 (Optional)
 * param groupBy: 是否对查询的tags分组 (Optional)
 * param rate: 是否求速率 (Optional)
 * param serviceCode: 资源的类型，取值vm, lb, ip, database 等 (Required)
 * param resourceId: 资源的uuid (Required)
 */
func NewDescribeMetricDataRequestWithAllParams(
    regionId string,
    metric string,
    aggrType *string,
    downSampleType *string,
    startTime *string,
    endTime *string,
    timeInterval *string,
    tags []monitor.TagFilter,
    groupBy *bool,
    rate *bool,
    serviceCode string,
    resourceId string,
) *DescribeMetricDataRequest {

    return &DescribeMetricDataRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/metrics/{metric}/metricData",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Metric: metric,
        AggrType: aggrType,
        DownSampleType: downSampleType,
        StartTime: startTime,
        EndTime: endTime,
        TimeInterval: timeInterval,
        Tags: tags,
        GroupBy: groupBy,
        Rate: rate,
        ServiceCode: serviceCode,
        ResourceId: resourceId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeMetricDataRequestWithoutParam() *DescribeMetricDataRequest {

    return &DescribeMetricDataRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/metrics/{metric}/metricData",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *DescribeMetricDataRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param metric: 监控项英文标识(id)(Required) */
func (r *DescribeMetricDataRequest) SetMetric(metric string) {
    r.Metric = metric
}

/* param aggrType: 聚合方式，默认等于downSampleType或avg，可选值参考:sum、avg、last、min、max(Optional) */
func (r *DescribeMetricDataRequest) SetAggrType(aggrType string) {
    r.AggrType = &aggrType
}

/* param downSampleType: 采样方式，默认等于aggrType或avg，可选值参考：sum、avg、last、min、max(Optional) */
func (r *DescribeMetricDataRequest) SetDownSampleType(downSampleType string) {
    r.DownSampleType = &downSampleType
}

/* param startTime: 查询时间范围的开始时间， UTC时间，格式：yyyy-MM-dd'T'HH:mm:ssZ(Optional) */
func (r *DescribeMetricDataRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param endTime: 查询时间范围的结束时间， UTC时间，格式：2016-12- yyyy-MM-dd'T'HH:mm:ssZ（为空时，将由startTime与timeInterval计算得出）(Optional) */
func (r *DescribeMetricDataRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

/* param timeInterval: 时间间隔：1h，6h，12h，1d，3d，7d，14d，固定时间间隔，timeInterval默认为1h，当前时间往 前1h(Optional) */
func (r *DescribeMetricDataRequest) SetTimeInterval(timeInterval string) {
    r.TimeInterval = &timeInterval
}

/* param tags: 自定义标签/tag；至少要传一个tag，且tag.Values不为空(Optional) */
func (r *DescribeMetricDataRequest) SetTags(tags []monitor.TagFilter) {
    r.Tags = tags
}

/* param groupBy: 是否对查询的tags分组(Optional) */
func (r *DescribeMetricDataRequest) SetGroupBy(groupBy bool) {
    r.GroupBy = &groupBy
}

/* param rate: 是否求速率(Optional) */
func (r *DescribeMetricDataRequest) SetRate(rate bool) {
    r.Rate = &rate
}

/* param serviceCode: 资源的类型，取值vm, lb, ip, database 等(Required) */
func (r *DescribeMetricDataRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = serviceCode
}

/* param resourceId: 资源的uuid(Required) */
func (r *DescribeMetricDataRequest) SetResourceId(resourceId string) {
    r.ResourceId = resourceId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeMetricDataRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeMetricDataResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeMetricDataResult `json:"result"`
}

type DescribeMetricDataResult struct {
    MetricDatas []monitor.MetricData `json:"metricDatas"`
}