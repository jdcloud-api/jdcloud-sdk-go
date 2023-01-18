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

type ExecStartRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Pod ID  */
    PodId string `json:"podId"`

    /* container name  */
    ContainerName string `json:"containerName"`

    /*   */
    ExecId string `json:"execId"`
}

/*
 * param regionId: Region ID (Required)
 * param podId: Pod ID (Required)
 * param containerName: container name (Required)
 * param execId:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewExecStartRequest(
    regionId string,
    podId string,
    containerName string,
    execId string,
) *ExecStartRequest {

	return &ExecStartRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/pods/{podId}/containers/{containerName}:execStart",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        PodId: podId,
        ContainerName: containerName,
        ExecId: execId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param podId: Pod ID (Required)
 * param containerName: container name (Required)
 * param execId:  (Required)
 */
func NewExecStartRequestWithAllParams(
    regionId string,
    podId string,
    containerName string,
    execId string,
) *ExecStartRequest {

    return &ExecStartRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/pods/{podId}/containers/{containerName}:execStart",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PodId: podId,
        ContainerName: containerName,
        ExecId: execId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewExecStartRequestWithoutParam() *ExecStartRequest {

    return &ExecStartRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/pods/{podId}/containers/{containerName}:execStart",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *ExecStartRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param podId: Pod ID(Required) */
func (r *ExecStartRequest) SetPodId(podId string) {
    r.PodId = podId
}
/* param containerName: container name(Required) */
func (r *ExecStartRequest) SetContainerName(containerName string) {
    r.ContainerName = containerName
}
/* param execId: (Required) */
func (r *ExecStartRequest) SetExecId(execId string) {
    r.ExecId = execId
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ExecStartRequest) GetRegionId() string {
    return r.RegionId
}

type ExecStartResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ExecStartResult `json:"result"`
}

type ExecStartResult struct {
}