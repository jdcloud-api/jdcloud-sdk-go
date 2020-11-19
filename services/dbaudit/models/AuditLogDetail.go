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


type AuditLogDetail struct {

    /* 客户端IP (Optional) */
    ClientIp string `json:"clientIp"`

    /* 客户端IP名称 (Optional) */
    ClientIpName string `json:"clientIpName"`

    /* 客户端端口 (Optional) */
    ClientPort int `json:"clientPort"`

    /* 客户端MAC (Optional) */
    ClientMac string `json:"clientMac"`

    /* 客户端主机名称 (Optional) */
    ClientHostName string `json:"clientHostName"`

    /* 客户端主机MAC地址 (Optional) */
    ClientMacAddr string `json:"clientMacAddr"`

    /* 操作的数据库名称 (Optional) */
    DbName string `json:"dbName"`

    /* 操作的数据库表名称 (Optional) */
    TableName string `json:"tableName"`

    /* 数据库用户名 (Optional) */
    DbUser string `json:"dbUser"`

    /* 数据库工具 (Optional) */
    DbTool string `json:"dbTool"`

    /* 查询语句标识 (Optional) */
    SqlIdentity string `json:"sqlIdentity"`

    /* 操作类型 (Optional) */
    SqlType string `json:"sqlType"`

    /* 操作对象 (Optional) */
    Target string `json:"target"`

    /* 影响行数 (Optional) */
    AffectLines int `json:"affectLines"`

    /* 执行时间,如5μs,3ms (Optional) */
    Duration string `json:"duration"`

    /* 捕获时间 (Optional) */
    CaptureTime string `json:"captureTime"`

    /* SQL详细语句 (Optional) */
    SqlQuery string `json:"sqlQuery"`

    /* SQL语句执行结果 (Optional) */
    SqlResult string `json:"sqlResult"`

    /* 风险级别: 0->无风险，1->低风险，2->中风险，3->高风险 (Optional) */
    RiskLevel int `json:"riskLevel"`

    /* 命中规则ID (Optional) */
    RiskRuleId string `json:"riskRuleId"`

    /* 命中规则名称 (Optional) */
    RiskRuleName string `json:"riskRuleName"`

    /* 命中规则类型 (Optional) */
    RiskRuleType string `json:"riskRuleType"`

    /* 命中规则所属规则组名称 (Optional) */
    RiskRuleGroup string `json:"riskRuleGroup"`

    /* 命中规则详细描述 (Optional) */
    RiskDesc string `json:"riskDesc"`

    /* 执行结果：
0：默认
1：未知
2：登录成功
3：登录失败
4：超时
5：执行成功
6：执行失败
7：语句不合法
8：注销
9：会话开始
10：阻断
11：会话断开 (Optional) */
    ExecuteResult int `json:"executeResult"`
}