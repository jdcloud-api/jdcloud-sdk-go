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
)

type EncryptRequest struct {

    core.JDCloudRequest

    /* 明文 (Optional) */
    Plaintext *string `json:"plaintext"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewEncryptRequest(
) *EncryptRequest {

	return &EncryptRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/encrypt",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param plaintext: 明文 (Optional)
 */
func NewEncryptRequestWithAllParams(
    plaintext *string,
) *EncryptRequest {

    return &EncryptRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/encrypt",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Plaintext: plaintext,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewEncryptRequestWithoutParam() *EncryptRequest {

    return &EncryptRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/encrypt",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param plaintext: 明文(Optional) */
func (r *EncryptRequest) SetPlaintext(plaintext string) {
    r.Plaintext = &plaintext
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r EncryptRequest) GetRegionId() string {
    return ""
}

type EncryptResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result EncryptResult `json:"result"`
}

type EncryptResult struct {
    Ciphertext string `json:"ciphertext"`
}