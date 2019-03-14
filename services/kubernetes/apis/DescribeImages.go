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
    kubernetes "github.com/jdcloud-api/jdcloud-sdk-go/services/kubernetes/models"
)

type DescribeImagesRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 集群的大版本，如 1.8.12 (Optional) */
    MasterVersion *string `json:"masterVersion"`

    /* master 的镜像编码，如 1.8.12-jke (Optional) */
    MasterImageCode *string `json:"masterImageCode"`
}

/*
 * param regionId: Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeImagesRequest(
    regionId string,
) *DescribeImagesRequest {

	return &DescribeImagesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/images",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param masterVersion: 集群的大版本，如 1.8.12 (Optional)
 * param masterImageCode: master 的镜像编码，如 1.8.12-jke (Optional)
 */
func NewDescribeImagesRequestWithAllParams(
    regionId string,
    masterVersion *string,
    masterImageCode *string,
) *DescribeImagesRequest {

    return &DescribeImagesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/images",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        MasterVersion: masterVersion,
        MasterImageCode: masterImageCode,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeImagesRequestWithoutParam() *DescribeImagesRequest {

    return &DescribeImagesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/images",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeImagesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param masterVersion: 集群的大版本，如 1.8.12(Optional) */
func (r *DescribeImagesRequest) SetMasterVersion(masterVersion string) {
    r.MasterVersion = &masterVersion
}

/* param masterImageCode: master 的镜像编码，如 1.8.12-jke(Optional) */
func (r *DescribeImagesRequest) SetMasterImageCode(masterImageCode string) {
    r.MasterImageCode = &masterImageCode
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeImagesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeImagesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeImagesResult `json:"result"`
}

type DescribeImagesResult struct {
    MasterImages []kubernetes.MasterImage `json:"masterImages"`
}