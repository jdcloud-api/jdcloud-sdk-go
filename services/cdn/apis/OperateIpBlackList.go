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

type OperateIpBlackListRequest struct {

    core.JDCloudRequest

    /* 用户域名  */
    Domain string `json:"domain"`

    /* ip黑名单状态取值[on,off] (Optional) */
    Status *string `json:"status"`
}

/*
 * param domain: 用户域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewOperateIpBlackListRequest(
    domain string,
) *OperateIpBlackListRequest {

	return &OperateIpBlackListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domain/{domain}/ipBlackList:operate",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        Domain: domain,
	}
}

/*
 * param domain: 用户域名 (Required)
 * param status: ip黑名单状态取值[on,off] (Optional)
 */
func NewOperateIpBlackListRequestWithAllParams(
    domain string,
    status *string,
) *OperateIpBlackListRequest {

    return &OperateIpBlackListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/ipBlackList:operate",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        Domain: domain,
        Status: status,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewOperateIpBlackListRequestWithoutParam() *OperateIpBlackListRequest {

    return &OperateIpBlackListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/ipBlackList:operate",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domain: 用户域名(Required) */
func (r *OperateIpBlackListRequest) SetDomain(domain string) {
    r.Domain = domain
}

/* param status: ip黑名单状态取值[on,off](Optional) */
func (r *OperateIpBlackListRequest) SetStatus(status string) {
    r.Status = &status
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r OperateIpBlackListRequest) GetRegionId() string {
    return ""
}

type OperateIpBlackListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result OperateIpBlackListResult `json:"result"`
}

type OperateIpBlackListResult struct {
}