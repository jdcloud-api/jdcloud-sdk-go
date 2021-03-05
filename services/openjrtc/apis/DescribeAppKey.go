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

type DescribeAppKeyRequest struct {

    core.JDCloudRequest

    /* 应用ID  */
    AppId string `json:"appId"`
}

/*
 * param appId: 应用ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeAppKeyRequest(
    appId string,
) *DescribeAppKeyRequest {

	return &DescribeAppKeyRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/applications/{appId}:describeAppKey",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        AppId: appId,
	}
}

/*
 * param appId: 应用ID (Required)
 */
func NewDescribeAppKeyRequestWithAllParams(
    appId string,
) *DescribeAppKeyRequest {

    return &DescribeAppKeyRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/applications/{appId}:describeAppKey",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        AppId: appId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeAppKeyRequestWithoutParam() *DescribeAppKeyRequest {

    return &DescribeAppKeyRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/applications/{appId}:describeAppKey",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param appId: 应用ID(Required) */
func (r *DescribeAppKeyRequest) SetAppId(appId string) {
    r.AppId = appId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeAppKeyRequest) GetRegionId() string {
    return ""
}

type DescribeAppKeyResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeAppKeyResult `json:"result"`
}

type DescribeAppKeyResult struct {
    AppId string `json:"appId"`
    AppKey string `json:"appKey"`
}