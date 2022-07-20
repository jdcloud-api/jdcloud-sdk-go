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

type GetPublicKeyRequest struct {

    core.JDCloudRequest

    /* 密钥ID  */
    KeyId string `json:"keyId"`
}

/*
 * param keyId: 密钥ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetPublicKeyRequest(
    keyId string,
) *GetPublicKeyRequest {

	return &GetPublicKeyRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/key/{keyId}:GetPublicKey",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        KeyId: keyId,
	}
}

/*
 * param keyId: 密钥ID (Required)
 */
func NewGetPublicKeyRequestWithAllParams(
    keyId string,
) *GetPublicKeyRequest {

    return &GetPublicKeyRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/key/{keyId}:GetPublicKey",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        KeyId: keyId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetPublicKeyRequestWithoutParam() *GetPublicKeyRequest {

    return &GetPublicKeyRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/key/{keyId}:GetPublicKey",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param keyId: 密钥ID(Required) */
func (r *GetPublicKeyRequest) SetKeyId(keyId string) {
    r.KeyId = keyId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetPublicKeyRequest) GetRegionId() string {
    return ""
}

type GetPublicKeyResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetPublicKeyResult `json:"result"`
}

type GetPublicKeyResult struct {
    PublicKeyBlob string `json:"publicKeyBlob"`
}