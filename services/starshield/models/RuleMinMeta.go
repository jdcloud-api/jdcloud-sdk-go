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


type RuleMinMeta struct {

    /* 规则标识。 (Optional) */
    Id *string `json:"id"`

    /* 命中规则的操作。 (Optional) */
    Action *string `json:"action"`

    /* 规则是否开启。 (Optional) */
    Enabled *bool `json:"enabled"`

    /* OWASP 异常情况分数阈值。 (Optional) */
    Score_threshold *int `json:"score_threshold"`
}
