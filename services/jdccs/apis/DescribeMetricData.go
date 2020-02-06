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
    jdccs "github.com/jdcloud-api/jdcloud-sdk-go/services/jdccs/models"
)

type DescribeMetricDataRequest struct {

    core.JDCloudRequest

    /* IDC机房ID  */
    Idc string `json:"idc"`

    /* 监控项英文标识(id)  */
    Metric string `json:"metric"`

    /* 资源ID  */
    ResourceId string `json:"resourceId"`

    /* 查询时间范围的开始时间， UNIX时间戳，（机柜电流最多支持最近90天数据查询、带宽流量最多支持最近30天数据查询）  */
    StartTime int `json:"startTime"`

    /* 查询时间范围的结束时间， UNIX时间戳，（机柜电流最多支持最近90天数据查询、带宽流量最多支持最近30天数据查询）  */
    EndTime int `json:"endTime"`

    /* 时间间隔：分钟m、小时h、天d，如： 10分钟=10m、1小时=1h，3天=3d；默认5m，最小支持5m，最大90d 目前带宽上、下行流量查询，会根据时间范围是否超过2小时，设定时间间隔为1m或5m (Optional) */
    TimeInterval *string `json:"timeInterval"`

    /* 交换机IP，指定ip时须同时指定port (Optional) */
    Ip *string `json:"ip"`

    /* 端口，指定port时须同时指定ip (Optional) */
    Port *string `json:"port"`
}

/*
 * param idc: IDC机房ID (Required)
 * param metric: 监控项英文标识(id) (Required)
 * param resourceId: 资源ID (Required)
 * param startTime: 查询时间范围的开始时间， UNIX时间戳，（机柜电流最多支持最近90天数据查询、带宽流量最多支持最近30天数据查询） (Required)
 * param endTime: 查询时间范围的结束时间， UNIX时间戳，（机柜电流最多支持最近90天数据查询、带宽流量最多支持最近30天数据查询） (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeMetricDataRequest(
    idc string,
    metric string,
    resourceId string,
    startTime int,
    endTime int,
) *DescribeMetricDataRequest {

	return &DescribeMetricDataRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/idcs/{idc}/metrics/{metric}/metricData",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        Idc: idc,
        Metric: metric,
        ResourceId: resourceId,
        StartTime: startTime,
        EndTime: endTime,
	}
}

/*
 * param idc: IDC机房ID (Required)
 * param metric: 监控项英文标识(id) (Required)
 * param resourceId: 资源ID (Required)
 * param startTime: 查询时间范围的开始时间， UNIX时间戳，（机柜电流最多支持最近90天数据查询、带宽流量最多支持最近30天数据查询） (Required)
 * param endTime: 查询时间范围的结束时间， UNIX时间戳，（机柜电流最多支持最近90天数据查询、带宽流量最多支持最近30天数据查询） (Required)
 * param timeInterval: 时间间隔：分钟m、小时h、天d，如： 10分钟=10m、1小时=1h，3天=3d；默认5m，最小支持5m，最大90d 目前带宽上、下行流量查询，会根据时间范围是否超过2小时，设定时间间隔为1m或5m (Optional)
 * param ip: 交换机IP，指定ip时须同时指定port (Optional)
 * param port: 端口，指定port时须同时指定ip (Optional)
 */
func NewDescribeMetricDataRequestWithAllParams(
    idc string,
    metric string,
    resourceId string,
    startTime int,
    endTime int,
    timeInterval *string,
    ip *string,
    port *string,
) *DescribeMetricDataRequest {

    return &DescribeMetricDataRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/idcs/{idc}/metrics/{metric}/metricData",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        Idc: idc,
        Metric: metric,
        ResourceId: resourceId,
        StartTime: startTime,
        EndTime: endTime,
        TimeInterval: timeInterval,
        Ip: ip,
        Port: port,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeMetricDataRequestWithoutParam() *DescribeMetricDataRequest {

    return &DescribeMetricDataRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/idcs/{idc}/metrics/{metric}/metricData",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param idc: IDC机房ID(Required) */
func (r *DescribeMetricDataRequest) SetIdc(idc string) {
    r.Idc = idc
}

/* param metric: 监控项英文标识(id)(Required) */
func (r *DescribeMetricDataRequest) SetMetric(metric string) {
    r.Metric = metric
}

/* param resourceId: 资源ID(Required) */
func (r *DescribeMetricDataRequest) SetResourceId(resourceId string) {
    r.ResourceId = resourceId
}

/* param startTime: 查询时间范围的开始时间， UNIX时间戳，（机柜电流最多支持最近90天数据查询、带宽流量最多支持最近30天数据查询）(Required) */
func (r *DescribeMetricDataRequest) SetStartTime(startTime int) {
    r.StartTime = startTime
}

/* param endTime: 查询时间范围的结束时间， UNIX时间戳，（机柜电流最多支持最近90天数据查询、带宽流量最多支持最近30天数据查询）(Required) */
func (r *DescribeMetricDataRequest) SetEndTime(endTime int) {
    r.EndTime = endTime
}

/* param timeInterval: 时间间隔：分钟m、小时h、天d，如： 10分钟=10m、1小时=1h，3天=3d；默认5m，最小支持5m，最大90d 目前带宽上、下行流量查询，会根据时间范围是否超过2小时，设定时间间隔为1m或5m(Optional) */
func (r *DescribeMetricDataRequest) SetTimeInterval(timeInterval string) {
    r.TimeInterval = &timeInterval
}

/* param ip: 交换机IP，指定ip时须同时指定port(Optional) */
func (r *DescribeMetricDataRequest) SetIp(ip string) {
    r.Ip = &ip
}

/* param port: 端口，指定port时须同时指定ip(Optional) */
func (r *DescribeMetricDataRequest) SetPort(port string) {
    r.Port = &port
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeMetricDataRequest) GetRegionId() string {
    return ""
}

type DescribeMetricDataResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeMetricDataResult `json:"result"`
}

type DescribeMetricDataResult struct {
    MetricData jdccs.MetricData `json:"metricData"`
}