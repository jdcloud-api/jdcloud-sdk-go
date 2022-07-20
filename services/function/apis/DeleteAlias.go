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
)

type DeleteAliasRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 别名所属函数名称  */
    FunctionName string `json:"functionName"`

    /* 别名名称  */
    AliasName string `json:"aliasName"`
}

/*
 * param regionId: Region ID (Required)
 * param functionName: 别名所属函数名称 (Required)
 * param aliasName: 别名名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteAliasRequest(
    regionId string,
    functionName string,
    aliasName string,
) *DeleteAliasRequest {

	return &DeleteAliasRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/functions/{functionName}/aliases/{aliasName}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        FunctionName: functionName,
        AliasName: aliasName,
	}
}

/*
 * param regionId: Region ID (Required)
 * param functionName: 别名所属函数名称 (Required)
 * param aliasName: 别名名称 (Required)
 */
func NewDeleteAliasRequestWithAllParams(
    regionId string,
    functionName string,
    aliasName string,
) *DeleteAliasRequest {

    return &DeleteAliasRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/functions/{functionName}/aliases/{aliasName}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        FunctionName: functionName,
        AliasName: aliasName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteAliasRequestWithoutParam() *DeleteAliasRequest {

    return &DeleteAliasRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/functions/{functionName}/aliases/{aliasName}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DeleteAliasRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param functionName: 别名所属函数名称(Required) */
func (r *DeleteAliasRequest) SetFunctionName(functionName string) {
    r.FunctionName = functionName
}

/* param aliasName: 别名名称(Required) */
func (r *DeleteAliasRequest) SetAliasName(aliasName string) {
    r.AliasName = aliasName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteAliasRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteAliasResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteAliasResult `json:"result"`
}

type DeleteAliasResult struct {
}