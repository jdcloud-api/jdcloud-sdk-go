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

type ModifyNetworkSecurityGroupRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* NetworkSecurityGroup ID  */
    NetworkSecurityGroupId string `json:"networkSecurityGroupId"`

    /* 安全组的名字。名称取值范围：1-32个中文、英文大小写的字母、数字和下划线分隔符 (Optional) */
    NetworkSecurityGroupName *string `json:"networkSecurityGroupName"`

    /* 安全组的描述，取值范围：0-256个UTF-8编码下的全部字符 (Optional) */
    Description *string `json:"description"`
}

/*
 * param regionId: Region ID (Required)
 * param networkSecurityGroupId: NetworkSecurityGroup ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyNetworkSecurityGroupRequest(
    regionId string,
    networkSecurityGroupId string,
) *ModifyNetworkSecurityGroupRequest {

	return &ModifyNetworkSecurityGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/networkSecurityGroups/{networkSecurityGroupId}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        NetworkSecurityGroupId: networkSecurityGroupId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param networkSecurityGroupId: NetworkSecurityGroup ID (Required)
 * param networkSecurityGroupName: 安全组的名字。名称取值范围：1-32个中文、英文大小写的字母、数字和下划线分隔符 (Optional)
 * param description: 安全组的描述，取值范围：0-256个UTF-8编码下的全部字符 (Optional)
 */
func NewModifyNetworkSecurityGroupRequestWithAllParams(
    regionId string,
    networkSecurityGroupId string,
    networkSecurityGroupName *string,
    description *string,
) *ModifyNetworkSecurityGroupRequest {

    return &ModifyNetworkSecurityGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/networkSecurityGroups/{networkSecurityGroupId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        NetworkSecurityGroupId: networkSecurityGroupId,
        NetworkSecurityGroupName: networkSecurityGroupName,
        Description: description,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyNetworkSecurityGroupRequestWithoutParam() *ModifyNetworkSecurityGroupRequest {

    return &ModifyNetworkSecurityGroupRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/networkSecurityGroups/{networkSecurityGroupId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *ModifyNetworkSecurityGroupRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param networkSecurityGroupId: NetworkSecurityGroup ID(Required) */
func (r *ModifyNetworkSecurityGroupRequest) SetNetworkSecurityGroupId(networkSecurityGroupId string) {
    r.NetworkSecurityGroupId = networkSecurityGroupId
}

/* param networkSecurityGroupName: 安全组的名字。名称取值范围：1-32个中文、英文大小写的字母、数字和下划线分隔符(Optional) */
func (r *ModifyNetworkSecurityGroupRequest) SetNetworkSecurityGroupName(networkSecurityGroupName string) {
    r.NetworkSecurityGroupName = &networkSecurityGroupName
}

/* param description: 安全组的描述，取值范围：0-256个UTF-8编码下的全部字符(Optional) */
func (r *ModifyNetworkSecurityGroupRequest) SetDescription(description string) {
    r.Description = &description
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyNetworkSecurityGroupRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyNetworkSecurityGroupResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyNetworkSecurityGroupResult `json:"result"`
}

type ModifyNetworkSecurityGroupResult struct {
}