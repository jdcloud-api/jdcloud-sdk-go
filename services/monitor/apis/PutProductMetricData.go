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
    monitor "github.com/lshuining/jdcloud-sdk-go/services/monitor/models"
)

type PutProductMetricDataRequest struct {

    core.JDCloudRequest

    /* 目前统一用jcloud  */
    AppCode string `json:"appCode"`

    /* 资源的类型，取值vm,ip,database,storage,disk,cdn,redis,balance,nat_gw,db_ro,vpn,ddos等,新接入的产品要求与opentapi命名的产品线名称一致  */
    ServiceCode string `json:"serviceCode"`

    /* 地域信息，如 cn-north-1 等  */
    Region string `json:"region"`

    /* 资源的唯一表示，一般为uuid  */
    ResourceId string `json:"resourceId"`

    /* 监控数据点  */
    DataPoints []monitor.DataPointX `json:"dataPoints"`
}

/*
 * param appCode: 目前统一用jcloud (Required)
 * param serviceCode: 资源的类型，取值vm,ip,database,storage,disk,cdn,redis,balance,nat_gw,db_ro,vpn,ddos等,新接入的产品要求与opentapi命名的产品线名称一致 (Required)
 * param region: 地域信息，如 cn-north-1 等 (Required)
 * param resourceId: 资源的唯一表示，一般为uuid (Required)
 * param dataPoints: 监控数据点 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewPutProductMetricDataRequest(
    appCode string,
    serviceCode string,
    region string,
    resourceId string,
    dataPoints []monitor.DataPointX,
) *PutProductMetricDataRequest {

	return &PutProductMetricDataRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/put",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        AppCode: appCode,
        ServiceCode: serviceCode,
        Region: region,
        ResourceId: resourceId,
        DataPoints: dataPoints,
	}
}

/*
 * param appCode: 目前统一用jcloud (Required)
 * param serviceCode: 资源的类型，取值vm,ip,database,storage,disk,cdn,redis,balance,nat_gw,db_ro,vpn,ddos等,新接入的产品要求与opentapi命名的产品线名称一致 (Required)
 * param region: 地域信息，如 cn-north-1 等 (Required)
 * param resourceId: 资源的唯一表示，一般为uuid (Required)
 * param dataPoints: 监控数据点 (Required)
 */
func NewPutProductMetricDataRequestWithAllParams(
    appCode string,
    serviceCode string,
    region string,
    resourceId string,
    dataPoints []monitor.DataPointX,
) *PutProductMetricDataRequest {

    return &PutProductMetricDataRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/put",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        AppCode: appCode,
        ServiceCode: serviceCode,
        Region: region,
        ResourceId: resourceId,
        DataPoints: dataPoints,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewPutProductMetricDataRequestWithoutParam() *PutProductMetricDataRequest {

    return &PutProductMetricDataRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/put",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param appCode: 目前统一用jcloud(Required) */
func (r *PutProductMetricDataRequest) SetAppCode(appCode string) {
    r.AppCode = appCode
}

/* param serviceCode: 资源的类型，取值vm,ip,database,storage,disk,cdn,redis,balance,nat_gw,db_ro,vpn,ddos等,新接入的产品要求与opentapi命名的产品线名称一致(Required) */
func (r *PutProductMetricDataRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = serviceCode
}

/* param region: 地域信息，如 cn-north-1 等(Required) */
func (r *PutProductMetricDataRequest) SetRegion(region string) {
    r.Region = region
}

/* param resourceId: 资源的唯一表示，一般为uuid(Required) */
func (r *PutProductMetricDataRequest) SetResourceId(resourceId string) {
    r.ResourceId = resourceId
}

/* param dataPoints: 监控数据点(Required) */
func (r *PutProductMetricDataRequest) SetDataPoints(dataPoints []monitor.DataPointX) {
    r.DataPoints = dataPoints
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r PutProductMetricDataRequest) GetRegionId() string {
    return ""
}

type PutProductMetricDataResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result PutProductMetricDataResult `json:"result"`
}

type PutProductMetricDataResult struct {
    Failed int `json:"failed"`
    Success int `json:"success"`
}