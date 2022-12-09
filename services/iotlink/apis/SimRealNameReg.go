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
    iotlink "github.com/jdcloud-api/jdcloud-sdk-go/services/iotlink/models"
)

type SimRealNameRegRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 物联网卡移动实名登记请求参数  */
    RequestParam string `json:"requestParam"`
}

/*
 * param regionId: Region ID (Required)
 * param requestParam: 物联网卡移动实名登记请求参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSimRealNameRegRequest(
    regionId string,
    requestParam string,
) *SimRealNameRegRequest {

	return &SimRealNameRegRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/simRealNameReg",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        RequestParam: requestParam,
	}
}

/*
 * param regionId: Region ID (Required)
 * param requestParam: 物联网卡移动实名登记请求参数 (Required)
 */
func NewSimRealNameRegRequestWithAllParams(
    regionId string,
    requestParam string,
) *SimRealNameRegRequest {

    return &SimRealNameRegRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/simRealNameReg",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        RequestParam: requestParam,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSimRealNameRegRequestWithoutParam() *SimRealNameRegRequest {

    return &SimRealNameRegRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/simRealNameReg",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *SimRealNameRegRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param requestParam: 物联网卡移动实名登记请求参数(Required) */
func (r *SimRealNameRegRequest) SetRequestParam(requestParam string) {
    r.RequestParam = requestParam
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SimRealNameRegRequest) GetRegionId() string {
    return r.RegionId
}

type SimRealNameRegResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SimRealNameRegResult `json:"result"`
}

type SimRealNameRegResult struct {
    Status string `json:"status"`
    Message string `json:"message"`
    Result iotlink.SimRealNameRegResp `json:"result"`
}