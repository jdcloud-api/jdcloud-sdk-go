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
)

type DescribeProductRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* IoT Engine实例ID信息  */
    InstanceId string `json:"instanceId"`

    /* 产品Key  */
    ProductKey string `json:"productKey"`
}

/*
 * param regionId: 地域ID (Required)
 * param instanceId: IoT Engine实例ID信息 (Required)
 * param productKey: 产品Key (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeProductRequest(
    regionId string,
    instanceId string,
    productKey string,
) *DescribeProductRequest {

	return &DescribeProductRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/products/{productKey}",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        ProductKey: productKey,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param instanceId: IoT Engine实例ID信息 (Required)
 * param productKey: 产品Key (Required)
 */
func NewDescribeProductRequestWithAllParams(
    regionId string,
    instanceId string,
    productKey string,
) *DescribeProductRequest {

    return &DescribeProductRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/products/{productKey}",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        ProductKey: productKey,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeProductRequestWithoutParam() *DescribeProductRequest {

    return &DescribeProductRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/products/{productKey}",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DescribeProductRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: IoT Engine实例ID信息(Required) */
func (r *DescribeProductRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param productKey: 产品Key(Required) */
func (r *DescribeProductRequest) SetProductKey(productKey string) {
    r.ProductKey = productKey
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeProductRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeProductResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeProductResult `json:"result"`
}

type DescribeProductResult struct {
    ProductName string `json:"productName"`
    ProductType int `json:"productType"`
    ProductKey string `json:"productKey"`
    ProductSecret string `json:"productSecret"`
    CreatedTime int64 `json:"createdTime"`
    DeviceCount int `json:"deviceCount"`
    DynamicRegister int `json:"dynamicRegister"`
    ProductDescription string `json:"productDescription"`
    TemplateName string `json:"templateName"`
}