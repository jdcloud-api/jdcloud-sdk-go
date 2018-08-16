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
    jmr "github.com/jdcloud-api/jdcloud-sdk-go/services/jmr/models"
)

type GetClusterCronJobCountRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 集群的Id  */
    ClusterId string `json:"clusterId"`
}

/*
 * param regionId: 地域ID (Required)
 * param clusterId: 集群的Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetClusterCronJobCountRequest(
    regionId string,
    clusterId string,
) *GetClusterCronJobCountRequest {

	return &GetClusterCronJobCountRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/clusterCronJob/{clusterId}:count",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ClusterId: clusterId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param clusterId: 集群的Id (Required)
 */
func NewGetClusterCronJobCountRequestWithAllParams(
    regionId string,
    clusterId string,
) *GetClusterCronJobCountRequest {

    return &GetClusterCronJobCountRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/clusterCronJob/{clusterId}:count",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ClusterId: clusterId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetClusterCronJobCountRequestWithoutParam() *GetClusterCronJobCountRequest {

    return &GetClusterCronJobCountRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/clusterCronJob/{clusterId}:count",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *GetClusterCronJobCountRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param clusterId: 集群的Id(Required) */
func (r *GetClusterCronJobCountRequest) SetClusterId(clusterId string) {
    r.ClusterId = clusterId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetClusterCronJobCountRequest) GetRegionId() string {
    return r.RegionId
}

type GetClusterCronJobCountResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetClusterCronJobCountResult `json:"result"`
}

type GetClusterCronJobCountResult struct {
    Status string `json:"status"`
    Message string `json:"message"`
    Data jmr.CronJobData `json:"data"`
}