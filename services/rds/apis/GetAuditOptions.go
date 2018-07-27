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

type GetAuditOptionsRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Instance ID  */
    InstanceId string `json:"instanceId"`

    /* 审计选项类别，大小写敏感  */
    Name string `json:""`
}

/*
 * param regionId: Region ID (Required)
 * param instanceId: Instance ID (Required)
 * param name: 审计选项类别，大小写敏感 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetAuditOptionsRequest(
    regionId string,
    instanceId string,
    name string,
) *GetAuditOptionsRequest {

	return &GetAuditOptionsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/audit:getOptions",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        Name: name,
	}
}

/*
 * param regionId: Region ID (Required)
 * param instanceId: Instance ID (Required)
 * param name: 审计选项类别，大小写敏感 (Required)
 */
func NewGetAuditOptionsRequestWithAllParams(
    regionId string,
    instanceId string,
    name string,
) *GetAuditOptionsRequest {

    return &GetAuditOptionsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/audit:getOptions",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        Name: name,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetAuditOptionsRequestWithoutParam() *GetAuditOptionsRequest {

    return &GetAuditOptionsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/audit:getOptions",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *GetAuditOptionsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: Instance ID(Required) */
func (r *GetAuditOptionsRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param name: 审计选项类别，大小写敏感(Required) */
func (r *GetAuditOptionsRequest) SetName(name string) {
    r.Name = name
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetAuditOptionsRequest) GetRegionId() string {
    return r.RegionId
}

type GetAuditOptionsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetAuditOptionsResult `json:"result"`
}

type GetAuditOptionsResult struct {
    Enabled []string `json:"enabled"`
    Disabled []string `json:"disabled"`
}