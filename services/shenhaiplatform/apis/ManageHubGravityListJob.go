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
    shenhaiplatform "github.com/jdcloud-api/jdcloud-sdk-go/services/shenhaiplatform/models"
)

type ManageHubGravityListJobRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 分页-页码 (Optional) */
    PageNum *int `json:"pageNum"`

    /* 分页-每页数量 (Optional) */
    PageSize *int `json:"pageSize"`

    /* 任务名称 (Optional) */
    CstJobName *string `json:"cstJobName"`

    /* 任务编码 (Optional) */
    JobName *string `json:"jobName"`

    /* 租户编码 (Optional) */
    CompanyCode *string `json:"companyCode"`

    /* 负责人pin (Optional) */
    ManagerPin *string `json:"managerPin"`

    /* 工作空间编码 (Optional) */
    WorkspaceCode *string `json:"workspaceCode"`

    /* 任务创建-开始日期: yyyy-MM-dd (Optional) */
    CreateTimeAfter *string `json:"createTimeAfter"`

    /* 任务创建-结束日期: yyyy-MM-dd (Optional) */
    CreateTimeBefore *string `json:"createTimeBefore"`

    /* 任务更新-开始日期: yyyy-MM-dd (Optional) */
    UpdateTimeAfter *string `json:"updateTimeAfter"`

    /* 任务更新-结束日期: yyyy-MM-dd (Optional) */
    UpdateTimeBefore *string `json:"updateTimeBefore"`

    /* 任务运行-开始时间: yyyy-MM-dd HH:mm (Optional) */
    JobStartRunTime *string `json:"jobStartRunTime"`

    /* 任务运行-结束时间: yyyy-MM-dd HH:mm (Optional) */
    JobEndRunTime *string `json:"jobEndRunTime"`

    /* 任务上下线状态: 1-上线 2-下线 (Optional) */
    Enable *int `json:"enable"`

    /* 任务周期 (Optional) */
    Cycle *string `json:"cycle"`

    /* 任务父类型 (Optional) */
    JobType *string `json:"jobType"`

    /* 任务子类型 (Optional) */
    JobChildType *string `json:"jobChildType"`

    /* 任务最后状态: not,Done,Pending,Running,Failed,Timeout,Clean,Ready (Optional) */
    LastJobStatus []string `json:"lastJobStatus"`

    /* 任务最后数据日期 (Optional) */
    LastTxdate *string `json:"lastTxdate"`

    /* 任务负责人 (Optional) */
    Manager *string `json:"manager"`

    /* 过滤模式: exact-精确过滤，否则为模糊过滤 (Optional) */
    QueryMode *string `json:"queryMode"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewManageHubGravityListJobRequest(
    regionId string,
    appName string,
) *ManageHubGravityListJobRequest {

	return &ManageHubGravityListJobRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/manageHubGravityListJob",
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
 * param pageNum: 分页-页码 (Optional)
 * param pageSize: 分页-每页数量 (Optional)
 * param cstJobName: 任务名称 (Optional)
 * param jobName: 任务编码 (Optional)
 * param companyCode: 租户编码 (Optional)
 * param managerPin: 负责人pin (Optional)
 * param workspaceCode: 工作空间编码 (Optional)
 * param createTimeAfter: 任务创建-开始日期: yyyy-MM-dd (Optional)
 * param createTimeBefore: 任务创建-结束日期: yyyy-MM-dd (Optional)
 * param updateTimeAfter: 任务更新-开始日期: yyyy-MM-dd (Optional)
 * param updateTimeBefore: 任务更新-结束日期: yyyy-MM-dd (Optional)
 * param jobStartRunTime: 任务运行-开始时间: yyyy-MM-dd HH:mm (Optional)
 * param jobEndRunTime: 任务运行-结束时间: yyyy-MM-dd HH:mm (Optional)
 * param enable: 任务上下线状态: 1-上线 2-下线 (Optional)
 * param cycle: 任务周期 (Optional)
 * param jobType: 任务父类型 (Optional)
 * param jobChildType: 任务子类型 (Optional)
 * param lastJobStatus: 任务最后状态: not,Done,Pending,Running,Failed,Timeout,Clean,Ready (Optional)
 * param lastTxdate: 任务最后数据日期 (Optional)
 * param manager: 任务负责人 (Optional)
 * param queryMode: 过滤模式: exact-精确过滤，否则为模糊过滤 (Optional)
 */
func NewManageHubGravityListJobRequestWithAllParams(
    regionId string,
    appName string,
    pageNum *int,
    pageSize *int,
    cstJobName *string,
    jobName *string,
    companyCode *string,
    managerPin *string,
    workspaceCode *string,
    createTimeAfter *string,
    createTimeBefore *string,
    updateTimeAfter *string,
    updateTimeBefore *string,
    jobStartRunTime *string,
    jobEndRunTime *string,
    enable *int,
    cycle *string,
    jobType *string,
    jobChildType *string,
    lastJobStatus []string,
    lastTxdate *string,
    manager *string,
    queryMode *string,
) *ManageHubGravityListJobRequest {

    return &ManageHubGravityListJobRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/manageHubGravityListJob",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
        PageNum: pageNum,
        PageSize: pageSize,
        CstJobName: cstJobName,
        JobName: jobName,
        CompanyCode: companyCode,
        ManagerPin: managerPin,
        WorkspaceCode: workspaceCode,
        CreateTimeAfter: createTimeAfter,
        CreateTimeBefore: createTimeBefore,
        UpdateTimeAfter: updateTimeAfter,
        UpdateTimeBefore: updateTimeBefore,
        JobStartRunTime: jobStartRunTime,
        JobEndRunTime: jobEndRunTime,
        Enable: enable,
        Cycle: cycle,
        JobType: jobType,
        JobChildType: jobChildType,
        LastJobStatus: lastJobStatus,
        LastTxdate: lastTxdate,
        Manager: manager,
        QueryMode: queryMode,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewManageHubGravityListJobRequestWithoutParam() *ManageHubGravityListJobRequest {

    return &ManageHubGravityListJobRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/manageHubGravityListJob",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *ManageHubGravityListJobRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *ManageHubGravityListJobRequest) SetAppName(appName string) {
    r.AppName = appName
}
/* param pageNum: 分页-页码(Optional) */
func (r *ManageHubGravityListJobRequest) SetPageNum(pageNum int) {
    r.PageNum = &pageNum
}
/* param pageSize: 分页-每页数量(Optional) */
func (r *ManageHubGravityListJobRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}
/* param cstJobName: 任务名称(Optional) */
func (r *ManageHubGravityListJobRequest) SetCstJobName(cstJobName string) {
    r.CstJobName = &cstJobName
}
/* param jobName: 任务编码(Optional) */
func (r *ManageHubGravityListJobRequest) SetJobName(jobName string) {
    r.JobName = &jobName
}
/* param companyCode: 租户编码(Optional) */
func (r *ManageHubGravityListJobRequest) SetCompanyCode(companyCode string) {
    r.CompanyCode = &companyCode
}
/* param managerPin: 负责人pin(Optional) */
func (r *ManageHubGravityListJobRequest) SetManagerPin(managerPin string) {
    r.ManagerPin = &managerPin
}
/* param workspaceCode: 工作空间编码(Optional) */
func (r *ManageHubGravityListJobRequest) SetWorkspaceCode(workspaceCode string) {
    r.WorkspaceCode = &workspaceCode
}
/* param createTimeAfter: 任务创建-开始日期: yyyy-MM-dd(Optional) */
func (r *ManageHubGravityListJobRequest) SetCreateTimeAfter(createTimeAfter string) {
    r.CreateTimeAfter = &createTimeAfter
}
/* param createTimeBefore: 任务创建-结束日期: yyyy-MM-dd(Optional) */
func (r *ManageHubGravityListJobRequest) SetCreateTimeBefore(createTimeBefore string) {
    r.CreateTimeBefore = &createTimeBefore
}
/* param updateTimeAfter: 任务更新-开始日期: yyyy-MM-dd(Optional) */
func (r *ManageHubGravityListJobRequest) SetUpdateTimeAfter(updateTimeAfter string) {
    r.UpdateTimeAfter = &updateTimeAfter
}
/* param updateTimeBefore: 任务更新-结束日期: yyyy-MM-dd(Optional) */
func (r *ManageHubGravityListJobRequest) SetUpdateTimeBefore(updateTimeBefore string) {
    r.UpdateTimeBefore = &updateTimeBefore
}
/* param jobStartRunTime: 任务运行-开始时间: yyyy-MM-dd HH:mm(Optional) */
func (r *ManageHubGravityListJobRequest) SetJobStartRunTime(jobStartRunTime string) {
    r.JobStartRunTime = &jobStartRunTime
}
/* param jobEndRunTime: 任务运行-结束时间: yyyy-MM-dd HH:mm(Optional) */
func (r *ManageHubGravityListJobRequest) SetJobEndRunTime(jobEndRunTime string) {
    r.JobEndRunTime = &jobEndRunTime
}
/* param enable: 任务上下线状态: 1-上线 2-下线(Optional) */
func (r *ManageHubGravityListJobRequest) SetEnable(enable int) {
    r.Enable = &enable
}
/* param cycle: 任务周期(Optional) */
func (r *ManageHubGravityListJobRequest) SetCycle(cycle string) {
    r.Cycle = &cycle
}
/* param jobType: 任务父类型(Optional) */
func (r *ManageHubGravityListJobRequest) SetJobType(jobType string) {
    r.JobType = &jobType
}
/* param jobChildType: 任务子类型(Optional) */
func (r *ManageHubGravityListJobRequest) SetJobChildType(jobChildType string) {
    r.JobChildType = &jobChildType
}
/* param lastJobStatus: 任务最后状态: not,Done,Pending,Running,Failed,Timeout,Clean,Ready(Optional) */
func (r *ManageHubGravityListJobRequest) SetLastJobStatus(lastJobStatus []string) {
    r.LastJobStatus = lastJobStatus
}
/* param lastTxdate: 任务最后数据日期(Optional) */
func (r *ManageHubGravityListJobRequest) SetLastTxdate(lastTxdate string) {
    r.LastTxdate = &lastTxdate
}
/* param manager: 任务负责人(Optional) */
func (r *ManageHubGravityListJobRequest) SetManager(manager string) {
    r.Manager = &manager
}
/* param queryMode: 过滤模式: exact-精确过滤，否则为模糊过滤(Optional) */
func (r *ManageHubGravityListJobRequest) SetQueryMode(queryMode string) {
    r.QueryMode = &queryMode
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ManageHubGravityListJobRequest) GetRegionId() string {
    return r.RegionId
}

type ManageHubGravityListJobResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ManageHubGravityListJobResult `json:"result"`
}

type ManageHubGravityListJobResult struct {
    Code string `json:"code"`
    Message string `json:"message"`
    Result shenhaiplatform.PageVoJobInfoVo `json:"result"`
}