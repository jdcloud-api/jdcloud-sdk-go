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
    jdfusion "github.com/lshuining/jdcloud-sdk-go/services/jdfusion/models"
)

type GetDeploymentsVersionRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* deployment ID  */
    Id string `json:"id"`

    /* application ID  */
    Version_id string `json:"version_id"`
}

/*
 * param regionId: 地域ID (Required)
 * param id: deployment ID (Required)
 * param version_id: application ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetDeploymentsVersionRequest(
    regionId string,
    id string,
    version_id string,
) *GetDeploymentsVersionRequest {

	return &GetDeploymentsVersionRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/deployments/{id}/versions/{version_id}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Id: id,
        Version_id: version_id,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param id: deployment ID (Required)
 * param version_id: application ID (Required)
 */
func NewGetDeploymentsVersionRequestWithAllParams(
    regionId string,
    id string,
    version_id string,
) *GetDeploymentsVersionRequest {

    return &GetDeploymentsVersionRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/deployments/{id}/versions/{version_id}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Id: id,
        Version_id: version_id,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetDeploymentsVersionRequestWithoutParam() *GetDeploymentsVersionRequest {

    return &GetDeploymentsVersionRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/deployments/{id}/versions/{version_id}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *GetDeploymentsVersionRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param id: deployment ID(Required) */
func (r *GetDeploymentsVersionRequest) SetId(id string) {
    r.Id = id
}

/* param version_id: application ID(Required) */
func (r *GetDeploymentsVersionRequest) SetVersion_id(version_id string) {
    r.Version_id = version_id
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetDeploymentsVersionRequest) GetRegionId() string {
    return r.RegionId
}

type GetDeploymentsVersionResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetDeploymentsVersionResult `json:"result"`
}

type GetDeploymentsVersionResult struct {
    Deployment jdfusion.DeploymentInfo `json:"deployment"`
}