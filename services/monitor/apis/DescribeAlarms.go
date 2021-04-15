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

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    monitor "github.com/jdcloud-api/jdcloud-sdk-go/services/monitor/models"
)

type DescribeAlarmsRequest struct {

    core.JDCloudRequest

    /* 当前所在页，默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 页面大小，默认为20；取值范围[1, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 数据所有者，1云监控控制台; 2云鼎。默认为1 (Optional) */
    DataOwner *int `json:"dataOwner"`

    /* 产品线标识，同一个产品线下可能存在多个product，如(redis下有redis2.8cluster、redis4.0) (Optional) */
    ServiceCode *string `json:"serviceCode"`

    /* 产品标识，如redis下分多个产品(redis2.8cluster、redis4.0)。同时指定serviceCode与product时，product优先生效 (Optional) */
    Product *string `json:"product"`

    /* 产品下的维度标识，指定dimension时必须指定product (Optional) */
    Dimension *string `json:"dimension"`

    /* 规则名称 (Optional) */
    RuleName *string `json:"ruleName"`

    /* 规则类型, 1表示资源监控，6表示站点监控,7表示可用性监控 (Optional) */
    RuleType *int `json:"ruleType"`

    /* 规则状态：1为启用，0为禁用 (Optional) */
    Enabled *int `json:"enabled"`

    /* 资源的规则状态  2：报警、4：数据不足 (Optional) */
    RuleStatus *int `json:"ruleStatus"`

    /* 服务码或资源Id列表
products - 产品product，精确匹配，支持多个
resourceIds - 资源Id，精确匹配，支持多个（必须指定serviceCode、product或dimension，否则该参数不生效）
alarmIds - 规则id，精确匹配，支持多个 (Optional) */
    Filters []monitor.Filter `json:"filters"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeAlarmsRequest(
) *DescribeAlarmsRequest {

	return &DescribeAlarmsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/groupAlarms",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
	}
}

/*
 * param pageNumber: 当前所在页，默认为1 (Optional)
 * param pageSize: 页面大小，默认为20；取值范围[1, 100] (Optional)
 * param dataOwner: 数据所有者，1云监控控制台; 2云鼎。默认为1 (Optional)
 * param serviceCode: 产品线标识，同一个产品线下可能存在多个product，如(redis下有redis2.8cluster、redis4.0) (Optional)
 * param product: 产品标识，如redis下分多个产品(redis2.8cluster、redis4.0)。同时指定serviceCode与product时，product优先生效 (Optional)
 * param dimension: 产品下的维度标识，指定dimension时必须指定product (Optional)
 * param ruleName: 规则名称 (Optional)
 * param ruleType: 规则类型, 1表示资源监控，6表示站点监控,7表示可用性监控 (Optional)
 * param enabled: 规则状态：1为启用，0为禁用 (Optional)
 * param ruleStatus: 资源的规则状态  2：报警、4：数据不足 (Optional)
 * param filters: 服务码或资源Id列表
products - 产品product，精确匹配，支持多个
resourceIds - 资源Id，精确匹配，支持多个（必须指定serviceCode、product或dimension，否则该参数不生效）
alarmIds - 规则id，精确匹配，支持多个 (Optional)
 */
func NewDescribeAlarmsRequestWithAllParams(
    pageNumber *int,
    pageSize *int,
    dataOwner *int,
    serviceCode *string,
    product *string,
    dimension *string,
    ruleName *string,
    ruleType *int,
    enabled *int,
    ruleStatus *int,
    filters []monitor.Filter,
) *DescribeAlarmsRequest {

    return &DescribeAlarmsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/groupAlarms",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        PageNumber: pageNumber,
        PageSize: pageSize,
        DataOwner: dataOwner,
        ServiceCode: serviceCode,
        Product: product,
        Dimension: dimension,
        RuleName: ruleName,
        RuleType: ruleType,
        Enabled: enabled,
        RuleStatus: ruleStatus,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeAlarmsRequestWithoutParam() *DescribeAlarmsRequest {

    return &DescribeAlarmsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/groupAlarms",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param pageNumber: 当前所在页，默认为1(Optional) */
func (r *DescribeAlarmsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 页面大小，默认为20；取值范围[1, 100](Optional) */
func (r *DescribeAlarmsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param dataOwner: 数据所有者，1云监控控制台; 2云鼎。默认为1(Optional) */
func (r *DescribeAlarmsRequest) SetDataOwner(dataOwner int) {
    r.DataOwner = &dataOwner
}

/* param serviceCode: 产品线标识，同一个产品线下可能存在多个product，如(redis下有redis2.8cluster、redis4.0)(Optional) */
func (r *DescribeAlarmsRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = &serviceCode
}

/* param product: 产品标识，如redis下分多个产品(redis2.8cluster、redis4.0)。同时指定serviceCode与product时，product优先生效(Optional) */
func (r *DescribeAlarmsRequest) SetProduct(product string) {
    r.Product = &product
}

/* param dimension: 产品下的维度标识，指定dimension时必须指定product(Optional) */
func (r *DescribeAlarmsRequest) SetDimension(dimension string) {
    r.Dimension = &dimension
}

/* param ruleName: 规则名称(Optional) */
func (r *DescribeAlarmsRequest) SetRuleName(ruleName string) {
    r.RuleName = &ruleName
}

/* param ruleType: 规则类型, 1表示资源监控，6表示站点监控,7表示可用性监控(Optional) */
func (r *DescribeAlarmsRequest) SetRuleType(ruleType int) {
    r.RuleType = &ruleType
}

/* param enabled: 规则状态：1为启用，0为禁用(Optional) */
func (r *DescribeAlarmsRequest) SetEnabled(enabled int) {
    r.Enabled = &enabled
}

/* param ruleStatus: 资源的规则状态  2：报警、4：数据不足(Optional) */
func (r *DescribeAlarmsRequest) SetRuleStatus(ruleStatus int) {
    r.RuleStatus = &ruleStatus
}

/* param filters: 服务码或资源Id列表
products - 产品product，精确匹配，支持多个
resourceIds - 资源Id，精确匹配，支持多个（必须指定serviceCode、product或dimension，否则该参数不生效）
alarmIds - 规则id，精确匹配，支持多个(Optional) */
func (r *DescribeAlarmsRequest) SetFilters(filters []monitor.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeAlarmsRequest) GetRegionId() string {
    return ""
}

type DescribeAlarmsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeAlarmsResult `json:"result"`
}

type DescribeAlarmsResult struct {
    AlarmList []monitor.DescribeGroupAlarm `json:"alarmList"`
    NumberPages int64 `json:"numberPages"`
    NumberRecords int64 `json:"numberRecords"`
    PageNumber int64 `json:"pageNumber"`
    PageSize int64 `json:"pageSize"`
}