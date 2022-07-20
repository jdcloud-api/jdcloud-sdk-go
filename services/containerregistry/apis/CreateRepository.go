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
    containerregistry "github.com/lshuining/jdcloud-sdk-go/services/containerregistry/models"
)

type CreateRepositoryRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 注册表名称  */
    RegistryName string `json:"registryName"`

    /* 镜像仓库名称。
可以专有模式如默认命名空间nginx-web-app；或者和命名空间一起将多个仓库聚集在一起如 project-a/nginx-web-app。
  */
    RepositoryName string `json:"repositoryName"`

    /* 注册表描述，<a href="https://www.jdcloud.com/help/detail/3870/isCatalog/1">参考公共参数规范</a>。
 (Optional) */
    Description *string `json:"description"`
}

/*
 * param regionId: Region ID (Required)
 * param registryName: 注册表名称 (Required)
 * param repositoryName: 镜像仓库名称。
可以专有模式如默认命名空间nginx-web-app；或者和命名空间一起将多个仓库聚集在一起如 project-a/nginx-web-app。
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateRepositoryRequest(
    regionId string,
    registryName string,
    repositoryName string,
) *CreateRepositoryRequest {

	return &CreateRepositoryRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/registries/{registryName}/repositories",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        RegistryName: registryName,
        RepositoryName: repositoryName,
	}
}

/*
 * param regionId: Region ID (Required)
 * param registryName: 注册表名称 (Required)
 * param repositoryName: 镜像仓库名称。
可以专有模式如默认命名空间nginx-web-app；或者和命名空间一起将多个仓库聚集在一起如 project-a/nginx-web-app。
 (Required)
 * param description: 注册表描述，<a href="https://www.jdcloud.com/help/detail/3870/isCatalog/1">参考公共参数规范</a>。
 (Optional)
 */
func NewCreateRepositoryRequestWithAllParams(
    regionId string,
    registryName string,
    repositoryName string,
    description *string,
) *CreateRepositoryRequest {

    return &CreateRepositoryRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/registries/{registryName}/repositories",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        RegistryName: registryName,
        RepositoryName: repositoryName,
        Description: description,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateRepositoryRequestWithoutParam() *CreateRepositoryRequest {

    return &CreateRepositoryRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/registries/{registryName}/repositories",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CreateRepositoryRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param registryName: 注册表名称(Required) */
func (r *CreateRepositoryRequest) SetRegistryName(registryName string) {
    r.RegistryName = registryName
}

/* param repositoryName: 镜像仓库名称。
可以专有模式如默认命名空间nginx-web-app；或者和命名空间一起将多个仓库聚集在一起如 project-a/nginx-web-app。
(Required) */
func (r *CreateRepositoryRequest) SetRepositoryName(repositoryName string) {
    r.RepositoryName = repositoryName
}

/* param description: 注册表描述，<a href="https://www.jdcloud.com/help/detail/3870/isCatalog/1">参考公共参数规范</a>。
(Optional) */
func (r *CreateRepositoryRequest) SetDescription(description string) {
    r.Description = &description
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateRepositoryRequest) GetRegionId() string {
    return r.RegionId
}

type CreateRepositoryResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateRepositoryResult `json:"result"`
}

type CreateRepositoryResult struct {
    Repository containerregistry.RepositoryShort `json:"repository"`
}