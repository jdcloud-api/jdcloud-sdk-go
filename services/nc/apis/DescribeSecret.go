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
    nc "github.com/jdcloud-api/jdcloud-sdk-go/services/nc/models"
)

type DescribeSecretRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Secret Name  */
    Name string `json:"name"`
}

/*
 * param regionId: Region ID (Required)
 * param name: Secret Name (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeSecretRequest(
    regionId string,
    name string,
) *DescribeSecretRequest {

	return &DescribeSecretRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/secrets/{name}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Name: name,
	}
}

/*
 * param regionId: Region ID (Required)
 * param name: Secret Name (Required)
 */
func NewDescribeSecretRequestWithAllParams(
    regionId string,
    name string,
) *DescribeSecretRequest {

    return &DescribeSecretRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/secrets/{name}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Name: name,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeSecretRequestWithoutParam() *DescribeSecretRequest {

    return &DescribeSecretRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/secrets/{name}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeSecretRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param name: Secret Name(Required) */
func (r *DescribeSecretRequest) SetName(name string) {
    r.Name = name
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeSecretRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeSecretResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeSecretResult `json:"result"`
}

type DescribeSecretResult struct {
    Secret nc.Secret `json:"secret"`
}