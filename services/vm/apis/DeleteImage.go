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

type DeleteImageRequest struct {

    core.JDCloudRequest

    /* 地域ID。  */
    RegionId string `json:"regionId"`

    /* 镜像ID。  */
    ImageId string `json:"imageId"`

    /* 删除镜像时是否删除关联的快照。默认为 `false`；如果指定为 `true`, 将会删除镜像关联的快照。
 (Optional) */
    DeleteSnapshot *bool `json:"deleteSnapshot"`
}

/*
 * param regionId: 地域ID。 (Required)
 * param imageId: 镜像ID。 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteImageRequest(
    regionId string,
    imageId string,
) *DeleteImageRequest {

	return &DeleteImageRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/images/{imageId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ImageId: imageId,
	}
}

/*
 * param regionId: 地域ID。 (Required)
 * param imageId: 镜像ID。 (Required)
 * param deleteSnapshot: 删除镜像时是否删除关联的快照。默认为 `false`；如果指定为 `true`, 将会删除镜像关联的快照。
 (Optional)
 */
func NewDeleteImageRequestWithAllParams(
    regionId string,
    imageId string,
    deleteSnapshot *bool,
) *DeleteImageRequest {

    return &DeleteImageRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/images/{imageId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ImageId: imageId,
        DeleteSnapshot: deleteSnapshot,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteImageRequestWithoutParam() *DeleteImageRequest {

    return &DeleteImageRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/images/{imageId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID。(Required) */
func (r *DeleteImageRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param imageId: 镜像ID。(Required) */
func (r *DeleteImageRequest) SetImageId(imageId string) {
    r.ImageId = imageId
}
/* param deleteSnapshot: 删除镜像时是否删除关联的快照。默认为 `false`；如果指定为 `true`, 将会删除镜像关联的快照。
(Optional) */
func (r *DeleteImageRequest) SetDeleteSnapshot(deleteSnapshot bool) {
    r.DeleteSnapshot = &deleteSnapshot
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteImageRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteImageResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteImageResult `json:"result"`
}

type DeleteImageResult struct {
}