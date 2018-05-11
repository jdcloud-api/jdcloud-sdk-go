// Copyright 2018-2025 JDCLOUD.COM
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

package client

import (
    . "github.com/jdcloud-api/jdcloud-sdk-go/core"
    . "github.com/jdcloud-api/jdcloud-sdk-go/services/monitor/apis"
    "encoding/json"
    "errors"
)

type MonitorClient struct {
    JDCloudClient
}

func NewMonitorClient(credential *Credential) *MonitorClient {
    if credential == nil {
        return nil
    }

    config := NewConfig()
    config.SetEndpoint("monitor.jdcloud-api.com")

    return &MonitorClient{
        JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "monitor",
            Revision:    "0.3.2",
            Logger:      NewDefaultLogger(LOG_INFO),
        }}
}

func (c *MonitorClient) SetConfig(config *Config) {
    c.Config = *config
}

func (c *MonitorClient) SetLogger(logger Logger) {
    c.Logger = logger
}

/* 批量删除规则 */
func (c *MonitorClient) DeleteAlarms(request *DeleteAlarmsRequest) (*DeleteAlarmsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DeleteAlarmsResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 查看某资源的监控数据 */
func (c *MonitorClient) DescribeMetricData(request *DescribeMetricDataRequest) (*DescribeMetricDataResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DescribeMetricDataResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 修改已创建的报警规则 */
func (c *MonitorClient) UpdateAlarm(request *UpdateAlarmRequest) (*UpdateAlarmResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &UpdateAlarmResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 查询报警历史 */
func (c *MonitorClient) DescribeAlarmHistory(request *DescribeAlarmHistoryRequest) (*DescribeAlarmHistoryResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DescribeAlarmHistoryResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 根据产品线查询可用监控项列表 */
func (c *MonitorClient) DescribeMetrics(request *DescribeMetricsRequest) (*DescribeMetricsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DescribeMetricsResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 禁用报警规则。报警规则禁用后，将停止探测实例的监控项数据。 */
func (c *MonitorClient) DisableAlarm(request *DisableAlarmRequest) (*DisableAlarmResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DisableAlarmResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 创建报警规则，可以为某一个实例创建报警规则，也可以为多个实例同时创建报警规则。 */
func (c *MonitorClient) CreateAlarm(request *CreateAlarmRequest) (*CreateAlarmResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &CreateAlarmResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 查询监控规则 */
func (c *MonitorClient) DescribeAlarms(request *DescribeAlarmsRequest) (*DescribeAlarmsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DescribeAlarmsResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 查询规则详情 */
func (c *MonitorClient) DescribeAlarmsByID(request *DescribeAlarmsByIDRequest) (*DescribeAlarmsByIDResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DescribeAlarmsByIDResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 查询可用创建监控规则的指标列表 */
func (c *MonitorClient) DescribeMetricsForCreateAlarm(request *DescribeMetricsForCreateAlarmRequest) (*DescribeMetricsForCreateAlarmResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DescribeMetricsForCreateAlarmResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 启用报警规则，当客户的报警规则处于停止状态时，可以使用此接口启用报警规则。 */
func (c *MonitorClient) EnableAlarm(request *EnableAlarmRequest) (*EnableAlarmResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &EnableAlarmResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

