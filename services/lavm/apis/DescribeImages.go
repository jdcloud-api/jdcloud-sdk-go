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
    lavm "github.com/jdcloud-api/jdcloud-sdk-go/services/lavm/models"
)

type DescribeImagesRequest struct {

    core.JDCloudRequest

    /* regionId
  */
    RegionId string `json:"regionId"`

    /* 镜像ID, `[\"img-m5s0****29\", \"img-m5s0****30\"]`, json array 字串。
 (Optional) */
    ImageIds *string `json:"imageIds"`

    /* 镜像类型。 可能取值：system, app, custom, 如果imageIds不传时，此项为必传。
 (Optional) */
    ImageType *string `json:"imageType"`

    /* 根据镜像的操作系统发行版查询。
取值范围：`Ubuntu、CentOS、Windows Server`。
 (Optional) */
    Platform *string `json:"platform"`
}

/*
 * param regionId: regionId
 (Required)
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
 * param regionId: regionId
 (Required)
 * param imageIds: 镜像ID, `[\"img-m5s0****29\", \"img-m5s0****30\"]`, json array 字串。
 (Optional)
 * param imageType: 镜像类型。 可能取值：system, app, custom, 如果imageIds不传时，此项为必传。
 (Optional)
 * param platform: 根据镜像的操作系统发行版查询。
取值范围：`Ubuntu、CentOS、Windows Server`。
 (Optional)
 */
func NewDescribeImagesRequestWithAllParams(
    regionId string,
    imageIds *string,
    imageType *string,
    platform *string,
) *DescribeImagesRequest {

    return &DescribeImagesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/images",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ImageIds: imageIds,
        ImageType: imageType,
        Platform: platform,
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

/* param regionId: regionId
(Required) */
func (r *DescribeImagesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param imageIds: 镜像ID, `[\"img-m5s0****29\", \"img-m5s0****30\"]`, json array 字串。
(Optional) */
func (r *DescribeImagesRequest) SetImageIds(imageIds string) {
    r.ImageIds = &imageIds
}
/* param imageType: 镜像类型。 可能取值：system, app, custom, 如果imageIds不传时，此项为必传。
(Optional) */
func (r *DescribeImagesRequest) SetImageType(imageType string) {
    r.ImageType = &imageType
}
/* param platform: 根据镜像的操作系统发行版查询。
取值范围：`Ubuntu、CentOS、Windows Server`。
(Optional) */
func (r *DescribeImagesRequest) SetPlatform(platform string) {
    r.Platform = &platform
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
    Images []lavm.Image `json:"images"`
}