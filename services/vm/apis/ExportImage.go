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

type ExportImageRequest struct {

    core.JDCloudRequest

    /* 地域ID。  */
    RegionId string `json:"regionId"`

    /* 镜像ID。  */
    ImageId string `json:"imageId"`

    /* 用户创建的服务角色名称。  */
    RoleName string `json:"roleName"`

    /* 存储导出镜像文件的 `oss bucket` 的域名，请填写以 https:// 开头的完整url。  */
    OssUrl string `json:"ossUrl"`

    /* 导出镜像文件名前缀，仅支持英文字母和数字，不能超过32个字符。 (Optional) */
    OssPrefix *string `json:"ossPrefix"`

    /* 用户导出镜像的幂等性保证。每次导出请传入不同的值，如果传值与某次的clientToken相同，则返还同一个请求结果，不能超过64个字符。 (Optional) */
    ClientToken *string `json:"clientToken"`
}

/*
 * param regionId: 地域ID。 (Required)
 * param imageId: 镜像ID。 (Required)
 * param roleName: 用户创建的服务角色名称。 (Required)
 * param ossUrl: 存储导出镜像文件的 `oss bucket` 的域名，请填写以 https:// 开头的完整url。 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewExportImageRequest(
    regionId string,
    imageId string,
    roleName string,
    ossUrl string,
) *ExportImageRequest {

	return &ExportImageRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/images/{imageId}:exportImage",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ImageId: imageId,
        RoleName: roleName,
        OssUrl: ossUrl,
	}
}

/*
 * param regionId: 地域ID。 (Required)
 * param imageId: 镜像ID。 (Required)
 * param roleName: 用户创建的服务角色名称。 (Required)
 * param ossUrl: 存储导出镜像文件的 `oss bucket` 的域名，请填写以 https:// 开头的完整url。 (Required)
 * param ossPrefix: 导出镜像文件名前缀，仅支持英文字母和数字，不能超过32个字符。 (Optional)
 * param clientToken: 用户导出镜像的幂等性保证。每次导出请传入不同的值，如果传值与某次的clientToken相同，则返还同一个请求结果，不能超过64个字符。 (Optional)
 */
func NewExportImageRequestWithAllParams(
    regionId string,
    imageId string,
    roleName string,
    ossUrl string,
    ossPrefix *string,
    clientToken *string,
) *ExportImageRequest {

    return &ExportImageRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/images/{imageId}:exportImage",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ImageId: imageId,
        RoleName: roleName,
        OssUrl: ossUrl,
        OssPrefix: ossPrefix,
        ClientToken: clientToken,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewExportImageRequestWithoutParam() *ExportImageRequest {

    return &ExportImageRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/images/{imageId}:exportImage",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID。(Required) */
func (r *ExportImageRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param imageId: 镜像ID。(Required) */
func (r *ExportImageRequest) SetImageId(imageId string) {
    r.ImageId = imageId
}

/* param roleName: 用户创建的服务角色名称。(Required) */
func (r *ExportImageRequest) SetRoleName(roleName string) {
    r.RoleName = roleName
}

/* param ossUrl: 存储导出镜像文件的 `oss bucket` 的域名，请填写以 https:// 开头的完整url。(Required) */
func (r *ExportImageRequest) SetOssUrl(ossUrl string) {
    r.OssUrl = ossUrl
}

/* param ossPrefix: 导出镜像文件名前缀，仅支持英文字母和数字，不能超过32个字符。(Optional) */
func (r *ExportImageRequest) SetOssPrefix(ossPrefix string) {
    r.OssPrefix = &ossPrefix
}

/* param clientToken: 用户导出镜像的幂等性保证。每次导出请传入不同的值，如果传值与某次的clientToken相同，则返还同一个请求结果，不能超过64个字符。(Optional) */
func (r *ExportImageRequest) SetClientToken(clientToken string) {
    r.ClientToken = &clientToken
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ExportImageRequest) GetRegionId() string {
    return r.RegionId
}

type ExportImageResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ExportImageResult `json:"result"`
}

type ExportImageResult struct {
    TaskId string `json:"taskId"`
}