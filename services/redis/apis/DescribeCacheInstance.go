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
    redis "github.com/jdcloud-api/jdcloud-sdk-go/services/redis/models"
)

type DescribeCacheInstanceRequest struct {

    core.JDCloudRequest

    /* 缓存Redis实例所在区域的Region ID。目前缓存Redis有华北、华南、华东区域，对应Region ID为cn-north-1、cn-south-1、cn-east-2  */
    RegionId string `json:"regionId"`

    /* 缓存Redis实例ID  */
    CacheInstanceId string `json:"cacheInstanceId"`
}

/*
 * param regionId: 缓存Redis实例所在区域的Region ID。目前缓存Redis有华北、华南、华东区域，对应Region ID为cn-north-1、cn-south-1、cn-east-2 
 * param cacheInstanceId: 缓存Redis实例ID 
 */
func NewDescribeCacheInstanceRequest(
    regionId string,
    cacheInstanceId string,
) *DescribeCacheInstanceRequest {

	return &DescribeCacheInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        CacheInstanceId: cacheInstanceId,
	}
}

func (r *DescribeCacheInstanceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *DescribeCacheInstanceRequest) SetCacheInstanceId(cacheInstanceId string) {
    r.CacheInstanceId = cacheInstanceId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeCacheInstanceRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeCacheInstanceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeCacheInstanceResult `json:"result"`
}

type DescribeCacheInstanceResult struct {
    CacheInstance redis.CacheInstance `json:"cacheInstance"`
}