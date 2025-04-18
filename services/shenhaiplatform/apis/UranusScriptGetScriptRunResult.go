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

type UranusScriptGetScriptRunResultRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 文件名称 taskId_${结果集数字}  第一个结果就是：1_1,第二个结果集就是：1_2 (Optional) */
    FileName *string `json:"fileName"`

    /* 文件业务code (Optional) */
    FileCode *string `json:"fileCode"`

    /* 文件偏移量 (Optional) */
    Pos *int64 `json:"pos"`

    /* 读取结果条数 (Optional) */
    Lines *int64 `json:"lines"`

    /* 运行ID (Optional) */
    HistoryId *int `json:"historyId"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUranusScriptGetScriptRunResultRequest(
    regionId string,
    appName string,
) *UranusScriptGetScriptRunResultRequest {

	return &UranusScriptGetScriptRunResultRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/uranusScriptGetScriptRunResult",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        AppName: appName,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param fileName: 文件名称 taskId_${结果集数字}  第一个结果就是：1_1,第二个结果集就是：1_2 (Optional)
 * param fileCode: 文件业务code (Optional)
 * param pos: 文件偏移量 (Optional)
 * param lines: 读取结果条数 (Optional)
 * param historyId: 运行ID (Optional)
 */
func NewUranusScriptGetScriptRunResultRequestWithAllParams(
    regionId string,
    appName string,
    fileName *string,
    fileCode *string,
    pos *int64,
    lines *int64,
    historyId *int,
) *UranusScriptGetScriptRunResultRequest {

    return &UranusScriptGetScriptRunResultRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/uranusScriptGetScriptRunResult",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
        FileName: fileName,
        FileCode: fileCode,
        Pos: pos,
        Lines: lines,
        HistoryId: historyId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUranusScriptGetScriptRunResultRequestWithoutParam() *UranusScriptGetScriptRunResultRequest {

    return &UranusScriptGetScriptRunResultRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/uranusScriptGetScriptRunResult",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *UranusScriptGetScriptRunResultRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *UranusScriptGetScriptRunResultRequest) SetAppName(appName string) {
    r.AppName = appName
}
/* param fileName: 文件名称 taskId_${结果集数字}  第一个结果就是：1_1,第二个结果集就是：1_2(Optional) */
func (r *UranusScriptGetScriptRunResultRequest) SetFileName(fileName string) {
    r.FileName = &fileName
}
/* param fileCode: 文件业务code(Optional) */
func (r *UranusScriptGetScriptRunResultRequest) SetFileCode(fileCode string) {
    r.FileCode = &fileCode
}
/* param pos: 文件偏移量(Optional) */
func (r *UranusScriptGetScriptRunResultRequest) SetPos(pos int64) {
    r.Pos = &pos
}
/* param lines: 读取结果条数(Optional) */
func (r *UranusScriptGetScriptRunResultRequest) SetLines(lines int64) {
    r.Lines = &lines
}
/* param historyId: 运行ID(Optional) */
func (r *UranusScriptGetScriptRunResultRequest) SetHistoryId(historyId int) {
    r.HistoryId = &historyId
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UranusScriptGetScriptRunResultRequest) GetRegionId() string {
    return r.RegionId
}

type UranusScriptGetScriptRunResultResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UranusScriptGetScriptRunResultResult `json:"result"`
}

type UranusScriptGetScriptRunResultResult struct {
    Code string `json:"code"`
    ErrorTitle string `json:"errorTitle"`
    ErrorMsg string `json:"errorMsg"`
    Result interface{} `json:"result"`
    SubCode string `json:"subCode"`
    Successed bool `json:"successed"`
}