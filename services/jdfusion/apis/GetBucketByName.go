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
    jdfusion "github.com/jdcloud-api/jdcloud-sdk-go/services/jdfusion/models"
)

type GetBucketByNameRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* OSS存储桶名称  */
    Name string `json:"name"`
}

/*
 * param regionId: 地域ID (Required)
 * param name: OSS存储桶名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetBucketByNameRequest(
    regionId string,
    name string,
) *GetBucketByNameRequest {

	return &GetBucketByNameRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/oss_buckets/{name}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Name: name,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param name: OSS存储桶名称 (Required)
 */
func NewGetBucketByNameRequestWithAllParams(
    regionId string,
    name string,
) *GetBucketByNameRequest {

    return &GetBucketByNameRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/oss_buckets/{name}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Name: name,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetBucketByNameRequestWithoutParam() *GetBucketByNameRequest {

    return &GetBucketByNameRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/oss_buckets/{name}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *GetBucketByNameRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param name: OSS存储桶名称(Required) */
func (r *GetBucketByNameRequest) SetName(name string) {
    r.Name = name
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetBucketByNameRequest) GetRegionId() string {
    return r.RegionId
}

type GetBucketByNameResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetBucketByNameResult `json:"result"`
}

type GetBucketByNameResult struct {
    Bucket jdfusion.OssBucketDetailInfo `json:"bucket"`
}