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


type RuleLogic struct {

    /* 是否启用 (Optional) */
    Enabled *bool `json:"enabled"`

    /* 逻辑: 0->不包含, 1->包含, 2->等于, 3->小于等于, 4->小于, 5->大于等于, 6->大于, 7->不等于, 8->正则 (Optional) */
    Logic *int `json:"logic"`

    /* 值 (Optional) */
    Value *string `json:"value"`
}
