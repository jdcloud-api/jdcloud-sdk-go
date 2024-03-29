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
    fc "github.com/jdcloud-api/jdcloud-sdk-go/services/fc/models"
)

type UpdateServiceRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Service Name  */
    ServiceName string `json:"serviceName"`

    /* 修改后的描述 (Optional) */
    Description *string `json:"description"`

    /* 修改后的日志配置，不修改的话传空值，关闭日志集logTopicUid传空值。 (Optional) */
    LogConfig *fc.LogConfig `json:"logConfig"`
}

/*
 * param regionId: Region ID (Required)
 * param serviceName: Service Name (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateServiceRequest(
    regionId string,
    serviceName string,
) *UpdateServiceRequest {

	return &UpdateServiceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/services/{serviceName}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ServiceName: serviceName,
	}
}

/*
 * param regionId: Region ID (Required)
 * param serviceName: Service Name (Required)
 * param description: 修改后的描述 (Optional)
 * param logConfig: 修改后的日志配置，不修改的话传空值，关闭日志集logTopicUid传空值。 (Optional)
 */
func NewUpdateServiceRequestWithAllParams(
    regionId string,
    serviceName string,
    description *string,
    logConfig *fc.LogConfig,
) *UpdateServiceRequest {

    return &UpdateServiceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/services/{serviceName}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ServiceName: serviceName,
        Description: description,
        LogConfig: logConfig,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateServiceRequestWithoutParam() *UpdateServiceRequest {

    return &UpdateServiceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/services/{serviceName}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *UpdateServiceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param serviceName: Service Name(Required) */
func (r *UpdateServiceRequest) SetServiceName(serviceName string) {
    r.ServiceName = serviceName
}
/* param description: 修改后的描述(Optional) */
func (r *UpdateServiceRequest) SetDescription(description string) {
    r.Description = &description
}
/* param logConfig: 修改后的日志配置，不修改的话传空值，关闭日志集logTopicUid传空值。(Optional) */
func (r *UpdateServiceRequest) SetLogConfig(logConfig *fc.LogConfig) {
    r.LogConfig = logConfig
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateServiceRequest) GetRegionId() string {
    return r.RegionId
}

type UpdateServiceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateServiceResult `json:"result"`
}

type UpdateServiceResult struct {
}