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


type TransTask struct {

    /* 任务ID (Optional) */
    TaskId int64 `json:"taskId"`

    /* 打包类型。取值范围：None, HLSPackage (Optional) */
    PackageType string `json:"packageType"`

    /* 转码模板ID。非打包转码，包含此字段。 (Optional) */
    TemplateId int64 `json:"templateId"`

    /* 转码模板组ID。若是以模板组方式提交作业，生成的转码任务中包含此字段。 (Optional) */
    TemplateGroupId string `json:"templateGroupId"`

    /* 模板ID列表。打包转码，包含一个模板组中的多个打包转码模板。 (Optional) */
    TemplateIds []int64 `json:"templateIds"`

    /* 封装格式 (Optional) */
    Format string `json:"format"`

    /* 清晰度 (Optional) */
    Definition string `json:"definition"`

    /* 任务状态。
submitted - 已提交
processing - 处理中
succeeded - 处理成功
failed - 处理失败
 (Optional) */
    Status string `json:"status"`

    /* 处理进度 (Optional) */
    Progress int `json:"progress"`

    /* 错误码，只有处理失败时，才有此字段 (Optional) */
    ErrorCode string `json:"errorCode"`

    /* 错误消息，只有处理失败时，才有此字段 (Optional) */
    ErrorMessage string `json:"errorMessage"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 完成时间 (Optional) */
    CompleteTime string `json:"completeTime"`
}
