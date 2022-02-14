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

type ValidateParserRequest struct {

    core.JDCloudRequest

    /* 解析类型。oneline - 单行，split - 分割， json - json， regexp - regexp  */
    ParserMode string `json:"parserMode"`

    /* 解析语法 (Optional) */
    ParserPattern *string `json:"parserPattern"`

    /* 日志样例 (Optional) */
    ParserSample *string `json:"parserSample"`
}

/*
 * param parserMode: 解析类型。oneline - 单行，split - 分割， json - json， regexp - regexp (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewValidateParserRequest(
    parserMode string,
) *ValidateParserRequest {

	return &ValidateParserRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/validateParser",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        ParserMode: parserMode,
	}
}

/*
 * param parserMode: 解析类型。oneline - 单行，split - 分割， json - json， regexp - regexp (Required)
 * param parserPattern: 解析语法 (Optional)
 * param parserSample: 日志样例 (Optional)
 */
func NewValidateParserRequestWithAllParams(
    parserMode string,
    parserPattern *string,
    parserSample *string,
) *ValidateParserRequest {

    return &ValidateParserRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/validateParser",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        ParserMode: parserMode,
        ParserPattern: parserPattern,
        ParserSample: parserSample,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewValidateParserRequestWithoutParam() *ValidateParserRequest {

    return &ValidateParserRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/validateParser",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param parserMode: 解析类型。oneline - 单行，split - 分割， json - json， regexp - regexp(Required) */
func (r *ValidateParserRequest) SetParserMode(parserMode string) {
    r.ParserMode = parserMode
}

/* param parserPattern: 解析语法(Optional) */
func (r *ValidateParserRequest) SetParserPattern(parserPattern string) {
    r.ParserPattern = &parserPattern
}

/* param parserSample: 日志样例(Optional) */
func (r *ValidateParserRequest) SetParserSample(parserSample string) {
    r.ParserSample = &parserSample
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ValidateParserRequest) GetRegionId() string {
    return ""
}

type ValidateParserResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ValidateParserResult `json:"result"`
}

type ValidateParserResult struct {
}