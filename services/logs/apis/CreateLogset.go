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
    logs "github.com/jdcloud-api/jdcloud-sdk-go/services/logs/models"
)

type CreateLogsetRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 日志集名称  */
    Name string `json:"name"`

    /* 日志集描述 (Optional) */
    Description *string `json:"description"`

    /* 保存周期，只能是 7， 15， 30 (Optional) */
    LifeCycle *int `json:"lifeCycle"`

    /* 标签列表 (Optional) */
    Tags []logs.Tag `json:"tags"`

    /* 资源组信息 (Optional) */
    ResourceGroupUID *string `json:"resourceGroupUID"`

    /* 是否是内部资源。 设置为 true，则外部租户不可见。 (Optional) */
    Inner *bool `json:"inner"`
}

/*
 * param regionId: 地域 Id (Required)
 * param name: 日志集名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateLogsetRequest(
    regionId string,
    name string,
) *CreateLogsetRequest {

	return &CreateLogsetRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/logsets",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Name: name,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param name: 日志集名称 (Required)
 * param description: 日志集描述 (Optional)
 * param lifeCycle: 保存周期，只能是 7， 15， 30 (Optional)
 * param tags: 标签列表 (Optional)
 * param resourceGroupUID: 资源组信息 (Optional)
 * param inner: 是否是内部资源。 设置为 true，则外部租户不可见。 (Optional)
 */
func NewCreateLogsetRequestWithAllParams(
    regionId string,
    name string,
    description *string,
    lifeCycle *int,
    tags []logs.Tag,
    resourceGroupUID *string,
    inner *bool,
) *CreateLogsetRequest {

    return &CreateLogsetRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/logsets",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Name: name,
        Description: description,
        LifeCycle: lifeCycle,
        Tags: tags,
        ResourceGroupUID: resourceGroupUID,
        Inner: inner,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateLogsetRequestWithoutParam() *CreateLogsetRequest {

    return &CreateLogsetRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/logsets",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *CreateLogsetRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param name: 日志集名称(Required) */
func (r *CreateLogsetRequest) SetName(name string) {
    r.Name = name
}
/* param description: 日志集描述(Optional) */
func (r *CreateLogsetRequest) SetDescription(description string) {
    r.Description = &description
}
/* param lifeCycle: 保存周期，只能是 7， 15， 30(Optional) */
func (r *CreateLogsetRequest) SetLifeCycle(lifeCycle int) {
    r.LifeCycle = &lifeCycle
}
/* param tags: 标签列表(Optional) */
func (r *CreateLogsetRequest) SetTags(tags []logs.Tag) {
    r.Tags = tags
}
/* param resourceGroupUID: 资源组信息(Optional) */
func (r *CreateLogsetRequest) SetResourceGroupUID(resourceGroupUID string) {
    r.ResourceGroupUID = &resourceGroupUID
}
/* param inner: 是否是内部资源。 设置为 true，则外部租户不可见。(Optional) */
func (r *CreateLogsetRequest) SetInner(inner bool) {
    r.Inner = &inner
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateLogsetRequest) GetRegionId() string {
    return r.RegionId
}

type CreateLogsetResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateLogsetResult `json:"result"`
}

type CreateLogsetResult struct {
    UID string `json:"uID"`
}