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

type UpdateCollectInfoRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 采集配置 UID  */
    CollectInfoUID string `json:"collectInfoUID"`

    /* 高可用组资源 (Optional) */
    AgResource []logs.AgResource `json:"agResource"`

    /* 采集状态，0-禁用，1-启用  */
    Enabled bool `json:"enabled"`

    /* 过滤器是否启用。当appcode为custom时必填 (Optional) */
    FilterEnabled *bool `json:"filterEnabled"`

    /* 自定义日志转发目的地, 只支持业务应用日志。支持类型："kafka"，"es" (Optional) */
    LogCustomTarget *string `json:"logCustomTarget"`

    /* 自定义日志转发目的地配置，KV 结构，具体配置参考 LogCustomTargetKafkaConf 和 LogCustomTargetEsConf (Optional) */
    LogCustomTargetConf *interface{} `json:"logCustomTargetConf"`

    /* 日志文件名。当appcode为custom时为必填。日志文件名支持正则表达式。 (Optional) */
    LogFile *string `json:"logFile"`

    /* 过滤器。设置过滤器后可根据用户设定的关键词采集部分日志，如仅采集 Error 的日志。目前最大允许5个。 (Optional) */
    LogFilters []string `json:"logFilters"`

    /* 日志路径。当appcode为custom时为必填。目前仅支持对 Linux 云主机上的日志进行采集，路径支持通配符“*”和“？”，文件路径应符合 Linux 的文件路径规则 (Optional) */
    LogPath *string `json:"logPath"`

    /* 目的地是否是日志服务logtopic，只支持业务应用日志 (Optional) */
    LogtopicEnabled *bool `json:"logtopicEnabled"`

    /* 首行正则 (Optional) */
    RegexpStr *string `json:"regexpStr"`

    /* 采集资源时选择的模式，1.正常的选择实例模式（默认模式）；2.选择标签tag模式 3.选择高可用组ag模式 (Optional) */
    ResourceMode *int64 `json:"resourceMode"`

    /* 采集实例类型, 只能是 all/part  当选择all时，传入的实例列表无效  */
    ResourceType string `json:"resourceType"`

    /* 采集实例列表（存在上限限制） (Optional) */
    Resources []logs.Resource `json:"resources"`

    /*  (Optional) */
    TagResource *logs.TagResource `json:"tagResource"`
}

/*
 * param regionId: 地域 Id (Required)
 * param collectInfoUID: 采集配置 UID (Required)
 * param enabled: 采集状态，0-禁用，1-启用 (Required)
 * param resourceType: 采集实例类型, 只能是 all/part  当选择all时，传入的实例列表无效 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateCollectInfoRequest(
    regionId string,
    collectInfoUID string,
    enabled bool,
    resourceType string,
) *UpdateCollectInfoRequest {

	return &UpdateCollectInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/collectinfos/{collectInfoUID}",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        CollectInfoUID: collectInfoUID,
        Enabled: enabled,
        ResourceType: resourceType,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param collectInfoUID: 采集配置 UID (Required)
 * param agResource: 高可用组资源 (Optional)
 * param enabled: 采集状态，0-禁用，1-启用 (Required)
 * param filterEnabled: 过滤器是否启用。当appcode为custom时必填 (Optional)
 * param logCustomTarget: 自定义日志转发目的地, 只支持业务应用日志。支持类型："kafka"，"es" (Optional)
 * param logCustomTargetConf: 自定义日志转发目的地配置，KV 结构，具体配置参考 LogCustomTargetKafkaConf 和 LogCustomTargetEsConf (Optional)
 * param logFile: 日志文件名。当appcode为custom时为必填。日志文件名支持正则表达式。 (Optional)
 * param logFilters: 过滤器。设置过滤器后可根据用户设定的关键词采集部分日志，如仅采集 Error 的日志。目前最大允许5个。 (Optional)
 * param logPath: 日志路径。当appcode为custom时为必填。目前仅支持对 Linux 云主机上的日志进行采集，路径支持通配符“*”和“？”，文件路径应符合 Linux 的文件路径规则 (Optional)
 * param logtopicEnabled: 目的地是否是日志服务logtopic，只支持业务应用日志 (Optional)
 * param regexpStr: 首行正则 (Optional)
 * param resourceMode: 采集资源时选择的模式，1.正常的选择实例模式（默认模式）；2.选择标签tag模式 3.选择高可用组ag模式 (Optional)
 * param resourceType: 采集实例类型, 只能是 all/part  当选择all时，传入的实例列表无效 (Required)
 * param resources: 采集实例列表（存在上限限制） (Optional)
 * param tagResource:  (Optional)
 */
func NewUpdateCollectInfoRequestWithAllParams(
    regionId string,
    collectInfoUID string,
    agResource []logs.AgResource,
    enabled bool,
    filterEnabled *bool,
    logCustomTarget *string,
    logCustomTargetConf *interface{},
    logFile *string,
    logFilters []string,
    logPath *string,
    logtopicEnabled *bool,
    regexpStr *string,
    resourceMode *int64,
    resourceType string,
    resources []logs.Resource,
    tagResource *logs.TagResource,
) *UpdateCollectInfoRequest {

    return &UpdateCollectInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/collectinfos/{collectInfoUID}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        CollectInfoUID: collectInfoUID,
        AgResource: agResource,
        Enabled: enabled,
        FilterEnabled: filterEnabled,
        LogCustomTarget: logCustomTarget,
        LogCustomTargetConf: logCustomTargetConf,
        LogFile: logFile,
        LogFilters: logFilters,
        LogPath: logPath,
        LogtopicEnabled: logtopicEnabled,
        RegexpStr: regexpStr,
        ResourceMode: resourceMode,
        ResourceType: resourceType,
        Resources: resources,
        TagResource: tagResource,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateCollectInfoRequestWithoutParam() *UpdateCollectInfoRequest {

    return &UpdateCollectInfoRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/collectinfos/{collectInfoUID}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *UpdateCollectInfoRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param collectInfoUID: 采集配置 UID(Required) */
func (r *UpdateCollectInfoRequest) SetCollectInfoUID(collectInfoUID string) {
    r.CollectInfoUID = collectInfoUID
}

/* param agResource: 高可用组资源(Optional) */
func (r *UpdateCollectInfoRequest) SetAgResource(agResource []logs.AgResource) {
    r.AgResource = agResource
}

/* param enabled: 采集状态，0-禁用，1-启用(Required) */
func (r *UpdateCollectInfoRequest) SetEnabled(enabled bool) {
    r.Enabled = enabled
}

/* param filterEnabled: 过滤器是否启用。当appcode为custom时必填(Optional) */
func (r *UpdateCollectInfoRequest) SetFilterEnabled(filterEnabled bool) {
    r.FilterEnabled = &filterEnabled
}

/* param logCustomTarget: 自定义日志转发目的地, 只支持业务应用日志。支持类型："kafka"，"es"(Optional) */
func (r *UpdateCollectInfoRequest) SetLogCustomTarget(logCustomTarget string) {
    r.LogCustomTarget = &logCustomTarget
}

/* param logCustomTargetConf: 自定义日志转发目的地配置，KV 结构，具体配置参考 LogCustomTargetKafkaConf 和 LogCustomTargetEsConf(Optional) */
func (r *UpdateCollectInfoRequest) SetLogCustomTargetConf(logCustomTargetConf interface{}) {
    r.LogCustomTargetConf = &logCustomTargetConf
}

/* param logFile: 日志文件名。当appcode为custom时为必填。日志文件名支持正则表达式。(Optional) */
func (r *UpdateCollectInfoRequest) SetLogFile(logFile string) {
    r.LogFile = &logFile
}

/* param logFilters: 过滤器。设置过滤器后可根据用户设定的关键词采集部分日志，如仅采集 Error 的日志。目前最大允许5个。(Optional) */
func (r *UpdateCollectInfoRequest) SetLogFilters(logFilters []string) {
    r.LogFilters = logFilters
}

/* param logPath: 日志路径。当appcode为custom时为必填。目前仅支持对 Linux 云主机上的日志进行采集，路径支持通配符“*”和“？”，文件路径应符合 Linux 的文件路径规则(Optional) */
func (r *UpdateCollectInfoRequest) SetLogPath(logPath string) {
    r.LogPath = &logPath
}

/* param logtopicEnabled: 目的地是否是日志服务logtopic，只支持业务应用日志(Optional) */
func (r *UpdateCollectInfoRequest) SetLogtopicEnabled(logtopicEnabled bool) {
    r.LogtopicEnabled = &logtopicEnabled
}

/* param regexpStr: 首行正则(Optional) */
func (r *UpdateCollectInfoRequest) SetRegexpStr(regexpStr string) {
    r.RegexpStr = &regexpStr
}

/* param resourceMode: 采集资源时选择的模式，1.正常的选择实例模式（默认模式）；2.选择标签tag模式 3.选择高可用组ag模式(Optional) */
func (r *UpdateCollectInfoRequest) SetResourceMode(resourceMode int64) {
    r.ResourceMode = &resourceMode
}

/* param resourceType: 采集实例类型, 只能是 all/part  当选择all时，传入的实例列表无效(Required) */
func (r *UpdateCollectInfoRequest) SetResourceType(resourceType string) {
    r.ResourceType = resourceType
}

/* param resources: 采集实例列表（存在上限限制）(Optional) */
func (r *UpdateCollectInfoRequest) SetResources(resources []logs.Resource) {
    r.Resources = resources
}

/* param tagResource: (Optional) */
func (r *UpdateCollectInfoRequest) SetTagResource(tagResource *logs.TagResource) {
    r.TagResource = tagResource
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateCollectInfoRequest) GetRegionId() string {
    return r.RegionId
}

type UpdateCollectInfoResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateCollectInfoResult `json:"result"`
}

type UpdateCollectInfoResult struct {
}