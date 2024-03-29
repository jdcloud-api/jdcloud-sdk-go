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


type Tag struct {

    /* 标签key。长度不能超过127字符，不能以 `jrn:` 或 `jdc-` 开头，仅支持中文、大/小写英文、数字及如下符号：`\_.,:\/=+-@`。 (Optional) */
    Key *string `json:"key"`

    /* 标签value。长度不能超过255字符，仅支持中文、大/小写英文、数字及如下符号：`\_.,:\/=+-@`。 (Optional) */
    Value *string `json:"value"`
}
