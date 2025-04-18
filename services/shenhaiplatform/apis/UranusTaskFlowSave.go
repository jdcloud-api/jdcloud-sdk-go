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

type UranusTaskFlowSaveRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 工作流名称  */
    FlowName string `json:"flowName"`

    /* 工作流描述 (Optional) */
    FlowDesc *string `json:"flowDesc"`

    /* 工作流所属目录  */
    CatalogCode string `json:"catalogCode"`

    /* 工作流协同人 (Optional) */
    Workers []string `json:"workers"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param flowName: 工作流名称 (Required)
 * param catalogCode: 工作流所属目录 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUranusTaskFlowSaveRequest(
    regionId string,
    appName string,
    flowName string,
    catalogCode string,
) *UranusTaskFlowSaveRequest {

	return &UranusTaskFlowSaveRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/uranusTaskFlowSave",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        AppName: appName,
        FlowName: flowName,
        CatalogCode: catalogCode,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param flowName: 工作流名称 (Required)
 * param flowDesc: 工作流描述 (Optional)
 * param catalogCode: 工作流所属目录 (Required)
 * param workers: 工作流协同人 (Optional)
 */
func NewUranusTaskFlowSaveRequestWithAllParams(
    regionId string,
    appName string,
    flowName string,
    flowDesc *string,
    catalogCode string,
    workers []string,
) *UranusTaskFlowSaveRequest {

    return &UranusTaskFlowSaveRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/uranusTaskFlowSave",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
        FlowName: flowName,
        FlowDesc: flowDesc,
        CatalogCode: catalogCode,
        Workers: workers,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUranusTaskFlowSaveRequestWithoutParam() *UranusTaskFlowSaveRequest {

    return &UranusTaskFlowSaveRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/uranusTaskFlowSave",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *UranusTaskFlowSaveRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *UranusTaskFlowSaveRequest) SetAppName(appName string) {
    r.AppName = appName
}
/* param flowName: 工作流名称(Required) */
func (r *UranusTaskFlowSaveRequest) SetFlowName(flowName string) {
    r.FlowName = flowName
}
/* param flowDesc: 工作流描述(Optional) */
func (r *UranusTaskFlowSaveRequest) SetFlowDesc(flowDesc string) {
    r.FlowDesc = &flowDesc
}
/* param catalogCode: 工作流所属目录(Required) */
func (r *UranusTaskFlowSaveRequest) SetCatalogCode(catalogCode string) {
    r.CatalogCode = catalogCode
}
/* param workers: 工作流协同人(Optional) */
func (r *UranusTaskFlowSaveRequest) SetWorkers(workers []string) {
    r.Workers = workers
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UranusTaskFlowSaveRequest) GetRegionId() string {
    return r.RegionId
}

type UranusTaskFlowSaveResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UranusTaskFlowSaveResult `json:"result"`
}

type UranusTaskFlowSaveResult struct {
    Code string `json:"code"`
    ErrorTitle string `json:"errorTitle"`
    ErrorMsg string `json:"errorMsg"`
    Result string `json:"result"`
    SubCode string `json:"subCode"`
    Successed bool `json:"successed"`
}