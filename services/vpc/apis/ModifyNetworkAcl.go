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

type ModifyNetworkAclRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* networkAclId ID  */
    NetworkAclId string `json:"networkAclId"`

    /* networkAcl名称,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符 (Optional) */
    NetworkAclName *string `json:"networkAclName"`

    /* 描述,允许输入UTF-8编码下的全部字符，不超过256字符 (Optional) */
    Description *string `json:"description"`
}

/*
 * param regionId: Region ID (Required)
 * param networkAclId: networkAclId ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyNetworkAclRequest(
    regionId string,
    networkAclId string,
) *ModifyNetworkAclRequest {

	return &ModifyNetworkAclRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/networkAcls/{networkAclId}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        NetworkAclId: networkAclId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param networkAclId: networkAclId ID (Required)
 * param networkAclName: networkAcl名称,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符 (Optional)
 * param description: 描述,允许输入UTF-8编码下的全部字符，不超过256字符 (Optional)
 */
func NewModifyNetworkAclRequestWithAllParams(
    regionId string,
    networkAclId string,
    networkAclName *string,
    description *string,
) *ModifyNetworkAclRequest {

    return &ModifyNetworkAclRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/networkAcls/{networkAclId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        NetworkAclId: networkAclId,
        NetworkAclName: networkAclName,
        Description: description,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyNetworkAclRequestWithoutParam() *ModifyNetworkAclRequest {

    return &ModifyNetworkAclRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/networkAcls/{networkAclId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *ModifyNetworkAclRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param networkAclId: networkAclId ID(Required) */
func (r *ModifyNetworkAclRequest) SetNetworkAclId(networkAclId string) {
    r.NetworkAclId = networkAclId
}

/* param networkAclName: networkAcl名称,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符(Optional) */
func (r *ModifyNetworkAclRequest) SetNetworkAclName(networkAclName string) {
    r.NetworkAclName = &networkAclName
}

/* param description: 描述,允许输入UTF-8编码下的全部字符，不超过256字符(Optional) */
func (r *ModifyNetworkAclRequest) SetDescription(description string) {
    r.Description = &description
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyNetworkAclRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyNetworkAclResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyNetworkAclResult `json:"result"`
}

type ModifyNetworkAclResult struct {
}