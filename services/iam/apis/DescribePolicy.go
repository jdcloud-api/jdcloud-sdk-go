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
    iam "github.com/lshuining/jdcloud-sdk-go/services/iam/models"
)

type DescribePolicyRequest struct {

    core.JDCloudRequest

    /* 策略名称  */
    PolicyName string `json:"policyName"`
}

/*
 * param policyName: 策略名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribePolicyRequest(
    policyName string,
) *DescribePolicyRequest {

	return &DescribePolicyRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/policy/{policyName}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        PolicyName: policyName,
	}
}

/*
 * param policyName: 策略名称 (Required)
 */
func NewDescribePolicyRequestWithAllParams(
    policyName string,
) *DescribePolicyRequest {

    return &DescribePolicyRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/policy/{policyName}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PolicyName: policyName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribePolicyRequestWithoutParam() *DescribePolicyRequest {

    return &DescribePolicyRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/policy/{policyName}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param policyName: 策略名称(Required) */
func (r *DescribePolicyRequest) SetPolicyName(policyName string) {
    r.PolicyName = policyName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribePolicyRequest) GetRegionId() string {
    return ""
}

type DescribePolicyResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribePolicyResult `json:"result"`
}

type DescribePolicyResult struct {
    Policy iam.PolicyDetail `json:"policy"`
}