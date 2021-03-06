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


type DomainTranslateTemplateConfig struct {

    /* 应用名称 (Optional) */
    AppName string `json:"appName"`

    /* 流名称 (Optional) */
    StreamName string `json:"streamName"`

    /* 翻译输出流扩展后缀 (Optional) */
    TemplateCode string `json:"templateCode"`

    /* 模板名称 (Optional) */
    TemplateName string `json:"templateName"`

    /* 翻译模式:
  - 0  中译英
  - 1  英译中
 (Optional) */
    TranslateMode int `json:"translateMode"`
}
