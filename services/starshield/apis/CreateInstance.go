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

type CreateInstanceRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 计费模式（CONFIG、FLOW、MONTHLY、ONCE） (Optional) */
    ChargeMode *string `json:"chargeMode"`

    /* 套餐类型（FREE、BASIC、PROFESSIONAL、ENTERPRISE、ULTIMATE、SMB_BASIC、SMB_BUSINESS） (Optional) */
    PackType *string `json:"packType"`

    /* 域名增量包数量 (Optional) */
    ZonePackNum *int `json:"zonePackNum"`

    /* 计费时长 (Optional) */
    Duration *int `json:"duration"`

    /* 计费时长单位（MONTH、YEAR） (Optional) */
    DurationUnit *string `json:"durationUnit"`

    /* 自动续费状态(OPEN->开通自动续费 CLOSE->关闭自动续费) (Optional) */
    AutoRenewStatus *string `json:"autoRenewStatus"`

    /* 实例名称 (Optional) */
    InstanceName *string `json:"instanceName"`

    /* 备注 (Optional) */
    Memo *string `json:"memo"`

    /* 支付成功返回路径 (Optional) */
    ReturnUrl *string `json:"returnUrl"`
}

/*
 * param regionId: 地域ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateInstanceRequest(
    regionId string,
) *CreateInstanceRequest {

	return &CreateInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param chargeMode: 计费模式（CONFIG、FLOW、MONTHLY、ONCE） (Optional)
 * param packType: 套餐类型（FREE、BASIC、PROFESSIONAL、ENTERPRISE、ULTIMATE、SMB_BASIC、SMB_BUSINESS） (Optional)
 * param zonePackNum: 域名增量包数量 (Optional)
 * param duration: 计费时长 (Optional)
 * param durationUnit: 计费时长单位（MONTH、YEAR） (Optional)
 * param autoRenewStatus: 自动续费状态(OPEN->开通自动续费 CLOSE->关闭自动续费) (Optional)
 * param instanceName: 实例名称 (Optional)
 * param memo: 备注 (Optional)
 * param returnUrl: 支付成功返回路径 (Optional)
 */
func NewCreateInstanceRequestWithAllParams(
    regionId string,
    chargeMode *string,
    packType *string,
    zonePackNum *int,
    duration *int,
    durationUnit *string,
    autoRenewStatus *string,
    instanceName *string,
    memo *string,
    returnUrl *string,
) *CreateInstanceRequest {

    return &CreateInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ChargeMode: chargeMode,
        PackType: packType,
        ZonePackNum: zonePackNum,
        Duration: duration,
        DurationUnit: durationUnit,
        AutoRenewStatus: autoRenewStatus,
        InstanceName: instanceName,
        Memo: memo,
        ReturnUrl: returnUrl,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateInstanceRequestWithoutParam() *CreateInstanceRequest {

    return &CreateInstanceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *CreateInstanceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param chargeMode: 计费模式（CONFIG、FLOW、MONTHLY、ONCE）(Optional) */
func (r *CreateInstanceRequest) SetChargeMode(chargeMode string) {
    r.ChargeMode = &chargeMode
}

/* param packType: 套餐类型（FREE、BASIC、PROFESSIONAL、ENTERPRISE、ULTIMATE、SMB_BASIC、SMB_BUSINESS）(Optional) */
func (r *CreateInstanceRequest) SetPackType(packType string) {
    r.PackType = &packType
}

/* param zonePackNum: 域名增量包数量(Optional) */
func (r *CreateInstanceRequest) SetZonePackNum(zonePackNum int) {
    r.ZonePackNum = &zonePackNum
}

/* param duration: 计费时长(Optional) */
func (r *CreateInstanceRequest) SetDuration(duration int) {
    r.Duration = &duration
}

/* param durationUnit: 计费时长单位（MONTH、YEAR）(Optional) */
func (r *CreateInstanceRequest) SetDurationUnit(durationUnit string) {
    r.DurationUnit = &durationUnit
}

/* param autoRenewStatus: 自动续费状态(OPEN->开通自动续费 CLOSE->关闭自动续费)(Optional) */
func (r *CreateInstanceRequest) SetAutoRenewStatus(autoRenewStatus string) {
    r.AutoRenewStatus = &autoRenewStatus
}

/* param instanceName: 实例名称(Optional) */
func (r *CreateInstanceRequest) SetInstanceName(instanceName string) {
    r.InstanceName = &instanceName
}

/* param memo: 备注(Optional) */
func (r *CreateInstanceRequest) SetMemo(memo string) {
    r.Memo = &memo
}

/* param returnUrl: 支付成功返回路径(Optional) */
func (r *CreateInstanceRequest) SetReturnUrl(returnUrl string) {
    r.ReturnUrl = &returnUrl
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateInstanceRequest) GetRegionId() string {
    return r.RegionId
}

type CreateInstanceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateInstanceResult `json:"result"`
}

type CreateInstanceResult struct {
    BuyId string `json:"buyId"`
}