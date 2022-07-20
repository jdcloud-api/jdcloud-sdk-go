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
    kubernetes "github.com/lshuining/jdcloud-sdk-go/services/kubernetes/models"
)

type DescribeUpgradableNodeVersionsRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 集群 ID  */
    ClusterId string `json:"clusterId"`

    /* 节点组 id (Optional) */
    NodeGroupIds []string `json:"nodeGroupIds"`
}

/*
 * param regionId: Region ID (Required)
 * param clusterId: 集群 ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeUpgradableNodeVersionsRequest(
    regionId string,
    clusterId string,
) *DescribeUpgradableNodeVersionsRequest {

	return &DescribeUpgradableNodeVersionsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/clusters/{clusterId}/upgradableNodeVersions",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ClusterId: clusterId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param clusterId: 集群 ID (Required)
 * param nodeGroupIds: 节点组 id (Optional)
 */
func NewDescribeUpgradableNodeVersionsRequestWithAllParams(
    regionId string,
    clusterId string,
    nodeGroupIds []string,
) *DescribeUpgradableNodeVersionsRequest {

    return &DescribeUpgradableNodeVersionsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/clusters/{clusterId}/upgradableNodeVersions",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ClusterId: clusterId,
        NodeGroupIds: nodeGroupIds,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeUpgradableNodeVersionsRequestWithoutParam() *DescribeUpgradableNodeVersionsRequest {

    return &DescribeUpgradableNodeVersionsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/clusters/{clusterId}/upgradableNodeVersions",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeUpgradableNodeVersionsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param clusterId: 集群 ID(Required) */
func (r *DescribeUpgradableNodeVersionsRequest) SetClusterId(clusterId string) {
    r.ClusterId = clusterId
}

/* param nodeGroupIds: 节点组 id(Optional) */
func (r *DescribeUpgradableNodeVersionsRequest) SetNodeGroupIds(nodeGroupIds []string) {
    r.NodeGroupIds = nodeGroupIds
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeUpgradableNodeVersionsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeUpgradableNodeVersionsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeUpgradableNodeVersionsResult `json:"result"`
}

type DescribeUpgradableNodeVersionsResult struct {
    NdoeVersions []kubernetes.NodeVersion `json:"ndoeVersions"`
}