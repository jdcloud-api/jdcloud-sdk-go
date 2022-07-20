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

type CreateAccessKeyRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 描述 (Optional) */
    Description *string `json:"description"`

    /* 密钥类型 (Optional) */
    AccessKeyType *string `json:"accessKeyType"`

    /* Access Key (Optional) */
    AccessKey *string `json:"accessKey"`
}

/*
 * param regionId: 地域ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateAccessKeyRequest(
    regionId string,
) *CreateAccessKeyRequest {

	return &CreateAccessKeyRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/accessKeys",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param description: 描述 (Optional)
 * param accessKeyType: 密钥类型 (Optional)
 * param accessKey: Access Key (Optional)
 */
func NewCreateAccessKeyRequestWithAllParams(
    regionId string,
    description *string,
    accessKeyType *string,
    accessKey *string,
) *CreateAccessKeyRequest {

    return &CreateAccessKeyRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/accessKeys",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Description: description,
        AccessKeyType: accessKeyType,
        AccessKey: accessKey,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateAccessKeyRequestWithoutParam() *CreateAccessKeyRequest {

    return &CreateAccessKeyRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/accessKeys",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *CreateAccessKeyRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param description: 描述(Optional) */
func (r *CreateAccessKeyRequest) SetDescription(description string) {
    r.Description = &description
}

/* param accessKeyType: 密钥类型(Optional) */
func (r *CreateAccessKeyRequest) SetAccessKeyType(accessKeyType string) {
    r.AccessKeyType = &accessKeyType
}

/* param accessKey: Access Key(Optional) */
func (r *CreateAccessKeyRequest) SetAccessKey(accessKey string) {
    r.AccessKey = &accessKey
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateAccessKeyRequest) GetRegionId() string {
    return r.RegionId
}

type CreateAccessKeyResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateAccessKeyResult `json:"result"`
}

type CreateAccessKeyResult struct {
    AccessKeyId string `json:"accessKeyId"`
}