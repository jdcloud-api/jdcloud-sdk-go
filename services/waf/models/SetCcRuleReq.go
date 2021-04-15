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


type SetCcRuleReq struct {

    /* WAF实例id  */
    WafInstanceId string `json:"wafInstanceId"`

    /* 域名  */
    Domain string `json:"domain"`

    /* 规则名称  */
    RuleName string `json:"ruleName"`

    /* uri以/开头  */
    Uri string `json:"uri"`

    /* 精确匹配0  前缀匹配1 包含匹配2 后缀匹配5 (Optional) */
    MatchType int `json:"matchType"`

    /* 检测周期，单位是秒，[30~600]  */
    DetectPeriod int `json:"detectPeriod"`

    /* ip访问次数，[1~9999999]  */
    SingleIpLimit int `json:"singleIpLimit"`

    /* 阻断类型 3:封禁，2:人机交互  */
    BlockType int `json:"blockType"`

    /* block 持续时间，单位为分钟 [1~9999999]  */
    BlockTime int `json:"blockTime"`

    /* blockType 为3 时，为自定义页面名称，缺省为default (Optional) */
    Redirection string `json:"redirection"`

    /* cc 统计维度，ip或cookie (Optional) */
    Dimension string `json:"dimension"`

    /* cookiename, 只有当 dimension 为 cookie 时才有效 (Optional) */
    Dmvalue string `json:"dmvalue"`
}
