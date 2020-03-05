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

type DeployAppRequest struct {

    core.JDCloudRequest

    /* 设备归属的实例ID  */
    InstanceId string `json:"instanceId"`

    /* 设备归属的实例所在区域  */
    RegionId string `json:"regionId"`

    /* 硬件版本  */
    HardwareId string `json:"hardwareId"`

    /* OSID  */
    OsId string `json:"osId"`

    /* App内部名称  */
    AppName string `json:"appName"`

    /* App版本  */
    AppVersion string `json:"appVersion"`

    /* edge名称  */
    EdgeName string `json:"edgeName"`

    /* App安装变量（默认为空字符串） (Optional) */
    Env *string `json:"env"`

    /* 是否特权模式运行(0-否 1-是)  */
    Privileged int `json:"privileged"`

    /* 硬盘操作卷 (Optional) */
    Volume *string `json:"volume"`

    /* App类型(1-设备服务 2-边缘应用)  */
    AppType int `json:"appType"`

    /* 是否使用Edge系统的Host配置(0-否 1-是)  */
    UseHost int `json:"useHost"`

    /* CPU权重（低-128 中-256 高-512）  */
    CpuShares int `json:"cpuShares"`

    /* 内存限制（0-无限 低-128 中-256 高-512）  */
    MemoryLimit string `json:"memoryLimit"`

    /* 重启策略(0-never 1-always 2-onFailure)  */
    RestartPolicy string `json:"restartPolicy"`

    /* 部署APP状态(3-更新 其余都是安装)  */
    DeployAppStatus int `json:"deployAppStatus"`

    /* App状态 (Optional) */
    AppStatus *int `json:"appStatus"`
}

/*
 * param instanceId: 设备归属的实例ID (Required)
 * param regionId: 设备归属的实例所在区域 (Required)
 * param hardwareId: 硬件版本 (Required)
 * param osId: OSID (Required)
 * param appName: App内部名称 (Required)
 * param appVersion: App版本 (Required)
 * param edgeName: edge名称 (Required)
 * param privileged: 是否特权模式运行(0-否 1-是) (Required)
 * param appType: App类型(1-设备服务 2-边缘应用) (Required)
 * param useHost: 是否使用Edge系统的Host配置(0-否 1-是) (Required)
 * param cpuShares: CPU权重（低-128 中-256 高-512） (Required)
 * param memoryLimit: 内存限制（0-无限 低-128 中-256 高-512） (Required)
 * param restartPolicy: 重启策略(0-never 1-always 2-onFailure) (Required)
 * param deployAppStatus: 部署APP状态(3-更新 其余都是安装) (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeployAppRequest(
    instanceId string,
    regionId string,
    hardwareId string,
    osId string,
    appName string,
    appVersion string,
    edgeName string,
    privileged int,
    appType int,
    useHost int,
    cpuShares int,
    memoryLimit string,
    restartPolicy string,
    deployAppStatus int,
) *DeployAppRequest {

	return &DeployAppRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/hardwareId/{hardwareId}/os/{osId}/edgeApp:deployApp",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        InstanceId: instanceId,
        RegionId: regionId,
        HardwareId: hardwareId,
        OsId: osId,
        AppName: appName,
        AppVersion: appVersion,
        EdgeName: edgeName,
        Privileged: privileged,
        AppType: appType,
        UseHost: useHost,
        CpuShares: cpuShares,
        MemoryLimit: memoryLimit,
        RestartPolicy: restartPolicy,
        DeployAppStatus: deployAppStatus,
	}
}

/*
 * param instanceId: 设备归属的实例ID (Required)
 * param regionId: 设备归属的实例所在区域 (Required)
 * param hardwareId: 硬件版本 (Required)
 * param osId: OSID (Required)
 * param appName: App内部名称 (Required)
 * param appVersion: App版本 (Required)
 * param edgeName: edge名称 (Required)
 * param env: App安装变量（默认为空字符串） (Optional)
 * param privileged: 是否特权模式运行(0-否 1-是) (Required)
 * param volume: 硬盘操作卷 (Optional)
 * param appType: App类型(1-设备服务 2-边缘应用) (Required)
 * param useHost: 是否使用Edge系统的Host配置(0-否 1-是) (Required)
 * param cpuShares: CPU权重（低-128 中-256 高-512） (Required)
 * param memoryLimit: 内存限制（0-无限 低-128 中-256 高-512） (Required)
 * param restartPolicy: 重启策略(0-never 1-always 2-onFailure) (Required)
 * param deployAppStatus: 部署APP状态(3-更新 其余都是安装) (Required)
 * param appStatus: App状态 (Optional)
 */
func NewDeployAppRequestWithAllParams(
    instanceId string,
    regionId string,
    hardwareId string,
    osId string,
    appName string,
    appVersion string,
    edgeName string,
    env *string,
    privileged int,
    volume *string,
    appType int,
    useHost int,
    cpuShares int,
    memoryLimit string,
    restartPolicy string,
    deployAppStatus int,
    appStatus *int,
) *DeployAppRequest {

    return &DeployAppRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/hardwareId/{hardwareId}/os/{osId}/edgeApp:deployApp",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        InstanceId: instanceId,
        RegionId: regionId,
        HardwareId: hardwareId,
        OsId: osId,
        AppName: appName,
        AppVersion: appVersion,
        EdgeName: edgeName,
        Env: env,
        Privileged: privileged,
        Volume: volume,
        AppType: appType,
        UseHost: useHost,
        CpuShares: cpuShares,
        MemoryLimit: memoryLimit,
        RestartPolicy: restartPolicy,
        DeployAppStatus: deployAppStatus,
        AppStatus: appStatus,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeployAppRequestWithoutParam() *DeployAppRequest {

    return &DeployAppRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/hardwareId/{hardwareId}/os/{osId}/edgeApp:deployApp",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param instanceId: 设备归属的实例ID(Required) */
func (r *DeployAppRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param regionId: 设备归属的实例所在区域(Required) */
func (r *DeployAppRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param hardwareId: 硬件版本(Required) */
func (r *DeployAppRequest) SetHardwareId(hardwareId string) {
    r.HardwareId = hardwareId
}

/* param osId: OSID(Required) */
func (r *DeployAppRequest) SetOsId(osId string) {
    r.OsId = osId
}

/* param appName: App内部名称(Required) */
func (r *DeployAppRequest) SetAppName(appName string) {
    r.AppName = appName
}

/* param appVersion: App版本(Required) */
func (r *DeployAppRequest) SetAppVersion(appVersion string) {
    r.AppVersion = appVersion
}

/* param edgeName: edge名称(Required) */
func (r *DeployAppRequest) SetEdgeName(edgeName string) {
    r.EdgeName = edgeName
}

/* param env: App安装变量（默认为空字符串）(Optional) */
func (r *DeployAppRequest) SetEnv(env string) {
    r.Env = &env
}

/* param privileged: 是否特权模式运行(0-否 1-是)(Required) */
func (r *DeployAppRequest) SetPrivileged(privileged int) {
    r.Privileged = privileged
}

/* param volume: 硬盘操作卷(Optional) */
func (r *DeployAppRequest) SetVolume(volume string) {
    r.Volume = &volume
}

/* param appType: App类型(1-设备服务 2-边缘应用)(Required) */
func (r *DeployAppRequest) SetAppType(appType int) {
    r.AppType = appType
}

/* param useHost: 是否使用Edge系统的Host配置(0-否 1-是)(Required) */
func (r *DeployAppRequest) SetUseHost(useHost int) {
    r.UseHost = useHost
}

/* param cpuShares: CPU权重（低-128 中-256 高-512）(Required) */
func (r *DeployAppRequest) SetCpuShares(cpuShares int) {
    r.CpuShares = cpuShares
}

/* param memoryLimit: 内存限制（0-无限 低-128 中-256 高-512）(Required) */
func (r *DeployAppRequest) SetMemoryLimit(memoryLimit string) {
    r.MemoryLimit = memoryLimit
}

/* param restartPolicy: 重启策略(0-never 1-always 2-onFailure)(Required) */
func (r *DeployAppRequest) SetRestartPolicy(restartPolicy string) {
    r.RestartPolicy = restartPolicy
}

/* param deployAppStatus: 部署APP状态(3-更新 其余都是安装)(Required) */
func (r *DeployAppRequest) SetDeployAppStatus(deployAppStatus int) {
    r.DeployAppStatus = deployAppStatus
}

/* param appStatus: App状态(Optional) */
func (r *DeployAppRequest) SetAppStatus(appStatus int) {
    r.AppStatus = &appStatus
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeployAppRequest) GetRegionId() string {
    return r.RegionId
}

type DeployAppResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeployAppResult `json:"result"`
}

type DeployAppResult struct {
    DeployStatus int `json:"deployStatus"`
}