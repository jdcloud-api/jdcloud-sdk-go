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


type ModelLabelReq struct {

    /* id,更新，删除时必传，创建时不传 (Optional) */
    Id int `json:"id"`

    /* 标签名称,创建时必填 (Optional) */
    LabelName string `json:"labelName"`

    /* 权重,创建时必填 (Optional) */
    Weights int `json:"weights"`

    /* add-添加,update-更新,delete-删除 (Optional) */
    LabelAction string `json:"labelAction"`
}
