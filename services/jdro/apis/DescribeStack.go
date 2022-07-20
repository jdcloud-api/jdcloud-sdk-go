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
    jdro "github.com/lshuining/jdcloud-sdk-go/services/jdro/models"
)

type DescribeStackRequest struct {

    core.JDCloudRequest

    /* 地域 ID  */
    RegionId string `json:"regionId"`

    /* 资源栈 ID  */
    StackId string `json:"stackId"`
}

/*
 * param regionId: 地域 ID (Required)
 * param stackId: 资源栈 ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeStackRequest(
    regionId string,
    stackId string,
) *DescribeStackRequest {

	return &DescribeStackRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/stacks/{stackId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        StackId: stackId,
	}
}

/*
 * param regionId: 地域 ID (Required)
 * param stackId: 资源栈 ID (Required)
 */
func NewDescribeStackRequestWithAllParams(
    regionId string,
    stackId string,
) *DescribeStackRequest {

    return &DescribeStackRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/stacks/{stackId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        StackId: stackId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeStackRequestWithoutParam() *DescribeStackRequest {

    return &DescribeStackRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/stacks/{stackId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 ID(Required) */
func (r *DescribeStackRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param stackId: 资源栈 ID(Required) */
func (r *DescribeStackRequest) SetStackId(stackId string) {
    r.StackId = stackId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeStackRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeStackResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeStackResult `json:"result"`
}

type DescribeStackResult struct {
    Stack jdro.StackOut `json:"stack"`
}