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

type GravityParticleSchedDictGetByTypeRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 数据字典id (Optional) */
    Id *int64 `json:"id"`

    /* 数据字典code (Optional) */
    DictCode *string `json:"dictCode"`

    /* 数据字典父级code (Optional) */
    PDictCode *string `json:"pDictCode"`

    /* 数据字典名称 (Optional) */
    DictName *string `json:"dictName"`

    /* 数据字典排序 (Optional) */
    DictSort *int `json:"dictSort"`

    /* 数据字典备份 (Optional) */
    DictBak *string `json:"dictBak"`

    /* 数据字典描述 (Optional) */
    DictDesc *string `json:"dictDesc"`

    /* 是否有效 1: 有效, 0: 无效 (Optional) */
    IsValid *string `json:"isValid"`

    /* 是否删除 1: 删除, 0: 未删除 (Optional) */
    IsDel *string `json:"isDel"`

    /* 创建时间 (Optional) */
    CreatedTime *string `json:"createdTime"`

    /* 修改时间 (Optional) */
    ModifiedTime *string `json:"modifiedTime"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGravityParticleSchedDictGetByTypeRequest(
    regionId string,
    appName string,
) *GravityParticleSchedDictGetByTypeRequest {

	return &GravityParticleSchedDictGetByTypeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/gravityParticleSchedDictGetByType",
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
 * param id: 数据字典id (Optional)
 * param dictCode: 数据字典code (Optional)
 * param pDictCode: 数据字典父级code (Optional)
 * param dictName: 数据字典名称 (Optional)
 * param dictSort: 数据字典排序 (Optional)
 * param dictBak: 数据字典备份 (Optional)
 * param dictDesc: 数据字典描述 (Optional)
 * param isValid: 是否有效 1: 有效, 0: 无效 (Optional)
 * param isDel: 是否删除 1: 删除, 0: 未删除 (Optional)
 * param createdTime: 创建时间 (Optional)
 * param modifiedTime: 修改时间 (Optional)
 */
func NewGravityParticleSchedDictGetByTypeRequestWithAllParams(
    regionId string,
    appName string,
    id *int64,
    dictCode *string,
    pDictCode *string,
    dictName *string,
    dictSort *int,
    dictBak *string,
    dictDesc *string,
    isValid *string,
    isDel *string,
    createdTime *string,
    modifiedTime *string,
) *GravityParticleSchedDictGetByTypeRequest {

    return &GravityParticleSchedDictGetByTypeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/gravityParticleSchedDictGetByType",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
        Id: id,
        DictCode: dictCode,
        PDictCode: pDictCode,
        DictName: dictName,
        DictSort: dictSort,
        DictBak: dictBak,
        DictDesc: dictDesc,
        IsValid: isValid,
        IsDel: isDel,
        CreatedTime: createdTime,
        ModifiedTime: modifiedTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGravityParticleSchedDictGetByTypeRequestWithoutParam() *GravityParticleSchedDictGetByTypeRequest {

    return &GravityParticleSchedDictGetByTypeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/gravityParticleSchedDictGetByType",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *GravityParticleSchedDictGetByTypeRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *GravityParticleSchedDictGetByTypeRequest) SetAppName(appName string) {
    r.AppName = appName
}
/* param id: 数据字典id(Optional) */
func (r *GravityParticleSchedDictGetByTypeRequest) SetId(id int64) {
    r.Id = &id
}
/* param dictCode: 数据字典code(Optional) */
func (r *GravityParticleSchedDictGetByTypeRequest) SetDictCode(dictCode string) {
    r.DictCode = &dictCode
}
/* param pDictCode: 数据字典父级code(Optional) */
func (r *GravityParticleSchedDictGetByTypeRequest) SetPDictCode(pDictCode string) {
    r.PDictCode = &pDictCode
}
/* param dictName: 数据字典名称(Optional) */
func (r *GravityParticleSchedDictGetByTypeRequest) SetDictName(dictName string) {
    r.DictName = &dictName
}
/* param dictSort: 数据字典排序(Optional) */
func (r *GravityParticleSchedDictGetByTypeRequest) SetDictSort(dictSort int) {
    r.DictSort = &dictSort
}
/* param dictBak: 数据字典备份(Optional) */
func (r *GravityParticleSchedDictGetByTypeRequest) SetDictBak(dictBak string) {
    r.DictBak = &dictBak
}
/* param dictDesc: 数据字典描述(Optional) */
func (r *GravityParticleSchedDictGetByTypeRequest) SetDictDesc(dictDesc string) {
    r.DictDesc = &dictDesc
}
/* param isValid: 是否有效 1: 有效, 0: 无效(Optional) */
func (r *GravityParticleSchedDictGetByTypeRequest) SetIsValid(isValid string) {
    r.IsValid = &isValid
}
/* param isDel: 是否删除 1: 删除, 0: 未删除(Optional) */
func (r *GravityParticleSchedDictGetByTypeRequest) SetIsDel(isDel string) {
    r.IsDel = &isDel
}
/* param createdTime: 创建时间(Optional) */
func (r *GravityParticleSchedDictGetByTypeRequest) SetCreatedTime(createdTime string) {
    r.CreatedTime = &createdTime
}
/* param modifiedTime: 修改时间(Optional) */
func (r *GravityParticleSchedDictGetByTypeRequest) SetModifiedTime(modifiedTime string) {
    r.ModifiedTime = &modifiedTime
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GravityParticleSchedDictGetByTypeRequest) GetRegionId() string {
    return r.RegionId
}

type GravityParticleSchedDictGetByTypeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GravityParticleSchedDictGetByTypeResult `json:"result"`
}

type GravityParticleSchedDictGetByTypeResult struct {
    Success int `json:"success"`
    Result []shenhaiplatform.GpsdalueDict `json:"result"`
    Code string `json:"code"`
    Msg string `json:"msg"`
    _REQ_SEQ_NO_ string `json:"_REQ_SEQ_NO_"`
}