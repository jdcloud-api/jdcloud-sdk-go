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

type SpectrumAppBandwidthDateHistogramRequest struct {

    core.JDCloudRequest

    /* 实例标识  */
    InstanceId string `json:"instanceId"`

    /* 域名标识  */
    ZoneId string `json:"zoneId"`

    /* 应用标识  */
    SpectrumAppId string `json:"spectrumAppId"`

    /* 开始时间  */
    Since string `json:"since"`

    /* 结束时间  */
    Until string `json:"until"`
}

/*
 * param instanceId: 实例标识 (Required)
 * param zoneId: 域名标识 (Required)
 * param spectrumAppId: 应用标识 (Required)
 * param since: 开始时间 (Required)
 * param until: 结束时间 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSpectrumAppBandwidthDateHistogramRequest(
    instanceId string,
    zoneId string,
    spectrumAppId string,
    since string,
    until string,
) *SpectrumAppBandwidthDateHistogramRequest {

	return &SpectrumAppBandwidthDateHistogramRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/instances/{instanceId}/zones/{zoneId}/spectrumApps/{spectrumAppId}/paBandwidthDateHistogram",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        InstanceId: instanceId,
        ZoneId: zoneId,
        SpectrumAppId: spectrumAppId,
        Since: since,
        Until: until,
	}
}

/*
 * param instanceId: 实例标识 (Required)
 * param zoneId: 域名标识 (Required)
 * param spectrumAppId: 应用标识 (Required)
 * param since: 开始时间 (Required)
 * param until: 结束时间 (Required)
 */
func NewSpectrumAppBandwidthDateHistogramRequestWithAllParams(
    instanceId string,
    zoneId string,
    spectrumAppId string,
    since string,
    until string,
) *SpectrumAppBandwidthDateHistogramRequest {

    return &SpectrumAppBandwidthDateHistogramRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/instances/{instanceId}/zones/{zoneId}/spectrumApps/{spectrumAppId}/paBandwidthDateHistogram",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        InstanceId: instanceId,
        ZoneId: zoneId,
        SpectrumAppId: spectrumAppId,
        Since: since,
        Until: until,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSpectrumAppBandwidthDateHistogramRequestWithoutParam() *SpectrumAppBandwidthDateHistogramRequest {

    return &SpectrumAppBandwidthDateHistogramRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/instances/{instanceId}/zones/{zoneId}/spectrumApps/{spectrumAppId}/paBandwidthDateHistogram",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param instanceId: 实例标识(Required) */
func (r *SpectrumAppBandwidthDateHistogramRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}
/* param zoneId: 域名标识(Required) */
func (r *SpectrumAppBandwidthDateHistogramRequest) SetZoneId(zoneId string) {
    r.ZoneId = zoneId
}
/* param spectrumAppId: 应用标识(Required) */
func (r *SpectrumAppBandwidthDateHistogramRequest) SetSpectrumAppId(spectrumAppId string) {
    r.SpectrumAppId = spectrumAppId
}
/* param since: 开始时间(Required) */
func (r *SpectrumAppBandwidthDateHistogramRequest) SetSince(since string) {
    r.Since = since
}
/* param until: 结束时间(Required) */
func (r *SpectrumAppBandwidthDateHistogramRequest) SetUntil(until string) {
    r.Until = until
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SpectrumAppBandwidthDateHistogramRequest) GetRegionId() string {
    return ""
}

type SpectrumAppBandwidthDateHistogramResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SpectrumAppBandwidthDateHistogramResult `json:"result"`
}

type SpectrumAppBandwidthDateHistogramResult struct {
    Sum float64 `json:"sum"`
    Max float64 `json:"max"`
    MaxTimestamp int `json:"maxTimestamp"`
    DataSeries []float64 `json:"dataSeries"`
    TimeSeries []int `json:"timeSeries"`
}