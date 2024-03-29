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
    flowlog "github.com/jdcloud-api/jdcloud-sdk-go/services/flowlog/models"
)

type CreateFlowLogRequest struct {

    core.JDCloudRequest

    /* 地域 ID  */
    RegionId string `json:"regionId"`

    /* 流日志名称，只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符  */
    FlowLogName string `json:"flowLogName"`

    /* 描述,允许输入UTF-8编码下的全部字符，不超过256字符 (Optional) */
    Description *string `json:"description"`

    /* 流日志类型
PORT：采集资源可为云主机、弹性网卡
  */
    FlowLogType string `json:"flowLogType"`

    /* 采集资源列表 (Optional) */
    CollectResources []flowlog.CollectResourceSpec `json:"collectResources"`

    /* 采集流量类型
ALL：记录指定资源的全部流量
ACCEPT：记录指定资源被安全组、网络ACL均接受的流量
REJECT：记录指定资源被安全组或网络ACL拒绝的流量     
 (Optional) */
    CollectTrafficType *string `json:"collectTrafficType"`

    /* 流日志采集时间间隔。单位：分钟。取值：1、5、10 (Optional) */
    CollectInterval *int `json:"collectInterval"`

    /* 流日志的存储服务类型，支持存储到日志服务，日志服务取值：LOG (Optional) */
    StorageType *string `json:"storageType"`

    /* 流日志的存储服务所在地域，如日志服务所属地域，如cn-north-1  */
    StorageRegionId string `json:"storageRegionId"`

    /* 流日志的存储服务ID
若storageType = LOG时，值取日志主题ID，如logtopic-xxxx
当flowLogType = PORT时，值需取 templateUID = eniflowlogs 的日志主题ID
  */
    StorageId string `json:"storageId"`
}

/*
 * param regionId: 地域 ID (Required)
 * param flowLogName: 流日志名称，只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符 (Required)
 * param flowLogType: 流日志类型
PORT：采集资源可为云主机、弹性网卡
 (Required)
 * param storageRegionId: 流日志的存储服务所在地域，如日志服务所属地域，如cn-north-1 (Required)
 * param storageId: 流日志的存储服务ID
若storageType = LOG时，值取日志主题ID，如logtopic-xxxx
当flowLogType = PORT时，值需取 templateUID = eniflowlogs 的日志主题ID
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateFlowLogRequest(
    regionId string,
    flowLogName string,
    flowLogType string,
    storageRegionId string,
    storageId string,
) *CreateFlowLogRequest {

	return &CreateFlowLogRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/flowLogs",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        FlowLogName: flowLogName,
        FlowLogType: flowLogType,
        StorageRegionId: storageRegionId,
        StorageId: storageId,
	}
}

/*
 * param regionId: 地域 ID (Required)
 * param flowLogName: 流日志名称，只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符 (Required)
 * param description: 描述,允许输入UTF-8编码下的全部字符，不超过256字符 (Optional)
 * param flowLogType: 流日志类型
PORT：采集资源可为云主机、弹性网卡
 (Required)
 * param collectResources: 采集资源列表 (Optional)
 * param collectTrafficType: 采集流量类型
ALL：记录指定资源的全部流量
ACCEPT：记录指定资源被安全组、网络ACL均接受的流量
REJECT：记录指定资源被安全组或网络ACL拒绝的流量     
 (Optional)
 * param collectInterval: 流日志采集时间间隔。单位：分钟。取值：1、5、10 (Optional)
 * param storageType: 流日志的存储服务类型，支持存储到日志服务，日志服务取值：LOG (Optional)
 * param storageRegionId: 流日志的存储服务所在地域，如日志服务所属地域，如cn-north-1 (Required)
 * param storageId: 流日志的存储服务ID
若storageType = LOG时，值取日志主题ID，如logtopic-xxxx
当flowLogType = PORT时，值需取 templateUID = eniflowlogs 的日志主题ID
 (Required)
 */
func NewCreateFlowLogRequestWithAllParams(
    regionId string,
    flowLogName string,
    description *string,
    flowLogType string,
    collectResources []flowlog.CollectResourceSpec,
    collectTrafficType *string,
    collectInterval *int,
    storageType *string,
    storageRegionId string,
    storageId string,
) *CreateFlowLogRequest {

    return &CreateFlowLogRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/flowLogs",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        FlowLogName: flowLogName,
        Description: description,
        FlowLogType: flowLogType,
        CollectResources: collectResources,
        CollectTrafficType: collectTrafficType,
        CollectInterval: collectInterval,
        StorageType: storageType,
        StorageRegionId: storageRegionId,
        StorageId: storageId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateFlowLogRequestWithoutParam() *CreateFlowLogRequest {

    return &CreateFlowLogRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/flowLogs",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 ID(Required) */
func (r *CreateFlowLogRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param flowLogName: 流日志名称，只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符(Required) */
func (r *CreateFlowLogRequest) SetFlowLogName(flowLogName string) {
    r.FlowLogName = flowLogName
}
/* param description: 描述,允许输入UTF-8编码下的全部字符，不超过256字符(Optional) */
func (r *CreateFlowLogRequest) SetDescription(description string) {
    r.Description = &description
}
/* param flowLogType: 流日志类型
PORT：采集资源可为云主机、弹性网卡
(Required) */
func (r *CreateFlowLogRequest) SetFlowLogType(flowLogType string) {
    r.FlowLogType = flowLogType
}
/* param collectResources: 采集资源列表(Optional) */
func (r *CreateFlowLogRequest) SetCollectResources(collectResources []flowlog.CollectResourceSpec) {
    r.CollectResources = collectResources
}
/* param collectTrafficType: 采集流量类型
ALL：记录指定资源的全部流量
ACCEPT：记录指定资源被安全组、网络ACL均接受的流量
REJECT：记录指定资源被安全组或网络ACL拒绝的流量     
(Optional) */
func (r *CreateFlowLogRequest) SetCollectTrafficType(collectTrafficType string) {
    r.CollectTrafficType = &collectTrafficType
}
/* param collectInterval: 流日志采集时间间隔。单位：分钟。取值：1、5、10(Optional) */
func (r *CreateFlowLogRequest) SetCollectInterval(collectInterval int) {
    r.CollectInterval = &collectInterval
}
/* param storageType: 流日志的存储服务类型，支持存储到日志服务，日志服务取值：LOG(Optional) */
func (r *CreateFlowLogRequest) SetStorageType(storageType string) {
    r.StorageType = &storageType
}
/* param storageRegionId: 流日志的存储服务所在地域，如日志服务所属地域，如cn-north-1(Required) */
func (r *CreateFlowLogRequest) SetStorageRegionId(storageRegionId string) {
    r.StorageRegionId = storageRegionId
}
/* param storageId: 流日志的存储服务ID
若storageType = LOG时，值取日志主题ID，如logtopic-xxxx
当flowLogType = PORT时，值需取 templateUID = eniflowlogs 的日志主题ID
(Required) */
func (r *CreateFlowLogRequest) SetStorageId(storageId string) {
    r.StorageId = storageId
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateFlowLogRequest) GetRegionId() string {
    return r.RegionId
}

type CreateFlowLogResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateFlowLogResult `json:"result"`
}

type CreateFlowLogResult struct {
    FlowLogId string `json:"flowLogId"`
}