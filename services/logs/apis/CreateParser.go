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

type CreateParserRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 日志主题 UID  */
    LogtopicUID string `json:"logtopicUID"`

    /* 行级索引分词符。 Rune 数组 (Optional) */
    IndexToken []string `json:"indexToken"`

    /*   */
    ParserFields []logs.ParserField `json:"parserFields"`

    /* 解析类型。oneline - 单行，split - 分割， json - json， regexp - regexp  */
    ParserMode string `json:"parserMode"`

    /* 解析语法 (Optional) */
    ParserPattern *string `json:"parserPattern"`

    /* 日志样例 (Optional) */
    ParserSample *string `json:"parserSample"`

    /* 预处理任务列表。按照数组的顺序执行。 (Optional) */
    Pipelines []logs.PipelineSpec `json:"pipelines"`

    /* 是否保留原始内容字段 (Optional) */
    ReserveOriginContent *bool `json:"reserveOriginContent"`
}

/*
 * param regionId: 地域 Id (Required)
 * param logtopicUID: 日志主题 UID (Required)
 * param parserFields:  (Required)
 * param parserMode: 解析类型。oneline - 单行，split - 分割， json - json， regexp - regexp (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateParserRequest(
    regionId string,
    logtopicUID string,
    parserFields []logs.ParserField,
    parserMode string,
) *CreateParserRequest {

	return &CreateParserRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/logtopics/{logtopicUID}/createParser",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        LogtopicUID: logtopicUID,
        ParserFields: parserFields,
        ParserMode: parserMode,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param logtopicUID: 日志主题 UID (Required)
 * param indexToken: 行级索引分词符。 Rune 数组 (Optional)
 * param parserFields:  (Required)
 * param parserMode: 解析类型。oneline - 单行，split - 分割， json - json， regexp - regexp (Required)
 * param parserPattern: 解析语法 (Optional)
 * param parserSample: 日志样例 (Optional)
 * param pipelines: 预处理任务列表。按照数组的顺序执行。 (Optional)
 * param reserveOriginContent: 是否保留原始内容字段 (Optional)
 */
func NewCreateParserRequestWithAllParams(
    regionId string,
    logtopicUID string,
    indexToken []string,
    parserFields []logs.ParserField,
    parserMode string,
    parserPattern *string,
    parserSample *string,
    pipelines []logs.PipelineSpec,
    reserveOriginContent *bool,
) *CreateParserRequest {

    return &CreateParserRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/logtopics/{logtopicUID}/createParser",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        LogtopicUID: logtopicUID,
        IndexToken: indexToken,
        ParserFields: parserFields,
        ParserMode: parserMode,
        ParserPattern: parserPattern,
        ParserSample: parserSample,
        Pipelines: pipelines,
        ReserveOriginContent: reserveOriginContent,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateParserRequestWithoutParam() *CreateParserRequest {

    return &CreateParserRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/logtopics/{logtopicUID}/createParser",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *CreateParserRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param logtopicUID: 日志主题 UID(Required) */
func (r *CreateParserRequest) SetLogtopicUID(logtopicUID string) {
    r.LogtopicUID = logtopicUID
}
/* param indexToken: 行级索引分词符。 Rune 数组(Optional) */
func (r *CreateParserRequest) SetIndexToken(indexToken []string) {
    r.IndexToken = indexToken
}
/* param parserFields: (Required) */
func (r *CreateParserRequest) SetParserFields(parserFields []logs.ParserField) {
    r.ParserFields = parserFields
}
/* param parserMode: 解析类型。oneline - 单行，split - 分割， json - json， regexp - regexp(Required) */
func (r *CreateParserRequest) SetParserMode(parserMode string) {
    r.ParserMode = parserMode
}
/* param parserPattern: 解析语法(Optional) */
func (r *CreateParserRequest) SetParserPattern(parserPattern string) {
    r.ParserPattern = &parserPattern
}
/* param parserSample: 日志样例(Optional) */
func (r *CreateParserRequest) SetParserSample(parserSample string) {
    r.ParserSample = &parserSample
}
/* param pipelines: 预处理任务列表。按照数组的顺序执行。(Optional) */
func (r *CreateParserRequest) SetPipelines(pipelines []logs.PipelineSpec) {
    r.Pipelines = pipelines
}
/* param reserveOriginContent: 是否保留原始内容字段(Optional) */
func (r *CreateParserRequest) SetReserveOriginContent(reserveOriginContent bool) {
    r.ReserveOriginContent = &reserveOriginContent
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateParserRequest) GetRegionId() string {
    return r.RegionId
}

type CreateParserResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateParserResult `json:"result"`
}

type CreateParserResult struct {
}