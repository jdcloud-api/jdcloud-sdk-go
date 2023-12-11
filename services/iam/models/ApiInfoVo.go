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

package models


type ApiInfoVo struct {

    /*   */
    ServiceNameCN string `json:"serviceNameCN"`

    /*   */
    ServiceName string `json:"serviceName"`

    /*   */
    ServiceOldCode string `json:"serviceOldCode"`

    /*   */
    ResourceType string `json:"resourceType"`

    /*   */
    SubResourceType string `json:"subResourceType"`

    /*   */
    ActionName string `json:"actionName"`

    /*   */
    ActionDescription string `json:"actionDescription"`

    /*   */
    ActionType string `json:"actionType"`

    /*   */
    Method string `json:"method"`

    /*   */
    Region int `json:"region"`

    /*   */
    Version int `json:"version"`

    /*   */
    AccessLevel int `json:"accessLevel"`

    /*   */
    ShowRule []ApiInfoShowRule `json:"showRule"`

    /*   */
    Supply int `json:"supply"`

    /*   */
    IgnoreDeny int `json:"ignoreDeny"`

    /*   */
    TagResourceLevel int `json:"tagResourceLevel"`
}
