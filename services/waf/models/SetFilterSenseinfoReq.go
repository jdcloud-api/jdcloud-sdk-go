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


type SetFilterSenseinfoReq struct {

    /* 更新时必须且非0，新增时可选若传需传0 (Optional) */
    Id int `json:"id"`

    /* WAF实例id  */
    WafInstanceId string `json:"wafInstanceId"`

    /* 域名  */
    Domain string `json:"domain"`

    /* equal prefix contain之一  */
    MatchLogic string `json:"matchLogic"`

    /* 从 idcard creditcard phone选，多选时使用','分隔  */
    MatchType string `json:"matchType"`

    /* notice或filter  */
    MatchAction string `json:"matchAction"`

    /* 规则名称  */
    RuleName string `json:"ruleName"`

    /* uri匹配值  */
    MatchVal string `json:"matchVal"`
}
