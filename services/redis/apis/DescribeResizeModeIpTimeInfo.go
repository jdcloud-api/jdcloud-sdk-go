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

type DescribeResizeModeIpTimeInfoRequest struct {

    core.JDCloudRequest

    /* 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2  */
    RegionId string `json:"regionId"`

    /* 缓存Redis实例ID，是访问实例的唯一标识  */
    CacheInstanceId string `json:"cacheInstanceId"`

    /* 新规格。不传参只返回是否支持新变配功能。 (Optional) */
    CacheInstanceClass *string `json:"cacheInstanceClass"`

    /* 自定义分片数。不传参只返回是否支持新变配功能。 (Optional) */
    ShardNumber *int `json:"shardNumber"`
}

/*
 * param regionId: 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2 (Required)
 * param cacheInstanceId: 缓存Redis实例ID，是访问实例的唯一标识 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeResizeModeIpTimeInfoRequest(
    regionId string,
    cacheInstanceId string,
) *DescribeResizeModeIpTimeInfoRequest {

	return &DescribeResizeModeIpTimeInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}/describeResizeModeIpTimeInfo",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        CacheInstanceId: cacheInstanceId,
	}
}

/*
 * param regionId: 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2 (Required)
 * param cacheInstanceId: 缓存Redis实例ID，是访问实例的唯一标识 (Required)
 * param cacheInstanceClass: 新规格。不传参只返回是否支持新变配功能。 (Optional)
 * param shardNumber: 自定义分片数。不传参只返回是否支持新变配功能。 (Optional)
 */
func NewDescribeResizeModeIpTimeInfoRequestWithAllParams(
    regionId string,
    cacheInstanceId string,
    cacheInstanceClass *string,
    shardNumber *int,
) *DescribeResizeModeIpTimeInfoRequest {

    return &DescribeResizeModeIpTimeInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}/describeResizeModeIpTimeInfo",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        CacheInstanceId: cacheInstanceId,
        CacheInstanceClass: cacheInstanceClass,
        ShardNumber: shardNumber,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeResizeModeIpTimeInfoRequestWithoutParam() *DescribeResizeModeIpTimeInfoRequest {

    return &DescribeResizeModeIpTimeInfoRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}/describeResizeModeIpTimeInfo",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2(Required) */
func (r *DescribeResizeModeIpTimeInfoRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param cacheInstanceId: 缓存Redis实例ID，是访问实例的唯一标识(Required) */
func (r *DescribeResizeModeIpTimeInfoRequest) SetCacheInstanceId(cacheInstanceId string) {
    r.CacheInstanceId = cacheInstanceId
}
/* param cacheInstanceClass: 新规格。不传参只返回是否支持新变配功能。(Optional) */
func (r *DescribeResizeModeIpTimeInfoRequest) SetCacheInstanceClass(cacheInstanceClass string) {
    r.CacheInstanceClass = &cacheInstanceClass
}
/* param shardNumber: 自定义分片数。不传参只返回是否支持新变配功能。(Optional) */
func (r *DescribeResizeModeIpTimeInfoRequest) SetShardNumber(shardNumber int) {
    r.ShardNumber = &shardNumber
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeResizeModeIpTimeInfoRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeResizeModeIpTimeInfoResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeResizeModeIpTimeInfoResult `json:"result"`
}

type DescribeResizeModeIpTimeInfoResult struct {
    SupportPScale bool `json:"supportPScale"`
    SupportParallel bool `json:"supportParallel"`
    SerialIpCount int `json:"serialIpCount"`
    SerialTimeCost int `json:"serialTimeCost"`
    ParallelIpCount int `json:"parallelIpCount"`
    ParallelTimeCost int `json:"parallelTimeCost"`
}