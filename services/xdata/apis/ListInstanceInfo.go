// Copyright 2018-2025 JDCLOUD.COM
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
    . "github.com/jdcloud-api/jdcloud-sdk-go/core"
    "reflect"
    xdata "github.com/jdcloud-api/jdcloud-sdk-go/services/xdata/models"
)

type ListInstanceInfoRequest struct {

    JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`
}

/*
 * param regionId: 地域ID 
 */
func NewListInstanceInfoRequest(
    regionId string,
) *ListInstanceInfoRequest {

	return &ListInstanceInfoRequest{
        JDCloudRequest: JDCloudRequest{
			URL:     "/regions/{regionId}/dwInstance",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

func (r *ListInstanceInfoRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ListInstanceInfoRequest) GetRegionId() string {
    fieldName := "RegionId"
    reqType := reflect.TypeOf(r)
    value := reflect.ValueOf(r)
    _, ok := reqType.FieldByName(fieldName)
    if ok {
        return value.FieldByName(fieldName).String()
    }

    return ""
}

type ListInstanceInfoResponse struct {
    RequestID string `json:"requestId"`
    Error ErrorResponse `json:"error"`
    Result ListInstanceInfoResult `json:"result"`
}

type ListInstanceInfoResult struct {
    Status bool `json:"status"`
    Message string `json:"message"`
    Data []xdata.DwInstance `json:"data"`
}