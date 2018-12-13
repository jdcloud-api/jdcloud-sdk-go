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
    clouddnsservice "github.com/jdcloud-api/jdcloud-sdk-go/services/clouddnsservice/models"
)

type SetLBRequest struct {

    core.JDCloudRequest

    /* 实例所属的地域ID  */
    RegionId string `json:"regionId"`

    /* 域名ID，请使用getDomains接口获取。  */
    DomainId string `json:"domainId"`

    /* 要设置解析记录的权重参数列表  */
    IdWeights []clouddnsservice.Setlb `json:"idWeights"`

    /* 这几条解析记录的类型。可以设置权重的类型有：A、AAAA、CNAME、JNAME  */
    Type string `json:"type"`
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param domainId: 域名ID，请使用getDomains接口获取。 (Required)
 * param idWeights: 要设置解析记录的权重参数列表 (Required)
 * param type_: 这几条解析记录的类型。可以设置权重的类型有：A、AAAA、CNAME、JNAME (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSetLBRequest(
    regionId string,
    domainId string,
    idWeights []clouddnsservice.Setlb,
    type_ string,
) *SetLBRequest {

	return &SetLBRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/domain/{domainId}/SetLB",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        DomainId: domainId,
        IdWeights: idWeights,
        Type: type_,
	}
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param domainId: 域名ID，请使用getDomains接口获取。 (Required)
 * param idWeights: 要设置解析记录的权重参数列表 (Required)
 * param type_: 这几条解析记录的类型。可以设置权重的类型有：A、AAAA、CNAME、JNAME (Required)
 */
func NewSetLBRequestWithAllParams(
    regionId string,
    domainId string,
    idWeights []clouddnsservice.Setlb,
    type_ string,
) *SetLBRequest {

    return &SetLBRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain/{domainId}/SetLB",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DomainId: domainId,
        IdWeights: idWeights,
        Type: type_,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSetLBRequestWithoutParam() *SetLBRequest {

    return &SetLBRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain/{domainId}/SetLB",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 实例所属的地域ID(Required) */
func (r *SetLBRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param domainId: 域名ID，请使用getDomains接口获取。(Required) */
func (r *SetLBRequest) SetDomainId(domainId string) {
    r.DomainId = domainId
}

/* param idWeights: 要设置解析记录的权重参数列表(Required) */
func (r *SetLBRequest) SetIdWeights(idWeights []clouddnsservice.Setlb) {
    r.IdWeights = idWeights
}

/* param type_: 这几条解析记录的类型。可以设置权重的类型有：A、AAAA、CNAME、JNAME(Required) */
func (r *SetLBRequest) SetType(type_ string) {
    r.Type = type_
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SetLBRequest) GetRegionId() string {
    return r.RegionId
}

type SetLBResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SetLBResult `json:"result"`
}

type SetLBResult struct {
}