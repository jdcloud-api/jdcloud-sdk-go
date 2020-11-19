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


type DomainConfig struct {

    /* 域名 (Optional) */
    Domain string `json:"domain"`

    /* cname域名 (Optional) */
    Cname string `json:"cname"`

    /* 绑定的证书名称 (Optional) */
    CertName string `json:"certName"`

    /* 网站lb配置 (Optional) */
    LbConf LbConf `json:"lbConf"`

    /* 网站dns配置 (Optional) */
    DnsStatus DnsStatus `json:"dnsStatus"`

    /* 网站waf防护配置 (Optional) */
    WafConf WafConf `json:"wafConf"`

    /* 网站cc防护配置 (Optional) */
    CcConf CcConf `json:"ccConf"`

    /* 网站acl防护配置 (Optional) */
    AclConf AclConf `json:"aclConf"`

    /* 网站恶意ip防护配置 (Optional) */
    IpbanConf IpbanConf `json:"ipbanConf"`

    /* 网站过去两天攻击情况 (Optional) */
    LastAttackReport LastAttackReport `json:"lastAttackReport"`

    /* 网站防爬虫防护配置 (Optional) */
    AntispiderConf EnableConf `json:"antispiderConf"`

    /* 网站防篡改防护配置 (Optional) */
    WebcacheConf EnableConf `json:"webcacheConf"`

    /* 是否关闭waf， 1表示是 (Optional) */
    DisableWaf int `json:"disableWaf"`

    /* 网站白名单防护配置 (Optional) */
    SkipConf SkipConf `json:"skipConf"`

    /* 网站黑名单防护配置 (Optional) */
    DenyConf DenyConf `json:"denyConf"`

    /* 网站web自定义规则防护配置 (Optional) */
    WebUserdefConf WebUserdefConf `json:"webUserdefConf"`

    /* 网站限速规则防护配置 (Optional) */
    RatelimitConf RatelimitConf `json:"ratelimitConf"`

    /* 网站自定义页面配置 (Optional) */
    UserDefPageConf UserDefPageConf `json:"userDefPageConf"`

    /* 网站请求头管理防护配置 (Optional) */
    FilterHeaderConf FilterHeaderConf `json:"filterHeaderConf"`

    /* 网站敏感信息防护配置 (Optional) */
    FilterSenseConf FilterSenseConf `json:"filterSenseConf"`

    /* 网站智能语义引擎防护配置 (Optional) */
    IntSemConf IntSemConf `json:"intSemConf"`

    /* 网站uri重写规则配置 (Optional) */
    UriRewriteConf UriRewriteConf `json:"uriRewriteConf"`

    /* 网站威胁情报防护配置 (Optional) */
    ThreatinfoConf EnableConf `json:"threatinfoConf"`

    /* proxy缓存配置 (Optional) */
    ProxycacheConf EnableConf `json:"proxycacheConf"`

    /* bot配置 (Optional) */
    BotConf EnableConf `json:"botConf"`
}