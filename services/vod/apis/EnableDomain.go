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

type EnableDomainRequest struct {

    core.JDCloudRequest

    /* 域名ID  */
    DomainId int `json:"domainId"`
}

/*
 * param domainId: 域名ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewEnableDomainRequest(
    domainId int,
) *EnableDomainRequest {

	return &EnableDomainRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domains/{domainId}:enable",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        DomainId: domainId,
	}
}

/*
 * param domainId: 域名ID (Required)
 */
func NewEnableDomainRequestWithAllParams(
    domainId int,
) *EnableDomainRequest {

    return &EnableDomainRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domains/{domainId}:enable",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        DomainId: domainId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewEnableDomainRequestWithoutParam() *EnableDomainRequest {

    return &EnableDomainRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domains/{domainId}:enable",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domainId: 域名ID(Required) */
func (r *EnableDomainRequest) SetDomainId(domainId int) {
    r.DomainId = domainId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r EnableDomainRequest) GetRegionId() string {
    return ""
}

type EnableDomainResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result EnableDomainResult `json:"result"`
}

type EnableDomainResult struct {
}