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
    starshield "github.com/lshuining/jdcloud-sdk-go/services/starshield/models"
)

type PurgeAllFilesRequest struct {

    core.JDCloudRequest

    /*   */
    Identifier string `json:"identifier"`

    /* 指示星盾缓存中的所有资源都应该删除的标志。
注意，执行此操作后，可能会对源服务器负载产生显著影响。
  */
    Purge_everything bool `json:"purge_everything"`
}

/*
 * param identifier:  (Required)
 * param purge_everything: 指示星盾缓存中的所有资源都应该删除的标志。
注意，执行此操作后，可能会对源服务器负载产生显著影响。
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewPurgeAllFilesRequest(
    identifier string,
    purge_everything bool,
) *PurgeAllFilesRequest {

	return &PurgeAllFilesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/zones/{identifier}/purge_cache__purgeAllFiles",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Identifier: identifier,
        Purge_everything: purge_everything,
	}
}

/*
 * param identifier:  (Required)
 * param purge_everything: 指示星盾缓存中的所有资源都应该删除的标志。
注意，执行此操作后，可能会对源服务器负载产生显著影响。
 (Required)
 */
func NewPurgeAllFilesRequestWithAllParams(
    identifier string,
    purge_everything bool,
) *PurgeAllFilesRequest {

    return &PurgeAllFilesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{identifier}/purge_cache__purgeAllFiles",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Identifier: identifier,
        Purge_everything: purge_everything,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewPurgeAllFilesRequestWithoutParam() *PurgeAllFilesRequest {

    return &PurgeAllFilesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{identifier}/purge_cache__purgeAllFiles",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param identifier: (Required) */
func (r *PurgeAllFilesRequest) SetIdentifier(identifier string) {
    r.Identifier = identifier
}

/* param purge_everything: 指示星盾缓存中的所有资源都应该删除的标志。
注意，执行此操作后，可能会对源服务器负载产生显著影响。
(Required) */
func (r *PurgeAllFilesRequest) SetPurge_everything(purge_everything bool) {
    r.Purge_everything = purge_everything
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r PurgeAllFilesRequest) GetRegionId() string {
    return ""
}

type PurgeAllFilesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result PurgeAllFilesResult `json:"result"`
}

type PurgeAllFilesResult struct {
    Data starshield.Zone `json:"data"`
}