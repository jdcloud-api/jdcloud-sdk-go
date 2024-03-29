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
    starshield "github.com/jdcloud-api/jdcloud-sdk-go/services/starshield/models"
)

type BandwidthTopKRequest struct {

    core.JDCloudRequest

    /*   */
    Zone_identifier string `json:"zone_identifier"`

    /*  (Optional) */
    ZoneName *string `json:"zoneName"`

    /*  (Optional) */
    Since *string `json:"since"`

    /*  (Optional) */
    Until *string `json:"until"`

    /*  (Optional) */
    TopK *int `json:"topK"`

    /*  (Optional) */
    Filters []starshield.AnalyticsReportingFilter `json:"filters"`
}

/*
 * param zone_identifier:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewBandwidthTopKRequest(
    zone_identifier string,
) *BandwidthTopKRequest {

	return &BandwidthTopKRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/zones/{zone_identifier}/analytics$$bandwidthTopK",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Zone_identifier: zone_identifier,
	}
}

/*
 * param zone_identifier:  (Required)
 * param zoneName:  (Optional)
 * param since:  (Optional)
 * param until:  (Optional)
 * param topK:  (Optional)
 * param filters:  (Optional)
 */
func NewBandwidthTopKRequestWithAllParams(
    zone_identifier string,
    zoneName *string,
    since *string,
    until *string,
    topK *int,
    filters []starshield.AnalyticsReportingFilter,
) *BandwidthTopKRequest {

    return &BandwidthTopKRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{zone_identifier}/analytics$$bandwidthTopK",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Zone_identifier: zone_identifier,
        ZoneName: zoneName,
        Since: since,
        Until: until,
        TopK: topK,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewBandwidthTopKRequestWithoutParam() *BandwidthTopKRequest {

    return &BandwidthTopKRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{zone_identifier}/analytics$$bandwidthTopK",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param zone_identifier: (Required) */
func (r *BandwidthTopKRequest) SetZone_identifier(zone_identifier string) {
    r.Zone_identifier = zone_identifier
}
/* param zoneName: (Optional) */
func (r *BandwidthTopKRequest) SetZoneName(zoneName string) {
    r.ZoneName = &zoneName
}
/* param since: (Optional) */
func (r *BandwidthTopKRequest) SetSince(since string) {
    r.Since = &since
}
/* param until: (Optional) */
func (r *BandwidthTopKRequest) SetUntil(until string) {
    r.Until = &until
}
/* param topK: (Optional) */
func (r *BandwidthTopKRequest) SetTopK(topK int) {
    r.TopK = &topK
}
/* param filters: (Optional) */
func (r *BandwidthTopKRequest) SetFilters(filters []starshield.AnalyticsReportingFilter) {
    r.Filters = filters
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r BandwidthTopKRequest) GetRegionId() string {
    return ""
}

type BandwidthTopKResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result BandwidthTopKResult `json:"result"`
}

type BandwidthTopKResult struct {
    TopkAnalytics starshield.TopkAnalytics `json:"topkAnalytics"`
}