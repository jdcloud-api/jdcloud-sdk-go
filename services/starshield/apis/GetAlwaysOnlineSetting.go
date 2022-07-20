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

type GetAlwaysOnlineSettingRequest struct {

    core.JDCloudRequest

    /*   */
    Zone_identifier string `json:"zone_identifier"`
}

/*
 * param zone_identifier:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetAlwaysOnlineSettingRequest(
    zone_identifier string,
) *GetAlwaysOnlineSettingRequest {

	return &GetAlwaysOnlineSettingRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/zones/{zone_identifier}/settings$$always_online",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        Zone_identifier: zone_identifier,
	}
}

/*
 * param zone_identifier:  (Required)
 */
func NewGetAlwaysOnlineSettingRequestWithAllParams(
    zone_identifier string,
) *GetAlwaysOnlineSettingRequest {

    return &GetAlwaysOnlineSettingRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{zone_identifier}/settings$$always_online",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        Zone_identifier: zone_identifier,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetAlwaysOnlineSettingRequestWithoutParam() *GetAlwaysOnlineSettingRequest {

    return &GetAlwaysOnlineSettingRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{zone_identifier}/settings$$always_online",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param zone_identifier: (Required) */
func (r *GetAlwaysOnlineSettingRequest) SetZone_identifier(zone_identifier string) {
    r.Zone_identifier = zone_identifier
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetAlwaysOnlineSettingRequest) GetRegionId() string {
    return ""
}

type GetAlwaysOnlineSettingResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetAlwaysOnlineSettingResult `json:"result"`
}

type GetAlwaysOnlineSettingResult struct {
    Data starshield.AlwaysOnlineMode `json:"data"`
}