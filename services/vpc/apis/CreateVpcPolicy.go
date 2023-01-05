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

type CreateVpcPolicyRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 访问控制策略名称。取值范围：1-32个中文、英文大小写的字母、数字和下划线分隔符  */
    VpcPolicyName string `json:"vpcPolicyName"`

    /* 访问控制策略所属vpc id  */
    VpcId string `json:"vpcId"`

    /* 具体策略内容(格式参考样例)  */
    PolicyDocument string `json:"policyDocument"`
}

/*
 * param regionId: Region ID (Required)
 * param vpcPolicyName: 访问控制策略名称。取值范围：1-32个中文、英文大小写的字母、数字和下划线分隔符 (Required)
 * param vpcId: 访问控制策略所属vpc id (Required)
 * param policyDocument: 具体策略内容(格式参考样例) (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateVpcPolicyRequest(
    regionId string,
    vpcPolicyName string,
    vpcId string,
    policyDocument string,
) *CreateVpcPolicyRequest {

	return &CreateVpcPolicyRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/vpcPolicy/",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        VpcPolicyName: vpcPolicyName,
        VpcId: vpcId,
        PolicyDocument: policyDocument,
	}
}

/*
 * param regionId: Region ID (Required)
 * param vpcPolicyName: 访问控制策略名称。取值范围：1-32个中文、英文大小写的字母、数字和下划线分隔符 (Required)
 * param vpcId: 访问控制策略所属vpc id (Required)
 * param policyDocument: 具体策略内容(格式参考样例) (Required)
 */
func NewCreateVpcPolicyRequestWithAllParams(
    regionId string,
    vpcPolicyName string,
    vpcId string,
    policyDocument string,
) *CreateVpcPolicyRequest {

    return &CreateVpcPolicyRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vpcPolicy/",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        VpcPolicyName: vpcPolicyName,
        VpcId: vpcId,
        PolicyDocument: policyDocument,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateVpcPolicyRequestWithoutParam() *CreateVpcPolicyRequest {

    return &CreateVpcPolicyRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vpcPolicy/",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CreateVpcPolicyRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param vpcPolicyName: 访问控制策略名称。取值范围：1-32个中文、英文大小写的字母、数字和下划线分隔符(Required) */
func (r *CreateVpcPolicyRequest) SetVpcPolicyName(vpcPolicyName string) {
    r.VpcPolicyName = vpcPolicyName
}
/* param vpcId: 访问控制策略所属vpc id(Required) */
func (r *CreateVpcPolicyRequest) SetVpcId(vpcId string) {
    r.VpcId = vpcId
}
/* param policyDocument: 具体策略内容(格式参考样例)(Required) */
func (r *CreateVpcPolicyRequest) SetPolicyDocument(policyDocument string) {
    r.PolicyDocument = policyDocument
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateVpcPolicyRequest) GetRegionId() string {
    return r.RegionId
}

type CreateVpcPolicyResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateVpcPolicyResult `json:"result"`
}

type CreateVpcPolicyResult struct {
    VpcPolicyId string `json:"vpcPolicyId"`
}