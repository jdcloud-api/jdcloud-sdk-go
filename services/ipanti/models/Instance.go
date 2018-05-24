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


type Instance struct {

    /* 实例id (Optional) */
    InstanceId int64 `json:"instanceId"`

    /* 实例名称 (Optional) */
    Name string `json:"name"`

    /* 线路，UNICOM、TELECOM (Optional) */
    Carrier string `json:"carrier"`

    /* 触发弹性带宽的次数 (Optional) */
    ElasticTriggerCount int `json:"elasticTriggerCount"`

    /* 超峰次数 (Optional) */
    AbovePeakCount int `json:"abovePeakCount"`

    /* 保底带宽 (Optional) */
    Bp int `json:"bp"`

    /* 弹性带宽 (Optional) */
    Ep int `json:"ep"`

    /* 业务带宽大小 (Optional) */
    Bw int `json:"bw"`

    /* cc阈值大小 (Optional) */
    CcThreshold int `json:"ccThreshold"`

    /* 非网站类规则数 (Optional) */
    RuleCount int `json:"ruleCount"`

    /* 网站类规则数 (Optional) */
    WebRuleCount int `json:"webRuleCount"`

    /* PAID|ARREARS|EXPIRED (Optional) */
    ChargeStatus string `json:"chargeStatus"`

    /* SAFE|CLEANING|BLOCKING (Optional) */
    SecurityStatus string `json:"securityStatus"`

    /* 实例的创建的时间 (Optional) */
    CreateTime int64 `json:"createTime"`

    /* 实例的过期时间 (Optional) */
    ExpireTime int64 `json:"expireTime"`

    /* 资源id，升级和续费时使用 (Optional) */
    ResourceId string `json:"resourceId"`

    /* cc防护模式，0正常、1紧急、2宽松、3自定义 (Optional) */
    CcProtectMode int `json:"ccProtectMode"`

    /* cc开关状态，0关闭，1开启 (Optional) */
    CcProtectStatus int `json:"ccProtectStatus"`

    /* cc防护模式为自定义时的限速大小 (Optional) */
    CcSpeedLimit int `json:"ccSpeedLimit"`

    /* cc防护模式为自定义时的限速周期 (Optional) */
    CcSpeedPeriod int `json:"ccSpeedPeriod"`

    /* ip黑名单列表 (Optional) */
    IpBlackList []string `json:"ipBlackList"`

    /* ip黑名单状态，0关闭，1开启 (Optional) */
    IpBlackStatus int `json:"ipBlackStatus"`

    /* ip白名单列表 (Optional) */
    IpWhiteList []string `json:"ipWhiteList"`

    /* ip白名单状态，0关闭，1开启 (Optional) */
    IpWhiteStatus int `json:"ipWhiteStatus"`

    /* 用户pin (Optional) */
    Pin string `json:"pin"`

    /* url白名单列表 (Optional) */
    UrlWhitelist []string `json:"urlWhitelist"`

    /* url白名单状态，0关闭，1开启 (Optional) */
    UrlWhitelistStatus int `json:"urlWhitelistStatus"`
}
