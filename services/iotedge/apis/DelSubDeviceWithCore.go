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
    iotedge "github.com/jdcloud-api/jdcloud-sdk-go/services/iotedge/models"
)

type DelSubDeviceWithCoreRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* IoTCore实例编号  */
    InstanceId string `json:"instanceId"`

    /* Edge名称  */
    EdgeName string `json:"edgeName"`

    /* Edge对应的产品key  */
    ProductKey string `json:"productKey"`

    /* 设备名称 (Optional) */
    DelDevices []iotedge.DelSubDevices `json:"delDevices"`
}

/*
 * param regionId: 地域ID (Required)
 * param instanceId: IoTCore实例编号 (Required)
 * param edgeName: Edge名称 (Required)
 * param productKey: Edge对应的产品key (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDelSubDeviceWithCoreRequest(
    regionId string,
    instanceId string,
    edgeName string,
    productKey string,
) *DelSubDeviceWithCoreRequest {

	return &DelSubDeviceWithCoreRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/products/{productKey}/edges/{edgeName}:delSubDevice",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        EdgeName: edgeName,
        ProductKey: productKey,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param instanceId: IoTCore实例编号 (Required)
 * param edgeName: Edge名称 (Required)
 * param productKey: Edge对应的产品key (Required)
 * param delDevices: 设备名称 (Optional)
 */
func NewDelSubDeviceWithCoreRequestWithAllParams(
    regionId string,
    instanceId string,
    edgeName string,
    productKey string,
    delDevices []iotedge.DelSubDevices,
) *DelSubDeviceWithCoreRequest {

    return &DelSubDeviceWithCoreRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/products/{productKey}/edges/{edgeName}:delSubDevice",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        EdgeName: edgeName,
        ProductKey: productKey,
        DelDevices: delDevices,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDelSubDeviceWithCoreRequestWithoutParam() *DelSubDeviceWithCoreRequest {

    return &DelSubDeviceWithCoreRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/products/{productKey}/edges/{edgeName}:delSubDevice",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DelSubDeviceWithCoreRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: IoTCore实例编号(Required) */
func (r *DelSubDeviceWithCoreRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param edgeName: Edge名称(Required) */
func (r *DelSubDeviceWithCoreRequest) SetEdgeName(edgeName string) {
    r.EdgeName = edgeName
}

/* param productKey: Edge对应的产品key(Required) */
func (r *DelSubDeviceWithCoreRequest) SetProductKey(productKey string) {
    r.ProductKey = productKey
}

/* param delDevices: 设备名称(Optional) */
func (r *DelSubDeviceWithCoreRequest) SetDelDevices(delDevices []iotedge.DelSubDevices) {
    r.DelDevices = delDevices
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DelSubDeviceWithCoreRequest) GetRegionId() string {
    return r.RegionId
}

type DelSubDeviceWithCoreResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DelSubDeviceWithCoreResult `json:"result"`
}

type DelSubDeviceWithCoreResult struct {
}