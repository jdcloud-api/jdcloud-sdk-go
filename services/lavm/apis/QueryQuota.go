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
    lavm "github.com/jdcloud-api/jdcloud-sdk-go/services/lavm/models"
)

type QueryQuotaRequest struct {

    core.JDCloudRequest

    /* 地域ID。  */
    RegionId string `json:"regionId"`

    /* 资源名列表，可取值:
general_instance 通用型套餐实例
custom_image 自定义镜像
  */
    ResourceTypes []string `json:"resourceTypes"`
}

/*
 * param regionId: 地域ID。 (Required)
 * param resourceTypes: 资源名列表，可取值:
general_instance 通用型套餐实例
custom_image 自定义镜像
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryQuotaRequest(
    regionId string,
    resourceTypes []string,
) *QueryQuotaRequest {

	return &QueryQuotaRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/quotas",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ResourceTypes: resourceTypes,
	}
}

/*
 * param regionId: 地域ID。 (Required)
 * param resourceTypes: 资源名列表，可取值:
general_instance 通用型套餐实例
custom_image 自定义镜像
 (Required)
 */
func NewQueryQuotaRequestWithAllParams(
    regionId string,
    resourceTypes []string,
) *QueryQuotaRequest {

    return &QueryQuotaRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/quotas",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ResourceTypes: resourceTypes,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryQuotaRequestWithoutParam() *QueryQuotaRequest {

    return &QueryQuotaRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/quotas",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID。(Required) */
func (r *QueryQuotaRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param resourceTypes: 资源名列表，可取值:
general_instance 通用型套餐实例
custom_image 自定义镜像
(Required) */
func (r *QueryQuotaRequest) SetResourceTypes(resourceTypes []string) {
    r.ResourceTypes = resourceTypes
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryQuotaRequest) GetRegionId() string {
    return r.RegionId
}

type QueryQuotaResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryQuotaResult `json:"result"`
}

type QueryQuotaResult struct {
    Quotas []lavm.Quota `json:"quotas"`
}