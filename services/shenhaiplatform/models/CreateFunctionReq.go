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


type CreateFunctionReq struct {

    /* 租户code (Optional) */
    CompanyCode string `json:"companyCode"`

    /* 命名空间code (Optional) */
    WorkspaceCode string `json:"workspaceCode"`

    /* 函数名称（支持字母、数字、下划线，不超过50个字符） (Optional) */
    FunctionName string `json:"functionName"`

    /* 分类目录id (Optional) */
    FunctionCatalogId int `json:"functionCatalogId"`

    /* 函数主类全路径 (Optional) */
    ClassName string `json:"className"`

    /* 负责人 (Optional) */
    Managers []string `json:"managers"`

    /* 依赖的资源code（包括资源文件/资源目录） (Optional) */
    ResourceCodes []string `json:"resourceCodes"`

    /* 函数格式 (Optional) */
    FunctionFormat string `json:"functionFormat"`

    /* 函数使用说明 (Optional) */
    FunctionUsage string `json:"functionUsage"`

    /* 关联引擎 (Optional) */
    RelativeEngine string `json:"relativeEngine"`
}
