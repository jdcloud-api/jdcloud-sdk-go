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
    antipro "github.com/lshuining/jdcloud-sdk-go/services/antipro/models"
)

type DescribeGeoAreasRequest struct {

    core.JDCloudRequest
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeGeoAreasRequest(
) *DescribeGeoAreasRequest {

	return &DescribeGeoAreasRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/describeGeoAreas",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 */
func NewDescribeGeoAreasRequestWithAllParams(
) *DescribeGeoAreasRequest {

    return &DescribeGeoAreasRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/describeGeoAreas",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeGeoAreasRequestWithoutParam() *DescribeGeoAreasRequest {

    return &DescribeGeoAreasRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/describeGeoAreas",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeGeoAreasRequest) GetRegionId() string {
    return ""
}

type DescribeGeoAreasResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeGeoAreasResult `json:"result"`
}

type DescribeGeoAreasResult struct {
    DataList []antipro.Country `json:"dataList"`
}