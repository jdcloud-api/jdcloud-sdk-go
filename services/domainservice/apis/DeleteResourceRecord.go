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
)

type DeleteResourceRecordRequest struct {

    core.JDCloudRequest

    /* 实例所属的地域ID  */
    RegionId string `json:"regionId"`

    /* 域名ID，请使用describeDomains接口获取。  */
    DomainId string `json:"domainId"`

    /* 解析记录ID，请使用describeResourceRecord接口获取。  */
    ResourceRecordId string `json:"resourceRecordId"`
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param domainId: 域名ID，请使用describeDomains接口获取。 (Required)
 * param resourceRecordId: 解析记录ID，请使用describeResourceRecord接口获取。 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteResourceRecordRequest(
    regionId string,
    domainId string,
    resourceRecordId string,
) *DeleteResourceRecordRequest {

	return &DeleteResourceRecordRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/domain/{domainId}/ResourceRecord/{resourceRecordId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        DomainId: domainId,
        ResourceRecordId: resourceRecordId,
	}
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param domainId: 域名ID，请使用describeDomains接口获取。 (Required)
 * param resourceRecordId: 解析记录ID，请使用describeResourceRecord接口获取。 (Required)
 */
func NewDeleteResourceRecordRequestWithAllParams(
    regionId string,
    domainId string,
    resourceRecordId string,
) *DeleteResourceRecordRequest {

    return &DeleteResourceRecordRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain/{domainId}/ResourceRecord/{resourceRecordId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        DomainId: domainId,
        ResourceRecordId: resourceRecordId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteResourceRecordRequestWithoutParam() *DeleteResourceRecordRequest {

    return &DeleteResourceRecordRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain/{domainId}/ResourceRecord/{resourceRecordId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 实例所属的地域ID(Required) */
func (r *DeleteResourceRecordRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param domainId: 域名ID，请使用describeDomains接口获取。(Required) */
func (r *DeleteResourceRecordRequest) SetDomainId(domainId string) {
    r.DomainId = domainId
}

/* param resourceRecordId: 解析记录ID，请使用describeResourceRecord接口获取。(Required) */
func (r *DeleteResourceRecordRequest) SetResourceRecordId(resourceRecordId string) {
    r.ResourceRecordId = resourceRecordId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteResourceRecordRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteResourceRecordResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteResourceRecordResult `json:"result"`
}

type DeleteResourceRecordResult struct {
}