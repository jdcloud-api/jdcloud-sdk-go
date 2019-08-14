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

type DeleteMonitorRequest struct {

    core.JDCloudRequest

    /* 实例所属的地域ID  */
    RegionId string `json:"regionId"`

    /* 域名ID，请使用describeDomains接口获取。  */
    DomainId string `json:"domainId"`

    /* 监控项ID，请使用describeMonitor接口获取。  */
    MonitorId string `json:"monitorId"`
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param domainId: 域名ID，请使用describeDomains接口获取。 (Required)
 * param monitorId: 监控项ID，请使用describeMonitor接口获取。 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteMonitorRequest(
    regionId string,
    domainId string,
    monitorId string,
) *DeleteMonitorRequest {

	return &DeleteMonitorRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/domain/{domainId}/monitor/{monitorId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        DomainId: domainId,
        MonitorId: monitorId,
	}
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param domainId: 域名ID，请使用describeDomains接口获取。 (Required)
 * param monitorId: 监控项ID，请使用describeMonitor接口获取。 (Required)
 */
func NewDeleteMonitorRequestWithAllParams(
    regionId string,
    domainId string,
    monitorId string,
) *DeleteMonitorRequest {

    return &DeleteMonitorRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain/{domainId}/monitor/{monitorId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        DomainId: domainId,
        MonitorId: monitorId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteMonitorRequestWithoutParam() *DeleteMonitorRequest {

    return &DeleteMonitorRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain/{domainId}/monitor/{monitorId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 实例所属的地域ID(Required) */
func (r *DeleteMonitorRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param domainId: 域名ID，请使用describeDomains接口获取。(Required) */
func (r *DeleteMonitorRequest) SetDomainId(domainId string) {
    r.DomainId = domainId
}

/* param monitorId: 监控项ID，请使用describeMonitor接口获取。(Required) */
func (r *DeleteMonitorRequest) SetMonitorId(monitorId string) {
    r.MonitorId = monitorId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteMonitorRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteMonitorResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteMonitorResult `json:"result"`
}

type DeleteMonitorResult struct {
}